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

package awx

import (
	"testing"

	"github.com/seborama/govcr"
)

func TestFilterHeader(t *testing.T) {
	result := filterHeader("password", []string{"foo1"})
	expected := "REDACTED"
	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	result = filterHeader("hello", []string{"foo"})
	expected = "foo"
	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFilterJSONBytes(t *testing.T) {
	input := []byte("{\"Password\":\"foo\"}")
	expected := []byte("{\"Password\":\"REDACTED\"}")
	result := filterJSONBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = []byte("{\"aaa\":{\"a\":\"a\",\"password\":\"foo\"}}")
	expected = []byte("{\"aaa\":{\"a\":\"a\",\"password\":\"REDACTED\"}}")
	result = filterJSONBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = []byte("{\"aaa\":[{\"password\":\"foo\"},\"bar\"]}")
	expected = []byte("{\"aaa\":[{\"password\":\"REDACTED\"},\"bar\"]}")
	result = filterJSONBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = []byte("[{\"password\":\"foo\"},\"bar\"]")
	expected = []byte("[{\"password\":\"REDACTED\"},\"bar\"]")
	result = filterJSONBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// When the api/o endpoint is not available, the server should accquire a token
// through api/v2/users/<name>/personal_tokens
func TestOAUTH2Token(t *testing.T) {
	connection, err := NewConnectionBuilder().
		URL("http://localhost:9100/api").
		Username("admin").
		Password("password").
		Build()
	if err != nil {
		t.Error(err)
	}
	defer connection.Close()
	vcr := govcr.NewVCR("connection_oauth2",
		&govcr.VCRConfig{
			Client:           connection.client,
			DisableRecording: true,
		})
	// Replace our HTTPClient with a vcr client wrapping it
	connection.client = vcr.Client
	projectsResource := connection.Projects()

	// Trigger the auth flow.
	getProjectsRequest := projectsResource.Get()
	if len(connection.token) != 0 || len(connection.bearer) != 0 {
		t.Errorf("Connection should have no tokens. token: '%s', bearer: '%s'",
			connection.token,
			connection.bearer)
	}
	_, err = getProjectsRequest.Send()
	if err != nil {
		panic(err)
	}
	if len(connection.token) != 0 || len(connection.bearer) == 0 {
		t.Errorf("Connection should have only a bearer token. token: '%s', bearer: '%s'",
			connection.token,
			connection.bearer)
	}
}

//
// When the api/o endpoint is not available, the server should accquire a token
// through api/v2/authtoken/
func TestPreOAUTH2(t *testing.T) {
	//
	// Password manuall edited in cassete:
	// Basic = printf "admin:PASSWORD"| base64
	// Body = printf '{"username":"admin","password":"PASSWORD"}'|base64
	connection, err := NewConnectionBuilder().
		URL("https://tower.private/api").
		Username("admin").
		Password("PASSWORD").
		Insecure(true).
		Build()
	if err != nil {
		t.Errorf("Error creating connection: %s", err)
	}
	defer connection.Close()
	vcr := govcr.NewVCR("connection_pre_oauth2",
		&govcr.VCRConfig{
			Client:           connection.client,
			DisableRecording: true,
		})
	// Replace our HTTPClient with a vcr client wrapping it
	connection.client = vcr.Client
	projectsResource := connection.Projects()

	// Trigger the auth flow.
	getProjectsRequest := projectsResource.Get()
	if len(connection.token) != 0 || len(connection.bearer) != 0 {
		t.Errorf("Connection should have no tokens. token: '%s', bearer: '%s'",
			connection.token,
			connection.bearer)
	}
	_, err = getProjectsRequest.Send()
	if err != nil {
		t.Errorf("Error sending project request: %s", err)
	}
	if len(connection.token) == 0 || len(connection.bearer) != 0 {
		t.Errorf("Connection should have only an auth token. token: '%s', bearer: '%s'",
			connection.token,
			connection.bearer)
	}
}
