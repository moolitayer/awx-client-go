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

// This file contains the implementation of the resource that manages the collection of
// Users.

package awx

import (
	"fmt"

	//"github.com/moolitayer/awx-client-go/awx/internal/data"
	"./internal/data"
)

type UsersResource struct {
	Resource
}

func NewUsersResource(connection *Connection, path string) *UsersResource {
	resource := new(UsersResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *UsersResource) Get() *UsersGetRequest {
	request := new(UsersGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *UsersResource) Id(id int) *UserResource {
	return NewUserResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type UsersGetRequest struct {
	Request
}

func (r *UsersGetRequest) Filter(name string, value interface{}) *UsersGetRequest {
	r.addFilter(name, value)
	return r
}

func (r *UsersGetRequest) Send() (response *UsersGetResponse, err error) {
	output := new(data.UsersGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(UsersGetResponse)
	response.count = output.Count
	response.previous = output.Previous
	response.next = output.Next
	response.results = make([]*User, len(output.Results))
	for i := 0; i < len(output.Results); i++ {
		response.results[i] = new(User)
		response.results[i].id = output.Results[i].Id
		response.results[i].username = output.Results[i].Username
		response.results[i].is_superuser = output.Results[i].Is_superuser
	}
	return
}

type UsersGetResponse struct {
	ListGetResponse

	results []*User
}

func (r *UsersGetResponse) Results() []*User {
	return r.results
}
