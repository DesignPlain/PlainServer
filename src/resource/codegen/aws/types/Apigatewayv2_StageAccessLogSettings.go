package types

type Apigatewayv2_StageAccessLogSettings struct {
	// ARN of the CloudWatch Logs log group to receive access logs. Any trailing `:-` is trimmed from the ARN.
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	// Single line [format](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#apigateway-cloudwatch-log-formats) of the access logs of data. Refer to log settings for [HTTP](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-logging-variables.html) or [Websocket](https://docs.aws.amazon.com/apigateway/latest/developerguide/websocket-api-logging.html).
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}
