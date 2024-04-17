package fcrequests

import (
	"io"
	"net/http"
	"strings"

)

type Header struct {
	Key   string
	Value string
}

// Realiza uma requisição HTTP POST. Os headers devem ser passados como parâmetros.
// Retorna uma resposta http.
func Post(url, body string, headers ...Header) (*http.Response, error) {
	// method := "Post"

	bodyBytes := strings.NewReader(body)
	rq, err := http.NewRequest(http.MethodPost, url, bodyBytes)
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		rq.Header.Add(header.Key, header.Value)
	}
	res, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Realiza uma requisição HTTP POST com autenticação básica. Os headers devem ser passados como parâmetros.
func PostWithBasicAuth(url, body, username, password string, headers ...Header) (*http.Response, error) {
	// method := "PostWithBasicAuth"

	bodyBytes := strings.NewReader(body)
	rq, err := http.NewRequest(http.MethodPost, url, bodyBytes)
	if err != nil {
		return nil, err
	}
	rq.SetBasicAuth(username, password)
	for _, header := range headers {
		rq.Header.Add(header.Key, header.Value)
	}
	res, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	requestReaderReplacer := strings.NewReader(body)
	resultReaderReplacer := strings.NewReader(string(resBody))
	res.Body = io.NopCloser(resultReaderReplacer)
	res.Request.Body = io.NopCloser(requestReaderReplacer)

	return res, nil
}

// Realiza uma requisição HTTP GET. Os headers devem ser passados como parâmetros.
func Get(url string, headers ...Header) (*http.Response, error) {
	// // method := "fcrequests.Get"
	rq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		rq.Header.Add(header.Key, header.Value)
	}
	res, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Realiza uma requisição HTTP PUT. Os headers devem ser passados como parâmetros.
func Put(url, body string, headers ...Header) (*http.Response, error) {
	// method := "Post"

	bodyBytes := strings.NewReader(body)
	rq, err := http.NewRequest(http.MethodPut, url, bodyBytes)
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		rq.Header.Add(header.Key, header.Value)
	}
	res, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Realiza uma requisição HTTP GET com autenticação básica. Os headers devem ser passados como parâmetros.
func GetWithBasicAuth(url, username, password string, headers ...Header) (*http.Response, error) {
	// method := "fcrequets.GetWithBasicAuth"
	rq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		rq.Header.Add(header.Key, header.Value)
	}
	rq.SetBasicAuth(username, password)

	res, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}
	return res, nil
}
