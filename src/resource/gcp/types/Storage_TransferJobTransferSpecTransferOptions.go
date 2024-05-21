package types

type Storage_TransferJobTransferSpecTransferOptions struct {
	// Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and `delete_objects_unique_in_sink` are mutually exclusive.
	DeleteObjectsFromSourceAfterTransfer bool `json:"deleteObjectsFromSourceAfterTransfer,omitempty" yaml:"deleteObjectsFromSourceAfterTransfer,omitempty"`

	/*
	   Whether objects that exist only in the sink should be deleted. Note that this option and
	   `delete_objects_from_source_after_transfer` are mutually exclusive.
	*/
	DeleteObjectsUniqueInSink bool `json:"deleteObjectsUniqueInSink,omitempty" yaml:"deleteObjectsUniqueInSink,omitempty"`

	// Whether overwriting objects that already exist in the sink is allowed.
	OverwriteObjectsAlreadyExistingInSink bool `json:"overwriteObjectsAlreadyExistingInSink,omitempty" yaml:"overwriteObjectsAlreadyExistingInSink,omitempty"`

	// When to overwrite objects that already exist in the sink. If not set, overwrite behavior is determined by `overwrite_objects_already_existing_in_sink`. Possible values: ALWAYS, DIFFERENT, NEVER.
	OverwriteWhen string `json:"overwriteWhen,omitempty" yaml:"overwriteWhen,omitempty"`
}
