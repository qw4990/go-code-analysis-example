#Purpose
This example uses go/ast and go/types to check 
wrong using of some Structs or Interfaces 
which should be used as global variables but 
are created in local every time when used.

#Description
In "test_code/foo_client", foo_client is defined, which should be used as a global variable.

In "test_code/main.go", we will call foo_client's Do() method in function useGlobal() and useLocal().

useGlobal() uses a global foo_client object to call Do(), where useLocal() creates a new object every time.

The goal of "check.go" is to find the wrong using of foo_client in useLocal().

#Output
<pre>
[Warn]: Use local foo client defined at ./test_code/main.go:25:2.
[Warn]: Use local foo client defined at ./test_code/main.go:26:2.
</pre>