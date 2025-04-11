package main //main go file er package hobe main

import(
	"fmt"
	"tasnia/mathlib"
)  //go er built-in package(fmt=format)

var v=10 

func main() {

	fmt.Println("Hello world") //Println er ln=line.ekta line e print korbe then enter diye dibe

	// variable declaration and initialization
	a := 10       //variable gula RAM e store hoy
	var x int = 5 //var name data_type
	var y = 6.79  //data-type na lekhleo go bujhe ney
	b := true
	c:=30
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
	sum:=sum(a,c)
	fmt.Println(sum)
	// block: {} jekono curly-braces and block er inside e joto variable thake shegulo local variable

	//package scope: go mod init tasnia
	mathlib.Add(5,6)
	fmt.Println(mathlib.Money)


	//shadowing
	if age>18{
		v:=47 //global scope eo v ache and ei local scope eo v initialize kora ache.erokom case holo variable shadowing and local scope e v er value shadow hoise 
		fmt.Println(v)
	}
	fmt.Println(v)
}


