package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const UTMPFile = "/var/run/utmpx"

func main() {
	file, err := os.Open(UTMPFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line []byte
	reader := bufio.NewReader(file)
	for {
		line, err = reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			} else {
				break
			}
		}
		fmt.Println(string(line))
	}
}
