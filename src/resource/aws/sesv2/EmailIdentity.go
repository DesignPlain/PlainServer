package sesv2

import types "DesignSphere_Server/src/resource/aws/types"

type EmailIdentity struct {
	// The configuration of the DKIM authentication settings for an email domain identity.
	DkimSigningAttributes types.Sesv2_EmailIdentityDkimSigningAttributes `json:"dkimSigningAttributes,omitempty" yaml:"dkimSigningAttributes,omitempty"`

	/*
	   The email address or domain to verify.

	   The following arguments are optional:
	*/
	EmailIdentity string `json:"emailIdentity,omitempty" yaml:"emailIdentity,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	ConfigurationSetName string `json:"configurationSetName,omitempty" yaml:"configurationSetName,omitempty"`
}
