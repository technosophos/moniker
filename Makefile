GOPACKAGE="moniker"

.PHONY: build
build:
	GOPACKAGE=$(GOPACKAGE) go run _generator/to_list.go ./animals.txt
	GOPACKAGE=$(GOPACKAGE) go run _generator/to_list.go ./descriptors.txt
