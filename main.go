package main

import (
	"fmt"
	"github.com/gabrielpsilva/glue/v1/cypher"
	"io/ioutil"
	"os"
)

func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}


func isFile(file string) bool {
	if s, err := os.Stat(file); s.IsDir() || err != nil {
		return false
	}

	return true
}

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage:")
		fmt.Println("-------")
		fmt.Println("")
		fmt.Println("dropencrypt [encrypt/decrypt] [key] [file]")
		os.Exit(0)
	}


	action := os.Args[1]
	key := os.Args[2]
	file := os.Args[3]

	if !isFile(file) {
		fmt.Println("File does not exist: ", file)
		os.Exit(1)
	}

	content, err := readFromFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if action == "encrypt"{
		encrypted := cypher.Encrypt(string(content), key)
		writeToFile(string(encrypted), file+".enc")
	}
	if action == "decrypt"{
		decrypted := cypher.Decrypt(content, key)
		writeToFile(decrypted, file[:len(file)-5])
	}
}

