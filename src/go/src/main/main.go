// according to https://stackoverflow.com/questions/30248259/swig-go-c-source-files-not-allowed-when-not-using-cgo
package main
import "C"

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	//"strings"
)

// Get the port to listen on
func getListenAddress(port string) string {
	return ":" + port
}

// Log the env variables required for a reverse proxy
func logSetup(port string, url string) {
	log.Printf("Server will run on: %s\n", getListenAddress(port))
	log.Printf("Proxying url: %s\n", url)
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	log.Printf("serveReverseProxy start")
	// parse the url
	url, _ := url.Parse(target)
	log.Printf(req.RequestURI)
	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
	
	for k, v := range res.Header() {
         log.Printf(k + " : " + v[0])
  }
	log.Printf("serveReverseProxy end")
}

type makeHandler struct {
    url string
}

func (m *makeHandler) ServeHTTP (res http.ResponseWriter, req *http.Request) {
  log.Printf("handler start")
	url := m.url

	serveReverseProxy(url, res, req)
	log.Printf("handler end")
}

//export runProxy
func runProxy(port string, url string) int {
  // Log setup values
	logSetup(port, url)

	// start server
	my_handler := &makeHandler{ url }

  http.Handle("/", my_handler)
  
	if err := http.ListenAndServe(getListenAddress(port), nil); err != nil {
		panic(err)
	}
	return 1
}


func main() {}
