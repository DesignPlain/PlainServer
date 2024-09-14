package types

type Monitoring_UptimeCheckConfigResourceGroup struct {
	// The group of resources being monitored. Should be the `name` of a group
	GroupId string `json:"groupId,omitempty" yaml:"groupId,omitempty"`

	/*
	   The resource type of the group members.
	   Possible values are: `RESOURCE_TYPE_UNSPECIFIED`, `INSTANCE`, `AWS_ELB_LOAD_BALANCER`.
	*/
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}
