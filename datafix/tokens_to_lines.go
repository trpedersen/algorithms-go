package main

import (
	"bitbucket.org/trpedersen/algorithms/io"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdIn := io.NewScanner(os.Stdin)

	line, err := stdIn.ReadLine()
	//fmt.Println(token, err)
	for err == nil {
		//fmt.Println(token)
		tokens := strings.FieldsFunc(line, func(r rune) bool {
			switch r {
			case ' ', '\t', '\n', '\r', '.', ',', '-', '"', '!', ';', '?', ':', '(', ')':
				return true
			}
			return false
		})
		for _, token := range tokens {
			fmt.Println(token)
		}
		line, err = stdIn.ReadLine()
	}
}
