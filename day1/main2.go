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
	a_read := -1
	a_sum := 0
	b_read := -2
	b_sum := 0
	c_read := -3
	c_sum := 0
	larger := 0
	prev := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		v, _ := strconv.Atoi(line)

		a_read += 1
		b_read += 1
		c_read += 1
		if a_read >= 0 {
			a_sum += v
		}
		if b_read >= 0 {
			b_sum += v
		}
		if c_read >= 0 {
			c_sum += v
		}

		if a_read == 2 {
			if prev != 0 && a_sum > prev {
				larger += 1
			}
			prev = a_sum
			a_sum = 0
			a_read = -1
			continue
		}
		if b_read == 2 {
			if b_sum > prev {
				larger += 1
			}
			prev = b_sum
			b_sum = 0
			b_read = -1
			continue
		}
		if c_read == 2 {
			if c_sum > prev {
				larger += 1
			}
			prev = c_sum
			c_sum = 0
			c_read = -1
			continue
		}
	}

	fmt.Printf("%v measurements larger", larger)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
