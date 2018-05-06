package main

func s1() string {
	var b []byte
	b = append(b, "hello"...)
	b = append(b, "world"...)
	return string(b) // BAD: string(b) escape
}

func s2(s string) []byte {
	return []byte(s) // BAD: []byte(s) escape
}

func main() {
	s1()
	s2("hello")
}
