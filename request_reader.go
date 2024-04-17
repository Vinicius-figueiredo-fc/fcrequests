package fcrequests

import (
	"io"
	"net/http"
	"strings"
)

const (
	loopbackIpv4Address string = "[::1]"
	loopbackIpv6Address string = "127.0.0.1"
	Authorization       string = "Authorization"
	ContentType         string = "Content-Type"
)

func ReadAndReplace(r io.ReadCloser) (io.ReadCloser, []byte, error) {
	if bytesVal, err := io.ReadAll(r); err != nil {
		return nil, bytesVal, err
	} else {
		r.Close()
		return io.NopCloser(strings.NewReader(string(bytesVal))), bytesVal, nil
	}

}

func ReadRequest(r *http.Request) (RequestInfo, error) {
	remoteAddr := strings.ReplaceAll(r.RemoteAddr, loopbackIpv4Address, loopbackIpv6Address)

	var err error
	var body []byte
	if r.Body, body, err = ReadAndReplace(r.Body); err != nil {
		return RequestInfo{}, err
	}
	result := RequestInfo{
		HostUrl:       r.Host + r.RequestURI,
		SenderUrl:     remoteAddr,
		RequestMethod: r.Method,
		Body:          string(body),
		ContentType:   r.Header.Get(ContentType),
		RequestObj:    r,
		Authorization: r.Header.Get(Authorization),
	}
	return result, nil
}

type RequestInfo struct {
	SenderUrl     string
	HostUrl       string
	RequestMethod string
	ContentType   string
	Body          string
	RequestObj    *http.Request
	Authorization string
}
