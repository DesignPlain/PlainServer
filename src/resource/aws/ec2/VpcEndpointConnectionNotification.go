package ec2

type VpcEndpointConnectionNotification struct {
	/*
	   One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.

	   > --NOTE:-- One of `vpc_endpoint_service_id` or `vpc_endpoint_id` must be specified.
	*/
	ConnectionEvents []string `json:"connectionEvents,omitempty" yaml:"connectionEvents,omitempty"`

	// The ARN of the SNS topic for the notifications.
	ConnectionNotificationArn string `json:"connectionNotificationArn,omitempty" yaml:"connectionNotificationArn,omitempty"`

	// The ID of the VPC Endpoint to receive notifications for.
	VpcEndpointId string `json:"vpcEndpointId,omitempty" yaml:"vpcEndpointId,omitempty"`

	// The ID of the VPC Endpoint Service to receive notifications for.
	VpcEndpointServiceId string `json:"vpcEndpointServiceId,omitempty" yaml:"vpcEndpointServiceId,omitempty"`
}
