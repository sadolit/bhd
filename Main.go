package main

import (
	"fmt"

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
	var number int
	if arguments["x"] == true {
		number = getFromHex(strNumber)
	} else if arguments["d"] == true {
		number = getFromDec(strNumber)
	} else if arguments["b"] == true {
		number = getFromBin(strNumber)
	} else {
		panic("Choose x, b, or d")
	}
	fmt.Println(number)
}
