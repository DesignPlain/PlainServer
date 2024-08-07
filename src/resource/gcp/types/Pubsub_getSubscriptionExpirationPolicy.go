package types

type Pubsub_getSubscriptionExpirationPolicy struct {
	/*
	   Specifies the "time-to-live" duration for an associated resource. The
	   resource expires if it is not active for a period of ttl.
	   If ttl is set to "", the associated resource never expires.
	   A duration in seconds with up to nine fractional digits, terminated by 's'.
	   Example - "3.5s".
	*/
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
