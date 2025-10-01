#
			## Architecture

			The service follows Clean Architecture principles with the following structure:

			```
			services/product-service/

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
			├── infrastructure/
			|       └── docker-compose.yml	
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
 		