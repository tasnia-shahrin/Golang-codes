package main

import "fmt"

var a=10

//standard or named function : that has name
func add(x int,y int){ //parameter : x,y--> value receive
	sum:=x+y
	fmt.Println(sum)
}
// init function: you cannot call this.computer calls this automatically.main function er age init function shobar age call hobe automatically.init() er kono input deya jayna .even eta kono o/p o return korena.init() er pore main() call hobe
func init(){
	fmt.Println("I'm the first function that is executed first")
	fmt.Println(a)
	a=20
	// a:=20 :variable shadowing
	// fmt.Println(a) 
}

func main(){
	add(2,4) //argument:2,4 =>value pass
	fmt.Println(a)

	// anonymous function:Jar name thake na
	func(a int,b int){
		c:=a+b
		fmt.Println(c)
	}(5,7) //go te anonymous function rekhe deya jayna.sthe sthe call/execute  korte hoy.etakei bole Immediately Invoked Function Expression(IIFE)

	//Function expression or Assign function in variable: jekono func/anonymous function e variable er modde assign korai holo function expression
	sub:=func(a int,b int){
		sub:=a-b
		fmt.Println(sub)
	}
	sub(20,15)
	
}