package types

type Quicksight_DashboardDashboardPublishOptions struct {
	// Ad hoc (one-time) filtering option. See ad_hoc_filtering_option.
	AdHocFilteringOption Quicksight_DashboardDashboardPublishOptionsAdHocFilteringOption `json:"adHocFilteringOption,omitempty" yaml:"adHocFilteringOption,omitempty"`

	// The drill-down options of data points in a dashboard. See data_point_drill_up_down_option.
	DataPointDrillUpDownOption Quicksight_DashboardDashboardPublishOptionsDataPointDrillUpDownOption `json:"dataPointDrillUpDownOption,omitempty" yaml:"dataPointDrillUpDownOption,omitempty"`

	// Determines if hidden fields are exported with a dashboard. See export_with_hidden_fields_option.
	ExportWithHiddenFieldsOption Quicksight_DashboardDashboardPublishOptionsExportWithHiddenFieldsOption `json:"exportWithHiddenFieldsOption,omitempty" yaml:"exportWithHiddenFieldsOption,omitempty"`

	// Sheet controls option. See sheet_controls_option.
	SheetControlsOption Quicksight_DashboardDashboardPublishOptionsSheetControlsOption `json:"sheetControlsOption,omitempty" yaml:"sheetControlsOption,omitempty"`

	// The axis sort options of a dashboard. See visual_axis_sort_option.
	VisualAxisSortOption Quicksight_DashboardDashboardPublishOptionsVisualAxisSortOption `json:"visualAxisSortOption,omitempty" yaml:"visualAxisSortOption,omitempty"`

	// The menu options of a visual in a dashboard. See visual_menu_option.
	VisualMenuOption Quicksight_DashboardDashboardPublishOptionsVisualMenuOption `json:"visualMenuOption,omitempty" yaml:"visualMenuOption,omitempty"`

	// The data point menu label options of a dashboard. See data_point_menu_label_option.
	DataPointMenuLabelOption Quicksight_DashboardDashboardPublishOptionsDataPointMenuLabelOption `json:"dataPointMenuLabelOption,omitempty" yaml:"dataPointMenuLabelOption,omitempty"`

	// The data point tool tip options of a dashboard. See data_point_tooltip_option.
	DataPointTooltipOption Quicksight_DashboardDashboardPublishOptionsDataPointTooltipOption `json:"dataPointTooltipOption,omitempty" yaml:"dataPointTooltipOption,omitempty"`

	// Export to .csv option. See export_to_csv_option.
	ExportToCsvOption Quicksight_DashboardDashboardPublishOptionsExportToCsvOption `json:"exportToCsvOption,omitempty" yaml:"exportToCsvOption,omitempty"`

	// The sheet layout maximization options of a dashboard. See sheet_layout_element_maximization_option.
	SheetLayoutElementMaximizationOption Quicksight_DashboardDashboardPublishOptionsSheetLayoutElementMaximizationOption `json:"sheetLayoutElementMaximizationOption,omitempty" yaml:"sheetLayoutElementMaximizationOption,omitempty"`
}
