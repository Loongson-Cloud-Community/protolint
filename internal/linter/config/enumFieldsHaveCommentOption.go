package config

// EnumFieldsHaveCommentOption represents the option for the ENUM_FIELDS_HAVE_COMMENT rule.
type EnumFieldsHaveCommentOption struct {
	ShouldFollowGolangStyle bool `yaml:"should_follow_golang_style"`
}
