package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	serviceName := flag.String("name", "", "Name of the service (e.g., user, payment)")
	flag.Parse()

	if *serviceName == "" {
		fmt.Println("Please provide a service name using -name flag")
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: cannot get current directory")
		os.Exit(1)
	}

	// parent of tools/ (repo root)
	rootPath := filepath.Dir(cwd)

	// create service directory structure
	basePath := filepath.Join(rootPath, *serviceName+"-service")
	dirs := []string{
		"cmd",
		"cmd/api",
		"internal",
		"internal/api",
		"internal/api/handlers",
		"internal/api/router",
		"internal/api/middlewares",
		"internal/models",
		"internal/repository",
		"internal/repository/mongodb",
		"internal/repository/sqlconnect",
		"pkg",
		"proto",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(basePath, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			fmt.Printf("Error: Creating the services folder")
			os.Exit(1)
		}
	}
	// create an empty README.md
	readmePath := filepath.Join(basePath, "README.md")
	readmeContent := fmt.Sprintf(`#
			## Architecture

			The service follows Clean Architecture principles with the following structure:

			`+"```"+`
			services/%s-service/

			├── cmd/ (This folder contains the main entry point of your application. Typically one per binary)
			│   └── api/
			│       ├── server.go  // Entry point
			│	  └── .env
			│
			├── internal/ (This folder contains the private application code, including your API handlers, models and repos)
			│   ├── api/ (This folder contains your API handlers and middleware functions)
			│   │   ├── handlers/
			│   │   ├── router/
			│   │   └── middlewares/
			│   ├── models/ (This folder contains your data models or domain entities)
			│   │   ├── user.go
			│   │   └── product.go
			│   └── repository/ (This folder contains the code for accessing the data store, eg. database, file system)
			│         ├── mongodb/
			│         │   └── mongoconnect.go
			│         └── sqlconnect/
			│               └── sqlconfig.go
			│   
			|
			|
			|
			├── pkg/ (This folder holds public libraries or packages that can be used by other projects)
			│   └── utils/
			│       ├── error_handling.go // Utility functions for error handling
			│       └── jwt_processing.go  // JWT processing functions
			│
			├── proto/ (This folder contains the protocol buffer source and generated files)
			│   └── main.proto // Protocol buffer file
			│
			├── go.mod
			└── go.sum
 		`, *serviceName)

	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created %s service structure in %s\n", *serviceName, basePath)
	fmt.Println("\nDirectory structure created:")
	fmt.Printf(`
					├── cmd/ (This folder contains the main entry point of your application. Typically one per binary)
					│   └── api/
					│       ├── server.go  // Entry point
					│	  └── .env
					│
					├── internal/ (This folder contains the private application code, including your API handlers, models and repos)
					│   ├── api/ (This folder contains your API handlers and middleware functions)
					│   │   ├── handlers/
					│   │   ├── router/
					│   │   └── middlewares/
					│   ├── models/ (This folder contains your data models or domain entities)
					│   │   ├── user.go
					│   │   └── product.go
					│   └── repository/ (This folder contains the code for accessing the data store, eg. database, file system)
					│         ├── mongodb/
					│         │   └── mongoconnect.go
					│         └── sqlconnect/
					│               └── sqlconfig.go
					│    
					|   
					|
					├── pkg/ (This folder holds public libraries or packages that can be used by other projects)
					│   └── utils/
					│       ├── error_handling.go // Utility functions for error handling
					│       └── jwt_processing.go  // JWT processing functions
					│
					├── proto/ (This folder contains the protocol buffer source and generated files)
					│   └── main.proto // Protocol buffer file
					│
					├── Dockerfile
					├── go.mod
					└── go.sum
		`)
}
