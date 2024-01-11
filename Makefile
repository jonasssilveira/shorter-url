format:
	goimports -w .

lint:
	golangci-lint run

test:
	go test -v ./...

test-coverage:
	go test -coverprofile coverage.out ./...
	@go tool cover -func=coverage.out | awk '/^total:/{ total += $$3; count++ } END { avg = total / count; if (count > 0) { printf "\033[1;33mAverage coverage: \033[0m"; if (avg > 95) printf "\033[1;32m%.2f%%\033[0m\n", avg; else printf "\033[1;31m%.2f%%\033[0m\n", avg; } else printf "\033[1;31mNo coverage information found\033[0m\n" }'
	go tool cover -html=coverage.out

vet:
	go vet ./...

staticcheck:
	staticcheck ./...

revive:
	revive -config ../revive.toml -formatter friendly ./...

remove-coverage-files:
	rm cover_report.out coverage.out

all-checks:
	-$(MAKE) format
	-$(MAKE) lint
	-$(MAKE) test-coverage
	-$(MAKE) vet
	-$(MAKE) staticcheck
	-$(MAKE) revive
	-$(MAKE) remove-coverage-files