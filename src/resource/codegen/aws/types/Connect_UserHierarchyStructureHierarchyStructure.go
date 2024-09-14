package types

type Connect_UserHierarchyStructureHierarchyStructure struct {
	// A block that defines the details of level two. The level block is documented below.
	LevelTwo Connect_UserHierarchyStructureHierarchyStructureLevelTwo `json:"levelTwo,omitempty" yaml:"levelTwo,omitempty"`

	/*
	   A block that defines the details of level five. The level block is documented below.

	   Each level block supports the following arguments:
	*/
	LevelFive Connect_UserHierarchyStructureHierarchyStructureLevelFive `json:"levelFive,omitempty" yaml:"levelFive,omitempty"`

	// A block that defines the details of level four. The level block is documented below.
	LevelFour Connect_UserHierarchyStructureHierarchyStructureLevelFour `json:"levelFour,omitempty" yaml:"levelFour,omitempty"`

	// A block that defines the details of level one. The level block is documented below.
	LevelOne Connect_UserHierarchyStructureHierarchyStructureLevelOne `json:"levelOne,omitempty" yaml:"levelOne,omitempty"`

	// A block that defines the details of level three. The level block is documented below.
	LevelThree Connect_UserHierarchyStructureHierarchyStructureLevelThree `json:"levelThree,omitempty" yaml:"levelThree,omitempty"`
}
