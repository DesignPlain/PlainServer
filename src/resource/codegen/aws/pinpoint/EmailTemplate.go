package pinpoint

import types "libds/aws/types"

type EmailTemplate struct {
	// Specifies the content and settings for a message template that can be used in messages that are sent through the email channel. See Email Template
	EmailTemplates []types.Pinpoint_EmailTemplateEmailTemplate `json:"emailTemplates,omitempty" yaml:"emailTemplates,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// name of the message template. A template name must start with an alphanumeric character and can contain a maximum of 128 characters. The characters can be alphanumeric characters, underscores (_), or hyphens (-). Template names are case sensitive.
	TemplateName string `json:"templateName,omitempty" yaml:"templateName,omitempty"`
}
