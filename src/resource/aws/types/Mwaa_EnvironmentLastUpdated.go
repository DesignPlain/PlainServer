package types

type Mwaa_EnvironmentLastUpdated struct {
	/*
	   The Created At date of the MWAA Environment
	   - `logging_configuration[0].<LOG_CONFIGURATION_TYPE>[0].cloud_watch_log_group_arn` - Provides the ARN for the CloudWatch group where the logs will be published
	*/
	CreatedAt string `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`

	//
	Errors []Mwaa_EnvironmentLastUpdatedError `json:"errors,omitempty" yaml:"errors,omitempty"`

	// The status of the Amazon MWAA Environment
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
