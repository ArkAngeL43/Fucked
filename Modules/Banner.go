package FUCKED

import "fmt"

var banner = `
____ _  _ ____ _  _ ____ ___ 
|--- |__| |___ |-:_ |=== |__>
	PCAP file mashing  {1.0}
-----------------------------
`

func Ban() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	fmt.Println(banner)
}
