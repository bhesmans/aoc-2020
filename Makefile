DAYS=$(shell ls -d d*)

.PHONY: $(DAYS)

$(DAYS):
	cd $@ && go build . && ./$@
