package apigateway

import types "libds/aws/types"

type Integration struct {
	// Integration input's [connectionType](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#connectionType). Valid values are `INTERNET` (default for connections through the public routable internet), and `VPC_LINK` (for private connections between API Gateway and a network load balancer in a VPC).
	ConnectionType string `json:"connectionType,omitempty" yaml:"connectionType,omitempty"`

	// Credentials required for the integration. For `AWS` integrations, 2 options are available. To specify an IAM Role for Amazon API Gateway to assume, use the role's ARN. To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::\-:user/\-`.
	Credentials string `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	/*
	   Integration HTTP method
	   (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONs`, `ANY`, `PATCH`) specifying how API Gateway will interact with the back end.
	   --Required-- if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	   Not all methods are compatible with all `AWS` integrations.
	   e.g., Lambda function [can only be invoked](https://github.com/awslabs/aws-apigateway-importer/issues/9#issuecomment-129651005) via `POST`.
	*/
	IntegrationHttpMethod string `json:"integrationHttpMethod,omitempty" yaml:"integrationHttpMethod,omitempty"`

	// Map of the integration's request templates.
	RequestTemplates map[string]string `json:"requestTemplates,omitempty" yaml:"requestTemplates,omitempty"`

	// List of cache key parameters for the integration.
	CacheKeyParameters []string `json:"cacheKeyParameters,omitempty" yaml:"cacheKeyParameters,omitempty"`

	/*
	   HTTP method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTION`, `ANY`)
	   when calling the associated resource.
	*/
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	// TLS configuration. See below.
	TlsConfig types.Apigateway_IntegrationTlsConfig `json:"tlsConfig,omitempty" yaml:"tlsConfig,omitempty"`

	// Integration's cache namespace.
	CacheNamespace string `json:"cacheNamespace,omitempty" yaml:"cacheNamespace,omitempty"`

	// How to handle request payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`. If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the passthroughBehaviors is configured to support payload pass-through.
	ContentHandling string `json:"contentHandling,omitempty" yaml:"contentHandling,omitempty"`

	/*
	   Map of request query string parameters and headers that should be passed to the backend responder.
	   For example: `request_parameters = { "integration.request.header.X-Some-Other-Header" = "method.request.header.X-Some-Header" }`
	*/
	RequestParameters map[string]string `json:"requestParameters,omitempty" yaml:"requestParameters,omitempty"`

	// ID of the associated REST API.
	RestApi string `json:"restApi,omitempty" yaml:"restApi,omitempty"`

	// Integration input's [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/). Valid values are `HTTP` (for HTTP backends), `MOCK` (not calling any real backend), `AWS` (for AWS services), `AWS_PROXY` (for Lambda proxy integration) and `HTTP_PROXY` (for HTTP proxy integration). An `HTTP` or `HTTP_PROXY` integration with a `connection_type` of `VPC_LINK` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   Input's URI. --Required-- if `type` is `AWS`, `AWS_PROXY`, `HTTP` or `HTTP_PROXY`.
	   For HTTP integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification . For AWS integrations, the URI should be of the form `arn:aws:apigateway:{region}:{subdomain.service|service}:{path|action}/{service_api}`. `region`, `subdomain` and `service` are used to determine the right endpoint.
	   e.g., `arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-1:012345678901:function:my-func/invocations`. For private integrations, the URI parameter is not used for routing requests to your endpoint, but is used for setting the Host header and for certificate validation.
	*/
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	// ID of the VpcLink used for the integration. --Required-- if `connection_type` is `VPC_LINK`
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	// Integration passthrough behavior (`WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, `NEVER`).  --Required-- if `request_templates` is used.
	PassthroughBehavior string `json:"passthroughBehavior,omitempty" yaml:"passthroughBehavior,omitempty"`

	// API resource ID.
	ResourceId string `json:"resourceId,omitempty" yaml:"resourceId,omitempty"`

	// Custom timeout between 50 and 300,000 milliseconds. The default value is 29,000 milliseconds. You need to raise a [Service Quota Ticket](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html) to increase time beyond 29,000 milliseconds.
	TimeoutMilliseconds int `json:"timeoutMilliseconds,omitempty" yaml:"timeoutMilliseconds,omitempty"`
}
