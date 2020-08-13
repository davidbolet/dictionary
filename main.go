package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var path = "./resources/texto.txt"

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
	if len(os.Args)>1 {
		file, err := os.Open("./resources/palabras_todas.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		line := 0
		fileOut, err := os.Create(path)
		defer fileOut.Close()
		if existeError(err) {
			return
		}
		for scanner.Scan() {
			//fmt.Println(line, scanner.Text())
			
			removed, has := removeCharacters(scanner.Text(), toRemove)
			if has {

				escribirArchivo(fileOut, scanner.Text(), removed)
				//fmt.Println(scanner.Text() + "->" + removed)
				line++
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		CorrigeTexto()
	}
	
}

func existeError(err error) bool {
	if err != nil {
	  fmt.Println(err.Error())
	}
	return (err != nil)
  }



func escribirArchivo(fileOut *os.File, sinAcento string, conAcento string) {
	
	_, err := fileOut.WriteString (conAcento +" "+ sinAcento + "\n")

	if existeError(err) {
		return
	  }

	  err = fileOut.Sync()
	  if existeError(err) {
		return
	  }

}


