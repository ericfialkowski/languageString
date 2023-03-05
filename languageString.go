package languageString

import (
	"errors"
	"fmt"
	"strings"
)

type LanguageString struct {
	language      string
	country       string
	hasCountry    bool
	AlwaysLower   bool
	UseUnderscore bool
}

func NewLanguageString(language string) (LanguageString, error) {
	return NewLanguageStringWithCountry(language, "")
}

func NewLanguageStringWithCountry(language string, country string) (LanguageString, error) {
	l := strings.TrimSpace(language)
	c := strings.TrimSpace(country)
	if len(l) == 0 {
		return LanguageString{}, errors.New("language cannot be empty")
	}
	return LanguageString{language: l, country: c, hasCountry: len(country) > 0}, nil
}

func Parse(input string) (LanguageString, error) {
	if len(input) == 0 {
		return LanguageString{}, errors.New("input cannot be empty")
	}
	if strings.Contains(input, "-") {
		s := strings.Split(input, "-")
		return NewLanguageStringWithCountry(s[0], s[1])
	}
	if strings.Contains(input, "_") {
		s := strings.Split(input, "_")
		return NewLanguageStringWithCountry(s[0], s[1])
	}
	return NewLanguageString(input)
}

func (ls *LanguageString) String() string {
	if ls.hasCountry {
		if ls.UseUnderscore {
			return fmt.Sprintf("%s_%s", ls.Language(), ls.justCountry())
		}
		return fmt.Sprintf("%s-%s", ls.Language(), ls.justCountry())
	}
	return ls.Language()
}

func (ls *LanguageString) Country() (string, bool) {
	if ls.AlwaysLower {
		return strings.ToLower(ls.country), ls.hasCountry
	}
	return ls.country, ls.hasCountry
}

// justCountry is used internally to just return the country portion
func (ls *LanguageString) justCountry() string {
	c, _ := ls.Country()
	return c
}

func (ls *LanguageString) Language() string {
	if ls.AlwaysLower {
		return strings.ToLower(ls.language)
	}
	return ls.language
}

func (ls *LanguageString) PriorityList() []string {
	if ls.hasCountry {
		return []string{ls.String(), ls.Language()}
	}
	return []string{ls.Language()}
}
