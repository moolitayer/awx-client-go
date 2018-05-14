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

// This file contains the implementation of the resource that manages a specific project.

package awx

import (
	"github.com/moolitayer/awx-client-go/awx/internal/data"
)

// ProjectResource is a Resource assciated with an AWX project.
type ProjectResource struct {
	Resource
}

// NewProjectResource creates a new ProjectResource.
func NewProjectResource(connection *Connection, path string) *ProjectResource {
	resource := new(ProjectResource)
	resource.connection = connection
	resource.path = path
	return resource
}

// Get returns a ProjectGetRequest for this ProjectResource so it could be sent.
func (r *ProjectResource) Get() *ProjectGetRequest {
	request := new(ProjectGetRequest)
	request.resource = &r.Resource
	return request
}

// ProjectGetRequest represents a GET request for an AWX project
type ProjectGetRequest struct {
	Request
}

// Send initiates a round trip against an AWX server.
func (r *ProjectGetRequest) Send() (response *ProjectGetResponse, err error) {
	output := new(data.ProjectGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(ProjectGetResponse)
	response.result = new(Project)
	response.result.id = output.ID
	response.result.name = output.Name
	response.result.scmType = output.SCMType
	response.result.scmURL = output.SCMURL
	response.result.scmBranch = output.SCMBranch
	return
}

// ProjectGetResponse represents a response for a GET request on a project.
type ProjectGetResponse struct {
	result *Project
}

// Result returns the Project result associated with this ProjectGetResponse.
func (r *ProjectGetResponse) Result() *Project {
	return r.result
}
