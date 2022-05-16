package wows

type EncyclopediaCrews struct {
	// Subnation index
	SubnationIndex int `json:"subnation_index,omitempty"`
	// Nation
	Nation string `json:"nation,omitempty"`
	// Commander training level purchased for credits
	MoneyTrainingLevel int `json:"money_training_level,omitempty"`
	// Training cost in credits
	MoneyTrainingHirePrice int `json:"money_training_hire_price,omitempty"`
	// Retraining cost in credits
	MoneyRetrainingPrice int `json:"money_retraining_price,omitempty"`
	// Commanders' last names
	LastNames []string `json:"last_names,omitempty"`
	// Indicates if the skill is preserved after retraining
	IsRetrainable bool `json:"is_retrainable,omitempty"`
	// The list of the Commander images:
	// 
	// 1—URL to the image of the Commander with 1–7 skill points;
	// 8—URL to the image of the Commander with 8–13 skill points;
	// 14—URL to the image of the Commander with 14–20 skill points;.
	// 
	// If only the value for the key 1 is specified, the Commander has not earned skill points yet.
	Icons map[string]string `json:"icons,omitempty"`
	// Commander training level purchased for doubloons
	GoldTrainingLevel int `json:"gold_training_level,omitempty"`
	// Training cost in gold
	GoldTrainingHirePrice int `json:"gold_training_hire_price,omitempty"`
	// Retraining cost in doubloons
	GoldRetrainingPrice int `json:"gold_retraining_price,omitempty"`
	// Commanders' first names
	FirstNames []string `json:"first_names,omitempty"`
	// Basic training level
	BaseTrainingLevel int `json:"base_training_level,omitempty"`
	// Basic training cost
	BaseTrainingHirePrice int `json:"base_training_hire_price,omitempty"`
} 
