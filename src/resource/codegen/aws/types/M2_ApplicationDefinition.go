package types

type M2_ApplicationDefinition struct {
	// JSON application definition. Either this or `s3_location` must be specified.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// Location of the application definition in S3. Either this or `content` must be specified.
	S3Location string `json:"s3Location,omitempty" yaml:"s3Location,omitempty"`
}
