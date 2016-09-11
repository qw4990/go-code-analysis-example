#go-code-analysis-example
I want to use packages go/types and go/ast to do some useful code static analysis.

I will note down my ideas here by examples.

#Examples
##global_variable_check
This example uses go/ast and go/types to check 
wrong using of some Structs or Interfaces 
which should be used as global variables but 
are created in local every time when used.