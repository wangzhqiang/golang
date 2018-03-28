package main

import "fmt"

type Any interface{}
type EvalFunc func(Any)(Any,Any)

func main() {
	evenFunc := func(state Any) (Any,Any) {
		os := state.(int)
		ns := os +2
		return os,ns
	}
	even := BuildLayIntEvaluator(evenFunc,0)

	for i :=0;i <10; i++ {
		fmt.Printf("%vth even: %v\n",i,even())
	}
}

func BuildLazyEvaluator(evalfunc EvalFunc,initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal,actState = evalfunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <- retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLayIntEvaluator(evalFunc EvalFunc,initState Any) func()int {
	ef := BuildLayIntEvaluator(evalFunc,initState)
	return func() int{
		return ef().(int)
	}
}