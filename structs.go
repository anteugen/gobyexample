package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ContactInfo struct {
	Email string
	PhoneNumber string
}

type Person struct {
	FirstName string
	LastName string
	Age int
	ContactInfo
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func generatePhoneNumber() string {
	rand.Seed(time.Now().UnixNano())

	areaCode := randInt(200, 1000)
	firstThree := randInt(200, 1000)
	lastFour := randInt(1000, 10000)

	return fmt.Sprintf("(%d) %d-%d", areaCode, firstThree, lastFour)
}

func NewPerson(firstName, lastName string, age int) Person {
	
	p := Person{FirstName: firstName, LastName: lastName, Age: age}
	email := fmt.Sprintf("%v%v@hotmail.com", firstName, lastName)
	number := generatePhoneNumber()
	c := ContactInfo{Email: email, PhoneNumber: number}
	p.ContactInfo = c
	return p
}

func FullName(p Person) string {
	return fmt.Sprintf("%v %v", p.FirstName, p.LastName)
}

func canVote(p Person) bool {
	if p.Age >= 18 {
		return true
	}
	return false
}

func incrementAge(p *Person) {
	p.Age++
}

func compareFirstName(p1, p2 Person) bool {
	if p1.FirstName == p2.FirstName {
		return true
	}
	return false
}

func main() {
	var persons []Person

	person1 := NewPerson("John", "Doe", 23)
	person2 := NewPerson("Dale", "Cooper", 14)
	person3 := NewPerson("John", "Williams", 46)

	persons = append(persons, person1, person2, person3)

	fmt.Println(person1)
	incrementAge(&person1)
	fmt.Println(person1, "\n")
	fmt.Println(persons, "\n")
	fmt.Println("Full name:", FullName(person1))
	fmt.Println("Can vote?:", canVote(person1))
	fmt.Println("Can vote?:", canVote(person2))
	fmt.Println("Same first name?:", compareFirstName(person1, person3))
}
