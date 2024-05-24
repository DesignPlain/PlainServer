package elb

type Attachment struct {
	// Instance ID to place in the ELB pool.
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// The name of the ELB.
	Elb string `json:"elb,omitempty" yaml:"elb,omitempty"`
}
