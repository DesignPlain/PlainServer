package types

type Cloudidentity_getGroupMembershipsMembershipMemberKey struct {
	/*
	   The ID of the entity. For Google-managed entities, the id is the email address of an existing
	   group or user. For external-identity-mapped entities, the id is a string conforming
	   to the Identity Source's requirements.
	*/
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   The namespace in which the entity exists.
	   If not populated, the EntityKey represents a Google-managed entity
	   such as a Google user or a Google Group.
	   If populated, the EntityKey represents an external-identity-mapped group.
	*/
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
