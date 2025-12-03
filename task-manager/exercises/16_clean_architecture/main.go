package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Robot struct {
	Name string
}

type Human struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: Hav hav!"
}

func (c Cat) Speak() string {
	return c.Name + " says: Miyav Miyav!"
}

func (r Robot) Speak() string {
	return r.Name + " says: Beep Beep!"
}

func (h Human) Speak() string {
	return h.Name + " says: Hello!"
}


func MakeSpeak (s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "Karabaş"}
	cat := Cat{Name: "Tekir"}
	robot := Robot{Name: "R2D2"}
	human := Human{Name: "John"}

	MakeSpeak(dog)
	MakeSpeak(cat)
	MakeSpeak(robot)
	MakeSpeak(human)
}