package types

type Pinpoint_EmailTemplateEmailTemplateHeader struct {
	// Name of the message header. The header name can contain up to 126 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value of the message header. The header value can contain up to 870 characters, including the length of any rendered attributes. For example if you add the {CreationDate} attribute, it renders as YYYY-MM-DDTHH:MM:SS.SSSZ and is 24 characters in length.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
