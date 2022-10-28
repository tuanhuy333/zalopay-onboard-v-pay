package authutil_test

import (
	"testing"

	"order-service/util/authutil"
)

// TestBuildMAC compare result with the output of from this website: https://www.freeformatter.com/hmac-generator.html
func TestBuildMAC(t *testing.T) {
	tests := map[string]struct {
		key    string
		params []interface{}

		want string
	}{
		"all params are string": {
			key:    "123",
			params: []interface{}{"abc", "xyz"},
			want:   "d16e6c46817f2bec91b1d73e37666aa0f3ecf3ae9ac586d3ecb01066cdd3e041",
		},

		"mixed string and integer": {
			key:    "foo",
			params: []interface{}{1, "bar"},
			want:   "b943195ae6e6ae245cb14b0f1eceb7dd943c16ad0c504cf095e6f0d3dc156e04",
		},

		"params contain float": {
			key:    "zab123",
			params: []interface{}{12.1, 13.95},
			want:   "e8d414f1a42fb919b1029b218c933592ff7b3ec16e27a861d4ea0a4e108c506d",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := authutil.BuildMAC(test.key, test.params...); got != test.want {
				t.Errorf("MAC not equal:\nwant: %s\ngot : %s", test.want, got)
			}
		})
	}
}

func TestValidMAC(t *testing.T) {
	if !authutil.ValidMAC("123", "d16e6c46817f2bec91b1d73e37666aa0f3ecf3ae9ac586d3ecb01066cdd3e041", "abc", "xyz") {
		t.Errorf("MAC should be valid")
	}

	if authutil.ValidMAC("123", "an-invalid-mac", "abc", "xyz") {
		t.Errorf("MAC should be not valid")
	}
}
