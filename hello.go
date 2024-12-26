package main

import (
	"fmt";
);

func sayHello() {
	var name string = "Aqsa Idris";
	var age int = 20;

	fmt.Println("Hello ", name, " and I am ", age, " years old");
};

func main() {
	fmt.Println("Hello World");
	sayHello();
};