package types

type Cleanrooms_CollaborationDataEncryptionMetadata struct {
	//
	PreserveNulls bool `json:"preserveNulls,omitempty" yaml:"preserveNulls,omitempty"`

	//
	AllowClearText bool `json:"allowClearText,omitempty" yaml:"allowClearText,omitempty"`

	//
	AllowDuplicates bool `json:"allowDuplicates,omitempty" yaml:"allowDuplicates,omitempty"`

	//
	AllowJoinsOnColumnsWithDifferentNames bool `json:"allowJoinsOnColumnsWithDifferentNames,omitempty" yaml:"allowJoinsOnColumnsWithDifferentNames,omitempty"`
}
