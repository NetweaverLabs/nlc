/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/gob"

	"github.com/NetweaverLabs/nlc/cmd"
	"github.com/NetweaverLabs/types"
)

func init() {
	gob.Register(types.User{})
}

func main() {
	cmd.Execute()
}
