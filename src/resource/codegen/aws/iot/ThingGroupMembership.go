package iot

type ThingGroupMembership struct {
	// Override dynamic thing groups with static thing groups when 10-group limit is reached. If a thing belongs to 10 thing groups, and one or more of those groups are dynamic thing groups, adding a thing to a static group removes the thing from the last dynamic group.
	OverrideDynamicGroup bool `json:"overrideDynamicGroup,omitempty" yaml:"overrideDynamicGroup,omitempty"`

	// The name of the group to which you are adding a thing.
	ThingGroupName string `json:"thingGroupName,omitempty" yaml:"thingGroupName,omitempty"`

	// The name of the thing to add to a group.
	ThingName string `json:"thingName,omitempty" yaml:"thingName,omitempty"`
}
