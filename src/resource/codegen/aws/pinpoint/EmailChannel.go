package pinpoint

type EmailChannel struct {
	// The email address used to send emails from. You can use email only (`user@example.com`) or friendly address (`User <user@example.com>`). This field comply with [RFC 5322](https://www.ietf.org/rfc/rfc5322.txt).
	FromAddress string `json:"fromAddress,omitempty" yaml:"fromAddress,omitempty"`

	// The ARN of an identity verified with SES.
	Identity string `json:"identity,omitempty" yaml:"identity,omitempty"`

	// -Deprecated- The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// The ARN of the Amazon SES configuration set that you want to apply to messages that you send through the channel.
	ConfigurationSet string `json:"configurationSet,omitempty" yaml:"configurationSet,omitempty"`

	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
