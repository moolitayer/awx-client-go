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

// This file contains the implementation of the resource that manages a specific Organization.

package awx

import (
	//"github.com/moolitayer/awx-client-go/awx/internal/data"
	"./internal/data"
)

type OrganizationResource struct {
	Resource
}

func NewOrganizationResource(connection *Connection, path string) *OrganizationResource {
	resource := new(OrganizationResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *OrganizationResource) Get() *OrganizationGetRequest {
	request := new(OrganizationGetRequest)
	request.resource = &r.Resource
	return request
}

type OrganizationGetRequest struct {
	Request
}

func (r *OrganizationGetRequest) Send() (response *OrganizationGetResponse, err error) {
	output := new(data.OrganizationGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(OrganizationGetResponse)
	response.result = new(Organization)
	response.result.id = output.Id
	response.result.name = output.Name
	response.result.execute_role_id = output.Summaryfields.Objectroles.Executeroles.Id
	return
}

type OrganizationGetResponse struct {
	result *Organization
}

func (r *OrganizationGetResponse) Result() *Organization {
	return r.result
}
