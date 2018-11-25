BUILDFLAGS=-ldflags '-X main.Version=$(shell git describe --tags)'
EXEC=$(shell basename "$(shell pwd)")
all: $(EXEC)

$(EXEC):
	cd front && yarn build --mode=production
	packr build ${BUILDFLAGS}

.PHONY: $(EXEC) run clean

clean:
	rm -rf front/dist gin-bin mme

run:
	gin --appPort 8064 --immediate rom.z64
