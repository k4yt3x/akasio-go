/*
Name: AKASIO (Go)
Creator: K4YT3X
Date Created: June 14, 2020
Last Modified: June 15, 2020

Licensed under the GNU General Public License Version 3 (GNU GPL v3),
    available at: https://www.gnu.org/licenses/gpl-3.0.txt
(C) 2020 K4YT3X
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	// RedirectTable defines redirect table path
	RedirectTable = "configs/redirect.json"

	// Version defines the version number of this application
	Version = "1.0.0"

	// Hostname defines the hostname of the server
	Hostname = "akas.io"
)

func readRedirectTable(uri string) string {
	// Open our jsonFile
	jsonFile, err := os.Open(RedirectTable)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read JSON file into byte stream
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// unmarshal JSON bytes into a map
	var objmap map[string]json.RawMessage
	err = json.Unmarshal(byteValue, &objmap)

	// get target URL to redirect to from the redirect table
	var targetURL string
	err = json.Unmarshal(objmap[uri], &targetURL)

	return targetURL
}

func requestHandler(responseWriter http.ResponseWriter, request *http.Request) {

	// print request information
	log.Printf("%s: %s", request.RemoteAddr, request.URL)

	// if hostname does not match, return 444
	// this prevents host spoofing
	if request.Host != Hostname {
		http.Error(responseWriter, "", 444)
		return
	}

	// get target URL from redirect table
	targetURL := readRedirectTable(request.RequestURI)

	// if URL not found in redirect table, return 404
	if targetURL == "" {
		http.Error(responseWriter, "", 404)
	} else {
		// send 301 response to client and redirect client to target URL
		http.Redirect(responseWriter, request, targetURL, 301)
	}
}

func main() {

	// let requestHandler handle all requests
	http.HandleFunc("/", requestHandler)

	// listen on port 8080
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
