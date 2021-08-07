package zendesk

import "time"

// Schedule represents a schedule in Zendesk.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/support/schedules
type Schedule struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	TimeZone  string      `json:"time_zone"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Intervals []Intervals `json:"intervals"`
}

type Intervals struct {
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
}

func (c *client) ListSchedules() ([]Schedule, error) {
	out := new(APIPayload)
	err := c.get("/api/v2/business_hours/schedules.json", out)
	return out.Schedules, err
}
