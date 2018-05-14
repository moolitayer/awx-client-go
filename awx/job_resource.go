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

// This file contains the implementation of the resource that manages launching of jobs from job
// templates.

package awx

import (
	"github.com/moolitayer/awx-client-go/awx/internal/data"
)

// JobResource represents AWX job.
type JobResource struct {
	Resource
}

// NewJobResource is used to create a JobResource.
func NewJobResource(connection *Connection, path string) *JobResource {
	resource := new(JobResource)
	resource.connection = connection
	resource.path = path
	return resource
}

// Get a JobGetRequest for this JobResource that could be sent.
func (r *JobResource) Get() *JobGetRequest {
	request := new(JobGetRequest)
	request.resource = &r.Resource
	return request
}

// JobGetRequest represents a GET request on a Job resource.
type JobGetRequest struct {
	Request
}

// Send initiates a round trip against an AWX server.
func (r *JobGetRequest) Send() (response *JobGetResponse, err error) {
	output := new(data.JobGetResponse)
	err = r.get(output)
	if err != nil {
		return nil, err
	}
	response = new(JobGetResponse)
	if output != nil {
		response.job = new(Job)
		response.job.id = output.ID
		response.job.status = (JobStatus)(output.Status)
	}
	return
}

// JobGetResponse represents a response from the server for a Job GET request.
type JobGetResponse struct {
	job *Job
}

// Job returns the Job associated with this JobGetResponse
func (r *JobGetResponse) Job() *Job {
	return r.job
}
