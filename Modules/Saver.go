package FUCKED

import (
	"os"

	"github.com/google/gopacket/layers"

	"github.com/google/gopacket/pcapgo"
)

func FUCKED_Saver(OUTF string) {
	f, _ := os.Create(OUTF)
	WRITER := pcapgo.NewWriter(f)
	WRITER.WriteFileHeader(1024, layers.LinkTypeEthernet)
	defer f.Close()
	for _, pkt := range Packets {
		WRITER.WritePacket(pkt.Metadata().CaptureInfo, pkt.Data())
	}

}
