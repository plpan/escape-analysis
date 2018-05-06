package main

type Z struct {
	a *int
	b *int
}

func sf(i int) *int {
	var z Z
	z.a = &i
	j := z.b
	return j
}

func main() {
	i := 1
	sf(i)
}
