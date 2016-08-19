package main

import (
	"strconv"
	"strings"
)

func getFromHex(hexStr string) (uint32, error) {
	if hexStr == "" {
		return 0x0, nil
	}
	hexStr = strings.TrimPrefix(hexStr, "0x")
	conv, err := strconv.ParseUint(hexStr, 16, 32)
	if err != nil {
		return 0, err
	}
	return uint32(conv), nil
}

func getFromDec(decStr string) uint32 {
	return 0
}

func getFromBin(binStr string) uint32 {
	return 0
}
