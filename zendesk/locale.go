package zendesk

import (
	"fmt"
	"time"
)

type Locale struct {
	ID        *int64     `json:"id,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Locale    *string    `json:"locale,omitempty"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c *client) ListLocales() ([]Locale, error) {
	out := new(APIPayload)
	err := c.get("/api/v2/locales.json", out)
	return out.Locales, err
}

func (c *client) ShowLocale(id int64) (*Locale, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/locales/%d.json", id), out)
	return out.Locale, err
}

func (c *client) ShowLocaleByCode(code string) (*Locale, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/locales/%s.json", code), out)
	return out.Locale, err
}
