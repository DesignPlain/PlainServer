package types

type S3control_MultiRegionAccessPointDetails struct {
	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	PublicAccessBlock S3control_MultiRegionAccessPointDetailsPublicAccessBlock `json:"publicAccessBlock,omitempty" yaml:"publicAccessBlock,omitempty"`

	//
	Regions []S3control_MultiRegionAccessPointDetailsRegion `json:"regions,omitempty" yaml:"regions,omitempty"`
}
