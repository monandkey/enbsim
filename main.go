package main

import (
	"log"
	"net"
	"time"
	"github.com/ishidawataru/sctp"
)

func setupSCTP() (conn *sctp.SCTPConn, info *sctp.SndRcvInfo) {
	log.Print("==================== SCTP Connection open ==========")

	const mmePort = 36412
	mmeAddr, _ := net.ResolveIPAddr("ip", "192.168.200.10")
	mmeips := []net.IPAddr{*mmeAddr}
	rAddr := &sctp.SCTPAddr{
		IPAddrs: mmeips,
		Port:   mmePort,
	}

	const enbPort = 36412
	enbAddr, _ := net.ResolveIPAddr("ip", "192.168.200.20")
	enbips := []net.IPAddr{*enbAddr}
	lAddr := &sctp.SCTPAddr{
		IPAddrs: enbips,
		Port:   enbPort,
	}

	conn, err := sctp.DialSCTP("sctp", lAddr, rAddr)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	log.Printf("Dail Local Address: %s; Remote Address: %s", conn.LocalAddr(), conn.RemoteAddr())

	ppid := 0
	info = &sctp.SndRcvInfo{
		Stream: uint16(ppid),
		PPID:   0x3c000000,
	}
	conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)
	return
}

func initRAN() {
	setupSCTP()
}

func init() {
	time.Sleep(time.Second * 1)
}

func main() {
	log.SetPrefix("[ENB] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Print("eNB Initialized")
	initRAN()
	for {
	}
}
