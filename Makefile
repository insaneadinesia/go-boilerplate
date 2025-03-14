.PHONY: build build-alpine clean test help default

GOTEST := go test -v
PACKAGES := $(shell go list ./internal/app/usecase/... ./internal/app/repository/... ./internal/app/wrapper/... ./internal/pkg/... | grep -v /mocks)

test:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go fmt ./... && $(GOTEST) -coverprofile coverage.cov -cover ${PACKAGES}
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.cov

mock-repository:
	mockery --dir=internal/app/repository --output=internal/app/repository/mocks/ --all

mock-wrapper:
	mockery --dir=internal/app/wrapper/location_svc --output=internal/app/wrapper/location_svc/mocks/ --all

mock-usecase:
	mockery --dir=internal/app/usecase/user --output=internal/app/usecase/user/mocks/ --all
	mockery --dir=internal/app/usecase/health_check --output=internal/app/usecase/health_check/mocks/ --all

mock: mock-repository mock-wrapper mock-usecase

swag-install:
	go install github.com/swaggo/swag/cmd/swag@latest
	
swag-init:
	swag init --parseDependency --parseInternal --parseDepth 1 --overridesFile .swaggo