/*
2024 Moopinger
*/

package lib

import (
	"crypto/tls"
)

type ScanJob struct {
	Target   *Target
	Conn     *tls.Conn
	Payloads []Payload
	StreamId uint32
}

func NewScanJob(target *Target, conn *tls.Conn, payloads []Payload) *ScanJob {
	return &ScanJob{
		Target:   target,
		Conn:     conn,
		Payloads: payloads,
		StreamId: 1,
	}
}

func (s *ScanJob) SetConn(conn *tls.Conn) {
	s.Conn = conn
}
