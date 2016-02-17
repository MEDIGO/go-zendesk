package zendesk

import "time"

type Locale struct {
	ID        *int64     `json:"id,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Locale    *string    `json:"locale,omitempty"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c *client) LocaleList() ([]Locale, error) {
	out := new(APIPayload)
	err := c.get("/api/v2/locales.json", out)
	return out.Locales, err
}
