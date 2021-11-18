package packetdump

import (
	"fmt"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

// RTPFormatCallback can be used to apply custom formatting to each dumped RTP
// packet. If new lines should be added after each packet, they must be included
// in the returned format.
type RTPFormatCallback func(*rtp.Packet) string

// RTCPFormatCallback can be used to apply custom formatting to each dumped RTCP
// packet. If new lines should be added after each packet, they must be included
// in the returned format.
type RTCPFormatCallback func([]rtcp.Packet) string

// DefaultRTPFormatter returns the default log format for RTP packets
func DefaultRTPFormatter(pkt *rtp.Packet) string {
	return fmt.Sprintf("%s\n", pkt)
}

// DefaultRTCPFormatter returns the default log format for RTCP packets
func DefaultRTCPFormatter(pkts []rtcp.Packet) string {
	return fmt.Sprintf("%s\n", pkts)
}