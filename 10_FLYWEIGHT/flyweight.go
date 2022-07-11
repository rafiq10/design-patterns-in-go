package main

import (
	"fmt"

	ff "design-patterns/flyweight/textformatting"
	un "design-patterns/flyweight/usernames"
)

func main() {
	text := "This is a brave new world"
	ft := ff.NewFormattedText(text)
	ft.Capitalize(10, 15)
	fmt.Println(ft.String())

	bft := ff.NewBetterFormattedText(text)
	bft.Range(16, 19).Capitalize = true
	fmt.Println(bft.String())

	// user names
	john := un.NewUser("John Doe")
	jane1 := un.NewUser("Jane Doe")
	jane2 := un.NewUser("Jane Smith")

	fmt.Println("Memory taken by users: ",
		len([]byte(john.FullName))+
			len([]byte(jane1.FullName))+
			len([]byte(jane2.FullName)),
	)

	john2 := un.NewUser2("John Doe")
	jane3 := un.NewUser2("Jane Doe")
	jane4 := un.NewUser2("Jane Smith")
	totalMem := 0
	for _, a := range un.AllNames {
		totalMem += len([]byte(a))
	}
	totalMem += len(john2.Names)
	totalMem += len(jane3.Names)
	totalMem += len(jane4.Names)
	fmt.Println("Memory taken by users: ", totalMem)

}
