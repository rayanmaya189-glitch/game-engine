package moderation

import (
	"strings"
)

// ProfanityConfig holds profanity filter configuration
type ProfanityConfig struct {
	Enabled         bool   `yaml:"enabled"`
	ReplacementChar string `yaml:"replacement_char"`
	FilterLevel     string `yaml:"filter_level"`
}

// ModerationConfig holds moderation configuration
type ModerationConfig struct {
	AutoMuteThreshold   int  `yaml:"auto_mute_threshold"`
	MuteDurationMinutes int  `yaml:"mute_duration_minutes"`
	BanDurationHours    int  `yaml:"ban_duration_hours"`
	RequiresModerator   bool `yaml:"requires_moderator"`
}

// FilterConfig holds all moderation-related config
type FilterConfig struct {
	ProfanityFilter ProfanityConfig  `yaml:"profanity_filter"`
	Moderation      ModerationConfig `yaml:"moderation"`
}

// ProfanityFilter filters inappropriate content
type ProfanityFilter struct {
	config          *FilterConfig
	replacementChar string
	badWords        map[string]bool
}

// NewProfanityFilter creates a new profanity filter
func NewProfanityFilter(config *FilterConfig) *ProfanityFilter {
	filter := &ProfanityFilter{
		config:          config,
		replacementChar: config.ProfanityFilter.ReplacementChar,
		badWords:        make(map[string]bool),
	}

	// Initialize with common bad words
	filter.initBadWords()

	return filter
}

// Filter filters inappropriate content from text
func (f *ProfanityFilter) Filter(text string) string {
	if !f.config.ProfanityFilter.Enabled {
		return text
	}

	result := text
	for word := range f.badWords {
		result = strings.ReplaceAll(result, word, f.getReplacement(word))
	}

	return result
}

func (f *ProfanityFilter) initBadWords() {
	// Common bad words list - in production this would be more comprehensive
	badWords := []string{
		"badword1", "badword2", // Placeholder - should be loaded from config/file
	}
	for _, word := range badWords {
		f.badWords[strings.ToLower(word)] = true
	}
}

func (f *ProfanityFilter) getReplacement(word string) string {
	if f.replacementChar == "" {
		f.replacementChar = "*"
	}
	return strings.Repeat(f.replacementChar, len(word))
}

// IsEnabled returns whether the filter is enabled
func (f *ProfanityFilter) IsEnabled() bool {
	return f.config.ProfanityFilter.Enabled
}
