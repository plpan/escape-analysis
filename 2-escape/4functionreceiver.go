package main

type Y struct {
	s *string
}

func (y *Y) f(s *string) {
	y.s = s
}

func main() {
	y := new(Y)
	s := "Hello, world"
	y.f(&s)
}
