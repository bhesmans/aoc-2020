FORCE:

d%: FORCE
	cd $@ && go build . && ./$@
