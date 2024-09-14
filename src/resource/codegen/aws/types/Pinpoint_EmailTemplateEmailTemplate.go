package types

type Pinpoint_EmailTemplateEmailTemplate struct {
	//
	Headers []Pinpoint_EmailTemplateEmailTemplateHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// The message body, in HTML format, to use in email messages that are based on the message template. We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message.
	HtmlPart string `json:"htmlPart,omitempty" yaml:"htmlPart,omitempty"`

	// The unique identifier for the recommender model to use for the message template. Amazon Pinpoint uses this value to determine how to retrieve and process data from a recommender model when it sends messages that use the template, if the template contains message variables for recommendation data.
	RecommenderId string `json:"recommenderId,omitempty" yaml:"recommenderId,omitempty"`

	// Subject line, or title, to use in email messages that are based on the message template.
	Subject string `json:"subject,omitempty" yaml:"subject,omitempty"`

	// Message body, in plain text format, to use in email messages that are based on the message template. We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices.
	TextPart string `json:"textPart,omitempty" yaml:"textPart,omitempty"`

	// JSON object that specifies the default values to use for message variables in the message template. This object is a set of key-value pairs. Each key defines a message variable in the template. The corresponding value defines the default value for that variable. When you create a message that's based on the template, you can override these defaults with message-specific and address-specific variables and values.
	DefaultSubstitutions string `json:"defaultSubstitutions,omitempty" yaml:"defaultSubstitutions,omitempty"`

	//
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
