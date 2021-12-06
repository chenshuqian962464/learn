//package calc
//
//import "github.com/gin-gonic/gin"
//
//func calc() {
//	engine := gin.Default()
//	engine.POST("/user/create", func(c *gin.Context) {
//		name := c.PostForm("name")
//		age := c.PostForm("age")
//	})
//
//	engine.GET("/index", func(c *gin.Context) {
//		//c.String(200,"<h1>come on </h1>")
//		c.Header("Content-Type", "text/html; charset=utf-8")
//		c.String(200, "<h1>aaaa</h1>")
//	})
//	engine.Run(":8080")
//}



package calc

import (
"fmt"
"github.com/gin-gonic/gin"
_ "github.com/go-sql-driver/mysql"
"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int64  `db:"id" json:"id,omitempty"`
	Firstname string `db:"firstname" json:"firstname,omitempty"`
	Age       int    `db:"age" json:"age,omitempty"`
	High      int    `db:"high" json:"high,omitempty"`
}

func calc() {
	engine := gin.Default()
	engine.POST("/user/create", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.PostForm("age")
		//fmt.Println(name, age)
		DB, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			c.String(200, err.Error())
			return
		}
		_, err = DB.Exec(`insert into user(firstname,age,high) values (?,?,180)`, name, age)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		c.String(200, "ok")
	})
	engine.GET("/index", func(c *gin.Context) {
		//c.String(200,"<h1>come on </h1>")
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, "<h1>aaaa</h1>")
	})
	engine.GET("/user/:ID", func(c *gin.Context) {
		ID := c.Param("ID")
		fmt.Println(ID)
		DB, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			c.String(200, err.Error())
			return
		}
		user := new(User)
		err = DB.Get(user, "select * from user where id=?", ID)
		if err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, user)

	})

	engine.Run(":8080")
}

//func main() {
//	var numbers []int
//	printSlice(numbers)
//
//	/* 允许追加空切片 */
//	numbers = append(numbers, 0)
//	printSlice(numbers)
//
//	/* 向切片添加一个元素 */
//	numbers = append(numbers, 1)
//	printSlice(numbers)
//
//	/* 同时添加多个元素 */
//	numbers = append(numbers, 2, 3, 4)
//	printSlice(numbers)
//
//	/* 创建切片 numbers1 是之前切片的两倍容量*/
//	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
//
//	/* 拷贝 numbers 的内容到 numbers1 */
//	copy(numbers1, numbers)
//	printSlice(numbers1)
//}
//
//func printSlice(x []int) {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}

//import (
//	"fmt"
//	//"github.com/chenshuqian962464/learn/calc"
//
//	//"github.com/chenshuqian962464/learn/calc"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//	"log"
//)
////
////func main() {
////	a := calc.Sum(1, 2)
////	fmt.Println(a)
////}
//
//type User struct {
//	ID        int64  `db:"id"`
//	Firstname string `db:"firstname"`
//	Age       int    `db:"age"`
//	High      int    `db:"high"`
//}
//
//
//func main() {
//	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(db.Ping())
//	users := make([]*User, 0, 4)
//	err = db.Select(&users, "select * from user where id in(?,?)", 1, 3)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(users)
//	for i, v := range users {
//		fmt.Printf("index:%d user:%#v \n", i, v)

	}
//
//}

//func main (){
//	var a int = 100
//	var b int = 200
//	var ret int
//
//	ret = max(a,b)
//	fmt.Printf("最大值是：%d\n",ret)
//
//}
//
//func max(num1,num2 int)int{
//	var result int
//
//	if num1>num2{
//		result = num1
//	}else{
//		result = num2
//	}
//	return result
//}
//
//

//func main(){
//	r := gin.Default()
//	r.GET("/index", func(c *gin.Context) {
//		night := c.Query("night")
//		c.String(200,night)
//
//	})
//	r.POST("/index", func(c *gin.Context) {
//		c.String(200,"ok")
//	})
//    r.Run(":8080")
//}

