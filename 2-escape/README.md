### 本节我们将列举变量逃逸的情况
1. closure call (闭包调用)
    - 在闭包调用的过程中，闭包属性会逃逸到堆上
2. indirect assignment
	- 对间接引用的赋值，会使变量逃逸