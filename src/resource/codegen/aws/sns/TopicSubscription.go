package sns

type TopicSubscription struct {
	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
	ConfirmationTimeoutInMinutes int `json:"confirmationTimeoutInMinutes,omitempty" yaml:"confirmationTimeoutInMinutes,omitempty"`

	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy string `json:"deliveryPolicy,omitempty" yaml:"deliveryPolicy,omitempty"`

	// Endpoint to send data to. The contents vary with the protocol. See details below.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
	EndpointAutoConfirms bool `json:"endpointAutoConfirms,omitempty" yaml:"endpointAutoConfirms,omitempty"`

	// Whether the `filter_policy` applies to `MessageAttributes` (default) or `MessageBody`.
	FilterPolicyScope string `json:"filterPolicyScope,omitempty" yaml:"filterPolicyScope,omitempty"`

	// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
	SubscriptionRoleArn string `json:"subscriptionRoleArn,omitempty" yaml:"subscriptionRoleArn,omitempty"`

	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy string `json:"filterPolicy,omitempty" yaml:"filterPolicy,omitempty"`

	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
	RawMessageDelivery bool `json:"rawMessageDelivery,omitempty" yaml:"rawMessageDelivery,omitempty"`

	// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
	RedrivePolicy string `json:"redrivePolicy,omitempty" yaml:"redrivePolicy,omitempty"`

	// JSON String with the archived message replay policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-archiving-and-replay-subscriber.html) for more details.
	ReplayPolicy string `json:"replayPolicy,omitempty" yaml:"replayPolicy,omitempty"`

	/*
	   ARN of the SNS topic to subscribe to.

	   The following arguments are optional:
	*/
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
}
