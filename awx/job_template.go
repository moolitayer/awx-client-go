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

// This file contains the implementation of the job template type.

package awx

// JobTemplate represents a single AWX job_template.
type JobTemplate struct {
	id               int
	name             string
	askLimitOnLaunch bool
	askVarsOnLaunch  bool
}

// ID is the identifier assigned to this JobTemplate by the AWX server.
func (t *JobTemplate) ID() int {
	return t.id
}

// Name is the name given to this JobTemplate by the AWX server.
func (t *JobTemplate) Name() string {
	return t.name
}

// AskLimitOnLaunch is true if this JobTemplate allows to override the limit
// parameter when running this template.
func (t *JobTemplate) AskLimitOnLaunch() bool {
	return t.askLimitOnLaunch
}

// AskVarsOnLaunch is true if this JobTemplate allows to override the extra_vars
// when running this template.
func (t *JobTemplate) AskVarsOnLaunch() bool {
	return t.askVarsOnLaunch
}
