EXEC=$(shell basename "$(shell pwd)")
all: $(EXEC)

$(EXEC):
	go build

.PHONY: $(EXEC) run

run:
	gin --immediate rom.z64
