package networkmanager

import types "DesignSphere_Server/src/resource/aws/types"

type Link struct {
	// A description of the link.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the global network.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`

	// The provider of the link.
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	// The ID of the site.
	SiteId string `json:"siteId,omitempty" yaml:"siteId,omitempty"`

	// Key-value tags for the link. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of the link.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The upload speed and download speed in Mbps. Documented below.
	Bandwidth types.Networkmanager_LinkBandwidth `json:"bandwidth,omitempty" yaml:"bandwidth,omitempty"`
}
