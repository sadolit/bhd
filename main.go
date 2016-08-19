package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

var usage = `bhd. Converts number between numerical systems. Input number can be hexadecimal (x), decimal (d) or binary (b).

Usage:
  bhd x <number>
  bhd d <number>
  bhd b <number>

  bhd -h | --help
  bhd --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

func main() {
	var strNumber string
	arguments, _ := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)
	strNumber = arguments["<number>"].(string)
	var number uint32
	var err error
	if arguments["x"] == true {
		number, err = getFromHex(strNumber)
	} else if arguments["d"] == true {
		number, err = getFromDec(strNumber)
	} else if arguments["b"] == true {
		number, err = getFromBin(strNumber)
	} else {
		panic("Choose x, b, or d")
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("\nDecimal: %d\nHexadecimal: 0x%X\nBinary: %b\n\n", number, number, number)
}
