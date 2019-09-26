/*
Copyright (c) 2018 Red Hat, Inc.

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

// This example shows how to list the Organizations.
//
// Use the following command to build and run it with all the debug output sent to the standard
// error output:
//
//	go run list_Organizations.go \
//		-url "https://awx.example.com/api" \
//		-username "admin" \
//		-password "..." \
//		-ca-file "ca.pem" \
//		-logtostderr \
//		-v=2

package main

import (
	"flag"
	"fmt"

	"../awx"
)

var (
	url      string
	username string
	password string
	proxy    string
	insecure bool
	caFile   string
)

func init() {
	flag.StringVar(&url, "url", "https://awx.example.com/api", "API URL.")
	flag.StringVar(&username, "username", "admin", "API user name.")
	flag.StringVar(&password, "password", "password", "API user password.")
	flag.StringVar(&proxy, "proxy", "", "API proxy URL.")
	flag.BoolVar(&insecure, "insecure", false, "Don't verify server certificate.")
	flag.StringVar(&caFile, "ca-file", "", "Trusted CA certificates.")
}

func main() {
	// Parse the command line:
	flag.Parse()

	// Connect to the server, and remember to close the connection:
	connection, err := awx.NewConnectionBuilder().
		URL(url).
		Username(username).
		Password(password).
		Proxy(proxy).
		CAFile(caFile).
		Insecure(insecure).
		Build()
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// Find the resource that manages the collection of Organizations:
	organizationsResource := connection.Organizations()

	// Send the request to get the list of Organizations:
	getOrganizationsRequest := organizationsResource.Get()
	getOrganizationsResponse, err := getOrganizationsRequest.Send()
	if err != nil {
		panic(err)
	}

	// Print the results:
	organizations := getOrganizationsResponse.Results()
	for _, organization := range organizations {
		fmt.Printf("%d: %s \n", organization.Id(), organization.Name())
	}
}
