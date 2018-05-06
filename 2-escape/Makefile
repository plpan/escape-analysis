DEPS=1closurecall

all: clean alldeps

alldeps:
	$(foreach dep,$(DEPS),	go build -gcflags="-N -l -m" -o $(dep) $(dep).go)

clean:
	$(foreach dep,$(DEPS),	-@rm $(dep))

.PHONY: all clean alldeps
