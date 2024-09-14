package quicksight

import types "libds/aws/types"

type DataSet struct {
	// Configures the combination and transformation of the data from the physical tables. Maximum of 1 entry. See logical_table_map.
	LogicalTableMaps []types.Quicksight_DataSetLogicalTableMap `json:"logicalTableMaps,omitempty" yaml:"logicalTableMaps,omitempty"`

	/*
	   Declares the physical tables that are available in the underlying data sources. See physical_table_map.

	   The following arguments are optional:
	*/
	PhysicalTableMaps []types.Quicksight_DataSetPhysicalTableMap `json:"physicalTableMaps,omitempty" yaml:"physicalTableMaps,omitempty"`

	// The row-level security configuration for the data that you want to create. See row_level_permission_data_set.
	RowLevelPermissionDataSet types.Quicksight_DataSetRowLevelPermissionDataSet `json:"rowLevelPermissionDataSet,omitempty" yaml:"rowLevelPermissionDataSet,omitempty"`

	// The folder that contains fields and nested subfolders for your dataset. See field_folders.
	FieldFolders []types.Quicksight_DataSetFieldFolder `json:"fieldFolders,omitempty" yaml:"fieldFolders,omitempty"`

	// Indicates whether you want to import the data into SPICE. Valid values are `SPICE` and `DIRECT_QUERY`.
	ImportMode string `json:"importMode,omitempty" yaml:"importMode,omitempty"`

	// Display name for the dataset.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// Identifier for the data set.
	DataSetId string `json:"dataSetId,omitempty" yaml:"dataSetId,omitempty"`

	// The refresh properties for the data set. --NOTE--: Only valid when `import_mode` is set to `SPICE`. See refresh_properties.
	RefreshProperties types.Quicksight_DataSetRefreshProperties `json:"refreshProperties,omitempty" yaml:"refreshProperties,omitempty"`

	// The configuration of tags on a dataset to set row-level security. Row-level security tags are currently supported for anonymous embedding only. See row_level_permission_tag_configuration.
	RowLevelPermissionTagConfiguration types.Quicksight_DataSetRowLevelPermissionTagConfiguration `json:"rowLevelPermissionTagConfiguration,omitempty" yaml:"rowLevelPermissionTagConfiguration,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Groupings of columns that work together in certain Amazon QuickSight features. Currently, only geospatial hierarchy is supported. See column_groups.
	ColumnGroups []types.Quicksight_DataSetColumnGroup `json:"columnGroups,omitempty" yaml:"columnGroups,omitempty"`

	// A set of 1 or more definitions of a [ColumnLevelPermissionRule](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnLevelPermissionRule.html). See column_level_permission_rules.
	ColumnLevelPermissionRules []types.Quicksight_DataSetColumnLevelPermissionRule `json:"columnLevelPermissionRules,omitempty" yaml:"columnLevelPermissionRules,omitempty"`

	// The usage configuration to apply to child datasets that reference this dataset as a source. See data_set_usage_configuration.
	DataSetUsageConfiguration types.Quicksight_DataSetDataSetUsageConfiguration `json:"dataSetUsageConfiguration,omitempty" yaml:"dataSetUsageConfiguration,omitempty"`

	// A set of resource permissions on the data source. Maximum of 64 items. See permissions.
	Permissions []types.Quicksight_DataSetPermission `json:"permissions,omitempty" yaml:"permissions,omitempty"`
}
