package osenv_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-osenv/osenv"
)

func Test_osEnv_Getenv(t *testing.T) {
	t.Parallel()
	o := osenv.New()
	o.Getenv("HOME")
}

func Test_mock_Getenv(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		env  map[string]string
		key  string
		exp  string
	}{
		{
			name: "nil",
			key:  "foo",
		},
		{
			name: "exist",
			key:  "foo",
			env: map[string]string{
				"foo": "bar",
			},
			exp: "bar",
		},
		{
			name: "not exist",
			key:  "foo",
			env: map[string]string{
				"bar": "bar",
			},
			exp: "",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			osEnv := osenv.NewMock(d.env)
			v := osEnv.Getenv(d.key)
			if v != d.exp {
				t.Fatalf("wanted %v, got %v", d.exp, v)
			}
		})
	}
}
