EXEC=$(shell basename "$(shell pwd)")
all: $(EXEC)

$(EXEC):
	go build

.PHONY: $(EXEC)
