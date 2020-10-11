/*
Name: AKASIO
Creator: K4YT3X
Date Created: June 14, 2020
Last Modified: October 11, 2020

Licensed under the GNU General Public License Version 3 (GNU GPL v3),
    available at: https://www.gnu.org/licenses/gpl-3.0.txt
(C) 2020 K4YT3X
*/

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"

	"go.uber.org/zap"
)

const (
	// Version defines the version number of this application
	Version = "1.0.0"
)

// define command line flags
var bindAddress = flag.String("b", "127.0.0.1:8000", "binding address (IP:port)")
var debug = flag.Bool("d", false, "enable debugging mode, which disables security checks")
var hostname = flag.String("n", "akas.io", "server hostname")
var redirectTablePath = flag.String("r", "/etc/akasio.json", "redirect table path")

// readRedirectTable returns the target URL the URI corresponds to
func readRedirectTable(uri string) (string, error) {
	// open redirect table
	jsonFile, err := os.Open(*redirectTablePath)

	// if os.Open returns an error then log and return
	if err != nil {
		zap.S().Error(err)
		return "", err
	}

	// defer the closing of jsonFile so it can be parsed
	defer jsonFile.Close()

	// read JSON file into byte stream
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// unmarshal JSON bytes into a map
	var objmap map[string]json.RawMessage
	json.Unmarshal(byteValue, &objmap)

	// get target URL to redirect to from the redirect table
	var targetURL string
	json.Unmarshal(objmap[uri], &targetURL)

	return targetURL, nil
}

// requestHandler handles the incoming HTTP requests
func requestHandler(responseWriter http.ResponseWriter, request *http.Request) {

	// print request information
	zap.S().Infof("%s: %s%s", request.RemoteAddr, request.Host, request.URL)

	// if hostname does not match, return 401 unauthorized
	// this prevents host spoofing
	if request.Host != *hostname && !*debug {
		zap.S().Infof("Responding %s with code 401 (Unauthorized)", request.RemoteAddr)
		http.Error(responseWriter, "401 Unauthorized", http.StatusUnauthorized)
		return
	}

	// get target URL from redirect table
	targetURL, err := readRedirectTable(request.RequestURI)

	if targetURL == "" {
		// return 404 if URL not found in redirect table
		zap.S().Infof("Responding %s with code 404 (Not Found)", request.RemoteAddr)
		http.Error(responseWriter, "404 Not Found", http.StatusNotFound)
	} else if err != nil {
		// send 500 internal error if readRedirectTable returns an error
		zap.S().Infof("Responding %s with code 500 (Internal Server Error)", request.RemoteAddr)
		http.Error(responseWriter, "500 Internal Server Error", http.StatusInternalServerError)
	} else {
		// send 301 response to client and redirect client to target URL
		zap.S().Infof("Responding %s with code 301 (Moved Permanently) to %s", request.RemoteAddr, targetURL)
		http.Redirect(responseWriter, request, targetURL, http.StatusMovedPermanently)
	}
}

func main() {
	flag.Parse()

	// create new zap production logger and replace the global logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	// print some basic information
	zap.S().Info("Akasio server started")
	zap.S().Infof("Listening on %s", *bindAddress)
	zap.S().Infof("Using redirect table at: %s", *redirectTablePath)

	// let requestHandler handle all requests
	http.HandleFunc("/", requestHandler)

	// listen on port 8000
	// there should be a front-end server like Apache or Caddy in front of this server
	err := http.ListenAndServe(*bindAddress, nil)

	if err != nil {
		zap.S().Fatal("ListenAndServe: ", err)
	}
}
