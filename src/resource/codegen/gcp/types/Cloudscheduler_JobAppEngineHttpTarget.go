package types

type Cloudscheduler_JobAppEngineHttpTarget struct {
	/*
	   HTTP request body.
	   A request body is allowed only if the HTTP method is POST or PUT.
	   It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.
	   A base64-encoded string.
	*/
	Body string `json:"body,omitempty" yaml:"body,omitempty"`

	/*
	   HTTP request headers.
	   This map contains the header field names and values.
	   Headers can be set when the job is created.
	*/
	Headers map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Which HTTP method to use for the request.
	HttpMethod string `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty"`

	/*
	   The relative URI.
	   The relative URL must begin with "/" and must be a valid HTTP relative URL.
	   It can contain a path, query string arguments, and \# fragments.
	   If the relative URL is empty, then the root path "/" will be used.
	   No spaces are allowed, and the maximum length allowed is 2083 characters
	*/
	RelativeUri string `json:"relativeUri,omitempty" yaml:"relativeUri,omitempty"`

	/*
	   App Engine Routing setting for the job.
	   Structure is documented below.
	*/
	AppEngineRouting Cloudscheduler_JobAppEngineHttpTargetAppEngineRouting `json:"appEngineRouting,omitempty" yaml:"appEngineRouting,omitempty"`
}
