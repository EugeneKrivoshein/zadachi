.PHONY: run test

run:
	@echo "Запуск всех задач main.go..."
	@for dir in $(shell find . -maxdepth 1 -type d -regex './[0-9]'); do \
		echo "Запуск $$dir/main.go..."; \
		go run $$dir/main.go || exit 1; \
	done

test:
	@echo "Запуск всех тестов main_test.go..."
	@for dir in $(shell find . -maxdepth 1 -type d -regex './[0-9]'); do \
		echo "Запуск тестов в $$dir..."; \
		go test $$dir || exit 1; \
	done