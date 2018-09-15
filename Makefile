OUT_DIR = builds
BIN_NAME = web-dav-server

build:
	govendor build -o $(OUT_DIR)/$(BIN_NAME)
dep:
	govendor sync
run: dep build
	chmod +x $(OUT_DIR)/$(BIN_NAME); ./$(OUT_DIR)/$(BIN_NAME)