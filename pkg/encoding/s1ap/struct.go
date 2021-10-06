package s1ap

// type GNB struct {
// 	GlobalGNBID     GlobalGNBID
// 	SupportedTAList []SupportedTA
// 	PagingDRX       string
// 	RANUENGAPID     uint32
// 	ULInfoNR        UserLocationInformationNR
// 	NGAPPeerAddr    string
// 	GTPuLocalAddr   string
// 	GTPuIFname      string
// 	GTPuTEID        uint32
// 	UE              nas.UE // base parameter to be used for each UE

// 	Recv struct {
// 		GTPuPeerAddr net.IP
// 		GTPuPeerTEID uint32
// 	}

// 	Camper []*Camper

// 	DecodeError error
// 	dbgLevel    int
// 	indent      int // indent for debug print.
// }

// type Camper struct {
// 	GNB          *GNB // camped in this gNB
// 	UE           *nas.UE
// 	GTPu	     *gtp.GTP
// 	AmfId        uint32
// 	RanId        uint32
// 	RRCstate     int
// 	PDUSessionID uint8
// 	QosFlowID    uint8

// 	SendMsg *[]byte
// 	RecvMsg *[]byte

// 	camperType int
// }
