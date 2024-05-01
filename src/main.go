package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	flagHelp := flag.Bool("h", false, "Print Usage")
	flagMulti := flag.Bool("m", false, "Multi-line")
	flagEncode := flag.Bool("e", false, "Encoder")
	flag.Parse()

	args := flag.Args()

	if *flagHelp || len(args) != 1 {
		fmt.Println("Decoder Usage: go run . \"[5 #][5 -_]-[5 #]\"")
		fmt.Println("Multi-line Usage: go run . -m <inputfilelocation>")
		fmt.Println("Encoder Usage: go run . -e <enter text>")
		return
	}

	var (
		output   string
		err      error
		prePrint string
	)

	switch {
	case *flagEncode && *flagMulti:
		output, err = Multiline(args[0], true)
		prePrint = "Multiline encoded output:"
	case *flagEncode:
		output, err = Encoder(args[0])
		prePrint = "Encoded output:"
	case *flagMulti:
		output, err = Multiline(args[0], false)
		prePrint = "Multiline decoded output:"
	default:
		output, err = Decoder(args[0])
		prePrint = "Decoded output:"
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(prePrint, "\n"+output)
}
