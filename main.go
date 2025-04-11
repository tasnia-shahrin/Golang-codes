package main //main go file er package hobe main

import "fmt" //go er built-in package(fmt=format)

func main() {

	fmt.Println("Hello world") //Println er ln=line.ekta line e print korbe then enter diye dibe

	// variable declaration and initialization
	a := 10  //variable gula RAM e store hoy
	var x int = 5 //var name data_type
	var y=6.79 //data-type na lekhleo go bujhe ney
	b:=true
	b=false //first e kono value declaration er time e colon dite hoy.next e oi var e value assign korte gele colon dite hobe na
	const p="tasnia"
	//p="orin"
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(b)
	fmt.Println(p)

	// if-else
	age:=24
	
	if age>18{
		fmt.Println("you're eligible to be married")
	}else if age<18 {
		fmt.Println("you're not eligible to be married but you can love someone")
	}else{
		fmt.Println("you're not eligible to be married. you are just a teenager")
	}
	// operator: >,<,>=,<=,==,&&,||,!
}
