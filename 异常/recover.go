package main


func Try(fun func(),handle func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handle(err)
		}
	}()
	fun()
}

func main(){
	Try(func() {
		panic("foo")
	}, func(e interface{}) {
		print(e)
	})
}
