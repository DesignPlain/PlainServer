package networksecurity

import types "DesignSphere_Server/src/resource/gcp/types"

type ServerTlsPolicy struct {
	/*
	   Name of the ServerTlsPolicy resource.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
	   Structure is documented below.
	*/
	ServerCertificate types.Networksecurity_ServerTlsPolicyServerCertificate `json:"serverCertificate,omitempty" yaml:"serverCertificate,omitempty"`

	/*
	   This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies.
	   Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility.
	   Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
	*/
	AllowOpen bool `json:"allowOpen,omitempty" yaml:"allowOpen,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Set of label tags associated with the ServerTlsPolicy resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of the server tls policy.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director.
	   Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If allowOpen and mtlsPolicy are set, server allows both plain text and mTLS connections.
	   Structure is documented below.
	*/
	MtlsPolicy types.Networksecurity_ServerTlsPolicyMtlsPolicy `json:"mtlsPolicy,omitempty" yaml:"mtlsPolicy,omitempty"`
}
