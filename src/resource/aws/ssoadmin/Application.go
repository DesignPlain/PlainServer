package ssoadmin

import types "DesignSphere_Server/src/resource/aws/types"

type Application struct {
	/*
	   Name of the application.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Options for the portal associated with an application. See `portal_options` below.
	PortalOptions types.Ssoadmin_ApplicationPortalOptions `json:"portalOptions,omitempty" yaml:"portalOptions,omitempty"`

	// Status of the application. Valid values are `ENABLED` and `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ARN of the application provider.
	ApplicationProviderArn string `json:"applicationProviderArn,omitempty" yaml:"applicationProviderArn,omitempty"`

	// A unique, case-sensitive ID that you provide to ensure the idempotency of the request. AWS generates a random value when not provided.
	ClientToken string `json:"clientToken,omitempty" yaml:"clientToken,omitempty"`

	// Description of the application.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ARN of the instance of IAM Identity Center.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`
}
