package route53

type ResolverQueryLogConfig struct {
	/*
	   The ARN of the resource that you want Route 53 Resolver to send query logs.
	   You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
	*/
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	// The name of the Route 53 Resolver query logging configuration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
