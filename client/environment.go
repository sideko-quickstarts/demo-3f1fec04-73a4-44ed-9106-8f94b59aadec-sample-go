package client

type Env string

const (
	Environment Env = "https://petstore3.swagger.io/api/v3"
	MockServer  Env = "https://api.sideko-staging.dev/v1/mock/demo/sample/0.1.0"
)

// String returns the environment as a string
func (e Env) String() string {
	return string(e)
}
