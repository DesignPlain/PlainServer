package types

type Netapp_VolumeReplicationTransferStat struct {
	/*
	   (Output)
	   A message describing the cause of the last transfer failure.
	*/
	LastTransferError string `json:"lastTransferError,omitempty" yaml:"lastTransferError,omitempty"`

	/*
	   (Output)
	   Total time taken so far during current transfer.
	*/
	TotalTransferDuration string `json:"totalTransferDuration,omitempty" yaml:"totalTransferDuration,omitempty"`

	/*
	   (Output)
	   Number of bytes transferred so far in current transfer.
	*/
	TransferBytes string `json:"transferBytes,omitempty" yaml:"transferBytes,omitempty"`

	/*
	   (Output)
	   Time when progress was updated last. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	*/
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	/*
	   (Output)
	   The elapsed time since the creation of the snapshot on the source volume that was last replicated
	   to the destination volume. Lag time represents the difference in age of the destination volume
	   data in relation to the source volume data.
	*/
	LagDuration string `json:"lagDuration,omitempty" yaml:"lagDuration,omitempty"`

	/*
	   (Output)
	   Size of last completed transfer in bytes.
	*/
	LastTransferBytes string `json:"lastTransferBytes,omitempty" yaml:"lastTransferBytes,omitempty"`

	/*
	   (Output)
	   Time taken during last completed transfer.
	*/
	LastTransferDuration string `json:"lastTransferDuration,omitempty" yaml:"lastTransferDuration,omitempty"`

	/*
	   (Output)
	   Time when last transfer completed. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	*/
	LastTransferEndTime string `json:"lastTransferEndTime,omitempty" yaml:"lastTransferEndTime,omitempty"`
}
