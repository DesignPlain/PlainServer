package types

type Iam_DenyPolicyRuleDenyRule struct {
	/*
	   The permissions that are explicitly denied by this rule. Each permission uses the format `{service-fqdn}/{resource}.{verb}`,
	   where `{service-fqdn}` is the fully qualified domain name for the service. For example, `iam.googleapis.com/roles.list`.
	*/
	DeniedPermissions []string `json:"deniedPermissions,omitempty" yaml:"deniedPermissions,omitempty"`

	// The identities that are prevented from using one or more permissions on Google Cloud resources.
	DeniedPrincipals []string `json:"deniedPrincipals,omitempty" yaml:"deniedPrincipals,omitempty"`

	/*
	   Specifies the permissions that this rule excludes from the set of denied permissions given by deniedPermissions.
	   If a permission appears in deniedPermissions and in exceptionPermissions then it will not be denied.
	   The excluded permissions can be specified using the same syntax as deniedPermissions.
	*/
	ExceptionPermissions []string `json:"exceptionPermissions,omitempty" yaml:"exceptionPermissions,omitempty"`

	/*
	   The identities that are excluded from the deny rule, even if they are listed in the deniedPrincipals.
	   For example, you could add a Google group to the deniedPrincipals, then exclude specific users who belong to that group.
	*/
	ExceptionPrincipals []string `json:"exceptionPrincipals,omitempty" yaml:"exceptionPrincipals,omitempty"`

	/*
	   User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header.
	   Structure is documented below.
	*/
	DenialCondition Iam_DenyPolicyRuleDenyRuleDenialCondition `json:"denialCondition,omitempty" yaml:"denialCondition,omitempty"`
}
