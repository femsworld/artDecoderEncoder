package main

import (
	"flag"
	"fmt"
)

func main() {
	output := ""
	var err error
	prePrint := ""

	// Arguments: usage/flag.
	flagHelp := flag.Bool("h", false, "Print Usage")
	flagMulti := flag.Bool("m", false, "Multi-line")
	flagEncode := flag.Bool("e", false, "Encoder")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 || *flagHelp {
		fmt.Println("Decoder Usage:\ngo run . \"[5 #][5 -_]-[5 #]\"\n#####-_-_-_-_-_-#####")
		fmt.Println("Multi-line Usage:\ngo run . -m <inputfilelocation>")
		fmt.Println("Encoder Usage:\ngo run . -e <enter text>")
		return
	}

	if *flagEncode && *flagMulti {
		output, err = Multiline(args[0], true)
		prePrint = "Multiline encoded output:"
	} else if *flagEncode {
		output, err = Encoder(args[0])
		prePrint = "Encoded output:"
	} else if *flagMulti {
		output, err = Multiline(args[0], false)
		prePrint = "Multiline decoded output:"
	} else {
		output, err = Decoder(args[0])
		prePrint = "Decoded output:"
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(prePrint, "\n"+output)
}
