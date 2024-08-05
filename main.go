package main

import (
	"fmt"

	"github.com/nauraizmushtaq/go-cryptographic-algorithms/decrypt"
	"github.com/nauraizmushtaq/go-cryptographic-algorithms/encrypt" // to import locak pkg, use ModuleName/PkgName
)

func main() {
	encryptedStr := encrypt.NimbusEnrypt("ABCD")
	fmt.Println("Encrypted String: ", encryptedStr)
	decryptedStr := decrypt.NimbusDecrypt(encryptedStr)
	fmt.Println("Decrypted String: ", decryptedStr)
}

// go mod  edit -replcae githubModuleName=localModuleName -> Incase if module isn't published yet and need to use locally
// run go mod tidy - After replcing, replace tag will exists in go.mod file
// go mod init -> create a new module and initializes go.mod, make the directory module in which it exists
// go mod tidy -> ensure go.mod matchs with code, add missing module/dependencies
//go run file.go -> to run a go file, it compiles the code(internally), and run a code
// go build -> cmd compilies the package name by import path, with dependencies in a executable file
// go install -> compile and move executable to $GOPATH/bin, so that file can be executed from any path
// go env GOPATH -> shows GOPATH env var value
//go get ->update go.mod and download/downgrade/update a package
//https://kodekloud.com/kk-media/image/upload/v1704815883/course-resource-new/advanced-golang.pdf __READ FROM HERE
