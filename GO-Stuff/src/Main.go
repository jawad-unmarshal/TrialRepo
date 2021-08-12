package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type compact struct {
	name, email string
	id          int
}

type emp struct {
	data   compact
	addr   string
	salary float64
	int
	string
}

func (det *compact) print() {
	fmt.Println("Name: ", det.name)
	det.name = "New Name"
	fmt.Println("Email: ", det.email)
	fmt.Println("ID: ", det.id)
}

func (employee *emp) print() {
	employee.data.print()
	fmt.Println("Address: ", employee.addr)
	employee.addr = "Lallantaap"
	fmt.Println("Salary: ", employee.salary)
	// fmt.Println("ID: ",det.id
}

func main() {
	// 	fmt.Println("Hello, playground")
	// 	x := 5.32
	// 	result := math.Pow(x, 8)
	// 	fmt.Println(x, "to the power of ", 8, " = ", result)
	// 	var balance = 10.0
	// 	_ = balance
	// 	const(
	// 		a = 10
	// 		b
	// 	)
	// 	// const dd [2]int= {5,6}
	// 	// fmt.Printf("%T -> %d\n",a,b)
	// 	var k rune
	// 	l := &x
	// 	_ = k
	// 	_ = l
	// 	fmt.Printf("%T",l)
	// 	fmt.Println("\n",os.Args)
	// 	var res , _= strconv.ParseFloat(os.Args[1],64)
	// 	fmt.Printf("%T\n",res)
	// 	for ;balance >=0; balance-- {

	// 		fmt.Println("MEOW")
	// 	}
	// 	arr := [2][3]int{
	// 		{0,2,3},
	// 		{4,5,6},
	// 	}
	// 	fmt.Printf("%v",arr)
	// 	grades := [...]int{
	// 		1:10,
	// 		0:5,
	// 		7:7,
	// 	}
	// 	_ = grades
	// abd := [...]int{10,11,12}
	// for i := range abd {
	// 	fmt.Println(i)
	// }
	// abd[0] = 3
	// fmt.Println(abd)
	// s:= "mamă"
	// r := strings.Split(s,"")
	// fmt.Printf("%#v",r)
	// fmt.Println(strings.Trim(s,"-"))
	// fmt.Println(strings.Repeat("-##-",5))
	// for _,r := range s {
	// 	fmt.Printf("%c\n",r)
	// }
	var n float64 = 0
	// n = n + (.01*10)
	n += 0.01
	n += 0.01
	n += 0.01
	n += 0.01
	n += 0.01
	n += 0.01
	n += 0.01
	n += 0.01
	n += 0.01
	n += 0.01

	// for i := 0; i < 1000; i++ {
	// 	if i %10 == 0 {
	// 		fmt.Println(n)
	// 	}
	// 	n += .01
	// 	if i % 100 == 0 && i!=0 {
	// 		break
	// 	}
	// }
	// fmt.Println(math.Round(n*10)/10 == 0.1)
	// newFile, _ := os.Create("Henlo.txt")
	// newFile.Close()

	file, _ := os.Create("Henlo.txt")
	file.Close()
	err := os.Remove("Henlo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// _ = os.Truncate("Henlo.txt",5)

	var s = make([]int, 5, 10)
	var a = []int{1, 2, 4}
	fmt.Println(a)
	CheckIt(a)
	fmt.Println(a)
	s[4] = 4
	// fmt.Println(len(a))
	_, _ = os.OpenFile("gg.txt", os.O_CREATE|os.O_WRONLY, 0)
	// fmt.Println(cap(a))
	fmt.Printf("%p\t%d\t%d\n", &s, cap(s), len(s))

	john := emp{
		addr:   "Street 20, London",
		salary: 3000,
		data: compact{
			id:    0,
			name:  "John Keller",
			email: "jkeller@company.com",
		},
		int:    88,
		string: "ggdd",
	}
	// fmt.Printf("%#v\n",john)
	john.print()
	fmt.Println("Printing after change")
	fmt.Printf("%#v\n", john)
	john.print()
	fmt.Println(MyFn("9"))
	var b int64 = 3322
	something := Bhaiyyaji()
	fmt.Println(*something)
	// &b = something
	_ = b

	// fmt.Println(f1(2,3))

	// s=append(s,10)
}
func CheckIt(myArr []int) {
	myArr[0] = 19
	fmt.Println(myArr)
	myArr = append(myArr, 23)
	fmt.Println("Inside function:", myArr)
	// return
	// fmt.Print("meow")

}
func MyFn(val string) int {
	ans, _ := strconv.Atoi(val)
	v1, _ := strconv.Atoi(val + val)
	ans += v1
	v1, _ = strconv.Atoi(val + val + val)
	ans += v1

	return ans

}
func Bhaiyyaji() *int64 {
	var MyLocalVal int64 = 23
	return &MyLocalVal
}
