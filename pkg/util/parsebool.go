package util

import (
	"log"
	"os"
	"strconv"
)

func ParseBool(str string) bool {
	boolStr, err := strconv.ParseBool(str)
	if err != nil {
		log.Printf("An error ocurred while convert debug var: %v", err)
		os.Exit(1)
	}

	return boolStr
}
