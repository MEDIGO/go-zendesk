package zendesk

type Attachment struct {
	ID          *int64  `json:"id,omitempty"`
	FileName    *string `json:"file_name,omitempty"`
	ContentURL  *string `json:"content_url,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Size        *string `json:"size,omitempty"`
	Inline      *bool   `json:"inline,omitempty"`
}
