package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Decoder(input string) (string, error) {
	if input == "" {
		return "", errors.New("input is empty")
	}

	var output string //Return string for append

	//reg := regexp.MustCompile(`\[[0-9]+ [^\[\]]+\]|.|\n`)
	reg := regexp.MustCompile(`\[[^\]]*\]|.|\n`)
	matches := reg.FindAllStringSubmatch(input, -1)

	//Loop for input
	for _, match := range matches {

		if len(match[0]) == 1 {
			output += match[0]
			if match[0] == "[" || match[0] == "]" {
				return "", errors.New("square brackets are unbalanced")
			}
			continue
		} else if !strings.ContainsRune(match[0], ' ') {
			return "", errors.New("arguments are not separated by a space")
		}

		ind := strings.IndexRune(match[0], ' ')
		multiplierStr := match[0][1:ind]
		characters := match[0][ind+1 : len(match[0])-1]
		loopStr, err := strconv.Atoi(multiplierStr)

		if err != nil {
			return "", errors.New("the first argument is not a number")
		}
		if characters == "" {
			return "", errors.New("there is no second argument")
		}

		for i := 0; i < loopStr; i++ {
			output += characters
		}
	}

	return output, nil
}
