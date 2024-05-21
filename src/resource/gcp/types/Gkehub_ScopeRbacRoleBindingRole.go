package types

type Gkehub_ScopeRbacRoleBindingRole struct {
	/*
	   PredefinedRole is an ENUM representation of the default Kubernetes Roles
	   Possible values are: `UNKNOWN`, `ADMIN`, `EDIT`, `VIEW`.

	   - - -
	*/
	PredefinedRole string `json:"predefinedRole,omitempty" yaml:"predefinedRole,omitempty"`
}
