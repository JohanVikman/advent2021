package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)
	horiz := 0
	depth := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Printf("%v\n", line)
		parts := strings.Split(line, " ")
		direction := parts[0]
		moveAmount, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "forward":
			{
				horiz += moveAmount
			}
		case "up":
			{
				depth -= moveAmount
			}
		case "down":
			{
				depth += moveAmount
			}
		}
	}
	fmt.Printf("%v, %v = %v", horiz, depth, horiz*depth)
}
