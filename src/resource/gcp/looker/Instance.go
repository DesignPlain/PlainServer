package looker

import types "DesignSphere_Server/src/resource/gcp/types"

type Instance struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Platform editions for a Looker instance. Each edition maps to a set of instance features, like its size. Must be one of these values:
	   - LOOKER_CORE_TRIAL: trial instance
	   - LOOKER_CORE_STANDARD: pay as you go standard instance
	   - LOOKER_CORE_STANDARD_ANNUAL: subscription standard instance
	   - LOOKER_CORE_ENTERPRISE_ANNUAL: subscription enterprise instance
	   - LOOKER_CORE_EMBED_ANNUAL: subscription embed instance
	   Default value is `LOOKER_CORE_TRIAL`.
	   Possible values are: `LOOKER_CORE_TRIAL`, `LOOKER_CORE_STANDARD`, `LOOKER_CORE_STANDARD_ANNUAL`, `LOOKER_CORE_ENTERPRISE_ANNUAL`, `LOOKER_CORE_EMBED_ANNUAL`.
	*/
	PlatformEdition string `json:"platformEdition,omitempty" yaml:"platformEdition,omitempty"`

	/*
	   Name of a reserved IP address range within the consumer network, to be used for
	   private service access connection. User may or may not specify this in a request.
	*/
	ReservedRange string `json:"reservedRange,omitempty" yaml:"reservedRange,omitempty"`

	/*
	   Metadata about users for a Looker instance.
	   These settings are only available when platform edition LOOKER_CORE_STANDARD is set.
	   There are ten Standard and two Developer users included in the cost of the product.
	   You can allocate additional Standard, Viewer, and Developer users for this instance.
	   It is an optional step and can be modified later.
	   With the Standard edition of Looker (Google Cloud core), you can provision up to 50
	   total users, distributed across Viewer, Standard, and Developer.
	   Structure is documented below.
	*/
	UserMetadata types.Looker_InstanceUserMetadata `json:"userMetadata,omitempty" yaml:"userMetadata,omitempty"`

	/*
	   Looker instance Admin settings.
	   Structure is documented below.
	*/
	AdminSettings types.Looker_InstanceAdminSettings `json:"adminSettings,omitempty" yaml:"adminSettings,omitempty"`

	/*
	   Custom domain settings for a Looker instance.
	   Structure is documented below.
	*/
	CustomDomain types.Looker_InstanceCustomDomain `json:"customDomain,omitempty" yaml:"customDomain,omitempty"`

	/*
	   Looker instance encryption settings.
	   Structure is documented below.
	*/
	EncryptionConfig types.Looker_InstanceEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	/*
	   The ID of the instance or a fully qualified identifier for the instance.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Network name in the consumer project in the format of: projects/{project}/global/networks/{network}
	   Note that the consumer network may be in a different GCP project than the consumer
	   project that is hosting the Looker Instance.
	*/
	ConsumerNetwork string `json:"consumerNetwork,omitempty" yaml:"consumerNetwork,omitempty"`

	/*
	   Maintenance denial period for this instance.
	   You must allow at least 14 days of maintenance availability
	   between any two deny maintenance periods.
	   Structure is documented below.
	*/
	DenyMaintenancePeriod types.Looker_InstanceDenyMaintenancePeriod `json:"denyMaintenancePeriod,omitempty" yaml:"denyMaintenancePeriod,omitempty"`

	// The name of the Looker region of the instance.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Maintenance window for an instance.
	   Maintenance of your instance takes place once a month, and will require
	   your instance to be restarted during updates, which will temporarily
	   disrupt service.
	   Structure is documented below.
	*/
	MaintenanceWindow types.Looker_InstanceMaintenanceWindow `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`

	/*
	   Looker Instance OAuth login settings.
	   Structure is documented below.
	*/
	OauthConfig types.Looker_InstanceOauthConfig `json:"oauthConfig,omitempty" yaml:"oauthConfig,omitempty"`

	// Whether private IP is enabled on the Looker instance.
	PrivateIpEnabled bool `json:"privateIpEnabled,omitempty" yaml:"privateIpEnabled,omitempty"`

	// Whether public IP is enabled on the Looker instance.
	PublicIpEnabled bool `json:"publicIpEnabled,omitempty" yaml:"publicIpEnabled,omitempty"`
}
