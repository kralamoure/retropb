.PHONY: all
all:
	#protoc --proto_path=proto --go_out=. --go_opt=paths=source_relative $$(find proto -type f -iname "*.proto")
	buf generate proto
