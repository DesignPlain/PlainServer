package types

type Glue_CatalogTableOpenTableFormatInput struct {
	// Configuration block for iceberg table config. See `iceberg_input` below.
	IcebergInput Glue_CatalogTableOpenTableFormatInputIcebergInput `json:"icebergInput,omitempty" yaml:"icebergInput,omitempty"`
}
