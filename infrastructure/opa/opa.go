// Package opa Open Policy Agent
// Roll-Based Access Control
package opa

import (
	"context"
	"errors"
	"opa-echo-test/internal/chk"
	"sync"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
)

// compilerをcache
var compilerMap map[string]*ast.Compiler = make(map[string]*ast.Compiler) // key: filename, value: compiler

// thread safeにする
var mu *sync.RWMutex = &sync.RWMutex{}

// Setup .
func Setup(fileName string, modules []byte) {

	c, err := ast.CompileModules(map[string]string{
		fileName: string(modules),
	})
	chk.SE(err)

	mu.Lock()
	compilerMap[fileName] = c
	mu.Unlock()
}

// getCompiler .
func getCompiler(fileName string) *ast.Compiler {
	mu.RLock()
	defer mu.RUnlock()
	return compilerMap[fileName]
}

// // Input
// type Input struct {
// 	User              string   `opa:"user" json:"user"`                       // user名
// 	Roles             []string `opa:"roles" json:"roles"`                     // 所持しているrole
// 	AllowResourceList []string `opa:"allow_resources" json:"allow_resources"` // 許可されているresource 正規表現も可

// 	Method         string `opa:"method" json:"method"`                   // 今回アクセスするmethod
// 	Path           string `opa:"path" json:"path"`                       // 今回アクセスするpath
// 	AccessResource string `opa:"access_resource" json:"access_resource"` // 今回アクセスするresource
// }

// Eval query = "data.rbac.authz.allow"
// func Eval(ctx context.Context, fileName string, query string, input interface{}) bool {

// 	r := rego.New(
// 		rego.Query(query),
// 		rego.Compiler(compiler),
// 		rego.Input(input),
// 	)

// 	rs, err := r.Eval(ctx)
// 	chk.SE(err)

// 	return rs.Allowed()
// }

// EvalAllowed .
func EvalAllowed(ctx context.Context, fileName string, query string, input interface{}, s storage.Store) bool {

	compiler := getCompiler(fileName)
	if compiler == nil {
		chk.SE(errors.New("compilerが見つかりませんでした"))
	}

	r := rego.New(
		rego.Query(query),
		rego.Compiler(compiler),
		rego.Store(s),
		rego.Input(input),
	)

	rs, err := r.Eval(ctx)
	chk.SE(err)
	return rs.Allowed()
}
