package finspace

import types "libds/aws/types"

type KxDataview struct {
	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The option to specify whether you want to apply all the future additions and corrections automatically to the dataview, when you ingest new changesets. The default value is false.
	AutoUpdate bool `json:"autoUpdate,omitempty" yaml:"autoUpdate,omitempty"`

	/*
	   The number of availability zones you want to assign per cluster. This can be one of the following:
	   - `SINGLE` - Assigns one availability zone per cluster.
	   - `MULTI` - Assigns all the availability zones per cluster.
	*/
	AzMode string `json:"azMode,omitempty" yaml:"azMode,omitempty"`

	// A unique identifier of the changeset of the database that you want to use to ingest data.
	ChangesetId string `json:"changesetId,omitempty" yaml:"changesetId,omitempty"`

	// The name of the database where you want to create a dataview.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	/*
	   The option to specify whether you want to make the dataview writable to perform database maintenance. The following are some considerations related to writable dataviews.
	   - You cannot create partial writable dataviews. When you create writeable dataviews you must provide the entire database path. You cannot perform updates on a writeable dataview. Hence, `auto_update` must be set as `false` if `read_write` is `true` for a dataview.
	   - You must also use a unique volume for creating a writeable dataview. So, if you choose a volume that is already in use by another dataview, the dataview creation fails.
	   - Once you create a dataview as writeable, you cannot change it to read-only. So, you cannot update the `read_write` parameter later.
	*/
	ReadWrite bool `json:"readWrite,omitempty" yaml:"readWrite,omitempty"`

	// The configuration that contains the database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume. If you do not explicitly specify any database path for a volume, they are accessible from the cluster through the default S3/object store segment. See segment_configurations below.
	SegmentConfigurations []types.Finspace_KxDataviewSegmentConfiguration `json:"segmentConfigurations,omitempty" yaml:"segmentConfigurations,omitempty"`

	// The identifier of the availability zones. If attaching a volume, the volume must be in the same availability zone as the dataview that you are attaching to.
	AvailabilityZoneId string `json:"availabilityZoneId,omitempty" yaml:"availabilityZoneId,omitempty"`

	// A description for the dataview.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Unique identifier for the KX environment.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	/*
	   A unique identifier for the dataview.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
