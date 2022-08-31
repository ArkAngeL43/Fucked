package FUCKED

import "github.com/google/gopacket"

var (
	Filepaths  []string // All filepaths loaded from a file
	Good_Files []string // Files that were found and validated
	Bad_Files  []string // Files that were not found or located
	Packets    []gopacket.Packet
)
