package types

type Apigateway_StageAccessLogSettings struct {
	// ARN of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-`. Automatically removes trailing `:-` if present.
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	/*
	   Formatting and values recorded in the logs.
	   For more information on configuring the log format rules visit the AWS [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html)
	*/
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}
