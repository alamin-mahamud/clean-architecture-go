package main

import "fmt"

type Color string

const (
	RED   Color = "red"
	GREEN       = "green"
)

type Creature struct {
	Name string
	Real bool
}

type FlyingCreature struct {
	Creature
	Wingspan int
}

type Unicorn struct {
	Creature
}

type Pterodactyl struct {
	FlyingCreature
}

func NewPetrodactyle(wingspan int) *Pterodactyl {
	return &Pterodactyl{
		FlyingCreature{
			Creature{
				"Pterodactyl",
				true,
			},
			wingspan,
		},
	}
}

type Door struct {
	Thickness int
	Color     string
}

func Dump(c *Creature) {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

func (c Creature) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

func RunForrestRun() {
	creature := &Creature{
		"Some Creature",
		false,
	}

	uni := Unicorn{
		Creature{
			"Unicorn",
			false,
		},
	}

	pet1 := &Pterodactyl{
		FlyingCreature{
			Creature{
				"Pterodactyl",
				true,
			},
			20,
		},
	}

	pet2 := NewPetrodactyle(8)

	door := &Door{
		10,
		RED,
	}

	Dump(creature)
	creature.Dump()

	uni.Dump()
	pet1.Dump()
	pet2.Dump()

	creatures := []Creature{
		*creature,
		uni.Creature,
		pet1.Creature,
		pet2.Creature,
	}

	fmt.Println("Dump() through Creature embedded type")

	for _, creature := range creatures {
		creature.Dump()
	}

	dumpers := []Dumper{creature, uni, pet1, pet2, door}
	fmt.Println("Dump() through Dumper interface")

	for _, dumper := range dumpers {
		dumper.Dump()
	}
}
