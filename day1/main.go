package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	prev := 0
	larger := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		v, _ := strconv.Atoi(line)
		if prev == 0 {
			fmt.Printf("%v (N/A) - no previous measurement\n", v)
			prev = v
			continue
		}

		if v < prev {
			fmt.Printf("%v < %v (decreased)\n", prev, v)
			prev = v
			continue
		}
		if v > prev {
			fmt.Printf("%v (increased)\n", v)
			prev = v
			larger += 1
			continue
		}

		if prev == v {
			fmt.Printf("%v (no change)\n", v)
			continue
		}
	}

	fmt.Printf("%v measurements larger", larger)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
