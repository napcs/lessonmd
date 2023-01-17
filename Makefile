.PHONY: all

VERSION = $(shell go run bin/lessonmd.go -v | cut -c11- )

all: windows_64 mac_silicon mac_intel linux_64 release

windows_64:
	mkdir -p dist/windows_64
	env GOOS=windows GOARCH=amd64 go build -o dist/windows_64/lessonmd.exe bin/lessonmd.go


mac_intel:
	mkdir -p dist/mac_intel
	env GOOS=darwin GOARCH=amd64 go build -o dist/mac_intel/lessonmd bin/lessonmd.go

mac_silicon:
	mkdir -p dist/mac_silicon
	env GOOS=darwin GOARCH=arm64 go build -o dist/mac_silicon/lessonmd bin/lessonmd.go

linux_64:
	mkdir -p dist/linux_64
	env GOOS=linux GOARCH=amd64 go build -o dist/linux_64/lessonmd bin/lessonmd.go

release:
	cd dist/windows_64 && zip lessonmd_${VERSION}_windows64.zip lessonmd.exe && mv lessonmd_${VERSION}_windows64.zip ../
	cd dist/mac_intel && zip lessonmd_${VERSION}_mac_intel.zip lessonmd && mv lessonmd_${VERSION}_mac_intel.zip ../
	cd dist/mac_silicon && zip lessonmd_${VERSION}_mac_silicon.zip lessonmd && mv lessonmd_${VERSION}_mac_silicon.zip ../
	cd dist/linux_64 && zip lessonmd_${VERSION}_linux64.zip lessonmd && mv lessonmd_${VERSION}_linux64.zip ../
	cd dist/linux_64 && tar -czvf lessonmd_${VERSION}_linux64.tar.gz lessonmd && mv lessonmd_${VERSION}_linux64.tar.gz ../

.PHONY: clean
clean:
	rm -r dist/


