package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	//scanner.Split(bufio.ScanWords)
	line, err := reader.ReadString('\n')
	for err == nil {
		//token := scanner.Text()
		// tokens := strings.FieldsFunc(line, func(r rune) bool {
		// 	switch r {
		// 	case ' ', '\t', '\n', '\r', '.', ',', '-', '"', '!', ';', '?', ':', '(', ')':
		// 		return true
		// 	}
		// 	return false
		// })

		fmt.Println(line)
		line, err = reader.ReadString(0xA)

	}
}
