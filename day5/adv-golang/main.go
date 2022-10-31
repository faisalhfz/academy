package main

import (
	"adv-golang/test"
	"fmt"
	"sync"
)

type student struct {
	name string
}

func (s student) sayHello() {
	fmt.Println("\nHello ", s.name)
}

func (s *student) gantiNama(namaBaru string) {
	s.name = namaBaru
}

type say interface {
	intro() string
	// ask() string
}

type employee struct {
	name string
}

func (s student) intro() {
	fmt.Println("I am a student")
}

func (e employee) intro() {
	fmt.Println("I am an employee")
}

func greet(i say) {
	fmt.Println(i.intro())
	fmt.Println(i.intro())
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var s1 = student{
		name: "golang",
	}

	s1.sayHello()
	s1.gantiNama("sir")
	s1.sayHello()

	s1.intro()

	var e1 = employee{name: "john"}

	e1.intro()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("a")
		wg.Done()
	}()

	go func() {
		fmt.Println("b")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(test.S)
	test.Greet()

	// greet(s1)

	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
