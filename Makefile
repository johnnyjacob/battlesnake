# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"

# Define the target executable
TARGET = snake

# Default target: build the executable
all: $(TARGET)

# Rule to build the target executable
$(TARGET): main.go
	$(GO) build $(GOFLAGS) -o $(TARGET)

# Clean target: remove the target executable
clean:
	rm -f $(TARGET)

# Run target: build and run the target executable
run: $(TARGET)
	./$(TARGET)

# Test target: run Go tests for the project
test:
	$(GO) test ./...