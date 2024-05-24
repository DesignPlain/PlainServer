package directoryservice

type LogService struct {
	// ID of directory.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// Name of the cloudwatch log group to which the logs should be published. The log group should be already created and the directory service principal should be provided with required permission to create stream and publish logs. Changing this value would delete the current subscription and create a new one. A directory can only have one log subscription at a time.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`
}
