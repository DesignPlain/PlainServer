package networkmanager

import types "DesignSphere_Server/src/resource/aws/types"

type Site struct {
	// Description of the Site.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the Global Network to create the site in.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`

	// The site location as documented below.
	Location types.Networkmanager_SiteLocation `json:"location,omitempty" yaml:"location,omitempty"`

	// Key-value tags for the Site. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
