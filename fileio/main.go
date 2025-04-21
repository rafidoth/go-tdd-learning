package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(filePath string) ([]string, error) {
	var err error
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rbuf := bufio.NewReader(f)
	// fmt.Println(rbuf)

	var lines []string
	for {
		line, err := rbuf.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func charByChar(filePath string) error {
	lines, err := lineByLine(filePath)
	if err != nil {
		return err
	}
	for _, line := range lines {
		for _, ch := range line {
			fmt.Printf("%v ", string(ch))
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: linebyline <file1> [<file2> ...]")
		return
	}

	for _, file := range flag.Args() {
		fmt.Println(file)
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}

	}
}
