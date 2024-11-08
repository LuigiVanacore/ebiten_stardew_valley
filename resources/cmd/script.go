package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func generateGoVarName(filePath string) string {
    // Replace non-alphanumeric characters with underscores
    varName := strings.ReplaceAll(filePath, string(os.PathSeparator), "_")
    varName = strings.ReplaceAll(varName, ".", "_")
    return varName
}

func traverseDirectory(rootDir string) ([]string, error) {
    var filePaths []string
    err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            filePaths = append(filePaths, path)
        }
        return nil
    })
    return filePaths, err
}

func writeGoFile(filePaths []string, outputFile string) error {
    file, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString("var (\n")
    if err != nil {
        return err
    }

    for _, filePath := range filePaths {
        varName := generateGoVarName(filePath)
        _, err := file.WriteString(fmt.Sprintf("\t//go:embed %s\n\t%s []byte\n\n", filePath, varName))
        if err != nil {
            return err
        }
    }

    _, err = file.WriteString(")\n")
    return err
}

func main() {
    rootDirectory := "graphics"
    outputFile := "output.go"

    filePaths, err := traverseDirectory(rootDirectory)
    if err != nil {
        fmt.Println("Error traversing directory:", err)
        return
    }

    err = writeGoFile(filePaths, outputFile)
    if err != nil {
        fmt.Println("Error writing Go file:", err)
    }
}