package types

type Datasync_TaskOptions struct {
	// Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
	VerifyMode string `json:"verifyMode,omitempty" yaml:"verifyMode,omitempty"`

	// Determines the type of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide. Valid values: `OFF`, `BASIC`, `TRANSFER`. Default: `OFF`.
	LogLevel string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`

	// Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
	PreserveDeletedFiles string `json:"preserveDeletedFiles,omitempty" yaml:"preserveDeletedFiles,omitempty"`

	// Determines whether tasks should be queued before executing the tasks. Valid values: `ENABLED`, `DISABLED`. Default `ENABLED`.
	TaskQueueing string `json:"taskQueueing,omitempty" yaml:"taskQueueing,omitempty"`

	// User identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
	Uid string `json:"uid,omitempty" yaml:"uid,omitempty"`

	// A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
	Atime string `json:"atime,omitempty" yaml:"atime,omitempty"`

	// Specifies whether object tags are maintained when transferring between object storage systems. If you want your DataSync task to ignore object tags, specify the NONE value. Valid values: `PRESERVE`, `NONE`. Default value: `PRESERVE`.
	ObjectTags string `json:"objectTags,omitempty" yaml:"objectTags,omitempty"`

	// Determines whether files at the destination should be overwritten or preserved when copying files. Valid values: `ALWAYS`, `NEVER`. Default: `ALWAYS`.
	OverwriteMode string `json:"overwriteMode,omitempty" yaml:"overwriteMode,omitempty"`

	// Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
	PosixPermissions string `json:"posixPermissions,omitempty" yaml:"posixPermissions,omitempty"`

	// Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
	PreserveDevices string `json:"preserveDevices,omitempty" yaml:"preserveDevices,omitempty"`

	// Determines which components of the SMB security descriptor are copied from source to destination objects. This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. Valid values: `NONE`, `OWNER_DACL`, `OWNER_DACL_SACL`. Default: `OWNER_DACL`.
	SecurityDescriptorCopyFlags string `json:"securityDescriptorCopyFlags,omitempty" yaml:"securityDescriptorCopyFlags,omitempty"`

	// Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
	BytesPerSecond int `json:"bytesPerSecond,omitempty" yaml:"bytesPerSecond,omitempty"`

	// A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
	Mtime string `json:"mtime,omitempty" yaml:"mtime,omitempty"`

	// Group identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
	Gid string `json:"gid,omitempty" yaml:"gid,omitempty"`

	// Determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location. Valid values: `CHANGED`, `ALL`. Default: `CHANGED`
	TransferMode string `json:"transferMode,omitempty" yaml:"transferMode,omitempty"`
}
