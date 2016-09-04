package main

import (
	"fmt"
	"os"
)

func main() {
	alex := contactItem{"Alex", "Moscow"}
	vlad := contactItem{"Vlad", "Berlin"}
	alice := contactItem{"Alice", "Ivanovo"}

	concactsList := []string{alex.name, vlad.name, alice.name}

	printContacts(concactsList...)

	contactsDump := saveContacts("contacts.txt")
	defer closeFile(contactsDump)
	writeToFile(contactsDump, concactsList...)
}

type contactItem struct {
	name string
	city string
}

func getInfo(name string, city string) string {
	return name + " from " + city
}

func printContacts(contacts ...string) {
	for _, concactName := range contacts {
		fmt.Println(concactName)
	}
}

func saveContacts(path string) *os.File {
	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	return file
}

func writeToFile(file *os.File, data ...string) {
	fmt.Fprintln(file, data)
}

func closeFile(file *os.File) {
	file.Close()
}
