package iam

type Group struct {
	// The group's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. Group names are not distinguished by case. For example, you cannot create groups named both "ADMINS" and "admins".
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Path in which to create the group.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
