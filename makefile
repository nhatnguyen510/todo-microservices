.PHONY: all proto-gen clean build

all: proto-gen build

proto-gen:
	@echo "Generating Go files from proto definitions..."
	@chmod +x ./scripts/proto-gen.sh
	@./scripts/proto-gen.sh
	@echo "Done."

build:
	@echo "Building the project..."
	@chmod +x ./scripts/build.sh
	@./scripts/build.sh
	@echo "Done."

stop:
	@echo "Stopping the project..."
	@chmod +x ./scripts/stop.sh
	@./scripts/stop.sh
	@echo "Done."

clean:
	rm -rf api/v1/pb/*