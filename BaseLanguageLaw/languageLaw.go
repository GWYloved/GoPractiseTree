package main

import "fmt"

//import "fmt"

//import (
//	"fmt"
//	"unsafe"
//)

//import (
//	"go/types"
//	"unsafe"
//)

//import "fmt"

func main()  {
	//fmt.Println("helloworld")//换行不需要分号
	//fmt.Printf("NONONO")

	//单行注释
	/*
	多行注释
	 */

	 //变量
	//var age int
	//age = 12
	//print(age)
	//
	//var num = 123
	//var str = "123"
	//print(num)
	//print(str)
	/*
	fmt的打印是输出标准输出流，可以进行部分更改，包括打印出类型，打印出值等等
	而print的打印是普通打印，打印功能简单
	 */

	//var a = true
	//var b = 21
	//var c = a
	//
	//d:= a// 使用:=不需要申明，但是左边的应该是没有声明的，否则会编译报错

	//var num1, num2, num3 = 1, 2, 3

	// 常量
	//const const1 = "abc"
	//const const2 = "abc"
	//
	//const (
	//	a = iota //iota = 0
	//	b  //iota = 1
	//	c  //iota = 2
	//	d  //iota = 3
	//	e = 100  //iota = 4
	//	f  //iota = 5
	//	h = iota //iota = 6
	//)
	////println(d)
	////println(f)
	////println(h)
	//
	//const (
	//	i = 1 <<iota //1左移0位
	//	j = 3 << iota//3左移1位 0x11 = 3 -> 0x110 = 6
	//	k  //3左移2位 0x11 = 3 -> 0x1100 = 12
	//	l  //3左移3位 0x11 -> 0x11000 = 24
	//)
	////println(i)
	////println(j)
	////println(k)
	////println(l)
	//var address = "1"
	//println(&address)	//取地址
	//var ptr = &address
	//println(*ptr)		//指针寻值

	//条件判断

	//var a = 1
	//
	////if a < 1 {//不需要括号
	////	println("a < 1")
	////}
	////if (a == 1) {//可以有括号
	////	println("a == 1")
	////}
	//
	//switch a {
	//case 1:
	//	println("1")
	//case 2:
	//	println("2")
	//}
	//var c char = "c"

	////闭包相关
	//nextNumber := getSequence()
	//println(nextNumber())
	//println(nextNumber())
	//nextNumber1 := getSequence()
	//println(nextNumber1())

	//指定接受者的函数相关
	//var c1 Circle
	//c1.radius = 10.00
	//println(c1.getArea())

	//pointer初始化
	//var p unsafe.Pointer //pointer 初始化为 nil
	//fmt.Printf(p)


	//数组相关
	//var balance[10] float32 //声明

	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0} //初始化
	//fmt.Println(balance)

	//var balance = [...]float32{1000.0, 2,3,4,5,5}//这种方式没有设置数组的大小，会根据实际多少来设置
	//var a = 10
	//var ip *int //指针，指向的应该是一个地址，ip = &a, 然后使用*ip的方式来寻址
	//fmt.Printf("变量的地址: %x\n", &a)
	//println(ip == nil)//nil 为空指针
	//println(*ip)

	//自定义结构体相关
	//var Book1 Books
	//Book1.author = "zuozhe"
	//Book1.book_id = "1"
	//Book1.title = "biaoti"
	//Book1.subject = "subject"
	//
	////fmt.Printf(Book1)
	//println(Book1.subject)
	//println(Book1.title)
	//println(Book1.book_id)
	//println(Book1.author)

	//切片相关
	//Go 语言切片是对数组的抽象。
	//s := []int {1,2,3,4,5,6,7}
	////println(len(s))
	////println(s
	//fmt.Printf("%v",s)
	//b := s[1:]//进行切片，[startIndex : endIndex]
	////println(len(b))
	//fmt.Printf("%v",b)
	//
	//b = s[2:4]
	//fmt.Printf("%v",b)
	//
	//b = make([]int, len(s), cap(s)) //切片创造方法
	//fmt.Printf("len = %d, cap = %d, slice = %v",len(s),cap(s),b)

	////range相关
	//s := []int {1, 2, 3, 4, 5, 6, 7}
	//for index, num := range s{
	//	fmt.Printf("%d-%d \n",index,num)
	//}
	////range 将一个数组分为index - value结构，这样假如不要index ,可以使用_符号
	//for i,c := range "hehehhe"{
	//	fmt.Println(i,c)
	//}

	//map相关
	//var countryCapitalMap map[string] string
	//countryCapitalMap = make(map[string] string)
	//
	//countryCapitalMap["France"] = "Paris"
	//countryCapitalMap["Italy"] = "Rome"
	//countryCapitalMap["Japan"] = "Tokyo"
	//countryCapitalMap["India"] = "New Delhi"
	//
	//for country:= range countryCapitalMap{
	//	fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	//}
	//
	//capital, ok := countryCapitalMap["United States"] //判断键值对是否存在， ok是返回的结果
	//if ok {
	//	fmt.Println("Capital of United States is", capital)
	//}else {
	//	fmt.Println("Capital of United States is not present")
	//}
	//
	//delete(countryCapitalMap, "France")//删除元素
	//for country,capital := range countryCapitalMap{
	//	fmt.Println("Capital of",country,"is",capital)
	//}

	////递归相关
	//var i = 15
	//fmt.Println(Factorial(uint64(i)))
	//fmt.Println(fibonacci(i))

	////语言类型转换
	//var sum = 17
	//var count = 5
	//var mean float32
	//
	//mean = float32(sum)/float32(count)
	//fmt.Println(mean)

	////接口相关
	//var phone Phone
	//phone = new(NokiaPhone)
	//phone.call()
	//
	//phone = new(Iphone)
	//phone.call()


	//错误处理相关

	//correct
	//result, errorMsg := Divide(100, 10)
	//if errorMsg == "" {
	//	fmt.Println("100/10=",result)
	//}
	//
	////fault
	//result, errorMsg = Divide(100, 0)
	//if  errorMsg != "" {
	//	fmt.Println("errorMsg is :", errorMsg)
	//}
}
//语言函数

