package types

type Accesscontextmanager_ServicePerimetersServicePerimeterStatusIngressPolicyIngressFrom struct {
	/*
	   A list of identities that are allowed access through this ingress policy.
	   Should be in the format of email address. The email address should represent
	   individual user or service account only.
	*/
	Identities []string `json:"identities,omitempty" yaml:"identities,omitempty"`

	/*
	   Specifies the type of identities that are allowed access from outside the
	   perimeter. If left unspecified, then members of `identities` field will be
	   allowed access.
	   Possible values are: `IDENTITY_TYPE_UNSPECIFIED`, `ANY_IDENTITY`, `ANY_USER_ACCOUNT`, `ANY_SERVICE_ACCOUNT`.
	*/
	IdentityType string `json:"identityType,omitempty" yaml:"identityType,omitempty"`

	/*
	   Sources that this `IngressPolicy` authorizes access from.
	   Structure is documented below.
	*/
	Sources []Accesscontextmanager_ServicePerimetersServicePerimeterStatusIngressPolicyIngressFromSource `json:"sources,omitempty" yaml:"sources,omitempty"`
}
