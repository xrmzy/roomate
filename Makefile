test:
	$(eval PACKAGE := $(shell powershell -Command "& { Read-Host 'Enter Package' }"))
	go test -coverprofile coverage.out ./$(PACKAGE)/...

read:
	go tool cover -html coverage.out

.PHONY: test read