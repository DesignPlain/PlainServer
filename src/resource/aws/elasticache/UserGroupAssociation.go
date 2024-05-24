package elasticache

type UserGroupAssociation struct {
	// ID of the user group.
	UserGroupId string `json:"userGroupId,omitempty" yaml:"userGroupId,omitempty"`

	// ID of the user to associated with the user group.
	UserId string `json:"userId,omitempty" yaml:"userId,omitempty"`
}