//func main (){
//	var a int = 100
//	var b int = 200
//	fmt.Printf("交换前a的值为：%d\n",a)
//	fmt.Printf("交换前b的值为：%d\n",b)
//
//	swap(&a,&b)
//
//	fmt.Printf("交换后a的值为：%d\n",a )
//	fmt.Printf("交换后b的值为：%d\n",b)
//}
//func swap(x *int,y *int){
//	var temp int
//	temp = *x
//	*x = *y
//	*y = temp
//}
//func main (){
//	var a int = 100
//	var b int = 200
//	fmt.Printf("交换前a的值为：%d\n",a)
//	fmt.Printf("交换前b的值为：%d\n",b)
//	swap(a,b)
//	fmt.Printf("交换后a的值为：%d\n",a )
//	fmt.Printf("交换后b的值为：%d\n",b)
//}
//func swap(x,y int)int{
//	var temp int
//	temp = x
//	x = y
//	y = temp
//	return temp;
//
//}
//func swap(x,y string)(string,string){
//	return y,x
//}
//func main (){
//    a,b:=swap("Google","Runoob")
//	fmt.Println(a,b)
//}

//func main (){
//	var a int = 100
//	var b int = 200
//	var r int
//
//	r = max(a,b)
//	fmt.Printf("最大值是：%d\n",r)
//}
//
//func max(num1,num2 int)int{
//	var result int
//
//	if (num1 > num2){
//		result = num1
//	}else{
//		result = num2
//	}
//	return result
//}
//func main (){
//	for m:=1;m<10;m++{
//		for n:=1;n<10;n++{
//			fmt.Printf("%d * %d = %d\n",n,m,m*n)
//		}
//		fmt.Println("")
//	}
//}
//func main (){
//	Nine(9)
//}
//
//func Nine (num int){
//	for i:=1;i<=num;i++{
//		for j:=1;j<=num;j++{
//			if j>i{
//				break
//			}
//			if i==j{
//				fmt.Printf("%d * %d = %d\n",j,i,i*j)
//            }else{
//	            fmt.Printf("%d * %d = %d\t",j,i,i*j)
//}
//}
//}
//}
//func main(){
//	switch{
//	case false :
//		fmt.Printf("1、case条件语句为false")
//		fallthrough
//	case true :
//		fmt.Printf("2、case条件语句为true")
//		fallthrough
//	case false :
//		fmt.Printf("3、case条件语句为false")
//		fallthrough
//	case true:
//		fmt.Printf("4、case条件语句为true")
//	case true:
//		fmt.Printf("5、case条件语句为true")
//	default:
//		fmt.Printf("6、默认case")
//
//	}
//}
//func main (){
//	var x interface{}
//
//	switch i:= x.(type){
//	case nil:
//		fmt.Printf("x的类型：%T\n",i)
//	case int:
//		fmt.Printf("x是int型")
//	case float64:
//		fmt.Printf("x是flost64型")
//	case func(int)float64:
//		fmt.Printf("x是func(int)型")
//	case bool,string:
//		fmt.Printf("x是bool或者string型")
//	default:
//		fmt.Printf("未知型")
//	}
//
//}
//func main (){
//	var grade string ="B"
//	var marks int = 90
//
//	switch marks{
//	case 90: grade = "A"
//	case 80: grade = "B"
//	case 50,60,70: grade = "C"
//	default: grade = "D"
//	}
//	switch {
//	case grade == "A":
//		fmt.Printf("优秀\n")
//	case grade == "B", grade == "C":
//		fmt.Printf("良好\n")
//	case grade == "D":
//		fmt.Printf("及格\n")
//	case grade == "F":
//		fmt.Printf("不及格\n")
//	default:
//		fmt.Printf("差\n")
//	}
//	fmt.Printf("你的等级是：%s\n",grade)
//}

