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

// This file contains the implementation of the User type.

package awx

// User represents an AWX User.
//
type User struct {
	id           int
	username     string
	is_superuser bool
}

// Id returns the unique identifier of the User.
//
func (p *User) Id() int {
	return p.id
}

// Name returns the username of the User.
//
func (p *User) Username() string {
	return p.username
}

// SCMType returns the source code management system type of the User.
//
func (p *User) Is_superuser() bool {
	return p.is_superuser
}
