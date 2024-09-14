package route53

type DelegationSet struct {
	/*
	   This is a reference name used in Caller Reference
	   (helpful for identifying single delegation set amongst others)
	*/
	ReferenceName string `json:"referenceName,omitempty" yaml:"referenceName,omitempty"`
}
