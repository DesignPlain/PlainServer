package cloudidentity

import types "libds/gcp/types"

type Group struct {
	/*
	   One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.
	   Google Groups are the default type of group and have a label with a key of cloudidentity.googleapis.com/groups.discussion_forum and an empty value.
	   Existing Google Groups can have an additional label with a key of cloudidentity.googleapis.com/groups.security and an empty value added to them. This is an immutable change and the security label cannot be removed once added.
	   Dynamic groups have a label with a key of cloudidentity.googleapis.com/groups.dynamic.
	   Identity-mapped groups for Cloud Search have a label with a key of system/groups/external and an empty value.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The resource name of the entity under which this Group resides in the
	   Cloud Identity resource hierarchy.
	   Must be of the form identitysources/{identity_source_id} for external-identity-mapped
	   groups or customers/{customer_id} for Google Groups.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   An extended description to help users determine the purpose of a Group.
	   Must not be longer than 4,096 characters.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The display name of the Group.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   EntityKey of the Group.
	   Structure is documented below.
	*/
	GroupKey types.Cloudidentity_GroupGroupKey `json:"groupKey,omitempty" yaml:"groupKey,omitempty"`

	/*
	   The initial configuration options for creating a Group.
	   See the
	   [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	   for possible values.
	   Default value is `EMPTY`.
	   Possible values are: `INITIAL_GROUP_CONFIG_UNSPECIFIED`, `WITH_INITIAL_OWNER`, `EMPTY`.
	*/
	InitialGroupConfig string `json:"initialGroupConfig,omitempty" yaml:"initialGroupConfig,omitempty"`
}
