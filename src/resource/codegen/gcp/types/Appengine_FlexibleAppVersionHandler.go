package types

type Appengine_FlexibleAppVersionHandler struct {
	/*
	   Actions to take when the user is not logged in.
	   Possible values are: `AUTH_FAIL_ACTION_REDIRECT`, `AUTH_FAIL_ACTION_UNAUTHORIZED`.
	*/
	AuthFailAction string `json:"authFailAction,omitempty" yaml:"authFailAction,omitempty"`

	/*
	   Methods to restrict access to a URL based on login status.
	   Possible values are: `LOGIN_OPTIONAL`, `LOGIN_ADMIN`, `LOGIN_REQUIRED`.
	*/
	Login string `json:"login,omitempty" yaml:"login,omitempty"`

	/*
	   30x code to use when performing redirects for the secure field.
	   Possible values are: `REDIRECT_HTTP_RESPONSE_CODE_301`, `REDIRECT_HTTP_RESPONSE_CODE_302`, `REDIRECT_HTTP_RESPONSE_CODE_303`, `REDIRECT_HTTP_RESPONSE_CODE_307`.
	*/
	RedirectHttpResponseCode string `json:"redirectHttpResponseCode,omitempty" yaml:"redirectHttpResponseCode,omitempty"`

	/*
	   Executes a script to handle the requests that match this URL pattern.
	   Only the auto value is supported for Node.js in the App Engine standard environment, for example "script:" "auto".
	   Structure is documented below.
	*/
	Script Appengine_FlexibleAppVersionHandlerScript `json:"script,omitempty" yaml:"script,omitempty"`

	/*
	   Security (HTTPS) enforcement for this URL.
	   Possible values are: `SECURE_DEFAULT`, `SECURE_NEVER`, `SECURE_OPTIONAL`, `SECURE_ALWAYS`.
	*/
	SecurityLevel string `json:"securityLevel,omitempty" yaml:"securityLevel,omitempty"`

	/*
	   Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files.
	   Static file handlers describe which files in the application directory are static files, and which URLs serve them.
	   Structure is documented below.
	*/
	StaticFiles Appengine_FlexibleAppVersionHandlerStaticFiles `json:"staticFiles,omitempty" yaml:"staticFiles,omitempty"`

	/*
	   URL prefix. Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings.
	   All URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.
	*/
	UrlRegex string `json:"urlRegex,omitempty" yaml:"urlRegex,omitempty"`
}
