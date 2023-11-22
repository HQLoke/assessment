# Project details
NAME		:=	bitcoin

# Operating system
OS			:=	windows
ARCH		:=	amd64

# Source codes
MAIN_DIR	:=	src/
MAIN_SRC	:=	$(addprefix $(MAIN_DIR), \
								 main.go \
								 controller.go \
								 math.go \
								 parse.go)

# Testing
TEST_DIR	:=	./test
TEST_SRC	:=	test.go

all:
	@go run $(MAIN_SRC)

test:
	@go run $(TEST_DIR)/$(TEST_SRC)

build:
	@go build $(MAIN_DIR)/$(MAIN_SRC)

.PHONY: test build