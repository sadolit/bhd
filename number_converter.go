package main

import (
	"errors"
	"strconv"
	"strings"
)

func getFromHex(hexStr string) (uint32, error) {
	if hexStr == "" {
		return 0x0, nil
	}
	hexStr = strings.TrimPrefix(hexStr, "0x")
	if strings.HasPrefix(hexStr, "0x") || strings.HasPrefix(hexStr, "0d") || strings.HasPrefix(hexStr, "0b") {
		return 0, errors.New("Incorrect format")
	}
	conv, err := strconv.ParseUint(hexStr, 16, 32)
	if err != nil {
		return 0, err
	}
	return uint32(conv), nil
}

func getFromDec(decStr string) (uint32, error) {
	if decStr == "" {
		return 0x0, nil
	}
	decStr = strings.TrimPrefix(decStr, "0x")
	if strings.HasPrefix(decStr, "0x") || strings.HasPrefix(decStr, "0d") || strings.HasPrefix(decStr, "0b") {
		return 0, errors.New("Incorrect format")
	}
	conv, err := strconv.ParseUint(decStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(conv), nil
}

func getFromBin(binStr string) uint32 {
	return 0
}
