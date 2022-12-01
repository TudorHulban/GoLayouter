.PHONY: clean_artifacts
clean_artifacts:
	@rm -r temporary_files/*

.PHONY: test
test: ## Run tests with check race and coverage
	@go test -failfast -count=1 ./... -json -cover -race | tparse -smallscreen
	@make clean_artifacts
.PHONY: benchmark
benchmark:
	@go test -bench=.