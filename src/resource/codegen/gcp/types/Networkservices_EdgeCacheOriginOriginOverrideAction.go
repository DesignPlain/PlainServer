package types

type Networkservices_EdgeCacheOriginOriginOverrideAction struct {
	/*
	   The URL rewrite configuration for request that are
	   handled by this origin.
	   Structure is documented below.
	*/
	UrlRewrite Networkservices_EdgeCacheOriginOriginOverrideActionUrlRewrite `json:"urlRewrite,omitempty" yaml:"urlRewrite,omitempty"`

	/*
	   The header actions, including adding and removing
	   headers, for request handled by this origin.
	   Structure is documented below.
	*/
	HeaderAction Networkservices_EdgeCacheOriginOriginOverrideActionHeaderAction `json:"headerAction,omitempty" yaml:"headerAction,omitempty"`
}