//func main (){
//	var a int
//	var b int
//	fmt.Printf("请输入密码： \n")
//	fmt.Scan(&a)
//	if a ==5201314{
//		fmt.Printf("请再次输入密码：")
//		fmt.Scan(&b)
//		if b == 5201314{
//			fmt.Printf("密码正确，门锁已打开")
//		}else{
//			fmt.Printf("非法入侵，已自动报警")
//		}
//	}else{
//		fmt.Printf("非法入侵，已自动报警")
//	}
//}
//func main (){
//	var a int = 100
//	var b int = 200
//	if a == 100{
//		if b == 200{
//			fmt.Printf("a的值为100,b的值为200\n")
//		}
//	}
//	fmt.Printf("a的值为:%d\n",a)
//	fmt.Printf("b的值为:%d\n",b)
//}
//func main () {
//	var a int = 100
//	if a<20{
//		fmt.Printf("a小于20\n")
//	}else{
//		fmt.Printf("a不小于20\n")
//	}
//	fmt.Printf("a的值为%d\n",a)
//}
//

//func main () {
//	var a int = 10
//	if a < 20 {
//		fmt.Printf("a 小于20\n")
//	}
//	fmt.Printf("a的值为：%d\n",a)
//}

//func main (){
//	var a int = 20
//	var b int = 10
//	var c int
//
//	c = a + b
//	fmt.Printf("c的值为%d\n",c)
//	c = a - b
//	fmt.Printf("c的值为%d\n",c)
//	c = a * b
//	fmt.Printf("c的值为%d\n",c)
//	c = a / b
//	fmt.Printf("c的值为%d\n",c)
//
//}
//

//import "fmt"
//const(
//	i = 1<<iota
//	j = 3<<iota
//	k
//	l
//)
//func main (){
//	fmt.Println("i",i)
//	fmt.Println("j",j)
//	fmt.Println("k",k)
//	fmt.Println("l",l)
//}
//
//
//
//

import "github.com/gin-gonic/gin"

func main () {
	engine := gin.Default()
	engine.GET("/index", func(c *gin.Context) {
		//c.String(200,"<h1>come on </h1>")
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, "<h1>aaaa</h1>")
	})
	engine.Run(":8080")
}

//func main () {
//	engine :=gin.Default()
//	engine.POST("/index",func(*gin.Context){
//	})
//	engine.GET("/index",func(c *gin.Context){
//		a:= c.Query("name")
//		b:= c.Query("age")
//		c.Header("Content-Type","text/html;charset=utf-8")
//		c.String(200,"<h1>%s,%s</h1>",a,b)
//	})
//	engine.Run(":8080")
//}
//

//import "fmt"
//
//type Angel struct{
//    Name   string
//	 Sex    string
//	 Height int
//	 Weight int
//}
//func main () {
//	var person1 = new(Angel)  //声明并赋值给一个
//	//var person2 Angel
//
//	person1.Name = "three-nine"
//	person1.Sex = "man"
//	person1.Height = 183
//	person1.Weight = 145
//
//	//fmt.Printf("person1 Name : %s\n",person1.Name)
//	//fmt.Printf("person1 Sex : %s\n",person1.Sex)
//	//fmt.Printf("person1 Height : %dcm\n",person1.Height)
//	//fmt.Printf("person1 Weight : %dkg\n",person1.Weight)
//
//	person2:= person1
//	person2.Name = "张三"
//	fmt.Println(person1.Name,person2.Name)
//
//
//}
//
//
//
//
//
//
//
////type Books struct{
////	title string
////	author string
////	subject string
////	book_id int
////}
////func main () {
////	//创建一个新的结构体
////	fmt.Println(Books{"GO 语言","www.runoob.com","Go 语言教程",6495407})
////	//也可以使用key=>value格式
////	fmt.Println(Books{title:"Go 语言",author:"www.runoob.com",subject:"Go 语言教程",book_id:6495407})
////	//忽略的字段为0或者空
////	fmt.Println(Books{title:"Go 语言",author:"www.runoob.com"})
////}
//
//

