package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"os"
	"time"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: otp <secret key>")
		os.Exit(1)
	}

	secret := args[1]

	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		fmt.Println("Invalid secret key:", err)
		os.Exit(1)
	}

	//for len(code) < 6 {
	//	code = "0" + code
	//}

	fmt.Println(code)
}
