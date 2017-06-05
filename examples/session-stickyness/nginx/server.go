/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"io"
	"net/http"
	"log"
	"net"
	"fmt"
	"strings"
)

// HelloHttpListnerServer return the local ip to the response
func HelloHttpListnerServer(w http.ResponseWriter, req *http.Request) {
	requestAddr := req.RemoteAddr
	io.WriteString(w, "Request client is: " + requestAddr + "\n")
	io.WriteString(w, "Response backend server is: " + getLocalIps() + "\n")
}

// getLocalIps will return all ips which have ipv4 address
func getLocalIps() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return fmt.Sprintf("failed to get local ip info - %v", err)
	}

	addrInfos := []string{}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addrInfos = append(addrInfos, ipnet.IP.String() + ":9999")
			}
		}
	}
	return strings.Join(addrInfos, ", ")
}

// main will start a special http listener with the fixed port 9999
func main() {
	http.HandleFunc("/hello", HelloHttpListnerServer)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatalf("Failed to execute ListenAndServe: %v", err)
	}
}
