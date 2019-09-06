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
// Organizations.

package awx

import (
	"fmt"

	//"github.com/moolitayer/awx-client-go/awx/internal/data"
	"./internal/data"
)

type OrganizationsResource struct {
	Resource
}

func NewOrganizationsResource(connection *Connection, path string) *OrganizationsResource {
	resource := new(OrganizationsResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *OrganizationsResource) Get() *OrganizationsGetRequest {
	request := new(OrganizationsGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *OrganizationsResource) Id(id int) *OrganizationResource {
	return NewOrganizationResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type OrganizationsGetRequest struct {
	Request
}

func (r *OrganizationsGetRequest) Filter(name string, value interface{}) *OrganizationsGetRequest {
	r.addFilter(name, value)
	return r
}

func (r *OrganizationsGetRequest) Send() (response *OrganizationsGetResponse, err error) {
	output := new(data.OrganizationsGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(OrganizationsGetResponse)
	response.count = output.Count
	response.previous = output.Previous
	response.next = output.Next
	response.results = make([]*Organization, len(output.Results))
	for i := 0; i < len(output.Results); i++ {
		response.results[i] = new(Organization)
		response.results[i].id = output.Results[i].Id
		response.results[i].name = output.Results[i].Name
		response.results[i].execute_role_id = output.Results[i].Summaryfields.Objectroles.Executeroles.Id
	}
	return
}

type OrganizationsGetResponse struct {
	ListGetResponse

	results []*Organization
}

func (r *OrganizationsGetResponse) Results() []*Organization {
	return r.results
}
