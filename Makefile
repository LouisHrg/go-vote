run:
	@ go build && ./go-vote

analysis:
	@ $$GOPATH/bin/revive handler model provider

install-dep:
	@ $$GOPATH/bin/govendor init
	@ $$GOPATH/bin/govendor add +external

install-go-vendor:
	@ go get -u github.com/kardianos/govendor

start:
	@ ./go-vote

reset-db:
	@ rm test.db

deploy-dev:
	@ git push heroku master
