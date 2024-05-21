package types

type Monitoring_UptimeCheckConfigHttpCheckAuthInfo struct {
	/*
	   The password to authenticate.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The username to authenticate.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
