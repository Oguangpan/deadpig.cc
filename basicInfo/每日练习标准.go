/*2018年2月22日
每日标准练习初稿
收集日常编写程序可能用到的基础语法资料
每日练习有助于帮助理解掌握
数据来源于我自己对相关书籍的阅读理解,可能有不准确的地方,以后发现会修改.
*/

// 包名字的定义,一个程序只能有一个main包,作为所有包的起点
package main

// 引入包的定义, 标准库可以直接填写包名,第三方库或自己的库需要写全命名空间
// 中的完整路径. (GOPATH>src>包路径 ,比如src目录下的
// "github.com\StackExchange\wmi")
// 引入的包必须要在程序中使用,否则编译器会报错. 不用就不要引用
import (
	"errors" // 用于生成一个错误对象
	"fmt"
	"os"
	"strings"
	"time"
)

// ===========常量的定义=============
// 允许表达式 允许并行定义 知识点 iota 以后学习
// 允许申明但不使用
const Pi = 3.1415926
const Pi2 float64 = 3.1415927
const 字符串常量 string = "可以不通过指定类型,编译器可自行通过上下文判断"
const num = 12 + 14
const Pi3, Pi4, Pi5 = 3.1415, 3.55, "3.14159267"
const ( //枚举常量
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// ===========变量的定义=============
// 请遵循驼峰法命名规则,函数外定义的变量被称为包级变量.
// 赋值法定义或进行计算时编译器会自动判断所属类型.
// 包级变量使用var创立,函数级别变量使用 a := 1 方式建立更加简洁方便.
// 赋值后不可再次申明,但可更改值
var a, b *int

// var a, b, c = 4, 6, "abc" 全局
// a, b, c := 4, 6, "abc" 局部

var c int
var (
	d   bool
	str string
)
var e int = 15
var f = 15
var g bool = false
var h string = "定义的同时赋值给该变量.不过go编译器会自动判断类型."
var (
	i = 15
	j = 1.55311
	k = "它能自动识别这是字符串类型"
)

// ====================函数定义================

// 函数的主要类型为三种: 带名字的函数 匿名函数或lambda函数 还有方法methods.
// 函数的按值传递和按引用传递.go 默认使用按值传递,参数只是副本修改副本不会影响到原有变量.
// 按引用传递消耗更少内存可直接对引用对象进行操作.在函数调用时，像切片（slice）、字典（map）、
// 接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。

// 普通的函数定义
func MultiPly3Nums(a int, b int, c int) int {
	return a * b * c
}

// 函数命名返回值,构建函数的时候同时给返回值命名,可直接return不带参数,哪怕只有一个命名返回值,构建时也要用
// 小括号括起来.尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。
func getX2AndX3(intput int) (x2 int, x3 int) {
	x2 = 2 * intput
	x3 = 3 * intput
	// return x2, x3
	return
}

// 函数内通过指针的方式直接改变外部变量值, 在大型程序中可极大节约内存使用提高运行速度和效率.
// 但当程序复杂程度到达一定的数量级以后,掌控不当就会导致一些不确定的事情,需要的时候要多做注释让其他人了解.
func MultiPlyZ(a, b int, reply *int) {
	*reply = a * b
}

// 使用变长参数. 变长参数是指应对无法明确参数数量的一种参数,下限是0,就是没有参数,上限理论上是无限.
// go语言中称为 变参函数 ,定义方式是 函数最后一个参数为 ...type 的形式
// 如果要获取到的参数类型不固定,可以使用 interface{} 这种类型来获取万能类型
func 变参函数(a ...interface{}) (m int) {
	m = len(a)
	return
}

// 函数中对defer关键字的运用. defer关键字常用于释放打开的资源或对象,因为它是在函数的return之后才会执行.
// 很多时候我们会在return的时候也做一些计算,defer这时候就能起作用了.比如打开文件读写后,能安全的关闭文件.
// 解锁已经加锁的资源 打印最终报告 关闭数据库连接之类, 多个defer语句共存时采取栈式执行(先入后出).
// 它能让程序更加简洁.鼓励使用. 同时可简单的实现代码追踪功能和结合log库记录函数的参数和返回值功能.
func 追踪(s string)   { fmt.Println("进入函数: ", s) }
func 解除追踪(s string) { fmt.Println("\n离开追踪: ", s) }

func de() {
	追踪("de")
	defer 解除追踪("de")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("你以为我是最后一个显示的吗?")
	return
}

// 返回错误值的函数定义, 返回整数和错误代码。
// 并且演示了如何使用defer关键字记录参数和返回值
// 闭包 我们不希望给函数起名字的时候可以使用匿名函数,此类型函数不能独立存在但能赋值给一个变量. 然后通过
// 变量名字对函数进行调用,其实相当于aau里面定义函数的方式,

func 函数名字(l int) (m int, n error) {
	// 这是一个匿名函数(闭包)
	// 从这里可以看到,虽然我们写在前面,但是该函数还是获取到了return之后的值.
	defer func() {
		fmt.Printf("参数值: %d, 返回值: %d\n", l, m)
	}()
	l++
	fmt.Println(Pi, l)
	if l < 10 {
		return l, errors.New("l值小于10.")
	}
	m = 23
	// 这里返回的是l,而不是m,就是说无视了有名返回的定义而直接返回不相关的值.这是允许的.
	return l, nil
}

// 内置函数

// 递归函数 一个函数在自身体内调用自身,称为递归,很多算法会用到,会让程序更简练.但小心栈溢出的问题.
// 使用递归函数来计算斐波那契数列,但是这种计算是每次循环都要执行,浪费大量时间和运算力
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 使用内存优化技术,检查内存中是否已计算过的方式优化斐波那契算法. 极大缩短运算时间和硬件负担
var fibs [num]uint64

func fibonacci2(n int) (res uint64) {
	// 当检测到集合中存在计算过的元素,跳过计算,直接返回上次计算过的结果.
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	// 返回结果之前,先把计算好的结果放到集合中去.
	fibs[n] = res
	return
}

// 函数回调 函数可以作为其它函数的参数进行传递,然后在其它函数内调用执行.
func Add(a, b int) {
	fmt.Printf("此二值一为%d 一为%d 合为%d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
}

// 另外,函数除了参数可以是函数 ,返回值也可以是一个函数.
// 工厂函数概念 一个函数的返回值是另外一个函数,这样的函数称为工厂函数. 创建一系列相似的函数时非常有用.能极大
// 的精简代码, 比如下面这个例子和后面添加文件名后缀的例子.
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// 工厂函数方式给文件加后缀
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		// strings包里面的内置函数 判断字符串是否以suffix结尾.
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 验证闭包中的变量值会得到保留.
func Adder2() func(int) int {
	var x int // 初时,x值为零.但每次调用都不会清空.
	return func(delta int) int {
		x += delta
		return x
	}
}

// 特殊函数,初始化函数init,任何包中加载完毕后第一个执行的函数,常用来检查数据
// 是否正确,也可以在程序正式开始前做点什么事情. init函数和main函数都是不带参数和返回值的.
func init() {
	//TODO
}

// ================结构与方法===================
// go中没有类的感念,所以结构体是go中最重要的概念
type struct1 struct {
	i1  int
	f1  float32
	str string
}

// 定义二叉树链表. 节点包含临近节点的地址.
type struct2 struct {
	le   *struct2
	data float32
	ri   *struct2
}

// 结构体工厂
type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

// 结构体的方法
// Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。

func pt(p *struct1) {
	fmt.Printf("结构的方法测试:\n 整数是:%d\n 浮点数是:%f\n 字符串是:%s\n", p.i1, p.f1, p.str)
}

func (tn *File) ShowId() {
	fmt.Printf("当前文件ID是: %d\n", tn.fd)
}

// 特殊函数,主程序的入口,是整个程序,而不是特指这个包文件.
func main() {

	// ===========字符串操作相关=============
	str1, str2 := "my name is superman,", "true"
	// 获取字符串长度
	str1len := len(str1)
	// 拼接字符串
	var str3 = str1 + str2
	str3 += str1
	fmt.Println(str3, str1len)
	// 判断 string 里面 是否有 tr 返回布尔值
	strings.Contains("string", "tr")
	// 查找字符串返回第一次出现的位置 返回整数型
	strings.Index("string", "in")
	// 返回字符串最后一次出现的位置
	strings.LastIndex("string", "n")
	// 查找中文字符串
	// strings.IndexRune("我是无敌的", "无敌")
	// 字符串替换， 把string中所有tr替换为m
	strings.Replace("string", "tr", "m", -1)
	// 统计字符串出现次数
	strings.Count("strinsg", "s")
	// 重复生成字符串5次
	strings.Repeat("super", 5)
	// 转换成小写
	strings.ToLower("STRINGS")
	// 转换成大写
	strings.ToUpper("string")
	// 去头尾 去头 去尾 去中间
	strings.TrimSpace("  string  ")
	// strings.TrimLeft("  string")
	// strings.TrimRight("string  ")
	strings.Trim("string", "ri")
	// 切片字符串
	strings.Fields("我们 有 很多 空格 方便 用来 切割")
	切片 := strings.Split("string,sting,sing,sin,si,s", ",")
	// 拼接切片成一个字符串
	strings.Join(切片, ":")

	// 指针的操作
	var 针 int = 256
	var zhen *int = &针
	fmt.Println("针这个变量当前值是: ", *zhen)

	// ===========流程控制相关=============
	condition := true
	condition2 := false
	if condition {
		//TODO
	}

	if condition {
		//TODO
	} else {
		//TODO
	}

	if condition {
		//TODO
	} else if condition2 {
		//TODO
	} else {
		//TODO
	}
	// 可以在判断之前赋值变量或直接调用函数返回值。
	if 只在流程中有效的变量 := 5; 只在流程中有效的变量 < 6 {
		//TODO
	}
	// 对可能发生的错误进行预估,如果数小于10就报错退出程序。
	假装一个值, err := 函数名字(9)
	if err != nil {
		fmt.Println("错误: ", err)
		// 結束程序的標准做法，输入法多了切换真麻烦。
		os.Exit(1)
	} else {
		fmt.Println(假装一个值)
	}

	// 另外一种写法,不用换行，但不利于阅读。
	// if falseValue, err := 函数名字(9);err != nil {

	// 多分支流程处理
	// var selecVar int = 1
	switch selecVar := 3; {
	case selecVar < 2:
		//TODO
		fallthrough //继续下面的分支操作
	case selecVar > 2:
		//TODO
	default:
		fmt.Println("当没有任何匹配项，此项目作为默认")
	}

	// go语言中唯一的循环结构 for 也是最强大的 for
	// 基于计数器的迭代 ,可以使用多个计数器
	for i := 0; i < 5; i++ {
		fmt.Printf("这是第%d次\n", i)
	}
	// 基于条件判断的迭代
	var iterationVar1 int = 5
	for iterationVar1 >= 0 {
		iterationVar1--
		if iterationVar1 < 0 {
			continue //跳过下面的代码开始新的循环
		}
		fmt.Printf("当前还差%d次归零\n", iterationVar1+1)
	}
	fmt.Println("已归零。")

	// 无限循环 直接 for{} 就可以， barak或return跳出。
	var 无限 int = 1
	for {
		if 无限 > 5 {
			break
		}
		无限++
	}
	// 非常好用的for-range迭代结构，可迭代 集合 字典 字符串rune集合
	var 用来迭代的中文字符 string = "这是一段用来迭代的中文字符哦。"
	fmt.Printf("用来迭代的中文字符长度为%d\n", len(用来迭代的中文字符))
	for pos, char := range 用来迭代的中文字符 {
		fmt.Printf("第%d位的内容是: %c\n", pos, char)
	}

	// ===========函数相关=============

	// 调用最基础的函数
	fmt.Printf("Multiply 2*5*6 = %d\n", MultiPly3Nums(2, 5, 6))
	// 调用函数的空白符的使用. 如果函数返回多个参数,你只需要其中的一部分.可使用空白符抛弃.
	我要的, _ := getX2AndX3(5)
	fmt.Printf("我需要的只是%d\n", 我要的)
	// 直接通过指针改变函数外部变量的值.
	MultiPlyZ(12, 4, zhen)
	fmt.Println("现在 zhen 这个指针指向的变量的值已经是: ", *zhen)
	// 调用变参函数
	fmt.Printf("我一共传递了%d个参数到变参函数中.\n", 变参函数(1, 2, 3, 4, 5, 6, 7, 8))
	// 使用defer在return之后做点事情
	de()
	// 使用递归来计算斐波那契数列
	result := 0
	fmt.Println("下面开始计算斐波那契数列")
	start := time.Now()
	for i := 0; i <= 26; i++ {
		result = fibonacci(i)
		fmt.Println(result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("26次斐波那契数列计算完毕,总共消耗时间%d\n", delta)
	// 使用内存优化技术,优化斐波那契数列算法
	var result2 uint64 = 0
	start = time.Now()
	for i := 0; i < num; i++ {
		result2 = fibonacci2(i)
		fmt.Printf("优化斐波那契数列算法: %d\n", result2)
	}
	end = time.Now()
	delta = end.Sub(start)
	fmt.Printf("26次斐波那契数列计算完毕,总共消耗时间%d\n", delta)
	// 演示使用函数回调的方法
	callback(7, Add)
	// 演示闭包. 匿名函数作为返回值
	var TwoAdder = Adder(3)
	fmt.Printf("值为: %d\n", TwoAdder(5))
	// 验证闭包中的变量值会得到保留. 可以看到多次调用Adder2函数,内部变量值一直未清零.
	var SeAdder = Adder2()
	fmt.Print(SeAdder(1), " - ")
	fmt.Print(SeAdder(100), " - ")
	fmt.Println(SeAdder(1000))
	// 使用工厂函数的方式给文件名加上后缀.创建工厂函数可以大大简化很多类似函数的编写.节约编程时间,代码更加
	// 清晰简洁,增加阅读力和运行效率.
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("super"), addJpeg("sex-vedio"))

	// ===============数组和切片==============

	// 声明一个数组 数组是具有相同 唯一类型 的一组已编号且长度固定的数据项序列
	// 数组索引从0开始, 格式为 var 名字 [长度]类型
	// 数字长度过大时,传递给函数会出现性能障碍,所以最好是多利用切片的方式来使用数组.
	var 数组1 [5]int
	var 数组2 = [...]string{"长度", "可以", "使用", "...", "无视"}
	// var 多维数组1 [3][3]int
	// 切片的概念 切片是对数组一个连续片断的引用(通常是匿名的相关数组),切片可索引,切片长度可变.切片是引用,所以
	// 不需要额外的内存,使用比数组更有效率,所以比数组更常用.多个切片可以绑定同一个数组.
	// 声明格式是: var 名字 []类型 (不需要说明长度) 未初始化之前默认为nil,长度为0.
	// 初始化格式为 ; var 名字 []类型 = 数组[起点:终点]
	// 注意,不要用指针指向切片,因为切片本身就是指针.
	// 可以使用make来生成需要的切片和字典map　func make([]T, len, cap)　cap可以省略．
	var 切片1 = []int{2, 3, 4, 5, 6, 7, 8, 9}
	var 切片2 []int = make([]int, 10, 100)
	切片3 := make([]int, 20, 100) //数组100,切片只有20
	var 切片4 []int = 数组1[2:4]
	切片5 := 数组2[:]
	// 切片的重组和扩展
	切片2 = 切片2[0 : len(切片2)+1]
	切片1 = append(切片3, 5, 6, 7)
	fmt.Println(切片1, 切片2, 切片3, 切片4, 切片5)
	// map类型. map是一种特殊的数据结构. 是一种键值对的无序结合. 一个元素是键,一个元素是值. map在其他语言
	// 里面叫做字典(python) hash hashtable 等.... 声明的方式如下
	// 不要使用new,永远实用make来构造map
	var map1 map[string]int
	var map2 map[string]int
	map1 = map[string]int{"one": 1, "two": 2}
	map3 := make(map[string]float32)
	map2 = map1 // 可以直接拷贝
	map3["key1"] = 4.5
	map3["key2"] = Pi
	fmt.Println(map1["one"])
	fmt.Println(map2["one"])
	fmt.Println(map3["key2"])
	// 判断map中是否有对应的键值对
	if _, ok := map1["fuck"]; !ok {
		fmt.Println("map1中并没有fuck键值.")
	}
	// 删除键值对
	delete(map2, "one")
	// 对map进行切片操作,第一次make分配切片,第二次make切片每个map元素.
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("显示map切片的内容%v\n", items)

	// ===========结构体与方法============
	// 实例化一个结构体正常的写法
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	fmt.Printf("int is %d.\nfloat is %f.\nstring is %s\n", ms.i1, ms.f1, ms.str)
	// 更优雅的写法
	ms2 := &struct1{10, 15.5, "Chris"}
	ms3 := struct1{i1: 10, str: "Chris", f1: 15.5}
	fmt.Printf("m2.i1 = %d, m3.f1 = %f\n", ms2.i1, ms3.f1)
	// 结构体的方法
	pt(ms)
	pt(ms2)
	// 使用结构体工厂的方法实例结构体
	// 函数和方法的区别是
	// 1. 函数将变量作为参数
	// 2. 方法在变量上被调用
	f2 := NewFile(10, "./test.txt")
	f2.ShowId()
}
