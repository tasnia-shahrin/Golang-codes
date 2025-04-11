package main

import "fmt"

var a=10

//standard or named function : that has name
func add(x int,y int){
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
	add(2,4)
	fmt.Println(a)
}