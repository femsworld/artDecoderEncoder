package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Multiline(input string, encode bool) (string, error) {
	inputData, err := os.Open(input)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return "", err
	}
	defer inputData.Close()

	scanner := bufio.NewScanner(inputData)
	returnStr := ""
	var processedLine string
	var processErr error

	for scanner.Scan() {
		line := scanner.Text()
		if encode {
			processedLine, processErr = Encoder(line)
		} else {
			processedLine, processErr = Decoder(line)
		}

		if processErr != nil {
			return "", processErr
		}

		returnStr += processedLine + "\n"
	}

	return strings.TrimRight(returnStr, "\n"), nil
}
