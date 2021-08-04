include .env
export

.PHONY: run
run:
	go run main.go \
		-timeout=180 \
		-token=$(TG_TOKEN)

# not to fail if .env is not present
.env:
	touch $@