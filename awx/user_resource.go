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

// This file contains the implementation of the resource that manages a specific User.

package awx

import (
	//"github.com/moolitayer/awx-client-go/awx/internal/data"
	"./internal/data"
)

type UserResource struct {
	Resource
}

func NewUserResource(connection *Connection, path string) *UserResource {
	resource := new(UserResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *UserResource) Get() *UserGetRequest {
	request := new(UserGetRequest)
	request.resource = &r.Resource
	return request
}

type UserGetRequest struct {
	Request
}

func (r *UserGetRequest) Send() (response *UserGetResponse, err error) {
	output := new(data.UserGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(UserGetResponse)
	response.result = new(User)
	response.result.id = output.Id
	response.result.username = output.Username
	response.result.is_superuser = output.Is_superuser
	return
}

type UserGetResponse struct {
	result *User
}

func (r *UserGetResponse) Result() *User {
	return r.result
}
