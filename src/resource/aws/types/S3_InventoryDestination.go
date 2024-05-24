package types

type S3_InventoryDestination struct {
	// S3 bucket configuration where inventory results are published (documented below).
	Bucket S3_InventoryDestinationBucket `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
