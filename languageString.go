package languageString

import (
	"errors"
	"fmt"
	"strings"
)

type LanguageString struct {
	language   string
	country    string
	hasCountry bool
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
	return NewLanguageString(input)
}

func (ls *LanguageString) String() string {
	if ls.hasCountry {
		return fmt.Sprintf("%s-%s", ls.language, ls.country)
	}
	return ls.language
}

func (ls *LanguageString) Country() (string, bool) {
	return ls.country, ls.hasCountry
}

func (ls *LanguageString) Language() string {
	return ls.language
}

func (ls *LanguageString) PriorityList() []string {
	if ls.hasCountry {
		return []string{ls.String(), ls.language}
	}
	return []string{ls.language}
}
