package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type RegionSecurityPolicy struct {
	/*
	   Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies.
	   A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits.
	   Rules may then specify matching values for these fields.
	   Structure is documented below.
	*/
	UserDefinedFields []types.Compute_RegionSecurityPolicyUserDefinedField `json:"userDefinedFields,omitempty" yaml:"userDefinedFields,omitempty"`

	/*
	   Configuration for Google Cloud Armor DDOS Proctection Config.
	   Structure is documented below.
	*/
	DdosProtectionConfig types.Compute_RegionSecurityPolicyDdosProtectionConfig `json:"ddosProtectionConfig,omitempty" yaml:"ddosProtectionConfig,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035.
	   Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The Region in which the created Region Security Policy should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The type indicates the intended use of the security policy.
	   - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers.
	   - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
	   - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application.
	   This field can be set only at resource creation time.
	   Possible values are: `CLOUD_ARMOR`, `CLOUD_ARMOR_EDGE`, `CLOUD_ARMOR_NETWORK`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
