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

//###Higher order function:
	func processOperation(a int,b int,op func(p int,q int)){
		op(a,b)
	}
	func mul(x int,y int){
		z:=x*y
		fmt.Println(z)
	}
	func call() func(p int,q int){
		return mul
	}


func main(){
	add(2,4) //argument:2,4 => value pass
	fmt.Println(a)
	processOperation(2,4,mul)
	product:=call()
	product(2,3)
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

/*####logic:
		object,property,relation
	1.first order logic:object,property,relation niye kaj kore,
	2.Higher order logic:rules niye kaj kore

	#####First order function:simple jinish niye kaj korbe
	1.standard or named function
	2.anonymous function
	3.IIFE
	4.function expression
	######Higher order function/first class function
	a function will be higher order if any one of the following is true:
	a.func can be passed as parameter
	b.function can be returned
	c.both
	###callback function:jei function ke higher order function e parameter hishabe pass kori ,shetai callback function
	##first class citizen:jei variable er moddhe all type data assign kora jay.so jei function er vitor arek func pass korte pari,take bola hoy first class function and that's why higher order function ke bole first class function


	##2 phases: 1.compilation phase(compile time) 2.execution phase(run time)
	##go build main.go=>compile it=>create main file.this is binary executable file
	##then run it by: ./main
*/  