package zendesk

import "fmt"

// JobStatus represents a Zendesk JobStatus.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/job_statuses#json-format
type JobStatus struct {
	ID       *string  `json:"id,omitempty"`
	Message  *string  `json:"message,omitempty"`
	Progress *int64   `json:"progress,omitempty"`
	Results  []Result `json:"results,omitempty"`
	Status   *string  `json:"status,omitempty"`
	Total    *int64   `json:"total,omitempty"`
	URL      *string  `json:"url,omitempty"`
}

// Result represents the data from processed tasks within the Job Status
type Result struct {
	Action  *string `json:"action,omitempty"`
	Errors  *string `json:"errors,omitempty"`
	ID      *int64  `json:"id,omitempty"`
	Status  *string `json:"status,omitempty"`
	Success *string `json:"success,omitempty"`
	Title   *string `json:"title,omitempty"`
}

// Show Job Status shows the status of a background job
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/job_statuses#show-job-status
func (c *client) ShowJobStatus(id string) (*JobStatus, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/job_statuses/%s.json", id), out)
	return out.JobStatus, err
}
