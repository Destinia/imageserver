// Package httpsource provides a HTTP Source example.
//
// Try http://localhost:8080/large.jpg
// or any image available in https://github.com/Destinia/imageserver/tree/master/testdata
package main

import (
	"flag"
	"net/http"

	imageserver_http "github.com/Destinia/imageserver/http"
	_ "github.com/Destinia/imageserver/image/gif"
	_ "github.com/Destinia/imageserver/image/jpeg"
	_ "github.com/Destinia/imageserver/image/png"
	imageserver_source_http "github.com/Destinia/imageserver/source/http"
)

const (
	urlPrefix = "https://raw.githubusercontent.com/Destinia/imageserver/master/testdata/"
)

var (
	flagHTTP = ":8080"
)

func main() {
	parseFlags()
	startHTTPServer()
}

func parseFlags() {
	flag.StringVar(&flagHTTP, "http", flagHTTP, "HTTP")
	flag.Parse()
}

func startHTTPServer() {
	http.Handle("/", http.StripPrefix("/", newImageHTTPHandler()))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(flagHTTP, nil)
	if err != nil {
		panic(err)
	}
}

func newImageHTTPHandler() http.Handler {
	return &imageserver_http.Handler{
		Parser: &imageserver_http.SourcePrefixParser{
			Parser: &imageserver_http.SourcePathParser{},
			Prefix: urlPrefix,
		},
		Server: &imageserver_source_http.Server{},
	}
}
