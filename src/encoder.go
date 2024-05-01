package main

import (
	"strconv"
	"strings"
)

func Encoder(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var encodedLines []string

	for _, line := range lines {
		encodedLine := ""
		for i := 0; i < len(line); {
			char := line[i]
			count := 1
			for i+1 < len(line) && line[i+1] == char {
				count++
				i++
			}
			twoCharEncoded, ind := EncoderTwoChar(line[i:])
			if count >= 2 {
				encodedLine += "[" + strconv.Itoa(count) + " " + string(char) + "]"
			} else if ind > 0 {
				i += ind
				encodedLine += twoCharEncoded
			}
			if i > 0 && char != line[i-1] {
				encodedLine += string(char)
			}
			i++
		}
		encodedLines = append(encodedLines, encodedLine)
	}

	return strings.Join(encodedLines, " "), nil
}

func EncoderTwoChar(input string) (string, int) {
	var encodedTwo = ""
	count := 1

	for i := 1; i < len(input); i += 2 {
		if i+3 < len(input) && input[i-1:i+1] == input[i+1:i+3] {
			count++
		} else {
			if count > 1 {
				encodedTwo += "[" + strconv.Itoa(count) + " " + string(input[i-1:i+1]) + "]"
				return encodedTwo, i
			}
			count = 1
			return encodedTwo, 0
		}
	}
	if len(input)%2 == 1 {
		encodedTwo += string(input[len(input)-1])
	}
	return encodedTwo, len(input)
}
