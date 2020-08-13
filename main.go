package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func removeCharacters(input string, characters string) (string, bool) {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	res := strings.Map(filter, input)
	return res, res != input

}

const toRemove = "áéíóúñçÁÉÍÓÚüÜ"

func main() {

	file, err := os.Open("./resources/palabras_todas.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		//fmt.Println(line, scanner.Text())
		line++
		removed, has := removeCharacters(scanner.Text(), toRemove)
		if has {
			fmt.Println(scanner.Text() + "->" + removed)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
