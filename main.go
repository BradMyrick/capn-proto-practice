package main

import (
	"fmt"
	"log"
	"os"

	person "capn-proto-practice/person"

	capnp "zombiezen.com/go/capnproto2"
)

func main() {
	// Create a new AddressBook message
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		log.Fatal(err)
	}
	addressBook, err := person.NewRootAddressBook(seg)
	if err != nil {
		log.Fatal(err)
	}

	// Create and add a new Person
	newPerson, err := person.NewPerson(seg)
	if err != nil {
		log.Fatal(err)
	}
	newPerson.SetId(1)
	newPerson.SetName("Brad")
	newPerson.SetAge(36)
	newPerson.SetEmail("kodr@codemucho.com")

	// Create a new list of Person structs and set the first one
	people, err := addressBook.NewPeople(1)
	if err != nil {
		log.Fatal(err)
	}
	people.Set(0, newPerson)
	addressBook.SetPeople(people)

	// Serialize the message to a file
	f, err := os.Create("addressbook.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = capnp.NewEncoder(f).Encode(msg)
	if err != nil {
		log.Fatal(err)
	}

	// Read the message back from the file
	f2, err := os.Open("addressbook.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	msg2, err := capnp.NewDecoder(f2).Decode()
	if err != nil {
		log.Fatal(err)
	}
	addressBook2, err := person.ReadRootAddressBook(msg2)
	if err != nil {
		log.Fatal(err)
	}

	// Print the contents of the AddressBook
	people2, err := addressBook2.People()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < people2.Len(); i++ {
		person2 := people2.At(i)
		name, err := person2.Name()
		if err != nil {
			log.Fatal(err)
		}
		email, err := person2.Email()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Person %d: %s (%d years old, email: %s)\n", person2.Id(), name, person2.Age(), email)
	}
}
