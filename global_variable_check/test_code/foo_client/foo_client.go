package foo_client

// FooClient is used by other modules;
// FooClient should be used as a global variable.
type FooClient struct{}

// Do is called by other modules.
func (fc *FooClient) Do() {}

// New return a FooClient.
func New() *FooClient {
	return &FooClient{}
}
