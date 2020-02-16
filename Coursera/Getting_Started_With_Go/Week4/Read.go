//package main
//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"os"
//	"strings"
//)
//
//type Name struct {
//	firstname string
//	lastname  string
//}
//
//func content(s string) string {
//	r := []rune(s)
//	return string(r[0:20])
//
//}
//
//func main() {
//	var fileName string
//	fmt.Print("Enter the file name: ")
//	fmt.Scan(&fileName)
//	var person Name
//	personsli := make([]Name, 0)
//	file, err := os.Open(fileName)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		s := strings.Split(scanner.Text(), " ")
//		if len(s[0]) > 20 {
//			s[0] = content(s[0])
//		}
//		if len(s[1]) > 20 {
//			s[1] = content(s[1])
//		}
//		person.firstname, person.lastname = s[0], s[1]
//		personsli = append(personsli, person)
//	}
//	for _, r := range personsli {
//		fmt.Println(r.firstname, r.lastname)
//
//	}
//	if err := scanner.Err(); err != nil {
//
//		log.Fatal(err)
//	}
//
//}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxLength = 20

type name struct {
	fname string
	lname string
}

func (n *name) setName(firstName string, lastName string) {
	processedFirstName := firstName
	if len(firstName) > maxLength {
		runeFirstName := []rune(processedFirstName)
		processedFirstName = string(runeFirstName[:maxLength])
	}

	processedLastName := lastName
	if len(lastName) > maxLength {
		runeLastName := []rune(processedLastName)
		processedLastName = string(runeLastName[:maxLength])
	}

	n.fname = processedFirstName
	n.lname = processedLastName
}

func main() {
	// get file path from user
	fmt.Printf("Enter path to a file with names:\n")
	pathScanner := bufio.NewScanner(os.Stdin)
	_ = pathScanner.Scan()
	filePath := pathScanner.Text()

	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file
	scanner := bufio.NewScanner(file)
	var nameSlice []name
	for scanner.Scan() {
		var currentName name
		scannedNameArray := strings.Fields(scanner.Text())
		currentName.setName(scannedNameArray[0], scannedNameArray[1])
		nameSlice = append(nameSlice, currentName)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// print names
	for _, value := range nameSlice {
		fmt.Println(value.fname, value.lname)
	}
}
