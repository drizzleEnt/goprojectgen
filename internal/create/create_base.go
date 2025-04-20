package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func CreateBaseStruct(name string) {
	dirs := []string{
		filepath.Join(name, "cmd"),
		filepath.Join(name, "internal"),
		filepath.Join(name, "pkg"),
		filepath.Join(name, "internal/service"),
		filepath.Join(name, "internal/repository"),
		filepath.Join(name, "internal/app"),
		filepath.Join(name, "internal/domain"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Printf("failed create directory %s: %s\n", dir, err.Error())
			return
		}

		log.Printf("directory %s created\n", dir)
	}

	mainFile := filepath.Join(name, "cmd", "main.go")
	//os.ReadFile("./template/main.go.txt")

	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello from %s")
}		
	`
	if err := os.WriteFile(mainFile, []byte(mainContent), 0644); err != nil {
		log.Printf("failed create main file %s: %s\n", mainFile, err.Error())
		return
	}

	log.Printf("main file created %s\n", mainFile)

	goVer := runtime.Version()
	goVer = strings.TrimPrefix(goVer, "go")

	modFile := filepath.Join(name, "go.mod")
	modContent := fmt.Sprintf("module %s\n\ngo %s\n", name, goVer)

	if err := os.WriteFile(modFile, []byte(modContent), 0644); err != nil {
		log.Printf("failed create gomod file %s: %s\n", modFile, err.Error())
		return
	}
}
