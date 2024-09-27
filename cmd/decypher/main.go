package main

import (
	"os"

	"github.com/lambda-mena/cybersecurity/internal/cesar"
)

func main() {
	var cryptoSystem string = os.Args[1]

	if cryptoSystem == "cesar" {
		cesar.RequestInputs()
	}
}
