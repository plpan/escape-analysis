DEPS=1closurecall 2indirectassign

all: clean $(DEPS)

$(DEPS): %: %.go
	go build -gcflags="-N -l -m" -o $@ $<

clean:
	-@rm $(DEPS)

.PHONY: all clean $(DEPS)
