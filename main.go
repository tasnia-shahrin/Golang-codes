package main //main go file er package hobe main

import(
	"fmt"
	"tasnia/mathlib"
)  //go er built-in package(fmt=format)

var v=10 //global scope e v:=10 lekhle error dekhabe


var arr2=[4]string {"tasnia","rumaisha","ekra","tahura"}

//pointer:
func print(numbers *[3]int){
	fmt.Println(numbers)
}

type User struct{
	Name string
	Age int
	FavFoods []int
}

func main() {

	fmt.Println("Hello world") //Println er ln=line.ekta line e print korbe then enter diye dibe

	// variable declaration and initialization
	a := 10       //variable gula RAM e store hoy
	var x int = 5 //var name data_type
	var y = 6.79  //data-type na lekhleo go bujhe ney
	b := true
	//c:=30
	b = false //first e kono value declaration er time e colon dite hoy.next e oi var e value assign korte gele colon dite hobe na
	const p = "tasnia"
	//p="orin"
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(b)
	fmt.Println(p)

	// if-else
	age := 24

	if age > 18 {
		fmt.Println("you're eligible to be married")
	} else if age < 18 {
		fmt.Println("you're not eligible to be married but you can love someone")
	} else {
		fmt.Println("you're not eligible to be married. you are just a teenager")
	}
	// operator: >,<,>=,<=,==,&&,||,!

	// switch-case
	switch a {
	case 5:
		fmt.Println("a is 5")
	case 10, 15:
		fmt.Println("a is either 10 or 15")
	default:
		fmt.Println("a is neither 5 nor 10 or 15")
	}
	// sum:=sum(a,c)
	// fmt.Println(sum)
	// block: {} jekono curly-braces and block er inside e joto variable thake shegulo local variable

	//package scope: go mod init tasnia
	mathlib.Add(5,6)
	fmt.Println(mathlib.Money)


	//shadowing
	fmt.Println(v)
	if age>18{
		v:=47 //global scope eo v ache and ei local scope eo v initialize kora ache.erokom case holo variable shadowing and local scope e v er value shadow hoise.but ekhane v er value re-assign korle(v=47) lekhle ei scope er baire v er value 47 e hoto. it's not shadowing. It's keeping the value in the same memory location. When variables in two different memory locations are created with same name, and one shadows another, then it's variable shadowing.
		fmt.Println(v)
	}
	fmt.Println(v)

	//array:
	var arr[2] int
	//shorthand: arr:=[2]int{3,6}
	arr[0]=3
	arr[1]=6
	fmt.Println(arr)
	fmt.Println(arr2)

	//pointer ===> address of memory(RAM)
	d:=20
	addr:=&d //& : address of.so &x means addr of x.this address is not fixed.this is dynamic
	fmt.Println("value of d: ",d)
	fmt.Println("Address of variable d:",addr)
	fmt.Println("value at the address:",*addr) //* means value inside of that address

	//update the value of d through pointer:
	*addr=30
	fmt.Println("value of d: ",d)
	fmt.Println("Address of variable d:",addr)
	fmt.Println("value at the address:",*addr)

	array:=[3]int{100,200,300}
	print(&array) //array er each element ekta ekta kore pass na kore array tar address pathacchi.ete time beche jabe.ekhane value gula pass korchina.array er starting address pass korchi.etai holo pass by reference.jokhn array er value pass kori,sheta pass by value

	usr:=User{
		Name: "Tasnia",
		Age: 24,
	}
	//fmt.Println(usr)
	struct_addr:=&usr
	fmt.Println(struct_addr)
	fmt.Println(*struct_addr)
	fmt.Println(struct_addr.Name)

}

