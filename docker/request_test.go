package docker

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewRequest(t *testing.T) {
	k1 := "k1"
	k2 := "k2"
	v1 := "v1"
	v2 := "v2"

	tests := []struct {
		name string
		arg  map[string]string
		want Request
	}{
		{
			"simple",
			map[string]string{k1: v1, k2: v2},
			Request{k1: &v1, k2: &v2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRequest(tt.arg)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRequest_Unmarshal(t *testing.T) {
	type typ struct {
		Bool    bool
		Int64   int64
		Float64 float64
		String  string
	}
	tests := []struct {
		name string
		r    Request
		arg  interface{}
		want interface{}
	}{
		{
			"request is nil",
			Request(nil),
			nil,
			nil,
		},
		{
			"some values",
			NewRequest(map[string]string{
				"Bool":    "true",
				"Int64":   "1",
				"Float64": "3.14",
				"String":  "xxx",
			}),
			&typ{},
			&typ{
				Bool:    true,
				Int64:   1,
				Float64: 3.14,
				String:  "xxx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Unmarshal(tt.arg); err != nil {
				t.Errorf("error:%s", err)
				return
			}
			if diff := cmp.Diff(tt.want, tt.arg); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
