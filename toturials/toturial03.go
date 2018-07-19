package main

import (
	"math"
	"image/color"
	"fmt"
	"strconv"
)

type Point struct {
	X,Y float64
}

func (p Point)Distance (q Point) float64  {
	return math.Sqrt((q.X-p.X)*(q.X-p.X) + (q.Y-p.Y)*(q.Y-p.Y))
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

type Human struct {
	name, phone string
	age int
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) SayHi()  {

	fmt.Printf("Hi,I am %s you can call me on %s",h.name,h.phone)

}


func (h Human)Sing(lyrics string){
		fmt.Println("La la la la...",lyrics)

}

func (e Employee)SayHi()  {
	fmt.Printf("Hi,I am %s,I work ar %s. Call me on %s",e.name,e.company,e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func (h Human) String() string  {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

type Person struct {
	name string
	age int
}


type Element interface{}
type List [] Element

func (p Person)String() string  {

	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"

}


type Animal interface {
	Speak() string
}
type Dog struct {
}
func (d Dog) Speak() string {
	return "Woof!"
}
type Cat struct {
}
//1
func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}
func (l Llama) Speak() string {
	return "?????"
}
type JavaProgrammer struct {
}
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func doubleValue() (int,error) {

	return 1,nil

}





func main() {


	var animal Animal
	cat := Cat{}

	animal = &cat

	fmt.Println(animal.Speak())

	//animals := []Animal{Dog{}, &Cat{}, Llama{}, JavaProgrammer{}}
	//
	//
	//for _, animal := range animals {
	//	fmt.Println(animal.Speak())
	//}


	//list := make(List,3)
	//list[0] = 1
	//list[1] = "Hello"
	//list[2] = Person{"Dennis",70}
	//
	//
	//t := reflect.TypeOf(list[2])
	//tag := t.Elem().Field(0).Tag
	//fmt.Println(tag)

	//
	//for index,element := range list {
	//
	//	switch value := element.(type) {
	//	case int:
	//		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	//	case string:
	//		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	//	case Person:
	//		fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
	//	default:
	//		fmt.Println("list[%d] is of a different type", index)
	//	}
	//}

	//p := ColoredPoint{&Point{1,1},color.RGBA{111,122,133,1}}
	//q := ColoredPoint{&Point{5,4},color.RGBA{222,144,155,1}}
	//p.Point = q.Point
	//fmt.Println(p.Distance(*(q.Point)))

	//mike := Student{Human{"Mike","222-222-xxx",25},"MIT",0.00}
	//paul := Student{Human{"Paul","333-222-xxx",21},"MIT",100}
	//sam := Student{Human{"Sam","444-222-xxx",34},"MIT",1000}
	//bob := Student{Human{"Bob","555-222-xxx",32},"MIT",5000}
	//
	////定义Men类型的变量i
	//var i Men
	//
	////i能存储Student
	//i = mike
	//fmt.Println("This is Mike, a Student:")
	//i.SayHi()
	//i.Sing("November rain")
	//
	////i也能存储Employee
	//i = bob
	//fmt.Println("This is Tom, an Employee:")
	//i.SayHi()
	//i.Sing("Born to be wild")
	//
	////定义了slice Men
	//fmt.Println("Let's use a slice of Men and see what happens")
	//x := make([]Men, 3)
	////这三个都是不同类型的元素，但是他们实现了interface同一个接口
	//x[0], x[1], x[2] = paul, sam, mike
	//
	//for _, value := range x{
	//	value.SayHi()
	//}
	//Bob := Human{"Bob", "000-7777-XXX", 39}
	//fmt.Println("This Human is : ", Bob)

}
