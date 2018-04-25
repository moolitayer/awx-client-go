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

import "testing"

func TestIsSuccessful(t *testing.T) {
	for _, status := range []JobStatus{
		JobStatusNew, JobStatusPending, JobStatusWaiting, JobStatusRunning,
		JobStatusFailed, JobStatusError, JobStatusCancelled,
	} {
		if (&Job{0, status}).IsSuccessful() {
			t.Errorf("Job.IsSuccessful() Should return false for %s", status)
		}
	}
	if !(&Job{0, JobStatusSuccesful}).IsSuccessful() {
		t.Errorf("Job.IsSuccessful() Should return true for JobStatusSuccesful")
	}
}

func TestIsFinished(t *testing.T) {
	for _, status := range []JobStatus{
		JobStatusNew, JobStatusPending, JobStatusWaiting, JobStatusRunning,
	} {
		if (&Job{0, status}).IsFinished() {
			t.Errorf("Job.IsFinished() Should return false for %s", status)
		}
	}
	for _, status := range []JobStatus{
		JobStatusSuccesful, JobStatusFailed, JobStatusError, JobStatusCancelled,
	} {
		if !(&Job{0, status}).IsFinished() {
			t.Errorf("Job.IsFinished() Should return false for %s", status)
		}
	}
}
