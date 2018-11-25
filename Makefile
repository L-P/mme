VERSION=$(shell git describe --tags)
BUILDFLAGS=-ldflags '-X main.Version=${VERSION}'
EXEC=$(shell basename "$(shell pwd)")
all: $(EXEC)

$(EXEC):
	cd front && yarn build --mode=production
	mkdir -p "release/$(VERSION)"
	env GOOS=linux GOARCH=amd64 packr build ${BUILDFLAGS} -o "mme_$(VERSION)/mme"
	env GOOS=windows GOARCH=amd64 packr build ${BUILDFLAGS} -o "mme_$(VERSION)/mme.exe"
	env GOOS=darwin GOARCH=amd64 packr build ${BUILDFLAGS} -o "mme_$(VERSION)/mme.osx"
	tar czf "mme_$(VERSION).tgz" "mme_$(VERSION)"
	zip -r "mme_$(VERSION).zip" "mme_$(VERSION)"
	rm -rf "release" "mme_$(VERSION)"

.PHONY: $(EXEC) run clean

clean:
	rm -rf front/dist gin-bin mme release

run:
	gin --appPort 8064 --immediate rom.z64