//func max(num1, num2 int) int {
//	var result int
//	if num1 > num2 {
//		result = num1
//	}else {
//		result = num2
//	}
//	return result
//}
/*
func function_name([parameter list] <- 这是行参列表，可以很多个行参) [return_types]<- 这是返回类型{
	函数体
}

go是值传递，和java一样
 */

// 闭包

//func getSequence() func() int  {
//	i := 0
//	return func() int {
//		i +=1
//		return i
//	}
//}

// 函数结构体

//type Circle struct {
//	radius float64
//}
//
//func (c Circle)getArea() float64 {
//	return 3.14 * c.radius * c.radius
//}

/*
类似下面这种就是包含了接受者的函数
func (variable_name variable_data_type <- 这就是接收者) function_name() [return_type]{
}
 */
//type Books struct {
//	title string
//	author string
//	subject string
//	book_id string
//}

////递归阶乘
//func Factorial(n uint64)(result uint64)  {
//	if n > 0 {
//		result = n * Factorial(n - 1)
//		return result
//	}
//	return 1
//}
////斐波那契数列
//func fibonacci(n int) int  {
//	if n < 2 {
//		return n
//	}
//	return fibonacci(n - 2) + fibonacci(n - 1)
//}

//// 接口
//
//type Phone interface{
//	call()
//}
//type NokiaPhone struct {
//
//}
//
//func (nokiaPhone NokiaPhone) call() {
//	fmt.Println("I am Nokia, I can call you!")
//}
//
//type Iphone struct {
//
//}
//
//func (iphone Iphone) call()  {
//	fmt.Println("I am iphone, i can call you!")
//}

////错误处理
//type DivideError struct {
//	dividee int
//	divider int
//}
//
//func (de *DivideError)Error() string  {
//	strFormat := `
//	Cannot proceed, the divider is zero.
//	dividee: %d
//	divider: 0
//`
//	return fmt.Sprintf(strFormat, de.dividee)
//}
//func Divide(varDividee int, varDivider int)(result int, errorMsg string)  {
//	if varDivider == 0 {
//		dData := DivideError{
//			dividee:varDividee,
//			divider:varDivider,
//		}
//		errorMsg = dData.Error()
//		return
//	}else {
//		return varDividee / varDivider, ""
//	}
//}