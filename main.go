package main

import (
	"bufio"
	"functional/functions"
	"log"
	"os"
)

func main() {
	file, err := os.Open("c:/in/sources.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines functions.StringSlice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	println(lines.MkStringWithEnds("[", ", ", "]"))
}
