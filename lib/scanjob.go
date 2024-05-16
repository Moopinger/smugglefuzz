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
	Keyword  string
}

func NewScanJob(target *Target, conn *tls.Conn, payloads []Payload, keyword string) *ScanJob {
	return &ScanJob{
		Target:   target,
		Conn:     conn,
		Payloads: payloads,
		StreamId: 1,
		Keyword:  keyword,
	}
}

func (s *ScanJob) SetConn(conn *tls.Conn) {
	s.Conn = conn
}
