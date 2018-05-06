package main

type X struct {
	p *int
}

func ia1() {
	a := 1
	p := new(*int)
	*p = &a // BAD: a escape
	_ = p
}

func ia2() {
	i := 1
	var x X
	x.p = &i
}

func ia3() {
	j := 1
	x := new(X)
	x.p = &j // BAD: j escape
}

func main() {
	ia1()
	ia2()
	ia3()
}
