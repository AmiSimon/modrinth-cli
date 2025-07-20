package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO: Handle 410 Gone (Deprecated)

const Version = "0.0.1"

type ApiData struct {
	BuildInfo struct {
		Profile string
	}
	Version string
}

func GetApiData(url string) (string, error) {
	client := &http.Client{}
	
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("User-Agent", fmt.Sprintf("AmiSimon/modrinth-cli/%v (simon.leneveu@gmail.com)", Version))

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	if response.StatusCode == 410 {
		log.Fatal("API Returned 410 Gone, it is deprecated, switch to newer")
	}
}

func main() {

}