package main

import (
	"errors"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ParseVLESS(link string) (*VLESSServer, error) {

	if !strings.HasPrefix(link, "vless://") {
		return nil, errors.New("invalid VLESS link")
	}

	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	user := u.User.Username()

	if _, err := uuid.Parse(user); err != nil {
		return nil, errors.New("invalid UUID")
	}

	host, portString, err := net.SplitHostPort(u.Host)

	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(portString)

	if err != nil {
		return nil, err
	}

	query := u.Query()

	server := &VLESSServer{
		ID:          uuid.NewString(),
		Name:        host,
		Address:     host,
		Port:        port,
		UUID:        user,
		SNI:         query.Get("sni"),
		PublicKey:   query.Get("pbk"),
		ShortID:     query.Get("sid"),
		Fingerprint: query.Get("fp"),
		Flow:        query.Get("flow"),
		Enabled:     true,
		Ping:        -1,
	}

	return server, nil
}
