/*
2024 Moopinger
*/

package lib

import (
	"net"
)

type ScanJob struct {
	Target   *Target
	Conn     net.Conn
	Payloads []Payload
	StreamId int
}

func NewScanJob(target *Target, conn net.Conn, payloads []Payload) *ScanJob {
	return &ScanJob{
		Target:   target,
		Conn:     conn,
		Payloads: payloads,
		StreamId: 1,
	}
}

func (s *ScanJob) SetConn(conn net.Conn) {
	s.Conn = conn
}
