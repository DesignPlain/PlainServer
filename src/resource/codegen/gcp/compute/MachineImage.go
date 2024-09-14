package compute

import types "libds/gcp/types"

type MachineImage struct {
	// A text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
	   Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).
	*/
	GuestFlush bool `json:"guestFlush,omitempty" yaml:"guestFlush,omitempty"`

	/*
	   Encrypts the machine image using a customer-supplied encryption key.
	   After you encrypt a machine image with a customer-supplied key, you must
	   provide the same key if you use the machine image later (e.g. to create a
	   instance from the image)
	   Structure is documented below.
	*/
	MachineImageEncryptionKey types.Compute_MachineImageMachineImageEncryptionKey `json:"machineImageEncryptionKey,omitempty" yaml:"machineImageEncryptionKey,omitempty"`

	// Name of the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.


	   - - -
	*/
	SourceInstance string `json:"sourceInstance,omitempty" yaml:"sourceInstance,omitempty"`
}
