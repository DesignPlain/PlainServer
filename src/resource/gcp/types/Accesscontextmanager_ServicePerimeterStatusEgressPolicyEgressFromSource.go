package types

type Accesscontextmanager_ServicePerimeterStatusEgressPolicyEgressFromSource struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	AccessLevel string `json:"accessLevel,omitempty" yaml:"accessLevel,omitempty"`
}