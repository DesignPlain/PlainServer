package types

type Evidently_ProjectDataDeliveryCloudwatchLogs struct {
	/*
	   The name of the log group where the project stores evaluation events.

	   The `s3_destination` block supports the following arguments:
	*/
	LogGroup string `json:"logGroup,omitempty" yaml:"logGroup,omitempty"`
}
