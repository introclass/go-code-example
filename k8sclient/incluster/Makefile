#
# Makefile
# lijiaocn, 2020-03-19 23:37
#

all:
	@echo "Makefile needs your attention"
	GOOS=linux go build
	docker build -t incluster .
	docker tag incluster:latest lijiaocn/incluster:latest
	docker push lijiaocn/incluster:latest

# vim:ft=make
#
