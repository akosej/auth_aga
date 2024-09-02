install: ## Updates the dependencies
	go mod tidy && go mod vendor

build: ## Build the project
	go mod tidy && go mod vendor && CGO_ENABLED=0 go build -trimpath -a -ldflags "-extldflags '-static' -s -w"

run-build: ## Build and run the project
	go mod tidy && go mod vendor && CGO_ENABLED=0 go build -trimpath -a -ldflags "-extldflags '-static' -s -w" && ./aga