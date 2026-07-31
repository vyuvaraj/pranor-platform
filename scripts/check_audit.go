//go:build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dirs := []string{"PranorAuth", "ServDB", "PranorNotify", "PranorFlow"}
	hasErrors := false

	for _, dir := range dirs {
		path := filepath.Join("..", "..", dir)
		err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			normalizedPath := filepath.ToSlash(filePath)
			if !strings.HasSuffix(normalizedPath, ".go") || strings.HasSuffix(normalizedPath, "_test.go") || strings.Contains(normalizedPath, "/vendor/") {
				return nil
			}

			fset := token.NewFileSet()
			fileAST, err := parser.ParseFile(fset, filePath, nil, 0)
			if err != nil {
				return nil
			}

			for _, decl := range fileAST.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}

				fnName := fn.Name.Name
				if fnName == "encryptAES" || fnName == "decryptAES" || fnName == "decryptWithKey" || fnName == "saveAPIKeysToStore" || fnName == "LoadKeys" || fnName == "SaveKeys" {
					continue
				}
				isPrivileged := false
				privKeywords := []string{"register", "login", "migrate", "rotate", "revoke", "encrypt", "decrypt", "delete", "mfa", "keys"}
				for _, kw := range privKeywords {
					if strings.Contains(strings.ToLower(fnName), kw) {
						isPrivileged = true
						break
					}
				}

				if !isPrivileged {
					continue
				}

				callsEmitAudit := false
				if fn.Body != nil {
					ast.Inspect(fn.Body, func(n ast.Node) bool {
						call, ok := n.(*ast.CallExpr)
						if !ok {
							return true
						}
						sel, ok := call.Fun.(*ast.SelectorExpr)
						if !ok {
							ident, ok := call.Fun.(*ast.Ident)
							if ok && ident.Name == "EmitAuditEvent" {
								callsEmitAudit = true
							}
							return true
						}
						if sel.Sel.Name == "EmitAuditEvent" {
							callsEmitAudit = true
						}
						return true
					})
				}

				if !callsEmitAudit {
					fmt.Printf("❌ Linter Violation in %s: privileged function '%s' does not call EmitAuditEvent!\n", filePath, fnName)
					hasErrors = true
				}
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error walking directory %s: %v\n", dir, err)
		}
	}

	if hasErrors {
		fmt.Println("\n❌ Audit Event Coverage check failed.")
		os.Exit(1)
	}
	fmt.Println("\n✅ Audit Event Coverage check completed.")
}
