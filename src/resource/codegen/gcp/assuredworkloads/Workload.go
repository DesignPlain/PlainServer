package assuredworkloads

import types "libds/gcp/types"

type Workload struct {
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings []types.Assuredworkloads_WorkloadResourceSetting `json:"resourceSettings,omitempty" yaml:"resourceSettings,omitempty"`

	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, HITRUST, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS, ISR_REGIONS, ISR_REGIONS_AND_SUPPORT, CA_PROTECTED_B, IL5, IL2, JP_REGIONS_AND_SUPPORT
	ComplianceRegime string `json:"complianceRegime,omitempty" yaml:"complianceRegime,omitempty"`

	// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// --DEPRECATED-- Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
	KmsSettings types.Assuredworkloads_WorkloadKmsSettings `json:"kmsSettings,omitempty" yaml:"kmsSettings,omitempty"`

	/*
	   Optional. Labels applied to the workload.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Optional. Partner regime associated with this workload. Possible values: PARTNER_UNSPECIFIED, LOCAL_CONTROLS_BY_S3NS, SOVEREIGN_CONTROLS_BY_T_SYSTEMS, SOVEREIGN_CONTROLS_BY_SIA_MINSAIT, SOVEREIGN_CONTROLS_BY_PSN
	Partner string `json:"partner,omitempty" yaml:"partner,omitempty"`

	// Optional. Permissions granted to the AW Partner SA account for the customer workload
	PartnerPermissions types.Assuredworkloads_WorkloadPartnerPermissions `json:"partnerPermissions,omitempty" yaml:"partnerPermissions,omitempty"`

	// Optional. Indicates whether the e-mail notification for a violation is enabled for a workload. This value will be by default True, and if not present will be considered as true. This should only be updated via updateWorkload call. Any Changes to this field during the createWorkload call will not be honored. This will always be true while creating the workload.
	ViolationNotificationsEnabled bool `json:"violationNotificationsEnabled,omitempty" yaml:"violationNotificationsEnabled,omitempty"`

	// Optional. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
	BillingAccount string `json:"billingAccount,omitempty" yaml:"billingAccount,omitempty"`

	// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
	EnableSovereignControls bool `json:"enableSovereignControls,omitempty" yaml:"enableSovereignControls,omitempty"`

	/*
	   The organization for the resource



	   - - -
	*/
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	// Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
	ProvisionedResourcesParent string `json:"provisionedResourcesParent,omitempty" yaml:"provisionedResourcesParent,omitempty"`
}
