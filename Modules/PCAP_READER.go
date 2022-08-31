package FUCKED

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Load(Good_File string, OUTPUT string) {
	handle, x := pcap.OpenOffline(Good_File)
	if x != nil {
		log.Fatal(x)
	}
	defer handle.Close()
	pktsrc := gopacket.NewPacketSource(handle, handle.LinkType())
	for pkt := range pktsrc.Packets() {
		Packets = append(Packets, pkt)
	}
	FUCKED_Saver(OUTPUT)
}
