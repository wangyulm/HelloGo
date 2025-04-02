package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	////基础打印
	fmt.Println("hello world!")
	////基础的变量声明 方式
	// var a = 100
	// var b = "wy"
	// var c = 0.17
	// var d func() bool
	// var e struct {
	// 	x int
	// }
	////批量命名的方式
	// var(
	// 	a int
	// 	b string
	// 	c []float32
	// 	d func()bool
	// 	e struct{
	// 		x int
	// 	}
	// )

	////一个变量交换的小算法,似乎对n个变量也适用
	a := 100 //一种更为精简的变量声明方式（左值必须是没有声明过的变量）
	b := 200
	a, b = b, a
	fmt.Println(a, b)
	c := 300
	a, b, c = c, a, b
	fmt.Println(a, b, c)

	////匿名变量,使用时在变量声明的地方使用下划线替代即可，匿名变量不占用命名空间，不会分配内存，匿名变量之间也不会因多次声明而无法使用
	a, _ = GetData()
	_, b = GetData()
	fmt.Println(a, b)
	e := 10.2
	d := 20.66666
	fmt.Println(e, d)
	//fmt.Println(fmt.Sprintf("%.2f\n", math.Pi)) //go语言限制float精度的办法
	fmt.Printf("%.2f", math.Pi)

	////输出正弦图像
	const size = 300
	pic := image.NewGray(image.Rect(0, 0, size, size)) //根据给定大小创建灰度图
	//变量所有像素点
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255}) //填充为白色
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size  // 让sin值的范围在0-2Pi之间
		y := size/2 - math.Sin(s)*size/2      // sin的幅度为一般的像素，向下偏移一半在翻转
		pic.SetGray(x, int(y), color.Gray{0}) // 黑色绘制轨迹
	}

	file, err := os.Create("output.png") //在当前目录创建一个名为output.png的文件，如果存在则覆盖，file是一个指向文件的指针，err捕获创建文件的错误，没有则为nil
	if err != nil {
		panic(err) //检测到错误则用panic抛出并打印，生产环境中建议优雅的处理错误
	}
	defer file.Close()    //确保程序退出时关闭文件句柄，防止文件泄露或数据未写入，defer语句会在函数返回前执行，适合处理资源清理
	png.Encode(file, pic) //将图像pic编码为png格式，并写入到file指向的文件中

	////多行字符串,反引号·``之间的数值按原样打印，所有转义字符均无效，代码也不会被识别
	str := `第一行
	第二行
	第三行
	\r\n
	    `
	fmt.Println(str)

	////Go的字符
	var char1 byte = 'a'
	var char2 rune = '你'
	fmt.Printf("%d, %T\n", char1, char1) //97, uint8
	fmt.Printf("%d, %T\n", char2, char2) //20320, int32

	////切片的创建
	//var pie [5]int //创建一个整型的切片,pie为切片名，5为切片大小，int为切片类型
	pie := make([]int, 5) //另一种创建方式
	for i := 0; i < 5; i++ {
		pie[i] = i
	}
	for i := 0; i < 5; i++ {
		fmt.Println(pie[i])
	}
	//字符串也可以按切片的方式操作
	str1 := "hello world"
	fmt.Println(str1[6:]) //world
	a = 100
	b = 200
	pa := &a
	pb := &b
	fmt.Println(*pa, *pb)

	swap(pa, pb)
	fmt.Println(*pa, *pb)
	fmt.Println(split(100))

}
func swap(x, y *int) { //函数定义的方法之一，在括号中写变量1，变量2.... 类型
	//x, y = y, x  //在函数中无论是交换地址还是交换值都无法影响主函数中的变量
	t := *x
	*x = *y
	*y = t //经过如上代码后，x与y对应的值可以真正交换
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //该语句等效于return x,y
	//go的返回值可以被提前命名，如果这样做，则它们将被视为在函数顶部定义的变量，如本例子中的x、y，不带参数的返回语句，称为a "naked" return
	//nake return应只在短函数中使用，因为它们会影响长函数的可读性。
}
func GetData() (int, int) {
	return 1000, 2000
}
