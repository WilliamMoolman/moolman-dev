.PHONY: dev fmt golint test
PACKAGES=`go list ./...`

local:
	@go run -mod=vendor app/main.go

dev:
	@dev_appserver.py app-local.yaml --host 0.0.0.0 --enable_host_checking false --log_level=error --require_indexes=false --datastore_path=/tmp/hubble-datastore --search_indexes_path=/tmp/hubble-indexes --storage_path=/tmp/hubble-storage

fmt:
	@for pkg in ${PACKAGES}; do \
		go fmt $$pkg; \
	done;

deploy:
	@gcloud app deploy app-prod.yaml --project personal-278509

test:
	@RICHGO_FORCE_COLOR=1 ENV=test richgo test -v -mod=vendor -cover ./...
