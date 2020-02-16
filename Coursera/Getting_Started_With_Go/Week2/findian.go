package Week2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//(1) read input
	fmt.Print("Input: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.ToLower(scanner.Text())

	//(2) search and print message
	if strings.HasPrefix(text, "i") && strings.HasSuffix(text, "n") && strings.Contains(text, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
