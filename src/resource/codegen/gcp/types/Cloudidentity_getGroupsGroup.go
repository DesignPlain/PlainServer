package types

type Cloudidentity_getGroupsGroup struct {
	// An extended description to help users determine the purpose of a Group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// EntityKey of the Group.  Structure is documented below.
	GroupKeys []Cloudidentity_getGroupsGroupGroupKey `json:"groupKeys,omitempty" yaml:"groupKeys,omitempty"`

	/*
	   The initial configuration options for creating a Group.

	   See the
	   [API reference](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups/create#initialgroupconfig)
	   for possible values. Default value: "EMPTY" Possible values: ["INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY"]
	*/
	InitialGroupConfig string `json:"initialGroupConfig,omitempty" yaml:"initialGroupConfig,omitempty"`

	// The parent resource under which to list all Groups. Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_id} for Google Groups.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// Additional group keys associated with the Group
	AdditionalGroupKeys []Cloudidentity_getGroupsGroupAdditionalGroupKey `json:"additionalGroupKeys,omitempty" yaml:"additionalGroupKeys,omitempty"`

	// The display name of the Group.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The labels that apply to the Group.
	   Contains 'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
	   'system/groups/external': '' if the Group is an external-identity-mapped group.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Resource name of the Group in the format: groups/{group_id}, where `group_id` is the unique ID assigned to the Group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The time when the Group was last updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	// The time when the Group was created.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`
}