//const(
//	a = "abc"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//)
//func main (){
//	println(a,b,c)
//}

//func main () {
//	const LENGTH int = 10
//	const WIDTH int = 5
//	var area int
//	const a,b,c = 1,false,"str"
//
//	area = LENGTH * WIDTH
//	fmt.Println("面积为:%d\n",area)
//	println(a,b,c)
//}

//常量的数据类型只可以是布尔型、数字型（整型、浮点型和复数）和字符串型
//var x,y int
//var (   //这种因式分解关键字的写法一般用于声明全局变量
//	a int
//	b bool
//	c string
//)
//var d,e int = 1,2
//var f,g = 123,"hello"
//func main () {
//   h,i := 123,"hello"      //var的格式属于带声明的格式   这种操作符的格式只能在函数中体现
//	println(x,y,a,b,c,d,e,f,g,h,i)
//}

//使用赋值操作符 ：=  的时候，只能被用在函数体内，不可以用于全局变量的赋值和声明。
//在函数中声明局部变量的时候，必须在代码块中使用他；全局变量是允许声明但不使用的
//如果想交换两个变量的值，必须保证两个变量的值是相同的， a,b=b,a

//
//func main () {
//	var stockcode=123
//	var enddate="2020-12-31"
//	var url = "Code=%d&endDate=%s"
//	var target_url=fmt.Sprintf(url,stockcode,enddate)
//	fmt.Println(target_url)
//}
//%d  表示整型数字    %s  表示字符串
//布尔型的值 是属于 常量   true / false
// 数字类型   整型int  浮点型float32、float64

//func main() {
//	fmt.Println("Hello,world!")
//}
//第一行package main定义了包名，你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main，package main 表示一个
//可独立执行的程序，每个go应用程序都包含一个名为main的包
//import "fmt" 告诉Go编译器这个程序需要用fmt包（的函数，或其他元素），fmt包实现了格式化IO (输入/输出)的函数
//func main （）是程序开始执行的函数。main函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数
//如果有init（）函数，则会先执行该函数

//标识符=  A~Z  a~z  0-9 _   这几种符号组成的序列   第一个字符必须是  字母  /  下划线    不能是数字！！

//func main () {
//	numbers:=[]int{1,3,4,5,6,3,2,9,12}      // 创建一个切片
//
//	var(                                   //定义一个全局变量
//		indexes []int
//		items   []int
//	)
//
//    for i,value:=range numbers{
//		if i>0 && i<4{
//			indexes = append(indexes,i)
//			items =append(items,value)
//		}
//	}
//	fmt.Println(indexes,items)
//}
//func main () {
//	nine(9)
//}
//func nine(num int){
//	for i:=1;i<=num;i++{
//		for j:=1;j<=num;j++{
//			if j>i{
//				break
//			}
//			if j==i{
//				fmt.Printf("%d * %d = %d\n",j,i,i*j)
//			}else{
//				fmt.Printf("%d * %d = %d\t",j,i,j*i)
//			}
//		}
//	}
//}

