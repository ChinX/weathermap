package rest

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)



var (
	proxyURL  *url.URL
	noProxies []string
	TimeOut   = 10 * time.Second
	KeepAlive = 10 * time.Second
	unexpectedStatus = "expected status: %d, but result status: %d"
)

func Get(addr string, header map[string]string, params url.Values) (*http.Response, error) {
	if params != nil {
		if strings.Contains(addr, "?") {
			addr += "&"
		} else {
			addr += "?"
		}
		addr += params.Encode()
	}
	return request(http.MethodGet, addr, header, nil)
}

func Post(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodPost, addr, header, body)
}

func Put(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodPut, addr, header, body)
}

func Patch(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodPatch, addr, header, body)
}

func Delete(addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	return request(http.MethodDelete, addr, header, body)
}

func request(mothod, addr string, header map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(mothod, addr, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for key, val := range header {
			req.Header.Add(key, val)
		}
	}

	clien := http.Client{
		Transport: &http.Transport{
			Proxy:                 proxyFromConf,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}

	return clien.Do(req)
}

func ParseResponse(resp *http.Response, expected int, v interface{}) error {
	if resp.StatusCode != expected {
		return fmt.Errorf(unexpectedStatus, expected, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func FormatURL(addr string) (string, error) {
	raw, err := url.Parse(addr)
	if err != nil {
		return addr, err
	}

	if raw.Scheme != "http" && raw.Scheme != "https" {
		raw.Scheme = "https"
	}

	fu := raw.Scheme + "://" + raw.Host + raw.Path
	if !strings.HasSuffix(fu, "/") {
		fu += "/"
	}

	return fu, nil
}

func proxyFromConf(req *http.Request) (*url.URL, error) {
	if proxyURL == nil || inNoProxies(req.Host) {
		return nil, nil
	}
	return proxyURL, nil
}

func inNoProxies(addr string) bool {
	if len(noProxies) == 0 {
		return false
	}

	for i := range noProxies {
		if strings.Contains(addr, noProxies[i]) {
			return true
		}
	}

	return false
}

func SetProxyAddress(addr string) {
	proxy, err := url.Parse(addr)
	if err == nil {
		proxyURL = proxy
	}
}

func SetNoProxies(noProxiesAddr ...string) {
	for i := range noProxiesAddr {
		if noProxiesAddr[i] == "" {
			continue
		}
		noProxies = append(noProxies, noProxiesAddr[i])
	}
}
