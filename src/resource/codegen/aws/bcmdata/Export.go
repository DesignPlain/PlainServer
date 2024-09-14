package bcmdata

import types "libds/aws/types"

type Export struct {
	// The details of the export, including data query, name, description, and destination configuration.  See the `export` argument reference below.
	Export types.Bcmdata_ExportExport `json:"export,omitempty" yaml:"export,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Bcmdata_ExportTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
