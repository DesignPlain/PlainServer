package types

type Compute_SubnetworkLogConfig struct {
	/*
	   Can only be specified if VPC flow logging for this subnetwork is enabled.
	   Toggles the aggregation interval for collecting flow logs. Increasing the
	   interval time will reduce the amount of generated flow logs for long
	   lasting connections. Default is an interval of 5 seconds per connection.
	   Default value is `INTERVAL_5_SEC`.
	   Possible values are: `INTERVAL_5_SEC`, `INTERVAL_30_SEC`, `INTERVAL_1_MIN`, `INTERVAL_5_MIN`, `INTERVAL_10_MIN`, `INTERVAL_15_MIN`.
	*/
	AggregationInterval string `json:"aggregationInterval,omitempty" yaml:"aggregationInterval,omitempty"`

	/*
	   Export filter used to define which VPC flow logs should be logged, as as CEL expression. See
	   https://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.
	   The default value is 'true', which evaluates to include everything.
	*/
	FilterExpr string `json:"filterExpr,omitempty" yaml:"filterExpr,omitempty"`

	/*
	   Can only be specified if VPC flow logging for this subnetwork is enabled.
	   The value of the field must be in [0, 1]. Set the sampling rate of VPC
	   flow logs within the subnetwork where 1.0 means all collected logs are
	   reported and 0.0 means no logs are reported. Default is 0.5 which means
	   half of all collected logs are reported.
	*/
	FlowSampling float64 `json:"flowSampling,omitempty" yaml:"flowSampling,omitempty"`

	/*
	   Can only be specified if VPC flow logging for this subnetwork is enabled.
	   Configures whether metadata fields should be added to the reported VPC
	   flow logs.
	   Default value is `INCLUDE_ALL_METADATA`.
	   Possible values are: `EXCLUDE_ALL_METADATA`, `INCLUDE_ALL_METADATA`, `CUSTOM_METADATA`.
	*/
	Metadata string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	/*
	   List of metadata fields that should be added to reported logs.
	   Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" is set to CUSTOM_METADATA.
	*/
	MetadataFields []string `json:"metadataFields,omitempty" yaml:"metadataFields,omitempty"`
}
