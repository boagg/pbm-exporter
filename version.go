package main

import "fmt"

var version = "2.0.0"

func init() {
	if version == "" {
		version = "dev"
	}
}

func printVersion() {
	fmt.Printf("pbm-exporter version %s\n", version)
}
