package main

import (
	"errors"
	"strconv"
	"strings"
)

func getFromNumericalSystem(str string, system int, prefix string) (uint32, error) {
	if str == "" {
		return 0x0, nil
	}
	str = strings.TrimPrefix(str, prefix)
	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0d") || strings.HasPrefix(str, "0b") {
		return 0, errors.New("Incorrect format")
	}
	conv, err := strconv.ParseUint(str, system, 32)
	if err != nil {
		return 0, err
	}
	return uint32(conv), nil
}

func getFromHex(hexStr string) (uint32, error) {
	return getFromNumericalSystem(hexStr, 16, "0x")
}

func getFromDec(decStr string) (uint32, error) {
	return getFromNumericalSystem(decStr, 10, "0d")
}

func getFromBin(binStr string) uint32 {
	return 0
}
