package types

type Appengine_StandardAppVersionHandlerStaticFiles struct {
	// Regular expression that matches the file paths for all files that should be referenced by this handler.
	UploadPathRegex string `json:"uploadPathRegex,omitempty" yaml:"uploadPathRegex,omitempty"`

	/*
	   Whether files should also be uploaded as code data. By default, files declared in static file handlers are uploaded as
	   static data and are only served to end users; they cannot be read by the application. If enabled, uploads are charged
	   against both your code and static data storage resource quotas.
	*/
	ApplicationReadable bool `json:"applicationReadable,omitempty" yaml:"applicationReadable,omitempty"`

	/*
	   Time a static file served by this handler should be cached by web proxies and browsers.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example "3.5s".
	*/
	Expiration string `json:"expiration,omitempty" yaml:"expiration,omitempty"`

	/*
	   HTTP headers to use for all responses from these URLs.
	   An object containing a list of "key:value" value pairs.".
	*/
	HttpHeaders map[string]string `json:"httpHeaders,omitempty" yaml:"httpHeaders,omitempty"`

	/*
	   MIME type used to serve all files served by this handler.
	   Defaults to file-specific MIME types, which are derived from each file's filename extension.
	*/
	MimeType string `json:"mimeType,omitempty" yaml:"mimeType,omitempty"`

	// Path to the static files matched by the URL pattern, from the application root directory. The path can refer to text matched in groupings in the URL pattern.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Whether this handler should match the request if the file referenced by the handler does not exist.
	RequireMatchingFile bool `json:"requireMatchingFile,omitempty" yaml:"requireMatchingFile,omitempty"`
}
