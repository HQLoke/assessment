# Project details
NAME		:=	bitcoin

# Operating system
OS			:=	windows
ARCH		:=	amd64

# Source codes
MAIN_DIR	:=	./src
MAIN_SRC	:=	main.go

# Testing
TEST_DIR	:=	./test
TEST_SRC	:=	test.go

all:
	@go run $(MAIN_DIR)/$(MAIN_SRC)

test:
	@go run $(TEST_DIR)/$(TEST_SRC)

build:
	@go build $(MAIN_DIR)/$(MAIN_SRC)

.PHONY: test build