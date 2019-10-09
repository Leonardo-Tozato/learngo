package main

import (
	"fmt"
	"os"
)

func main() {
	vUser, vPass := "leonardo", "xxt123"

	if len(os.Args) != 3 {
		fmt.Println("invalid usage, pass args username and password")
		return
	}

	user, pass := os.Args[1], os.Args[2]

	if user != vUser || pass != vPass {
		fmt.Println("invalid username or password for", user, pass)
		return
	}

	fmt.Println("user granted for", user)

}
