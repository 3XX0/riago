VERSION := 2.0.0.17

TARGET = rpb_v$(VERSION).go

PROTO = riak_dt.proto \
		riak_kv.proto \
		riak.proto \
		riak_search.proto \
		riak_yokozuna.proto \

GOFILES	= $(PROTO:.proto=.pb.go)

define TEMPLATE
// +build $(VERSION)

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
