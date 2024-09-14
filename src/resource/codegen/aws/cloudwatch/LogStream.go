package cloudwatch

type LogStream struct {
	// The name of the log group under which the log stream is to be created.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	// The name of the log stream. Must not be longer than 512 characters and must not contain `:`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
