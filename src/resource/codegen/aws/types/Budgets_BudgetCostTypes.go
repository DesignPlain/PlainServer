package types

type Budgets_BudgetCostTypes struct {
	// A boolean value whether to include refunds in the cost budget. Defaults to `true`
	IncludeRefund bool `json:"includeRefund,omitempty" yaml:"includeRefund,omitempty"`

	// A boolean value whether to include support costs in the cost budget. Defaults to `true`
	IncludeSupport bool `json:"includeSupport,omitempty" yaml:"includeSupport,omitempty"`

	// A boolean value whether to include tax in the cost budget. Defaults to `true`
	IncludeTax bool `json:"includeTax,omitempty" yaml:"includeTax,omitempty"`

	// A boolean value whether to include upfront costs in the cost budget. Defaults to `true`
	IncludeUpfront bool `json:"includeUpfront,omitempty" yaml:"includeUpfront,omitempty"`

	// A boolean value whether to use blended costs in the cost budget. Defaults to `false`
	UseBlended bool `json:"useBlended,omitempty" yaml:"useBlended,omitempty"`

	// A boolean value whether to include credits in the cost budget. Defaults to `true`
	IncludeCredit bool `json:"includeCredit,omitempty" yaml:"includeCredit,omitempty"`

	// A boolean value whether to include other subscription costs in the cost budget. Defaults to `true`
	IncludeOtherSubscription bool `json:"includeOtherSubscription,omitempty" yaml:"includeOtherSubscription,omitempty"`

	// A boolean value whether to include recurring costs in the cost budget. Defaults to `true`
	IncludeRecurring bool `json:"includeRecurring,omitempty" yaml:"includeRecurring,omitempty"`

	// A boolean value whether to include subscriptions in the cost budget. Defaults to `true`
	IncludeSubscription bool `json:"includeSubscription,omitempty" yaml:"includeSubscription,omitempty"`

	// Whether a budget uses the amortized rate. Defaults to `false`
	UseAmortized bool `json:"useAmortized,omitempty" yaml:"useAmortized,omitempty"`

	// Whether a budget includes discounts. Defaults to `true`
	IncludeDiscount bool `json:"includeDiscount,omitempty" yaml:"includeDiscount,omitempty"`
}
