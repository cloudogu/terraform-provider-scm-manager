OS_ARCH=linux_amd64
VERSION=2.0.0
ARTIFACT_ID=terraform-provider-scm-manager${VERSION}_${OS_ARCH}

MAKEFILES_VERSION=4.6.0
HOSTNAME=cloudogu.com
NAMESPACE=tf
NAME=scm

TEST?=$$(go list ./... | grep -v 'vendor')

.DEFAULT_GOAL:=compile
ADDITIONAL_CLEAN=clean-test-cache

include build/make/variables.mk
include build/make/info.mk
include build/make/dependencies-gomod.mk
include build/make/build.mk
include build/make/test-common.mk
include build/make/test-unit.mk
include build/make/test-integration.mk
include build/make/static-analysis.mk
include build/make/clean.mk
include build/make/digital-signature.mk
include build/make/self-update.mk
include build/make/release.mk

install-local: compile
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

.PHONY: package
package: CHANGELOG.md LICENSE README.md $(BINARY)
	tar czf $(BINARY).tar.gz CHANGELOG.md LICENSE README.md $(BINARY)

PRE_INTEGRATIONTESTS=start-local-docker-compose wait-for-scm

.PHONY: clean-test-cache
clean-test-cache:
	@echo clean go testcache
	@go clean -testcache

.PHONY: testacc
testacc:
	@mkdir -p $(TARGET_DIR)/acceptance-tests
	TF_ACC=1 SCM_USERNAME=scmadmin SCM_PASSWORD=scmadmin \
	go test $(TEST) -coverprofile=$(TARGET_DIR)/acceptance-tests/coverage.out -timeout 120m

testacc-local: export SCM_URL = http://localhost:8080/scm

.PHONY: testacc-local
testacc-local: start-local-docker-compose
	SCM_URL=http://localhost:8080/scm
	@make testacc

.PHONY: wait-for-scm
wait-for-scm:
	@echo wait-for-scm by sleeping 10 seconds
	@sleep 10