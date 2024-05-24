package types

type Backup_SelectionCondition struct {
	//
	StringEquals []Backup_SelectionConditionStringEqual `json:"stringEquals,omitempty" yaml:"stringEquals,omitempty"`

	//
	StringLikes []Backup_SelectionConditionStringLike `json:"stringLikes,omitempty" yaml:"stringLikes,omitempty"`

	//
	StringNotEquals []Backup_SelectionConditionStringNotEqual `json:"stringNotEquals,omitempty" yaml:"stringNotEquals,omitempty"`

	//
	StringNotLikes []Backup_SelectionConditionStringNotLike `json:"stringNotLikes,omitempty" yaml:"stringNotLikes,omitempty"`
}
