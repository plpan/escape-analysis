### 本节我们将列举变量逃逸的情况
1. closure call (闭包调用)
    - 在闭包调用的过程中，闭包属性会逃逸到堆上
2. indirect assignment
	- 对间接引用的赋值，会使变量逃逸
3. slice map
	- 对slice和map的间接赋值也会导致变量逃逸
4. function receiver
	- 将变量赋值给函数receiver指针成员，会导致变量逃逸
5. interface call function
	- 接口类型调用方法，该接口变量会逃逸