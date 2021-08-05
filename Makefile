include .env
export

.PHONY: run
run:
	go run main.go \
		-timeout=30 \
		-token=$(TG_TOKEN) \
		-template=$(TEMPLATE_PATH) \
		-path=$(NOTES_PATH)

# not to fail if .env is not present
.env:
	touch $@