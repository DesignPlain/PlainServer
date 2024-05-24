package types

type Alb_getTargetGroupStickiness struct {
	//
	CookieDuration int `json:"cookieDuration,omitempty" yaml:"cookieDuration,omitempty"`

	//
	CookieName string `json:"cookieName,omitempty" yaml:"cookieName,omitempty"`

	//
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
