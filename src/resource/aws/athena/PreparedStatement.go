package athena

type PreparedStatement struct {
	// Brief explanation of prepared statement. Maximum length of 1024.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the prepared statement. Maximum length of 256.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The query string for the prepared statement.
	QueryStatement string `json:"queryStatement,omitempty" yaml:"queryStatement,omitempty"`

	// The name of the workgroup to which the prepared statement belongs.
	Workgroup string `json:"workgroup,omitempty" yaml:"workgroup,omitempty"`
}
