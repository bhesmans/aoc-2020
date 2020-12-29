DAYS=$(shell ls -d -v d*)
LAST_DAY=$(shell ls -d -v d* | tail -n 1)

.PHONY: $(DAYS)

default: $(LAST_DAY)

all: $(DAYS)

$(DAYS):
	cd $@ && go build . && ./$@
