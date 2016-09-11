package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"strings"
)

var (
	typesToBeCheck []string
	globalObjects  []types.Object
)

// getTypes parse config/config.go and get all types to be check.
func getTypes() {
	bytes, err := ioutil.ReadFile("./config/config")
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			typesToBeCheck = append(typesToBeCheck, line)
		}
	}
}

func printType(tp types.Type) {
	fmt.Println(tp.String())
}

func toBeCheck(tp types.Type) bool {
	for _, name := range typesToBeCheck {
		if tp.String() == name {
			return true
		}
	}
	return false
}

// getGlobalObjects parse test_code/main.go,
// and get all global objects whose type is contained in typesToBeCheck.
func getGlobalObjects(f *ast.File, info *types.Info) {
	for _, st := range f.Decls {
		decl, ok := st.(*ast.GenDecl)
		if ok && decl.Tok == token.VAR {
			for _, specDecl := range decl.Specs {
				varSpec, ok := specDecl.(*ast.ValueSpec)
				if ok {
					id := varSpec.Names[0]
					obj := info.Defs[id]
					if obj == nil {
						continue
					}
					if toBeCheck(obj.Type()) == true {
						globalObjects = append(globalObjects, obj)
					}
				}
			}
		}
	}
}

func notGlobal(obj types.Object) bool {
	for _, gobj := range globalObjects {
		if obj == gobj {
			return true
		}
	}
	return false
}

func check() {
	// Syntax parse.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./test_code/main.go", nil, 0)
	if err != nil {
		panic(err)
	}

	// Get types which is to be check.
	tconfig := &types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}
	_, err = tconfig.Check("tmp", fset, []*ast.File{f}, info)
	if err != nil {
		panic(err)
	}

	getGlobalObjects(f, info)

	// Do check
	fcc := &fooClientChecker{
		fset: fset,
		info: info,
	}

	ast.Walk(fcc, f)
}

type fooClientChecker struct {
	fset *token.FileSet
	info *types.Info
}

func (fcc *fooClientChecker) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}

	id, ok := node.(*ast.Ident)
	if ok == false {
		return fcc
	}

	var ob types.Object
	ob = fcc.info.Uses[id]
	if ob == nil {
		ob = fcc.info.Defs[id]
	}

	if ob == nil {
		return fcc
	}

	if toBeCheck(ob.Type()) == true {
		if notGlobal(ob) == false {
			fmt.Printf("[Warn]: Use local foo client defined at %v. \n", fcc.fset.Position(node.Pos()))
		}
	}

	return fcc
}

func main() {
	getTypes()
	check()
}
