package types

type Fsx_LustreFileSystemLogConfiguration struct {
	// The Amazon Resource Name (ARN) that specifies the destination of the logs. The name of the Amazon CloudWatch Logs log group must begin with the `/aws/fsx` prefix. If you do not provide a destination, Amazon FSx will create and use a log stream in the CloudWatch Logs `/aws/fsx/lustre` log group.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Sets which data repository events are logged by Amazon FSx. Valid values are `WARN_ONLY`, `FAILURE_ONLY`, `ERROR_ONLY`, `WARN_ERROR` and `DISABLED`. Default value is `DISABLED`.
	Level string `json:"level,omitempty" yaml:"level,omitempty"`
}
