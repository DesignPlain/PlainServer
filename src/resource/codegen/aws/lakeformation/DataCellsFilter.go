package lakeformation

import types "libds/aws/types"

type DataCellsFilter struct {
	//
	Timeouts types.Lakeformation_DataCellsFilterTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Information about the data cells filter. See Table Data below for details.
	TableData types.Lakeformation_DataCellsFilterTableData `json:"tableData,omitempty" yaml:"tableData,omitempty"`
}
