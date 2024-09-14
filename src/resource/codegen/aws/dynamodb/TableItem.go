package dynamodb

type TableItem struct {
	// Hash key to use for lookups and identification of the item
	HashKey string `json:"hashKey,omitempty" yaml:"hashKey,omitempty"`

	// JSON representation of a map of attribute name/value pairs, one for each attribute. Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
	Item string `json:"item,omitempty" yaml:"item,omitempty"`

	// Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
	RangeKey string `json:"rangeKey,omitempty" yaml:"rangeKey,omitempty"`

	// Name of the table to contain the item.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`
}
