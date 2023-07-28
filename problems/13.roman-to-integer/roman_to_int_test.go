package leeetcode0013

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantResult int
	}{
		{
			name:       "Example 1",
			input:      "III",
			wantResult: 3,
		},
		{
			name:       "Example 2",
			input:      "IV",
			wantResult: 4,
		},
		{
			name:       "Example 3",
			input:      "IX",
			wantResult: 9,
		},
		{
			name:       "Example 4",
			input:      "LVIII",
			wantResult: 58,
		},
		{
			name:       "Example 5",
			input:      "MCMXCIV",
			wantResult: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := romanToInt(tt.input); gotResult != tt.wantResult {
				t.Errorf("romanToInt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
