DEPS=1closurecall 2indirectassign 3container 4functionreceiver 5interface 6structfields

all: clean $(DEPS)

$(DEPS): %: %.go
	go build -gcflags="-N -l -m -m" -o $@ $<

clean:
	-@rm $(DEPS)

.PHONY: all clean $(DEPS)
