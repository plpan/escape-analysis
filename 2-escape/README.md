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
6. struct fields
	- 这里牵涉到引用层级的问题（这个后面再介绍），结构体的多个字段是作为一个整体，对一些字段的引用会导致其他字段的引用发生逃逸
7. byte slice/string transform
	- 字符串和字节数组的转化会导致变量逃逸