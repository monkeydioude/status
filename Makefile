.PHONY: watch
watch:
	@nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run bin/status/main.go -c config_example/config.json