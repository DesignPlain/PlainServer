package types

type Ssm_AssociationTarget struct {
	// A list of instance IDs or tag values. AWS currently limits this list size to one value.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Either `InstanceIds` or `tag:Tag Name` to specify an EC2 tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
