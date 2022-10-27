package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello World")

	nama := "faisal"
	var umur int = 12
	var decimal = 1.23456789

	fmt.Printf("%s is a %T \n", nama, nama)
	fmt.Printf("%d is a %T \n", umur, umur)
	fmt.Printf("%.3f is a %T \n", decimal, decimal)

	var exist bool = true
	fmt.Printf("exist? %t	\n", exist)

	msg := `Hello
	World`
	fmt.Println(msg)

	var nickname string = "john"

	var lastname string
	lastname = "wick"

	fmt.Printf("%s %s\n", nickname, lastname)

	name := "faisal hafizh"
	name = "faisal h"

	fmt.Println(name)

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Printf("%s %s %s \n", first, second, third)

	fourth, fifth, sixth := 4, 5, 6

	fmt.Printf("%d %d %d \n", fourth, fifth, sixth)

	a := 10
	_ = a // without this line, a wants to be used

	const pi = 3.14
	const negara string = "indonesia"

	// error
	// pi = 3
	// negara = "jepang"

	var arr = []string{"satu", "dua", "tiga", "empat"}

	for i, a := range arr {
		fmt.Printf("%d = %s \n", i, a)
	}

	aa, bb := "Hello", "World"
	fmt.Println(swap(aa, bb))

	type person struct {
		name string
		age  int
	}

	type student struct {
		person
		jurusan string
	}

	var student1 student

	student1.name = "budi"
	student1.age = 10
	student1.jurusan = "seni"

	student2 := student{
		person: person{
			name: "faisal",
			age:  20,
		},
		jurusan: "mesin",
	}
	_ = student2

	var x = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	y := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Println(x, y)

	var xx = make(map[string]string)
	xx["brand"] = "Ford"
	xx["model"] = "Mustang"
	xx["year"] = "1964"

	yy := make(map[string]int)
	yy["Oslo"] = 1
	yy["Bergen"] = 2
	yy["Trondheim"] = 3
	yy["Stavanger"] = 4

}
