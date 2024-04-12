package helplerx

import (
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	// 正面测试用例
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"basic", "hello", "5d41402abc4b2a76b9719d911017c592"},
		{"empty string", "", "d41d8cd98f00b204e9800998ecf8427e"},
		{"numeric", "123", "202cb962ac59075b964b07152d234b70"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EncodeMD5(tt.input)
			if result != tt.expectedOutput {
				t.Errorf("EncodeMD5(%s) = %s, want %s", tt.input, result, tt.expectedOutput)
			}
		})
	}

}
