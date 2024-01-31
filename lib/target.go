/*
2024 Moopinger
*/

package lib

import (
	"crypto/tls"
	"fmt"
	"net"
	neturl "net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Target struct {
	URL    neturl.URL
	Status string
}

func ReadTargetFile(fileName string) ([]string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	unifiedContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(unifiedContent, "\n")

	return lines, nil
}

func NewTarget(url string) (*Target, error) {

	if url == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	if strings.ToLower(url[:4]) != "http" {
		url = "https://" + url
	}
	u, err := neturl.Parse(url)

	if u.Scheme != "https" {
		return nil, fmt.Errorf("h2c not implemented")
	}

	if u.Port() == "" {
		u.Host = u.Host + ":443"
	}

	if u.Path == "" {
		u.Path = "/"
	}

	if err != nil {
		return nil, err
	}
	return &Target{
		URL:    *u,
		Status: "pending",
	}, nil
}

func (t *Target) GetConnection() (*tls.Conn, error) {
	var conn *tls.Conn

	portNumber, _ := strconv.Atoi(t.URL.Port())
	addr := fmt.Sprintf("%s:%d", t.URL.Hostname(), portNumber)

	dialer := &net.Dialer{
		Timeout: time.Second * 5,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         t.URL.Hostname(),
		NextProtos:         []string{"h2"},
	})

	if err != nil {
		return conn, fmt.Errorf("TLS Cannot be established")
	}

	//if the protocol is not h2, we return an error

	if conn.ConnectionState().NegotiatedProtocol != "h2" {
		return conn, fmt.Errorf("h2 is not supported")
	}

	return conn, nil

}
