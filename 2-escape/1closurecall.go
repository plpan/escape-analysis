package main

func cc1() {
	x := 0 // BAD: x escapes
	defer func(p *int) {
		*p = 1
	}(&x)
}

func cc2() {
	var y int // BAD: y escapes
	func(p *int, x int) {
		*p = x
	}(&y, 42)
}

func main() {
	cc1()
	cc2()
}
