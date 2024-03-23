package osenv

import "os"

type OSEnv interface {
	Getenv(k string) string
}

type osEnv struct{}

func New() OSEnv {
	return &osEnv{}
}

func (o *osEnv) Getenv(k string) string {
	return os.Getenv(k)
}

type mockOSEnv struct {
	m map[string]string
}

func NewMock(m map[string]string) OSEnv {
	return &mockOSEnv{
		m: m,
	}
}

func (o *mockOSEnv) Getenv(k string) string {
	return o.m[k]
}
