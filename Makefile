EXEC=$(shell basename "$(shell pwd)")
all: $(EXEC)

$(EXEC):
	cd front && yarn build --mode=production
	packr build

.PHONY: $(EXEC) run clean

clean:
	rm -rf front/dist

run:
	gin --appPort 8064 --immediate rom.z64
