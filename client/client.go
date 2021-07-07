package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	fmt.Println("Running the client")
	caCert, err := ioutil.ReadFile("../client_certs/ca.crt")
	if err != nil {
		fmt.Println(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair("../client_certs/client.crt", "../client_certs/client.key")
	if err != nil {
		fmt.Println(err)
	}

	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	transport := http2.Transport{
		TLSClientConfig: tlsConfig,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			// from https://github.com/golang/net/blob/04defd469f4e290175cd2fb95a0e5f235f9bf173/http2/not_go115.go
			cn, err := tls.Dial(network, addr, cfg)
			if err != nil {
				return nil, err
			}
			if err := cn.Handshake(); err != nil {
				return nil, err
			}
			if cfg.InsecureSkipVerify {
				return cn, nil
			}
			if err := cn.VerifyHostname(cfg.ServerName); err != nil {
				return nil, err
			}
			return cn, nil
		},
	}

	client := &http.Client{Transport: &transport}

	fmt.Println("Starting the GET...")
	resp, err := client.Get("https://localhost:61001")
	if err != nil {
		fmt.Printf("Error getting: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Success")
	fmt.Println(resp.StatusCode)
}
