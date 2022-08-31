package main

import (
	"fmt"
	FUCKED "main/Modules"
	"os"
)

func main() {
	FUCKED.Ban()
	f := os.Args[1] // list of filepaths
	o := os.Args[2] // output PCAP file
	if f == "" {
		fmt.Println("[!] Error: please list a list of filepaths, for example Filepaths.txt")
		os.Exit(0)
	}
	if o == "" {
		fmt.Println("[!] Error: Please list a output pcap filename example Output.pcap")
		os.Exit(0)
	}
	if f != "" {
		if o != "" {
			FUCKED.FUCKED_LOADER(f, o)
		}
	}
}
