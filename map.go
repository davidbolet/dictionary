package main

import (
	"bufio"
	"io/ioutil"
	"fmt"
	"log"
	"os"
	"strings"
)

func CorrigeTexto() {
	file, err := os.Open("./resources/texto.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var diccionario = make(map[string]string)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		diccionario[arr[0]] = arr[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(diccionario))

	textoCorregido, err := ioutil.ReadFile("./resources/ejemplo.txt")
	result := ""
	for _,word := range strings.Split(string(textoCorregido), " ") {
		val, ok := diccionario[word]
		if ok && len(word)>3 {
			result = result + " " + val
		} else {
			result = result + " " + word
		}
	}

	fmt.Println(result)
}
