package types

type Securitylake_SubscriberSource struct {
	// Amazon Security Lake supports log and event collection for natively supported AWS services. See `aws_log_source_resource` Block below.
	AwsLogSourceResource Securitylake_SubscriberSourceAwsLogSourceResource `json:"awsLogSourceResource,omitempty" yaml:"awsLogSourceResource,omitempty"`

	// Amazon Security Lake supports custom source types. See `custom_log_source_resource` Block below.
	CustomLogSourceResource Securitylake_SubscriberSourceCustomLogSourceResource `json:"customLogSourceResource,omitempty" yaml:"customLogSourceResource,omitempty"`
}
