APP_NAME = easymod

.PHONY: build run clean

install:
	go install

build:
	go build -o $(APP_NAME) .

run: build
	./$(APP_NAME)

clean:
	go clean
	rm -f $(APP_NAME)


	