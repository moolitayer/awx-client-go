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

// This example shows how to launch a job template.
//
// Use the following command to build and run it with all the debug output sent to the standard
// error output:
//
//	go run launch_job_template.go \
//		-url "https://awx.example.com/api" \
//		-username "admin" \
//      -proxy "http://proxy.com:3128" \
//		-password "..." \
//		-ca-file "ca.pem" \
//		-logtostderr \
//      -project "project" \
//		-template "template-name" \
//		-limit "node0.openshift.private" \
// 		-extra-vars "alertname=bla job=my-job complex={\"simple\":\"label\"}" \
//		-v=2

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/golang/glog"
	awx "github.com/moolitayer/awx-client-go/awx"
)

var (
	url           string
	username      string
	password      string
	proxy         string
	insecure      bool
	caFile        string
	project       string
	template      string
	limit         string
	extraVarsFlag string

//	extraVar map[string]interface{}
)

func init() {
	flag.StringVar(&url, "url", "https://awx.example.com/api", "API URL.")
	flag.StringVar(&username, "username", "admin", "API user name.")
	flag.StringVar(&password, "password", "password", "API user password.")
	flag.StringVar(&proxy, "proxy", "", "API proxy URL.")
	flag.BoolVar(&insecure, "insecure", false, "Don't verify server certificate.")
	flag.StringVar(&caFile, "ca-file", "", "Trusted CA certificates.")
	flag.StringVar(&project, "project", "", "Project Name.")
	flag.StringVar(&template, "template", "", "Template Name.")
	flag.StringVar(&limit, "limit", "", "Hosts limit")
	flag.StringVar(&extraVarsFlag, "extra-vars", "", "extra variables to the Job")
}

func main() {
	// Parse the command line:
	flag.Parse()

	var extraVars map[string]interface{}
	var err error
	if len(extraVarsFlag) > 0 {
		extraVars, err = parseExtraVars(extraVarsFlag)
		if err != nil {
			fmt.Printf("Failed to parse extra-vars %s: %v\n", extraVarsFlag, err)
			return
		}
	} else {
		// create default extraVars
		extraVars = map[string]interface{}{
			"node":  "example.com",
			"count": 4,
		}
	}

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

	// Get the template by name
	templatesResource := connection.JobTemplates()
	templatesResponse, err := templatesResource.Get().
		Filter("project__name", project).
		Filter("name", template).
		Send()

	if err != nil {
		fmt.Printf("Failed to get template resource %v\n", err)
		return
	}

	if templatesResponse.Count() == 0 {
		fmt.Printf(
			"Template '%s' not found in project '%s'\n",
			template,
			project,
		)
		return
	}

	// Launch all corresponding templated
	for _, t := range templatesResponse.Results() {
		launchResource := connection.JobTemplates().Id(t.Id()).Launch()

		if limit != "" && !t.AskLimitOnLaunch() {
			glog.Warningf("About to launch template '%s' with limit '%s', but 'prompt-on-launch' is false. Limit will be ignored",
				template, limit)
		}

		if extraVars != nil && !t.AskVarsOnLaunch() {
			glog.Warningf("About to launch template '%s' with extra-vars, but 'prompt-on-launch' is false. Extra Variables will be ignored",
				template)
		}

		response, err := launchResource.Post().
			Limit(limit).
			ExtraVars(extraVars).
			ExtraVar("my-var", "example-val").
			Send()
		if err != nil {
			fmt.Printf("Failed to get launch job %v\n", err)
			return
		}

		glog.Infof(
			"Request to launch AWX job from template '%s' has been sent, job identifier is '%v'",
			template,
			response.Job,
		)
	}
}

// Parse array of strings to extra vars json
// Expected input format: "a=b x=y c={\"v\":\"w\"}"
func parseExtraVars(input string) (output map[string]interface{}, err error) {
	variables := strings.Split(input, " ")
	if len(variables) > 0 {
		output = make(map[string]interface{})
	}
	for _, currVar := range variables {
		list := strings.SplitN(currVar, "=", 2)
		if len(list) != 2 {
			err = fmt.Errorf("bad format of extra-var")
			return
		}

		key := list[0]
		val := list[1]

		if val[0] == '{' {
			// handle complex json
			var parsedJson interface{}
			err = json.Unmarshal([]byte(val), &parsedJson)
			if err != nil {
				return
			}

			output[key] = parsedJson
		} else {
			output[key] = val
		}
	}
	return
}
