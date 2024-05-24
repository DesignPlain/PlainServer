package types

type Lambda_getFunctionUrlCor struct {
	//
	ExposeHeaders []string `json:"exposeHeaders,omitempty" yaml:"exposeHeaders,omitempty"`

	//
	MaxAge int `json:"maxAge,omitempty" yaml:"maxAge,omitempty"`

	//
	AllowCredentials bool `json:"allowCredentials,omitempty" yaml:"allowCredentials,omitempty"`

	//
	AllowHeaders []string `json:"allowHeaders,omitempty" yaml:"allowHeaders,omitempty"`

	//
	AllowMethods []string `json:"allowMethods,omitempty" yaml:"allowMethods,omitempty"`

	//
	AllowOrigins []string `json:"allowOrigins,omitempty" yaml:"allowOrigins,omitempty"`
}
