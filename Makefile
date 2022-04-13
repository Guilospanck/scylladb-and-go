test:
	if ! [ -d "coverage" ]; then \
		echo "Creating coverage folder" ; \
		mkdir coverage; \
	fi
	go test ./... -race -coverprofile=./coverage/coverage.out -covermode=atomic

test-cov:
	if ! [ -d "coverage" ]; then \
		echo "Creating coverage folder" ; \
		mkdir coverage; \
	fi
	go test ./... -race -coverprofile=./coverage/coverage.out -covermode=atomic && go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html