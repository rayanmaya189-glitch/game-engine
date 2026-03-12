package moderation

import (
	"strings"

	"github.com/game_engine/chat/internal/room"
)

// ProfanityFilter filters inappropriate content
type ProfanityFilter struct {
	config          *room.Config
	replacementChar string
	badWords        map[string]bool
}

// NewProfanityFilter creates a new profanity filter
func NewProfanityFilter(config *room.Config) *ProfanityFilter {
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

// IsClean checks if text is clean
func (f *ProfanityFilter) IsClean(text string) bool {
	if !f.config.ProfanityFilter.Enabled {
		return true
	}

	lowerText := strings.ToLower(text)
	for word := range f.badWords {
		if strings.Contains(lowerText, word) {
			return false
		}
	}

	return true
}

// AddWord adds a word to the filter
func (f *ProfanityFilter) AddWord(word string) {
	f.badWords[strings.ToLower(word)] = true
}

// RemoveWord removes a word from the filter
func (f *ProfanityFilter) RemoveWord(word string) {
	delete(f.badWords, strings.ToLower(word))
}

// getReplacement returns the replacement string for a word
func (f *ProfanityFilter) getReplacement(word string) string {
	return strings.Repeat(f.replacementChar, len(word))
}

// initBadWords initializes the bad words list
func (f *ProfanityFilter) initBadWords() {
	// Common bad words (this is a minimal example - in production, use a comprehensive list)
	words := []string{
		"badword1",
		"badword2",
		// Add more words as needed
	}

	for _, word := range words {
		f.badWords[word] = true
	}
}
