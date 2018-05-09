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

func TestFilterJsonBytes(t *testing.T) {
	input := []byte("{\"Password\":\"foo\"}")
	expected := []byte("{\"Password\":\"REDACTED\"}")
	result := filterJsonBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = []byte("{\"aaa\":{\"a\":\"a\",\"password\":\"foo\"}}")
	expected = []byte("{\"aaa\":{\"a\":\"a\",\"password\":\"REDACTED\"}}")
	result = filterJsonBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = []byte("{\"aaa\":[{\"password\":\"foo\"},\"bar\"]}")
	expected = []byte("{\"aaa\":[{\"password\":\"REDACTED\"},\"bar\"]}")
	result = filterJsonBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = []byte("[{\"password\":\"foo\"},\"bar\"]")
	expected = []byte("[{\"password\":\"REDACTED\"},\"bar\"]")
	result = filterJsonBytes(input)
	if string(result) != string(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}


}
