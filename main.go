package main

import "fmt"


// Fooer
type Fooer interface {
	Foo1()
	Foo2()
	Foo3()
}

type Foo struct {

}

func (f Foo) Foo1() {
	fmt.Println("Foo1()")
}

func (f Foo) Foo2() {
	fmt.Println("Foo2()")
}

func (f Foo) Foo3() {
	fmt.Println("Foo3()")
}

func NewFoo() Fooer {
	return &Foo{}
}

type SuperFooer struct {
	Fooer
}

func main()  {

	dragon := &FlyingCreature{
		Creature: Creature{
			Name: "Dragon",
			Real: false,
		},
		Wingspan: 15,
	}

	fmt.Println(dragon.Name)
	fmt.Println(dragon.Real)
	fmt.Println(dragon.Wingspan)

	f := NewFoo()
	f.Foo1()
	f.Foo2()
	f.Foo3()

	RunForrestRun();
}
