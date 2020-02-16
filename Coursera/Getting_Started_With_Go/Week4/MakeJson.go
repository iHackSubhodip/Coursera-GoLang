package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var mapHash map[string]string
	mapHash = make(map[string]string)

	fmt.Println("Please input your name:")
	var br = bufio.NewReader(os.Stdin)
	nameValue, _, _ := br.ReadLine()
	name := string(nameValue)
	mapHash["name"]=name

	fmt.Println("Then input your adress:")
	addressValue := bufio.NewReader(os.Stdin)
	address, _, _ := addressValue.ReadLine()
	addressString := string(address)
	mapHash["address"] = addressString

	var data, err = json.Marshal(mapHash)

	if err != nil {
		fmt.Println("Encoding failed")
	} else {
		fmt.Println("Encoded data : ")
		fmt.Println(data)
		fmt.Println("Decoded data : ")
		fmt.Println(string(data))
	}
}