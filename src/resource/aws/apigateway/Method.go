package apigateway

type Method struct {
	// Authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId string `json:"authorizerId,omitempty" yaml:"authorizerId,omitempty"`

	// HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	/*
	   Map of the API models used for the request's content type
	   where key is the content type (e.g., `application/json`)
	   and value is either `Error`, `Empty` (built-in models) or `aws.apigateway.Model`'s `name`.
	*/
	RequestModels map[string]string `json:"requestModels,omitempty" yaml:"requestModels,omitempty"`

	// API resource ID
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Specify if the method requires an API key
	ApiKeyRequired bool `json:"apiKeyRequired,omitempty" yaml:"apiKeyRequired,omitempty"`

	// Authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes []string `json:"authorizationScopes,omitempty" yaml:"authorizationScopes,omitempty"`

	// Function name that will be given to the method when generating an SDK through API Gateway. If omitted, API Gateway will generate a function name based on the resource path and HTTP verb.
	OperationName string `json:"operationName,omitempty" yaml:"operationName,omitempty"`

	/*
	   Map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	   For example: `request_parameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	*/
	RequestParameters map[string]bool `json:"requestParameters,omitempty" yaml:"requestParameters,omitempty"`

	// ID of a `aws.apigateway.RequestValidator`
	RequestValidatorId string `json:"requestValidatorId,omitempty" yaml:"requestValidatorId,omitempty"`

	// ID of the associated REST API
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization string `json:"authorization,omitempty" yaml:"authorization,omitempty"`
}
