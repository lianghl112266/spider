package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Limit the send rate to prevent being banned
var timeRate = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-timeRate

	//Create client object and set proxy
	client := http.Client{}
	//client.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	client.Transport = GetTransportFieldURL("http://localhost:7897")

	//Create request object
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	//Get corresponding
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("The request is incorrect", err)
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	res, _ := io.ReadAll(resp.Body)
	return res, nil
}

// GetTransportFieldURL get proxy from URL
func GetTransportFieldURL(proxyAddr string) (transport *http.Transport) {
	url_i := url.URL{}
	url_proxy, err := url_i.Parse(proxyAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	transport = &http.Transport{Proxy: http.ProxyURL(url_proxy)}
	return
}

// GetTransportFromEnvironment get the HTTP proxy address from the environment variable $http proxy or $HTTP proxy
func GetTransportFromEnvironment() (transport *http.Transport) {
	transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	return
}
