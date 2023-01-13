.PHONY: all
all: windows_64 mac_silicon mac_intel linux_64

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

.PHONY: clean
clean:
	rm -r dist/


