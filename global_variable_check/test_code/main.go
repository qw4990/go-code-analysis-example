package main

import "github.com/qw4990/go-code-analysis-example/global_variable_check/test_code/foo_client"

var (
	globalFooClient *foo_client.FooClient
)

func init() {
	globalFooClient = foo_client.New()
}

func main() {
	useGlobal()
	useLocal()
}

// useGlobal uses globalFooClient to do some operations.
func useGlobal() {
	globalFooClient.Do()
}

// useLocal creates a FooClient in local, which is wrong.
func useLocal() {
	localFooClient := foo_client.New()
	localFooClient.Do()
}
