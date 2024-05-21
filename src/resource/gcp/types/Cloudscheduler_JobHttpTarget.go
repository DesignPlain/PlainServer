package types

type Cloudscheduler_JobHttpTarget struct {
	/*
	   Contains information needed for generating an OAuth token.
	   This type of authorization should be used when sending requests to a GCP endpoint.
	   Structure is documented below.
	*/
	OauthToken Cloudscheduler_JobHttpTargetOauthToken `json:"oauthToken,omitempty" yaml:"oauthToken,omitempty"`

	/*
	   Contains information needed for generating an OpenID Connect token.
	   This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
	   Structure is documented below.
	*/
	OidcToken Cloudscheduler_JobHttpTargetOidcToken `json:"oidcToken,omitempty" yaml:"oidcToken,omitempty"`

	// The full URI path that the request will be sent to.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	/*
	   HTTP request body.
	   A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
	   It is an error to set body on a job with an incompatible HttpMethod.
	   A base64-encoded string.
	*/
	Body string `json:"body,omitempty" yaml:"body,omitempty"`

	/*
	   This map contains the header field names and values.
	   Repeated headers are not supported, but a header value can contain commas.
	*/
	Headers map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Which HTTP method to use for the request.
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`
}