//func main() {
//	var countryCapitalMap map[string]string /*创建集合 */
//	countryCapitalMap = make(map[string]string)
//
//	/* map插入key - value对,各个国家对应的首都 */
//	countryCapitalMap [ "France" ] = "巴黎"
//	countryCapitalMap [ "Italy" ] = "罗马"
//	countryCapitalMap [ "Japan" ] = "东京"
//	countryCapitalMap [ "India " ] = "新德里"
//
//	/*使用键输出地图值 */
//	for country := range countryCapitalMap {
//		fmt.Println(country, "首都是", countryCapitalMap [country])
//	}
//
//	/*查看元素在集合中是否存在 */
//	capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
//	/*fmt.Println(capital) */
//	/*fmt.Println(ok) */
//	if (ok) {
//		fmt.Println("American 的首都是", capital)
//	} else {
//		fmt.Println("American 的首都不存在")
//	}
//}
//func main () {
//	fmt.Println("---------m-------")
//	m:=make(map[string]string)
//	m["1"]="0"
//	fmt.Printf("m outer address %p,m=%v\n",m,m)
//	passMap(m)
//	fmt.Printf("post m outer address %p,m=%v\n",m,m)
//
//}
//func passMap(m map[string]string){
//	fmt.Printf("m inner address %p\n",m)
//	m["11111111"] = "11111111"
//	fmt.Printf("post m inner address %p\n",m)
//}

//当传参为map的时候，其实传递的是指针地址。函数内外map的地址都是一样的。
//函数内部的改变会透传到函数外部。

//var(
//	slice = make([]int,5)    //使用 make() 函数来创建切片
//	slice2 = slice [2:]
//)
//func main () {
//	printSlice()
//	slice2[0] = 1
//	printSlice()
//	slice2 = append(slice2,2)
//	printSlice()
//	slice2[0] = 3
//	printSlice()
//	slice[0] =4
//	printSlice()
//}
//func printSlice(){
//	fmt.Printf("len=%d cap=%d %v\n",len(slice),cap(slice),slice)
//	fmt.Printf("len=%d cap=%d %v\n",len(slice2),cap(slice2),slice2)
//}
// 第一步我们创建了一个长度容量为5的slice，然后从slice的index为2的元素开始切到最后生成slice2，第二步我们将slice2的第0个元素改为1，打印slice和slice2，发现slice的第2个元素也变成了1。这也验证了slice和slice2共用了一个底层数组。
//第二步我们对slice2进行扩容，增加一个元素为2，此时slice2的长度和容量分别为4和6。
//第三步我们修改slice2的第0个元素为3，打印slice和slice2，发现slice的元素并未发生改变。
//第四步修改了slice的第0个元素为4，打印后发现确实只有slice本身有变化，slice2没变化。
//这验证了切片进行扩容的时候，会将引用的底层数组自身范围内的元素值，拷贝到扩容后的新的底层数组里。

//func main () {
//	numbers := []int{3,2,8,7,1,6,5,4,9,10} //切片初始化公式  表示创建了一个切片
//	var (                                 //声明全局变量（因式分解关键字的写法）
//		indexes []int                     //表示从切片中索引位置
//		items   []int                     //表示从切片中索引值
//	)
//
//	for i,value:=range numbers{        //for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
//		if i > 0 && i < 4{               //range表示一个数据范围//i 是索引  value是值  i、value 的名称是可以自定义的
//			indexes = append(indexes,i)     // 接收值=把i增加到indexes的数组里
//			items = append(items,value)
//	    }
//	}
//	fmt.Println(indexes,items)
//}

//func main () {
//	var numbers =make([]int,3,5)
//
//	printSlice(numbers)
//}
//func printSlice(x []int){
//	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
//}

//func main () {
//	Nine (9)
//}
////func Nine(num int){
////	for i:=1;i<=num;i++{
////		if i>num{
////			break
////		}
////		for j:=1;j<=num;j++{
////			if j>i{
////				break
////			}
////			if j==i{
////				fmt.Printf("%d * %d = %d\n",j,i,i*j )
////			} else{
////				fmt.Printf("%d * %d = %d\t",j,i,i*j )
////		}
////
////		}
////
////	}
////}
//func Nine(num int){
//	for i:=1;i<=num;i++{
//		for j:=1;j<=num;j++{
//			if j>i{
//				break
//			}
//			if j==i{
//				fmt.Printf("%d * %d = %d\n",j,i,i*j)
//			}else{
//				fmt.Printf("%d * %d = %d\t",j,i,i*j)
//			}
//		}
//	}
//
//}
