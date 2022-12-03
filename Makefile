.PHONY: clean_artifacts
clean_artifacts:
	@rm -r temporary_files/*

.PHONY: test
test: ## Run tests with check race and coverage
	@make clean_artifacts
	@go test -failfast -count=1 ./... -json -cover -race | tparse -smallscreen

.PHONY: benchmark
benchmark:
	@go test -bench=.