package types

type Wafv2_WebAclAssociationConfigRequestBody struct {
	// Customizes the request body that your protected Amazon App Runner services forward to AWS WAF for inspection. Applicable only when `scope` is set to `REGIONAL`. See `app_runner_service` below for details.
	AppRunnerServices []Wafv2_WebAclAssociationConfigRequestBodyAppRunnerService `json:"appRunnerServices,omitempty" yaml:"appRunnerServices,omitempty"`

	// Customizes the request body that your protected Amazon CloudFront distributions forward to AWS WAF for inspection. Applicable only when `scope` is set to `REGIONAL`. See `cloudfront` below for details.
	Cloudfronts []Wafv2_WebAclAssociationConfigRequestBodyCloudfront `json:"cloudfronts,omitempty" yaml:"cloudfronts,omitempty"`

	// Customizes the request body that your protected Amazon Cognito user pools forward to AWS WAF for inspection. Applicable only when `scope` is set to `REGIONAL`. See `cognito_user_pool` below for details.
	CognitoUserPools []Wafv2_WebAclAssociationConfigRequestBodyCognitoUserPool `json:"cognitoUserPools,omitempty" yaml:"cognitoUserPools,omitempty"`

	// Customizes the request body that your protected AWS Verfied Access instances forward to AWS WAF for inspection. Applicable only when `scope` is set to `REGIONAL`. See `verified_access_instance` below for details.
	VerifiedAccessInstances []Wafv2_WebAclAssociationConfigRequestBodyVerifiedAccessInstance `json:"verifiedAccessInstances,omitempty" yaml:"verifiedAccessInstances,omitempty"`

	// Customizes the request body that your protected Amazon API Gateway REST APIs forward to AWS WAF for inspection. Applicable only when `scope` is set to `CLOUDFRONT`. See `api_gateway` below for details.
	ApiGateways []Wafv2_WebAclAssociationConfigRequestBodyApiGateway `json:"apiGateways,omitempty" yaml:"apiGateways,omitempty"`
}
