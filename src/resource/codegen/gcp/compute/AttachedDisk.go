package compute

type AttachedDisk struct {
	/*
	   Specifies a unique device name of your choice that is
	   reflected into the /dev/disk/by-id/google-- tree of a Linux operating
	   system running within the instance. This name can be used to
	   reference the device for mounting, resizing, and so on, from within
	   the instance.

	   If not specified, the server chooses a default device name to apply
	   to this disk, in the form persistent-disks-x, where x is a number
	   assigned by Google Compute Engine.
	*/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	/*
	   `name` or `self_link` of the disk that will be attached.


	   - - -
	*/
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`

	/*
	   `name` or `self_link` of the compute instance that the disk will be attached to.
	   If the `self_link` is provided then `zone` and `project` are extracted from the
	   self link. If only the name is used then `zone` and `project` must be defined
	   as properties on the resource or provider.
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   The mode in which to attach this disk, either READ_WRITE or
	   READ_ONLY. If not specified, the default is to attach the disk in
	   READ_WRITE mode.

	   Possible values:
	   "READ_ONLY"
	   "READ_WRITE"
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   The project that the referenced compute instance is a part of. If `instance` is referenced by its
	   `self_link` the project defined in the link will take precedence.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The zone that the referenced compute instance is located within. If `instance` is referenced by its
	   `self_link` the zone defined in the link will take precedence.
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
