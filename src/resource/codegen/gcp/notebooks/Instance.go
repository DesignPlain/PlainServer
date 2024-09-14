package notebooks

import types "libds/gcp/types"

type Instance struct {
	/*
	   The service account on this instance, giving access to other
	   Google Cloud services. You can use any service account within
	   the same project, but you must have the service account user
	   permission to use the instance. If not specified,
	   the Compute Engine default service account is used.
	*/
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   Use a Compute Engine VM image to start the notebook instance.
	   Structure is documented below.
	*/
	VmImage types.Notebooks_InstanceVmImage `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`

	/*
	   Possible disk types for notebook instances.
	   Possible values are: `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, `PD_BALANCED`, `PD_EXTREME`.
	*/
	DataDiskType string `json:"dataDiskType,omitempty" yaml:"dataDiskType,omitempty"`

	/*
	   The name of the VPC that this instance is in.
	   Format: projects/{project_id}/global/networks/{network_id}
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The type of vNIC driver.
	   Possible values are: `UNSPECIFIED_NIC_TYPE`, `VIRTIO_NET`, `GVNIC`.
	*/
	NicType string `json:"nicType,omitempty" yaml:"nicType,omitempty"`

	/*
	   Reservation Affinity for consuming Zonal reservation.
	   Structure is documented below.
	*/
	ReservationAffinity types.Notebooks_InstanceReservationAffinity `json:"reservationAffinity,omitempty" yaml:"reservationAffinity,omitempty"`

	// Desired state of the Notebook Instance. Set this field to `ACTIVE` to start the Instance, and `STOPPED` to stop the Instance.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	/*
	   The list of owners of this instance after creation.
	   Format: alias@example.com.
	   Currently supports one owner only.
	   If not specified, all of the service account users of
	   your VM instance's service account can use the instance.
	*/
	InstanceOwners []string `json:"instanceOwners,omitempty" yaml:"instanceOwners,omitempty"`

	/*
	   A reference to the zone where the machine resides.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Whether the end user authorizes Google Cloud to install GPU driver
	   on this instance. If this field is empty or set to false, the GPU driver
	   won't be installed. Only applicable to instances with GPUs.
	*/
	InstallGpuDriver bool `json:"installGpuDriver,omitempty" yaml:"installGpuDriver,omitempty"`

	/*
	   Labels to apply to this instance. These can be later modified by the setLabels method.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Path to a Bash script that automatically runs after a
	   notebook instance fully boots up. The path must be a URL
	   or Cloud Storage path (gs://path-to-file/file-name).
	*/
	PostStartupScript string `json:"postStartupScript,omitempty" yaml:"postStartupScript,omitempty"`

	/*
	   The hardware accelerator used on this instance. If you use accelerators,
	   make sure that your configuration has enough vCPUs and memory to support the
	   machineType you have selected.
	   Structure is documented below.
	*/
	AcceleratorConfig types.Notebooks_InstanceAcceleratorConfig `json:"acceleratorConfig,omitempty" yaml:"acceleratorConfig,omitempty"`

	/*
	   Disk encryption method used on the boot and data disks, defaults to GMEK.
	   Possible values are: `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, `CMEK`.
	*/
	DiskEncryption string `json:"diskEncryption,omitempty" yaml:"diskEncryption,omitempty"`

	/*
	   The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
	   Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
	*/
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	// Instance update time.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	/*
	   Use a container image to start the notebook instance.
	   Structure is documented below.
	*/
	ContainerImage types.Notebooks_InstanceContainerImage `json:"containerImage,omitempty" yaml:"containerImage,omitempty"`

	// The name specified for the Notebook instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A set of Shielded Instance options. Check [Images using supported Shielded VM features]
	   Not all combinations are valid
	   Structure is documented below.
	*/
	ShieldedInstanceConfig types.Notebooks_InstanceShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" yaml:"shieldedInstanceConfig,omitempty"`

	/*
	   The name of the subnet that this instance is in.
	   Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
	*/
	Subnet string `json:"subnet,omitempty" yaml:"subnet,omitempty"`

	// The Compute Engine tags to add to instance.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The size of the boot disk in GB attached to this instance,
	   up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
	   If not specified, this defaults to 100.
	*/
	BootDiskSizeGb int `json:"bootDiskSizeGb,omitempty" yaml:"bootDiskSizeGb,omitempty"`

	/*
	   Possible disk types for notebook instances.
	   Possible values are: `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, `PD_BALANCED`, `PD_EXTREME`.
	*/
	BootDiskType string `json:"bootDiskType,omitempty" yaml:"bootDiskType,omitempty"`

	// The notebook instance will not register with the proxy..
	NoProxyAccess bool `json:"noProxyAccess,omitempty" yaml:"noProxyAccess,omitempty"`

	/*
	   Optional. The URIs of service account scopes to be included in Compute Engine instances.
	   If not specified, the following scopes are defined:
	   - https://www.googleapis.com/auth/cloud-platform
	   - https://www.googleapis.com/auth/userinfo.email
	*/
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty" yaml:"serviceAccountScopes,omitempty"`

	// No public IP will be assigned to this instance.
	NoPublicIp bool `json:"noPublicIp,omitempty" yaml:"noPublicIp,omitempty"`

	// Instance creation time
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	/*
	   Specify a custom Cloud Storage path where the GPU driver is stored.
	   If not specified, we'll automatically choose from official GPU drivers.
	*/
	CustomGpuDriverPath string `json:"customGpuDriverPath,omitempty" yaml:"customGpuDriverPath,omitempty"`

	/*
	   The size of the data disk in GB attached to this instance,
	   up to a maximum of 64000 GB (64 TB).
	   You can choose the size of the data disk based on how big your notebooks and data are.
	   If not specified, this defaults to 100.
	*/
	DataDiskSizeGb int `json:"dataDiskSizeGb,omitempty" yaml:"dataDiskSizeGb,omitempty"`

	/*
	   Custom metadata to apply to this instance.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// A reference to a machine type which defines VM kind.
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// If true, the data disk will not be auto deleted when deleting the instance.
	NoRemoveDataDisk bool `json:"noRemoveDataDisk,omitempty" yaml:"noRemoveDataDisk,omitempty"`
}
