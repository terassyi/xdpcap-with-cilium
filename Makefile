GOCMD := go
GOBUILD := $(GOCMD) build
GOGENERATE := $(GOCMD) generate
GOCLEAN := $(GOCMD) clean

GO_BINARY := gonapt
GO_SOURCE := main.go
AUTO_GEN := xdpcapprog_*

all: bpf_build build

bpf_build:
	$(GOGENERATE)

build:
	$(GOBUILD) -v .

clean:
	$(GOCLEAN)
	rm $(AUTO_GEN)
