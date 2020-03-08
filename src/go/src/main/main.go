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

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Get the port to listen on
func getListenAddress() string {
	port := getEnv("PORT", "1338")
	return ":" + port
}

// Log the env variables required for a reverse proxy
func logSetup() {
	a_condtion_url := os.Getenv("A_CONDITION_URL")
	b_condtion_url := os.Getenv("B_CONDITION_URL")
	default_condtion_url := os.Getenv("DEFAULT_CONDITION_URL")

	log.Printf("Server will run on: %s\n", getListenAddress())
	log.Printf("Redirecting to A url: %s\n", a_condtion_url)
	log.Printf("Redirecting to B url: %s\n", b_condtion_url)
	log.Printf("Redirecting to Default url: %s\n", default_condtion_url)
}

type requestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condition"`
}

// Get a json decoder for a given requests body
func requestBodyDecoder(request *http.Request) *json.Decoder {
	log.Printf("requestBodyDecoder start")
	// Read body to buffer
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	// Because go lang is a pain if you read the body then any susequent calls
	// are unable to read the body again....
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	log.Printf("requestBodyDecoder end")
	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

// Parse the requests body
func parseRequestBody(request *http.Request) requestPayloadStruct {
	log.Printf("parseRequestBody start")
	decoder := requestBodyDecoder(request)

	var requestPayload requestPayloadStruct
	err := decoder.Decode(&requestPayload)

	if err != nil {
		panic(err)
	}

	log.Printf("parseRequestBody end")
	return requestPayload
}

// Log the typeform payload and redirect url
func logRequestPayload(requestionPayload requestPayloadStruct, proxyUrl string) {
	log.Printf("proxy_condition: %s, proxy_url: %s\n", requestionPayload.ProxyCondition, proxyUrl)
}

// Get the url for a given proxy condition
func getProxyUrl(proxyConditionRaw string) string {
	log.Printf("getProxyUrl start")
	//proxyCondition := strings.ToUpper(proxyConditionRaw)

	//a_condtion_url := os.Getenv("A_CONDITION_URL")
	//b_condtion_url := os.Getenv("B_CONDITION_URL")
	default_condtion_url := os.Getenv("DEFAULT_CONDITION_URL")

	//if proxyCondition == "A" {
	//	return a_condtion_url
	//}

	//if proxyCondition == "B" {
	//	return b_condtion_url
	//}

	log.Printf("getProxyUrl end")
	return default_condtion_url
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	log.Printf("serveReverseProxy start")
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
	log.Printf("serveReverseProxy end")
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	log.Printf("handleRequestAndRedirect start")
	//requestPayload := parseRequestBody(req)
	url := getProxyUrl("hi")//requestPayload.ProxyCondition)

	//logRequestPayload(requestPayload, url)

	serveReverseProxy(url, res, req)
	log.Printf("handleRequestAndRedirect end")
}


//export runProxy
func runProxy() int {
//  // Log setup values
//	logSetup()
//
//	// start server
//	http.HandleFunc("/", handleRequestAndRedirect)
//	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
//		panic(err)
//	}
	return 1
}


func main() {}
