package prettybytes

import "testing"

func TestConvert(t *testing.T) {
	t.Run("Converts bytes to human readable strings", func(t *testing.T) {
		tests := []struct {
			input float64
			want  string
		}{
			{0, "0 B"},
			{0.7, "0.7 B"},
			{10, "10 B"},
			{10.1, "10 B"},
			{999, "999 B"},
			{1001, "1 kB"},
			{1e16, "10 PB"},
			{1e30, "1000000 YB"},
		}

		for _, test := range tests {
			got := Convert(test.input)

			if got != test.want {
				t.Errorf("Expected %v to be equal %v", got, test.want)
			}
		}
	})

	t.Run("Supports negative numbers", func(t *testing.T) {
		tests := []struct {
			input float64
			want  string
		}{
			{-0.7, "-0.7 B"},
			{-10, "-10 B"},
			{-10.1, "-10 B"},
			{-999, "-999 B"},
			{-1001, "-1 kB"},
			{-1e16, "-10 PB"},
			{-1e30, "-1000000 YB"},
		}

		for _, test := range tests {
			got := Convert(test.input)

			if got != test.want {
				t.Errorf("Expected %v to be equal %v", got, test.want)
			}
		}
	})

	t.Run("Has the Signed option", func(t *testing.T) {
		tests := []struct {
			input float64
			want  string
		}{
			{0, "0 B"},
			{-13, "-13 B"},
			{42, "+42 B"},
		}

		for _, test := range tests {
			got := Convert(test.input, Options{
				Signed: true,
			})

			if got != test.want {
				t.Errorf("Expected %v to be equal %v", got, test.want)
			}
		}
	})

	t.Run("Has the Bits option", func(t *testing.T) {
		tests := []struct {
			input float64
			want  string
		}{
			{0, "0 b"},
			{0.4, "0.4 b"},
			{0.7, "0.7 b"},
			{10, "10 b"},
			{10.1, "10 b"},
			{999, "999 b"},
			{1000, "1 kbit"},
			{1001, "1 kbit"},
			{1e16, "10 Pbit"},
			{1e30, "1000000 Ybit"},
		}

		for _, test := range tests {
			got := Convert(test.input, Options{
				Bits: true,
			})

			if got != test.want {
				t.Errorf("Expected %v to be equal %v", got, test.want)
			}
		}
	})
}
