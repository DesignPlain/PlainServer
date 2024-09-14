package elb

type Attachment struct {
	// The name of the ELB.
	Elb string `json:"elb,omitempty" yaml:"elb,omitempty"`

	// Instance ID to place in the ELB pool.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`
}
