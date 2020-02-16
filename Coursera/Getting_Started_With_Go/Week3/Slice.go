package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	elements := make([]int,3)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ToLower(line) == "x" {
			break
		}else{
			if value, err := strconv.Atoi(line); err == nil {
				if index < len(elements) {
					elements[index] = value
				}else{
					elements= append(elements, value)
				}
				index++
			}
			swapElement := 0;

			for i := index -1; i >=0; i-- {
				if i == 0 {
					continue
				}
				if elements[i] < elements[i-1] {
					swapElement = elements[i]
					elements[i] = elements[i-1]
					elements[i-1] = swapElement
				}
			}
		}
		fmt.Println(elements)
	}
}