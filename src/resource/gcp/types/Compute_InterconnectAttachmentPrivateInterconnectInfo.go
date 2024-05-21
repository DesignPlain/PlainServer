package types

type Compute_InterconnectAttachmentPrivateInterconnectInfo struct {
	/*
	   (Output)
	   802.1q encapsulation tag to be used for traffic between
	   Google and the customer, going to and from this network and region.
	*/
	Tag8021q int `json:"tag8021q,omitempty" yaml:"tag8021q,omitempty"`
}
