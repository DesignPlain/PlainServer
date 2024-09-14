package types

type Firebase_HostingVersionConfigRedirect struct {
	// The user-supplied RE2 regular expression to match against the request URL path.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	// The status HTTP code to return in the response. It must be a valid 3xx status code.
	StatusCode int `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`

	// The user-supplied glob to match against the request URL path.
	Glob string `json:"glob,omitempty" yaml:"glob,omitempty"`

	/*
	   The value to put in the HTTP location header of the response.
	   The location can contain capture group values from the pattern using a : prefix to identify
	   the segment and an optional - to capture the rest of the URL. For example:
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
