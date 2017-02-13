package zendesk

import (
	"fmt"
	"io"

	"github.com/google/go-querystring/query"
)

// Attachment represents a Zendesk attachment for tickets and forum posts.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/attachments
type Attachment struct {
	ID          *int64  `json:"id,omitempty"`
	FileName    *string `json:"file_name,omitempty"`
	ContentURL  *string `json:"content_url,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Size        *int64  `json:"size,omitempty"`
	Inline      *bool   `json:"inline,omitempty"`
}

// Upload represents a Zendesk file upload.
type Upload struct {
	Token       *string      `json:"token"`
	Attachment  *Attachment  `json:"attachment"`
	Attachments []Attachment `json:"attachments"`
}

// ShowAttachment fetches an attachment by its ID.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/attachments#getting-attachments
func (c *client) ShowAttachment(id int64) (*Attachment, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/tickets/%d.json", id), out)
	return out.Attachment, err
}

// UploadFile uploads a file as a io.Reader.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/attachments#uploading-files
func (c *client) UploadFile(filename string, token *string, filecontent io.Reader) (*Upload, error) {
	params, err := query.Values(struct {
		Filename string  `url:"filename"`
		Token    *string `url:"token,omitempty"`
	}{filename, token})
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": "application/binary",
	}

	res, err := c.request("POST", fmt.Sprintf("/api/v2/uploads.json?%s", params.Encode()), headers, filecontent)
	if err != nil {
		return nil, err
	}

	out := new(APIPayload)
	err = unmarshall(res, out)
	return out.Upload, err
}
