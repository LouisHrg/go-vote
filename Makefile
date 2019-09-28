run:
	@ go build && ./go-vote

start: reset-db run
	@ echo "starting go app"

analysis:
	@ $$GOPATH/bin/revive handler model provider

install-dep:
	@ $$GOPATH/bin/govendor init
	@ $$GOPATH/bin/govendor add +external

install-go-vendor:
	@ go get -u github.com/kardianos/govendor

reset-db:
	@ rm test.db

deploy-dev:
	@ git push heroku master
