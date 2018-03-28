package main

func deferRet(x,y int) (z int) {
	defer z += 100
	z = x + y
	return z + 50
}

func mian() {
	i := deferRet(1,1)
	println(i)
}


