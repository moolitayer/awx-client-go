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

// This file contains the implementation of the Organization type.

package awx

// Organization represents an AWX Organization.
//
type Organization struct {
	id              int
	name            string
	execute_role_id int
}

// Id returns the unique identifier of the Organization.
//
func (p *Organization) Id() int {
	return p.id
}

// Name returns the name of the Organization.
//
func (p *Organization) Name() string {
	return p.name
}

// Name returns the name of the Organization.
//
func (p *Organization) Execute_role_Id() int {
	return p.execute_role_id
}
