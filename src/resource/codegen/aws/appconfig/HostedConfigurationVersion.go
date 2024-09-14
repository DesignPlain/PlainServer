package appconfig

type HostedConfigurationVersion struct {
	// Configuration profile ID.
	ConfigurationProfileId string `json:"configurationProfileId,omitempty" yaml:"configurationProfileId,omitempty"`

	// Content of the configuration or the configuration data.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// Standard MIME type describing the format of the configuration content. For more information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Description of the configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`
}
