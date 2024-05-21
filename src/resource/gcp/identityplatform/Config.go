package identityplatform

import types "DesignSphere_Server/src/resource/gcp/types"

type Config struct {
	/*
	   Configuration related to multi-tenant functionality.
	   Structure is documented below.
	*/
	MultiTenant types.Identityplatform_ConfigMultiTenant `json:"multiTenant,omitempty" yaml:"multiTenant,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Configuration related to quotas.
	   Structure is documented below.
	*/
	Quota types.Identityplatform_ConfigQuota `json:"quota,omitempty" yaml:"quota,omitempty"`

	/*
	   Configuration related to local sign in methods.
	   Structure is documented below.
	*/
	SignIn types.Identityplatform_ConfigSignIn `json:"signIn,omitempty" yaml:"signIn,omitempty"`

	// List of domains authorized for OAuth redirects.
	AuthorizedDomains []string `json:"authorizedDomains,omitempty" yaml:"authorizedDomains,omitempty"`

	/*
	   Configuration related to blocking functions.
	   Structure is documented below.
	*/
	BlockingFunctions types.Identityplatform_ConfigBlockingFunctions `json:"blockingFunctions,omitempty" yaml:"blockingFunctions,omitempty"`

	/*
	   Options related to how clients making requests on behalf of a project should be configured.
	   Structure is documented below.
	*/
	Client types.Identityplatform_ConfigClient `json:"client,omitempty" yaml:"client,omitempty"`

	/*
	   Configuration related to monitoring project activity.
	   Structure is documented below.
	*/
	Monitoring types.Identityplatform_ConfigMonitoring `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	// Whether anonymous users will be auto-deleted after a period of 30 days
	AutodeleteAnonymousUsers bool `json:"autodeleteAnonymousUsers,omitempty" yaml:"autodeleteAnonymousUsers,omitempty"`

	/*
	   Options related to how clients making requests on behalf of a project should be configured.
	   Structure is documented below.
	*/
	Mfa types.Identityplatform_ConfigMfa `json:"mfa,omitempty" yaml:"mfa,omitempty"`

	/*
	   Configures the regions where users are allowed to send verification SMS for the project or tenant. This is based on the calling code of the destination phone number.
	   Structure is documented below.
	*/
	SmsRegionConfig types.Identityplatform_ConfigSmsRegionConfig `json:"smsRegionConfig,omitempty" yaml:"smsRegionConfig,omitempty"`
}
