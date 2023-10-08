package main

import "log"

// A Simple for loop
func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
	// Create a slice and iterate through the array
	animals := []string{"dog", "Fish", "Rat", "Cat", "Bird"}
	// The underscore allows me to declare variables that may not be used later.
	for _, animal := range animals {
		log.Println(animal)
	}
	// Range a for lop over a map of elements
	pets := make(map[string]string)

	pets["dog"] = "Fido"
	pets["cat"] = "Frisky"
	pets["Fish"] = "Snuggles"
	pets["Snake"] = "Slinky"
	pets["Dolphin"] = "Flipper"

	// The underscore allows me to declare variables that may not be used later.
	for petsType, pet := range pets {
		log.Println(petsType, pet)
	}
	// Range a for loop over a string
	var firstLine = "I gave fifty racks in the club to a pimp (ju-juice)"
	for z, lines := range firstLine {
		log.Println(z, lines)
	}

	// Range a for loop over a customer object structcture
	// Create a struct
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	// Create a click using the customer struct
	var users []User
	users = append(users, User{"John", "Smith", "john@gmail.com", 31})
	users = append(users, User{"Jill", "Smith", "jill@gmail.com", 28})
	users = append(users, User{"Jack", "Smith", "jack@gmail.com", 19})
	users = append(users, User{"JoeAnn", "Smith", "joe@gmail.com", 16})
	users = append(users, User{"Jenifer", "Smith", "jen@gmail.com", 11})

	for y, l := range users {
		log.Println(y, l.FirstName, l.LastName, l.Email, l.Age)
	}
}
