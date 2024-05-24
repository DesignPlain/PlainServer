package rolesanywhere

import types "DesignSphere_Server/src/resource/aws/types"

type TrustAnchor struct {
	// Whether or not the Trust Anchor should be enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the Trust Anchor.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The source of trust, documented below
	Source types.Rolesanywhere_TrustAnchorSource `json:"source,omitempty" yaml:"source,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
