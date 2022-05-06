.PHONY: all
all: generate

.PHONY: generate
generate:
	buf generate proto
