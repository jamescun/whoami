package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	Addr    = flag.String("addr", "127.0.0.1:8080", "server listen address")
	TLS     = flag.Bool("tls", false, "enable transport security")
	TLSKey  = flag.String("tls-key", "cert.key", "tls private key")
	TLSCert = flag.String("tls-cert", "cert.pem", "tls certificate")
)

type Whoami []byte

func NewWhoami() (Whoami, error) {
	buf := bytes.NewBuffer([]byte(`{"addrs":[`))

	inets, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	first := true
	for _, inet := range inets {
		addrs, err := inet.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			if !first {
				fmt.Fprint(buf, ",")
			}
			first = false

			fmt.Fprintf(buf, `"%s"`, addr.String())
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(buf, `],"hostname":"%s"}`, hostname)

	return Whoami(buf.Bytes()), nil
}

func (who Whoami) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(who)
}

func main() {
	flag.Parse()

	w, err := NewWhoami()
	if err != nil {
		log.Fatalln("fatal: whoami:", err)
	}

	s := &http.Server{
		Addr:    *Addr,
		Handler: w,
	}

	if *TLS {
		log.Printf("info: listen: https://%s\n", s.Addr)
		err = s.ListenAndServeTLS(*TLSCert, *TLSKey)
	} else {
		log.Printf("info: listen: http://%s\n", s.Addr)
		err = s.ListenAndServe()
	}

	if err != nil {
		log.Fatalln("fatal: listen:", err)
	}
}
