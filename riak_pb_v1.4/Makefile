VERSION := 1.4.4.0

TARGET = rpb_v$(VERSION).go

PROTO = riak_kv.proto \
		riak.proto \
		riak_search.proto \

GOFILES	= $(PROTO:.proto=.pb.go)

define TEMPLATE
// +build $(shell echo $(VERSION) | cut -d'.' -f1,2)

package riago

import (
endef

export TEMPLATE

all: $(GOFILES)
	@echo "$$TEMPLATE" > $(TARGET)
	@for file in $(GOFILES); do \
		sed -n '/import riak/b; s/import //p' $$file >> $(TARGET); \
		sed -i '/^import/d' $$file; \
		sed -i '/^package/d' $$file; \
	done
	@echo ")" >> $(TARGET)
	@cat $(GOFILES) >> $(TARGET)
	@sed -i 's/riak.RpbPair/RpbPair/g' $(TARGET)
	@echo "formatting file..."
	@go fmt $(TARGET)
	@mv $(TARGET) ../

%.pb.go : %.proto
	protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go --go_out=./ $<

clean:
	rm -f $(GOFILES)
