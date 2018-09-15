OUT_DIR = builds
BIN_NAME = web-dav-server

build:
        govendor build -o $(OUT_DIR)/$(BIN_NAME)
install-dependencies:
        govendor sync
run: install-dependencies build
        chmod +x $(OUT_DIR)/$(BIN_NAME); ./$(OUT_DIR)/$(BIN_NAME)
execute:
         go run -race ./$(OUT_DIR)/$(BIN_NAME)