include .env
export

.PHONY: run
run:
	go run main.go \
		-timeout=30 \
		-token=$(TG_TOKEN) \
		-path=$(NOTES_PATH)

# not to fail if .env is not present
.env:
	touch $@