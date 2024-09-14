package types

type Gkehub_MembershipRbacRoleBindingRole struct {
	/*
	   PredefinedRole is an ENUM representation of the default Kubernetes Roles
	   Possible values are: `UNKNOWN`, `ADMIN`, `EDIT`, `VIEW`, `ANTHOS_SUPPORT`.

	   - - -
	*/
	PredefinedRole string `json:"predefinedRole,omitempty" yaml:"predefinedRole,omitempty"`
}
