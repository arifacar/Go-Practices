package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {

	file, err := os.Open("/Users/arifacar/Desktop/userinfoBps5.md")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

