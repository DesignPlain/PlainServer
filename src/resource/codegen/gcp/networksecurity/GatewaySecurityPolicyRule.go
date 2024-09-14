package networksecurity

type GatewaySecurityPolicyRule struct {
	/*
	   Profile which tells what the primitive action should be. Possible values are: - ALLOW - DENY.
	   Possible values are: `BASIC_PROFILE_UNSPECIFIED`, `ALLOW`, `DENY`.
	*/
	BasicProfile string `json:"basicProfile,omitempty" yaml:"basicProfile,omitempty"`

	// Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether the rule is enforced.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// CEL expression for matching on session criteria.
	SessionMatcher string `json:"sessionMatcher,omitempty" yaml:"sessionMatcher,omitempty"`

	/*
	   Flag to enable TLS inspection of traffic matching on. Can only be true if the
	   parent GatewaySecurityPolicy references a TLSInspectionConfig.
	*/
	TlsInspectionEnabled bool `json:"tlsInspectionEnabled,omitempty" yaml:"tlsInspectionEnabled,omitempty"`

	// CEL expression for matching on L7/application level criteria.
	ApplicationMatcher string `json:"applicationMatcher,omitempty" yaml:"applicationMatcher,omitempty"`

	/*
	   The name of the gatewat security policy this rule belongs to.


	   - - -
	*/
	GatewaySecurityPolicy string `json:"gatewaySecurityPolicy,omitempty" yaml:"gatewaySecurityPolicy,omitempty"`

	// The location of the gateway security policy.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule}
	   rule should match the pattern: (^a-z?$).
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Priority of the rule. Lower number corresponds to higher precedence.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
