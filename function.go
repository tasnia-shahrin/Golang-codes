package main

import "fmt"



func printWelcomeMsg() {
	fmt.Println("Welcome to our application")
}
func getName() string {
	var name string
	fmt.Println("Enter your name --")
	fmt.Scanln(&name)
	return name
}
func get2Numbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter 1st number---")
	fmt.Scanln(&num1)
	fmt.Println("Enter 2nd number---")
	fmt.Scanln(&num2)
	return num1, num2
}
func sum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func display(name string, sum int) {
	fmt.Println("Hello", name)
	fmt.Println("Summation = ", sum)
}
func sayGoodbye() {
	fmt.Println("Thanks for using the application")
	fmt.Println("Goodbye!")
}

func main2() {
	printWelcomeMsg()
	name := getName()
	n1, n2 := get2Numbers()
	sum := sum(n1, n2)
	display(name, sum)
	sayGoodbye()
	

}
