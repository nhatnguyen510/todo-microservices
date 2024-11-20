PROTO_DIR=api/v1/proto
OUTPUT_DIR=api/v1/pb

# Clean up previous generated files
echo "Cleaning up previous generated files..."
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Check if the proto directory exists, create it if it doesn't
if [ ! -d "$PROTO_DIR" ]; then
    echo "Directory $PROTO_DIR does not exist. Creating it now..."
    mkdir -p "$PROTO_DIR"
fi

echo "Generating Go files from proto definitions..."
protoc -I $PROTO_DIR $PROTO_DIR/*.proto --go_out=$OUTPUT_DIR --go-grpc_out=$OUTPUT_DIR 
echo "Done."
