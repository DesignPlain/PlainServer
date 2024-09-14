package types

type Firebase_HostingVersionConfigRewrite struct {
	// The function to proxy requests to. Must match the exported function name exactly.
	Function string `json:"function,omitempty" yaml:"function,omitempty"`

	// The user-supplied glob to match against the request URL path.
	Glob string `json:"glob,omitempty" yaml:"glob,omitempty"`

	// The user-supplied RE2 regular expression to match against the request URL path.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	/*
	   The request will be forwarded to Cloud Run.
	   Structure is documented below.
	*/
	Run Firebase_HostingVersionConfigRewriteRun `json:"run,omitempty" yaml:"run,omitempty"`
}
