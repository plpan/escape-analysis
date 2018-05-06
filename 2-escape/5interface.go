package main

type User interface {
	Echo()
}

type Student struct {
	Name string
}

func (s *Student) Echo() {

}

func main() {
	var s User
	s = new(Student)
	s.Echo() // BAD: s escape
	_ = s.(User)
}
