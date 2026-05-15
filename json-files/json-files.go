package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Phone struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Person struct {
	Name         string  `json:"name"`
	Age          int     `json:"age"`
	Address      Address `json:"address"`
	PhoneNumbers []Phone `json:"phoneNumbers"`
}

func newPerson(name string, age int, address Address, phoneNumbers []Phone) Person {
	return Person{
		Name:         name,
		Age:          age,
		Address:      address,
		PhoneNumbers: phoneNumbers,
	}
}

func parsePerson(jsonData *os.File) (Person, error) {
	var person Person

	byteData, _ := io.ReadAll(jsonData)

	err := json.Unmarshal(byteData, &person)
	return person, err
}

func main() {
	// This is where you would typically read a JSON file and unmarshal it into the Person struct.
	jsonFile, err := os.Open("./json/person.json")
	if err != nil {
		panic(err)
	}

	person, err := parsePerson(jsonFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed Person: %+v\n", person)

	defer jsonFile.Close()
}
