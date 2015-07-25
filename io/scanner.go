package io

import (
	//"text/scanner"
	//"github.com/trpedersen/scanner"
	"bufio"
	"io"
	//"os"
	"strconv"
	//"fmt"
)

type IScanner interface {
	IsEmpty() (bool, error)
	ReadInt() (int, error)
	ReadDouble() (float64, error)
	ReadFloat() (float32, error)
	ReadLong() (int64, error)
	ReadBoolean() (bool, error)
	//ReadChar() rune
	ReadByte() (byte, error)
	ReadString() (string, error)
	HasNextLine() (bool, error)
	ReadLine() (string, error)
	readAll() (string, error)
}

type Scanner struct {
	r       *bufio.Reader
	scanner *bufio.Scanner
}

/*
func NewScanner() *Scanner {
	s := &Scanner{}
	s.r = bufio.NewReader(os.Stdin)
	s.scanner = bufio.NewScanner(s.r)
	return s
}
*/
func NewScanner(reader io.Reader) *Scanner {
	s := &Scanner{}
	s.r = bufio.NewReader(reader)
	s.scanner = bufio.NewScanner(s.r)
	return s
}

func (this *Scanner) IsEmpty() bool {
	_, err := this.r.Peek(1)
	return err != nil
}

func (this *Scanner) ReadLine() (string, error) {

	var (
		result string
		err    error = nil
	)
	this.scanner.Split(bufio.ScanLines)
	if this.scanner.Scan() {
		result = this.scanner.Text()
	} else {
		result = ""
		err = this.scanner.Err()
		if err == nil {
			err = io.EOF
		}
	}
	return result, err
}

func (this *Scanner) ReadAllLines() (result []string, err error) {
	var line string
	line, err = this.ReadLine()
	for err == nil {
		result = append(result, line)
		line, err = this.ReadToken()
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func (this *Scanner) ReadToken() (token string, err error) {
	this.scanner.Split(bufio.ScanWords)
	if this.scanner.Scan() {
		token = this.scanner.Text()
	} else {
		token = ""
		err = this.scanner.Err()
		if err == nil {
			err = io.EOF
		}
	}
	return
}

func (this *Scanner) ReadAllTokens() (result []string, err error) {
	var token string
	token, err = this.ReadToken()
	for err == nil {
		result = append(result, token)
		token, err = this.ReadToken()
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func (this *Scanner) ReadInt() (result int, err error) {
	token, err := this.ReadToken()
	if err == nil {
		result, err = strconv.Atoi(token)
	}
	return
}

func (this *Scanner) ReadAllInts() (result []int, err error) {
	var i int
	i, err = this.ReadInt()
	for err == nil {
		result = append(result, i)
		i, err = this.ReadInt()
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func (this *Scanner) ReadDouble() (result float64, err error) {
	token, err := this.ReadToken()
	if err == nil {
		result, err = strconv.ParseFloat(token, 64)
	}
	return
}

func (this *Scanner) ReadFloat() (result float32, err error) {
	token, err := this.ReadToken()
	if err == nil {
		var f float64
		f, err = strconv.ParseFloat(token, 64)
		result = float32(f)
	}
	return
}

func (this *Scanner) ReadLong() (result int64, err error) {
	token, err := this.ReadToken()
	if err == nil {
		result, err = strconv.ParseInt(token, 10, 64)
	}
	return
}

func (this *Scanner) ReadBoolean() (result bool, err error) {
	token, err := this.ReadToken()
	if err == nil {
		result, err = strconv.ParseBool(token)
	}
	return
}

func (this *Scanner) ReadByte() (result byte, err error) {
	//var size int
	result, err = this.r.ReadByte()
	return
}

// func (this *Scanner) ReadString() (string, error) {
// 	return this.ReadToken()
// }

/*
func (this *Scanner) HasNextLine() bool {
	return this.scanner.HasNextLine()
}
*/
/*
func (this *Scanner) HasNextChar() bool {
	return this.scanner.HasNext() // todo: HasNextChar
}

func (this *Scanner) ReadLine() string{
	return this.scanner.NextLine()
}

func (this *Scanner) ReadChar() Rune

*/
