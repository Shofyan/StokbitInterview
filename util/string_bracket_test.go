package util

import "testing"

func TestFindFirstStringInBracket(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				str: "shofyan(cari bracket)",
			},
			wantOut: "cari bracket",
		},
		{
			name: "fail",
			args: args{
				str: "",
			},
			wantOut: "",
		},
		{
			name: "bracket not found",
			args: args{
				str: "shofyan",
			},
			wantOut: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FindFirstStringInBracket(tt.args.str); gotOut != tt.wantOut {
				t.Errorf("FindFirstStringInBracket() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
