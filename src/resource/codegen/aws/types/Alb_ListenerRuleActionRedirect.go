package types

type Alb_ListenerRuleActionRedirect struct {
	// The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port string `json:"port,omitempty" yaml:"port,omitempty"`

	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`

	// The HTTP redirect code. The redirect is either permanent (`HTTP_301`) or temporary (`HTTP_302`).
	StatusCode string `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}
