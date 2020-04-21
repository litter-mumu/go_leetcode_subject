package main

import "fmt"

func main() {
	day05()
}

// day01 下面这段代码输出的内容：
func day01() {
	defer func() { fmt.Println("打印前") }()
	defer func() {
		fmt.Println("打印中")
		// 看看捕获后和捕获前的打印顺序
		//err := recover()
		//if err != nil {
		//	fmt.Println(err)
		//}
	}()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

//参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic。

// day02 下面这段代码输出什么，说明原因。
func day02() {
	slice := []int{1, 2, 3, 4, 5}
	m := make(map[int]*int, 5)
	for i, v := range slice {
		//TODO for range 循环的时候会创建每个元素的副本，而不是元素的引用
		//应该改为
		//value := v
		//m[i] = &value
		m[i] = &v
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

//参考解析：这是新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，
//而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
//所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.

// day03
func day03() {
	//1、下面两段代码输出什么。
	fun1 := func() {
		s := make([]int, 5)
		s = append(s, 1, 2, 3)
		fmt.Println(s)
	}
	fun2 := func() {
		s := make([]int, 0)
		s = append(s, 1, 2, 3, 4)
		fmt.Println(s)
	}
	fun1()
	fun2()

	//2、下面的代码有什么缺陷
	fun3 := func(x, y int) (sum int, error) {
		return x + y, nil
	}

	//3.new() 与 make() 的区别
}

/*
1.两段代码分别输出：

1[0 0 0 0 0 1 2 3]
2[1 2 3 4]
参考解析：这道题考的是使用 append 向 slice 添加元素，第一段代码常见的错误是 [1 2 3]，需要注意。

2.
参考答案：第二个返回值没有命名。
参考解析：
在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。这里的第一个返回值有命名 sum，第二个没有命名，所以错误。

3.
参考答案：
new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.
*/

// day04
func day04() {
	//1.下面这段代码能否通过编译，不能的话原因是什么；如果能，输出什么。
	fun1 := func() {
		list := new([]int)
		list = append(list, 1)
	}
	fun1()

	//2.下面这段代码能否通过编译，如果可以，输出什么？
	fun2 := func() {
		s1 := []int{1, 2, 3}
		s2 := []int{4, 5}
		s1 = append(s1, s2)
		fmt.Println(s1)
	}
	fun2()

	//3.下面这段代码能否通过编译，如果可以，输出什么？
	var (
		size := 1024
		maxSize = size * 2
	)
	fun3 := func() {
		fmt.Println(size, maxSize)
	}
	fun3()
}

/**
1.
参考答案及解析：不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，
不能对指针执行 append 操作。可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

2.
参考答案及解析：不能通过编译。append() 的第二个参数不能直接使用 slice，需使用 … 操作符，
将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)。

3.
参考答案及解析：不能通过编译。这道题的主要知识点是变量声明的简短模式，形如：x := 100。但这种声明方式有限制：
	1)必须使用显示初始化；
	2)不能提供数据类型，编译器会自动推导；
	3)只能在函数内部使用简短模式；
*/

// day05 下面这段代码能否通过编译？不能的话，原因是什么？如果通过，输出什么？
func day05() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}

/**
参考答案及解析：编译不通过 invalid operation: sm1 == sm2

这道题目考的是结构体的比较，有几个需要注意的地方：

结构体只能比较是否相等，但是不能比较大小。
相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体；
	sn3:= struct {
	   name string
	   age  int
	}{age:11,name:"qq"}
如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
	那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，
	注意：像slice、map、func 等是不能比较的。
*/

type MyInt1 int
type MyInt2 = int

func day06_1() {
	var i int = 0
	var i1 MyInt1 = i //应该写成 var i1 MyInt1 = MyInt1(1)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func day06() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}
func ap(a []int)  { a = append(a, 10) }
func app(a []int) { a[0] = 1 }

/**
type MyInt1 int 是基于类型 int 创建了新类型 MyInt1，type MyInt2 = int 是创建了 int 的类型别名 MyInt2，注意类型别名的定义时 = 。
所以，第 10 行代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；
而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。
*/
