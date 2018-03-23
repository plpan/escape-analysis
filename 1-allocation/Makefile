all: godemo cdemo clean

godemo:
	go run hello.go

cdemo:
	gcc hello.c -o hello
	./hello

alloc:
	go build -gcflags="-N -l -m -m" hello.go
	-@rm hello

.PHONY: clean

clean:
	-@rm hello