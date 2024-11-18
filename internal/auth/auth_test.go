package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name     string
		input    http.Header
		setKey   string
		setValue string
		want     string
		wantErr  bool
	}{
		{
			name:     "Happy Path",
			input:    http.Header{},
			setKey:   "Authorization",
			setValue: "ApiKey TEST_VALUE",
			want:     "TEST_VALUE",
			wantErr:  false,
		}, {
			name:     "No Header",
			input:    http.Header{},
			setKey:   "",
			setValue: "",
			want:     ErrNoAuthHeaderIncluded.Error(),
			wantErr:  true,
		}, {
			name:     "Malformed header: No spaces",
			input:    http.Header{},
			setKey:   "Authorization",
			setValue: "ApiKey;TEST_VALUE;ANOTHER_VALUE",
			want:     "malformed authorization header",
			wantErr:  true,
		}, {
			name:     "Malformed header: no ApiKey",
			input:    http.Header{},
			setKey:   "Authorization",
			setValue: "Bearer TEST_VALUE",
			want:     "malformed authorization header",
			wantErr:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setKey != "" {
				tc.input.Set(tc.setKey, tc.setValue)
			}

			got, err := GetAPIKey(tc.input)

			if err != nil && !tc.wantErr {
				t.Fatalf("%v: An unexpected error occurred: %v", tc.name, err)
				return
			}

			if err != nil && tc.wantErr && err.Error() != tc.want {
				t.Fatalf("%v: Got Err: %v, Expected Err: %v", tc.name, got, tc.want)
				return
			}
			if got != tc.want && !tc.wantErr {
				t.Fatalf("%v: Got: %v, Expected: %v", tc.name, got, tc.want)
				return
			}
		})

	}

}
