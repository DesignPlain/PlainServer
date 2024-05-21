package types

type Organizations_getFoldersFolder struct {
	// The timestamp of when the folder was created
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// The timestamp of when the folder was requested to be deleted (if applicable)
	DeleteTime string `json:"deleteTime,omitempty" yaml:"deleteTime,omitempty"`

	// The display name of the folder
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Entity tag identifier of the folder
	Etag string `json:"etag,omitempty" yaml:"etag,omitempty"`

	// The id of the folder
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The parent id of the folder
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// The lifecycle state of the folder
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// The timestamp of when the folder was last modified
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
