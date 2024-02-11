package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func init() {
	if !checkTerraformBinary() {
		log.Fatal("Terraform binary or tf command does not exist.")
	}
}

func main() {
	(&cli.App{}).Run(os.Args)
}
