package main

import (
	"crypto/tls"
	"flag"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is a wonderful HTTPS setup.\n"))
}

func main() {
	listenAddr := flag.String("listen", ":8443", "")
	tlsCert := flag.String("cert", "server.crt", "")
	tlsKey := flag.String("key", "server.key", "")

	flag.Parse()

	http.HandleFunc("/", HelloHandler)

	cer, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
		MinVersion:   tls.VersionSSL30,
		MaxVersion:   tls.VersionTLS10,
	}

	l, err := tls.Listen("tcp", *listenAddr, config)
	defer l.Close()

	err = http.Serve(l, nil)
	if err != nil {
		panic(err)
	}
}
