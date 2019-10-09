package main

import (
	"fmt"
	"os"
	"strconv"
)

const invUsage = "invalid usage. go run main.go [feets]"
const errMsg = "Error:"
const magic = 0.3048
const succmsg = " meters"

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(errMsg, invUsage)
		return
	}

	feets, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println(errMsg, err)
		return
	}

	meters := feets * magic

	fmt.Println(meters, succmsg)
}
