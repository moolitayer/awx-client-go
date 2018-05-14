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

// JobStatus represents possible values for AWX job statuses.
type JobStatus string

const (
	// JobStatusNew represents and AWX job with status new
	JobStatusNew JobStatus = "new"
	// JobStatusPending represents and AWX job with status pending.
	JobStatusPending JobStatus = "pending"
	// JobStatusWaiting represents and AWX job with status waiting.
	JobStatusWaiting JobStatus = "waiting"
	// JobStatusRunning represents and AWX job with status running.
	JobStatusRunning JobStatus = "running"
	// JobStatusSuccesful represents and AWX job with status successful.
	JobStatusSuccesful JobStatus = "successful"
	// JobStatusFailed represents and AWX job with status failed.
	JobStatusFailed JobStatus = "failed"
	// JobStatusError represents and AWX job with status error.
	JobStatusError JobStatus = "error"
	// JobStatusCancelled represents and AWX job with status cancelled.
	JobStatusCancelled JobStatus = "cancelled"
)

// Job Represents an AWX job.
type Job struct {
	id     int
	status JobStatus
}

// ID is The AWX identifier given to this job.
func (j *Job) ID() int {
	return j.id
}

// Status is the status of this job in AWX. See JobStatus.
func (j *Job) Status() JobStatus {
	return j.status
}

// IsFinished returns true if the state of this job's state is considered finished
func (j *Job) IsFinished() bool {
	switch j.status {
	case
		JobStatusSuccesful,
		JobStatusFailed,
		JobStatusError,
		JobStatusCancelled:
		return true
	}
	return false
}

//IsSuccessful returns true if the state of this job is JobStatusSuccesful
func (j *Job) IsSuccessful() bool {
	return j.status == JobStatusSuccesful
}
