package ec2

type VolumeAttachment struct {
	/*
	   Set to `true` if you want to force the
	   volume to detach. Useful if previous attempts failed, but use this option only
	   as a last resort, as this can result in --data loss--. See
	   [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
	*/
	ForceDetach bool `json:"forceDetach,omitempty" yaml:"forceDetach,omitempty"`

	// ID of the Instance to attach to
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   Set this to true if you do not wish
	   to detach the volume from the instance to which it is attached at destroy
	   time, and instead just remove the attachment from this provider state. This is
	   useful when destroying an instance which has volumes created by some other
	   means attached.
	*/
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	/*
	   Set this to true to ensure that the target instance is stopped
	   before trying to detach the volume. Stops the instance, if it is not already stopped.
	*/
	StopInstanceBeforeDetaching bool `json:"stopInstanceBeforeDetaching,omitempty" yaml:"stopInstanceBeforeDetaching,omitempty"`

	// ID of the Volume to be attached
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`

	/*
	   The device name to expose to the instance (for
	   example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
	*/
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`
}
