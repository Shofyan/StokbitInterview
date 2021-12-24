package util

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	type args struct {
		in []string
	}
	dict := make(map[string][]string)
	dict["aikt"] = []string{"kita", "atik", "tika"}
	dict["aku"] = []string{"aku", "kua"}
	dict["aik"] = []string{"kia"}
	dict["aakmn"] = []string{"makan"}

	tests := []struct {
		name     string
		args     args
		wantDict map[string][]string
	}{
		{
			name: "success",
			args: args{
				in: []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"},
			},
			wantDict: dict,
		},
		{
			name: "fail",
			args: args{
				in: []string{},
			},
			wantDict: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDict := Anagram(tt.args.in); !reflect.DeepEqual(gotDict, tt.wantDict) {
				t.Errorf("Anagram() = %v, want %v", gotDict, tt.wantDict)
			}
		})
	}
}
