package types

type Lambda_FunctionEphemeralStorage struct {
	// The size of the Lambda function Ephemeral storage(`/tmp`) represented in MB. The minimum supported `ephemeral_storage` value defaults to `512`MB and the maximum supported value is `10240`MB.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`
}
