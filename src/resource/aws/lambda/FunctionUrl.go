package lambda

import types "DesignSphere_Server/src/resource/aws/types"

type FunctionUrl struct {
	// The type of authentication that the function URL uses. Set to `"AWS_IAM"` to restrict access to authenticated IAM users only. Set to `"NONE"` to bypass IAM authentication and create a public endpoint. See the [AWS documentation](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) for more details.
	AuthorizationType string `json:"authorizationType,omitempty" yaml:"authorizationType,omitempty"`

	// The [cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for the function URL. Documented below.
	Cors types.Lambda_FunctionUrlCors `json:"cors,omitempty" yaml:"cors,omitempty"`

	// The name (or ARN) of the Lambda function.
	FunctionName string `json:"functionName,omitempty" yaml:"functionName,omitempty"`

	// Determines how the Lambda function responds to an invocation. Valid values are `BUFFERED` (default) and `RESPONSE_STREAM`. See more in [Configuring a Lambda function to stream responses](https://docs.aws.amazon.com/lambda/latest/dg/configuration-response-streaming.html).
	InvokeMode string `json:"invokeMode,omitempty" yaml:"invokeMode,omitempty"`

	// The alias name or `"$LATEST"`.
	Qualifier string `json:"qualifier,omitempty" yaml:"qualifier,omitempty"`
}
