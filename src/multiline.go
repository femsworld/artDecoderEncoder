package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Multiline(input string, encode bool) (string, error) {
	file, err := os.Open(input)
	if err != nil {
		return "", fmt.Errorf("Error reading input file: %w", err)
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		var processedLine string
		if encode {
			processedLine, err = Encoder(line)
		} else {
			processedLine, err = Decoder(line)
		}

		if err != nil {
			return "", err
		}

		builder.WriteString(processedLine)
		builder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}
