.PHONY: test tag

tidy:
	@echo "Running tidy..."
	@go mod tidy
	@goimports -w .
	@go vet ./...

test:
	@echo "Running tests..."
	@go mod tidy
	@goimports -w .
	@go vet ./...
	@GOMAXPROCS=1 go test -p=1 ./... -v

tag:
	@echo "Fetching latest remote tag..."
	@latest=$$(git ls-remote --tags origin 'v*' | grep -v '\^{}' | sed 's|.*refs/tags/||' | sort -V | tail -n1); \
	[ -n "$$latest" ] || latest="v0.0.0"; \
	ver=$${latest#v}; \
	major=$$(echo "$$ver" | cut -d. -f1); \
	minor=$$(echo "$$ver" | cut -d. -f2); \
	patch=$$(echo "$$ver" | cut -d. -f3); \
	patch=$$((patch + 1)); \
	if [ $$patch -gt 9 ]; then patch=0; minor=$$((minor + 1)); fi; \
	if [ $$minor -gt 9 ]; then minor=0; major=$$((major + 1)); fi; \
	new="v$$major.$$minor.$$patch"; \
	echo "Latest tag: $$latest -> new tag: $$new"; \
	git tag "$$new" && git push origin "$$new"