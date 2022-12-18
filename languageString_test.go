package languageString

import (
	"reflect"
	"testing"
)

func TestLanguageString_Country(t *testing.T) {
	type fields struct {
		language   string
		country    string
		hasCountry bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  bool
	}{
		{"full string", fields{language: "en", country: "US", hasCountry: true}, "US", true},
		{"no country", fields{language: "en", country: "", hasCountry: false}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &LanguageString{
				language:   tt.fields.language,
				country:    tt.fields.country,
				hasCountry: tt.fields.hasCountry,
			}
			got, got1 := ls.Country()
			if got != tt.want {
				t.Errorf("Country() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Country() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLanguageString_Language(t *testing.T) {
	type fields struct {
		language   string
		country    string
		hasCountry bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"full string", fields{language: "en", country: "US", hasCountry: true}, "en"},
		{"no country", fields{language: "en", country: "", hasCountry: false}, "en"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &LanguageString{
				language:   tt.fields.language,
				country:    tt.fields.country,
				hasCountry: tt.fields.hasCountry,
			}
			if got := ls.Language(); got != tt.want {
				t.Errorf("Language() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguageString_String(t *testing.T) {
	type fields struct {
		language   string
		country    string
		hasCountry bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"full string", fields{language: "en", country: "US", hasCountry: true}, "en-US"},
		{"no country", fields{language: "en", country: "", hasCountry: false}, "en"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &LanguageString{
				language:   tt.fields.language,
				country:    tt.fields.country,
				hasCountry: tt.fields.hasCountry,
			}
			if got := ls.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    LanguageString
		wantErr bool
	}{
		{"full string", args{input: "en-US"}, LanguageString{language: "en", country: "US", hasCountry: true}, false},
		{"no country", args{input: "en"}, LanguageString{language: "en", country: "", hasCountry: false}, false},
		{"empty", args{input: ""}, LanguageString{}, true},
		{"empty language", args{input: "-US"}, LanguageString{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguageString_PriorityList(t *testing.T) {
	type fields struct {
		language   string
		country    string
		hasCountry bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"full string", fields{language: "en", country: "US", hasCountry: true}, []string{"en-US", "en"}},
		{"no country", fields{language: "en", country: "", hasCountry: false}, []string{"en"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &LanguageString{
				language:   tt.fields.language,
				country:    tt.fields.country,
				hasCountry: tt.fields.hasCountry,
			}
			if got := ls.PriorityList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriorityList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLanguageString(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name    string
		args    args
		want    LanguageString
		wantErr bool
	}{
		{"normal", args{language: "en"}, LanguageString{language: "en", country: "", hasCountry: false}, false},
		{"spaces", args{language: " en "}, LanguageString{language: "en", country: "", hasCountry: false}, false},
		{"empty string", args{language: ""}, LanguageString{language: "", country: "", hasCountry: false}, true},
		{"bunch of spaces", args{language: "     "}, LanguageString{language: "", country: "", hasCountry: false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLanguageString(tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLanguageString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLanguageString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLanguageStringWithCountry(t *testing.T) {
	type args struct {
		language string
		country  string
	}
	tests := []struct {
		name    string
		args    args
		want    LanguageString
		wantErr bool
	}{
		{"normal", args{language: "en", country: "US"}, LanguageString{language: "en", country: "US", hasCountry: true}, false},
		{"spaces", args{language: " en", country: "US "}, LanguageString{language: "en", country: "US", hasCountry: true}, false},
		{"empty country", args{language: "en", country: ""}, LanguageString{language: "en", country: "", hasCountry: false}, false},
		{"empty language", args{language: "", country: "US"}, LanguageString{language: "", country: "", hasCountry: false}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLanguageStringWithCountry(tt.args.language, tt.args.country)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLanguageStringWithCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLanguageStringWithCountry() got = %v, want %v", got, tt.want)
			}
		})
	}
}