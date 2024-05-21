package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type Image struct {
	/*
	   Name of the resource; provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and
	   match the regular expression `a-z?` which means
	   the first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the
	   last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   URL of the source snapshot used to create this image.
	   In order to create an image, you must provide the full or partial URL of one of the following:
	   - The selfLink URL
	   - This property
	   - The sourceImage URL
	   - The rawDisk.source URL
	   - The sourceDisk URL
	*/
	SourceSnapshot string `json:"sourceSnapshot,omitempty" yaml:"sourceSnapshot,omitempty"`

	/*
	   A list of features to enable on the guest operating system.
	   Applicable only for bootable images.
	   Structure is documented below.
	*/
	GuestOsFeatures []types.Compute_ImageGuestOsFeature `json:"guestOsFeatures,omitempty" yaml:"guestOsFeatures,omitempty"`

	/*
	   The parameters of the raw disk image.
	   Structure is documented below.
	*/
	RawDisk types.Compute_ImageRawDisk `json:"rawDisk,omitempty" yaml:"rawDisk,omitempty"`

	/*
	   URL of the source image used to create this image. In order to create an image, you must provide the full or partial
	   URL of one of the following:
	   - The selfLink URL
	   - This property
	   - The rawDisk.source URL
	   - The sourceDisk URL
	*/
	SourceImage string `json:"sourceImage,omitempty" yaml:"sourceImage,omitempty"`

	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb int `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`

	/*
	   Encrypts the image using a customer-supplied encryption key.
	   After you encrypt an image with a customer-supplied key, you must
	   provide the same key if you use the image later (e.g. to create a
	   disk from the image)
	   Structure is documented below.
	*/
	ImageEncryptionKey types.Compute_ImageImageEncryptionKey `json:"imageEncryptionKey,omitempty" yaml:"imageEncryptionKey,omitempty"`

	/*
	   Labels to apply to this Image.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Any applicable license URI.
	Licenses []string `json:"licenses,omitempty" yaml:"licenses,omitempty"`

	/*
	   The source disk to create this image based on.
	   You must provide either this property or the
	   rawDisk.source property but not both to create an image.
	*/
	SourceDisk string `json:"sourceDisk,omitempty" yaml:"sourceDisk,omitempty"`

	/*
	   Cloud Storage bucket storage location of the image
	   (regional or multi-regional).
	   Reference link: https://cloud.google.com/compute/docs/reference/rest/v1/images
	*/
	StorageLocations []string `json:"storageLocations,omitempty" yaml:"storageLocations,omitempty"`

	/*
	   The name of the image family to which this image belongs. You can
	   create disks by specifying an image family instead of a specific
	   image name. The image family always returns its latest image that is
	   not deprecated. The name of the image family must comply with
	   RFC1035.
	*/
	Family string `json:"family,omitempty" yaml:"family,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
