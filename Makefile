APP = blockchain
MAIN = main.go
GO = go

build:
	$(GO) build -o $(APP)

run:
	$(GO) run $(MAIN)

clean:
	$(GO) clean
