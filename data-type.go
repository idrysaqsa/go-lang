package main

import "fmt"
func dataType() {

	var name string = "Aqsa Idris";
	var age int = 20;
	var isMarried bool = false;
	var height float32 = 1.71;

	fmt.Println("Hello ", name, " and I am ", age, " years old");
	fmt.Println("Am I married? ", isMarried);
	fmt.Println("My height is ", height);
};


func main() {
	dataType();
};