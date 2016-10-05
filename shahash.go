package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("shahash - a simple sha1 file checker\n\t usage: shahash <filepath> \n\treturns: sha1 hash of the file specified.")
		return
	}

	fileBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	sha := sha1.New()
	sha.Write(fileBytes)
	fmt.Printf("sha1 %s: %x\n", os.Args[1], sha.Sum(nil))
}
