package apigee

import types "DesignSphere_Server/src/resource/gcp/types"

type Organization struct {
	/*
	   Flag that specifies whether the VPC Peering through Private Google Access should be
	   disabled between the consumer network and Apigee. Required if an `authorizedNetwork`
	   on the consumer project is not provided, in which case the flag should be set to `true`.
	   Valid only when `RuntimeType` is set to CLOUD. The value must be set before the creation
	   of any Apigee runtime instance and can be updated only when there are no runtime instances.
	*/
	DisableVpcPeering bool `json:"disableVpcPeering,omitempty" yaml:"disableVpcPeering,omitempty"`

	// The display name of the Apigee organization.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Properties defined in the Apigee organization profile.
	   Structure is documented below.
	*/
	Properties types.Apigee_OrganizationProperties `json:"properties,omitempty" yaml:"properties,omitempty"`

	/*
	   Optional. This setting is applicable only for organizations that are soft-deleted (i.e., BillingType
	   is not EVALUATION). It controls how long Organization data will be retained after the initial delete
	   operation completes. During this period, the Organization may be restored to its last known state.
	   After this period, the Organization will no longer be able to be restored.
	   Default value is `DELETION_RETENTION_UNSPECIFIED`.
	   Possible values are: `DELETION_RETENTION_UNSPECIFIED`, `MINIMUM`.
	*/
	Retention string `json:"retention,omitempty" yaml:"retention,omitempty"`

	/*
	   Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances.
	   Update is not allowed after the organization is created.
	   If not specified, a Google-Managed encryption key will be used.
	   Valid only when `RuntimeType` is CLOUD. For example: `projects/foo/locations/us/keyRings/bar/cryptoKeys/baz`.
	*/
	RuntimeDatabaseEncryptionKeyName string `json:"runtimeDatabaseEncryptionKeyName,omitempty" yaml:"runtimeDatabaseEncryptionKeyName,omitempty"`

	// Primary GCP region for analytics data storage. For valid values, see [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
	AnalyticsRegion string `json:"analyticsRegion,omitempty" yaml:"analyticsRegion,omitempty"`

	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType string `json:"billingType,omitempty" yaml:"billingType,omitempty"`

	// Description of the Apigee organization.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Compute Engine network used for Service Networking to be peered with Apigee runtime instances.
	   See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started).
	   Valid only when `RuntimeType` is set to CLOUD. The value can be updated only when there are no runtime instances. For example: "default".
	*/
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty" yaml:"authorizedNetwork,omitempty"`

	/*
	   The project ID associated with the Apigee organization.


	   - - -
	*/
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	/*
	   Runtime type of the Apigee organization based on the Apigee subscription purchased.
	   Default value is `CLOUD`.
	   Possible values are: `CLOUD`, `HYBRID`.
	*/
	RuntimeType string `json:"runtimeType,omitempty" yaml:"runtimeType,omitempty"`
}
