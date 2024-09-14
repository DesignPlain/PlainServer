package ec2

import types "libds/aws/types"

type FlowLog struct {
	// The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
	TrafficType string `json:"trafficType,omitempty" yaml:"trafficType,omitempty"`

	// Describes the destination options for a flow log. More details below.
	DestinationOptions types.Ec2_FlowLogDestinationOptions `json:"destinationOptions,omitempty" yaml:"destinationOptions,omitempty"`

	// Elastic Network Interface ID to attach to
	EniId string `json:"eniId,omitempty" yaml:"eniId,omitempty"`

	// The fields to include in the flow log record. Accepted format example: `"$${interface-id} $${srcaddr} $${dstaddr} $${srcport} $${dstport}"`.
	LogFormat string `json:"logFormat,omitempty" yaml:"logFormat,omitempty"`

	// ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
	DeliverCrossAccountRole string `json:"deliverCrossAccountRole,omitempty" yaml:"deliverCrossAccountRole,omitempty"`

	// The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`, `kinesis-data-firehose`. Default: `cloud-watch-logs`.
	LogDestinationType string `json:"logDestinationType,omitempty" yaml:"logDestinationType,omitempty"`

	// Transit Gateway ID to attach to
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Transit Gateway Attachment ID to attach to
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	// VPC ID to attach to
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// The ARN of the logging destination. Either `log_destination` or `log_group_name` must be set.
	LogDestination string `json:"logDestination,omitempty" yaml:"logDestination,omitempty"`

	// --Deprecated:-- Use `log_destination` instead. The name of the CloudWatch log group. Either `log_group_name` or `log_destination` must be set.
	LogGroupName string `json:"logGroupName,omitempty" yaml:"logGroupName,omitempty"`

	/*
	   The maximum interval of time
	   during which a flow of packets is captured and aggregated into a flow
	   log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
	   minutes). Default: `600`. When `transit_gateway_id` or `transit_gateway_attachment_id` is specified, `max_aggregation_interval` -must- be 60 seconds (1 minute).
	*/
	MaxAggregationInterval int `json:"maxAggregationInterval,omitempty" yaml:"maxAggregationInterval,omitempty"`

	// Subnet ID to attach to
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
