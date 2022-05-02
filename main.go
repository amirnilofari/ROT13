package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	ROT13(text)
}

func ROT13(msg string) string {
	str := []byte(msg)
	for i := 0; i < len(str); i++ {
		switch {
		case str[i] >= 65 && str[i] <= 77:
			str[i] += 13
		case str[i] >= 78 && str[i] <= 90:
			str[i] -= 13
		case str[i] >= 97 && str[i] <= 109:
			str[i] += 13
		case str[i] >= 110 && str[i] < 123:
			str[i] -= 13
		default:
			str[i] = str[i]
		}
	}
	fmt.Println(string(str))
	return string(str)
}
