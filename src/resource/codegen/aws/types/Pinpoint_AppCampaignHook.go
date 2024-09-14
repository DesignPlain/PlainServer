package types

type Pinpoint_AppCampaignHook struct {
	// Lambda function name or ARN to be called for delivery. Conflicts with `web_url`
	LambdaFunctionName string `json:"lambdaFunctionName,omitempty" yaml:"lambdaFunctionName,omitempty"`

	// What mode Lambda should be invoked in. Valid values for this parameter are `DELIVERY`, `FILTER`.
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with `lambda_function_name`
	WebUrl string `json:"webUrl,omitempty" yaml:"webUrl,omitempty"`
}
