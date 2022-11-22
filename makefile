generate:
		cd graph/ && rm -f *.resolvers.go && go generate
run:
		nodemon --exec go run server.go --delay 2.5 --signal SIGTERM

reload:
		rs