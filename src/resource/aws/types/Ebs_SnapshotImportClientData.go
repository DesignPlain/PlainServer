package types

type Ebs_SnapshotImportClientData struct {
	// A user-defined comment about the disk upload.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The time that the disk upload ends.
	UploadEnd string `json:"uploadEnd,omitempty" yaml:"uploadEnd,omitempty"`

	// The size of the uploaded disk image, in GiB.
	UploadSize float64 `json:"uploadSize,omitempty" yaml:"uploadSize,omitempty"`

	// The time that the disk upload starts.
	UploadStart string `json:"uploadStart,omitempty" yaml:"uploadStart,omitempty"`
}
