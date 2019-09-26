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

// This file contains the data structures used for sending and receiving Organizations.

package data

type ExecuteRoles struct {
	Id int `json:"id,omitempty"`
}

//Results []*Organization `json:"results,omitempty"`
type ObjectRoles struct {
	Executeroles ExecuteRoles `json:"execute_role,omitempty"`
}

type SummaryFields struct {
	Objectroles ObjectRoles `json:"object_roles,omitempty"`
}
type Organization struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	//Execute_role_Id int    `json:"id,omitempty"` //`json:"summary_fields.object_roles.execute_role.id,omitempty"`
	Summaryfields SummaryFields `json:"summary_fields,omitempty"`
}

type OrganizationGetResponse struct {
	Organization
}
