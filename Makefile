SCRIPTS_PATH      := scripts
TEST_SCRIPTS_PATH := test/scripts
PROJECT_FLOGO_CLI_VERSION := v0.9.1-rc.2
PROJECT_FLOGO_CORE_VERSION := v0.9.3
PROJECT_FLOGO_FLOW_VERSION := v0.9.3

.PHONY: all
all: install dovetail-tests

.PHONY: depend
depend: 
	$(SCRIPTS_PATH)/dependencies.sh -f

.PHONY: depend-noforce
depend-noforce: 
	@$(SCRIPTS_PATH)/dependencies.sh

.PHONY: install
install:
	go install github.com/TIBCOSoftware/dovetail-cli/cmd/dovetail
	chmod -R 755 ${GOPATH}/pkg/mod/github.com/project-flogo
	cp flogo-patch/cli/util/contrib.go ${GOPATH}/pkg/mod/github.com/project-flogo/cli@$(PROJECT_FLOGO_CLI_VERSION)/util/
	cp flogo-patch/core/support/ref.go ${GOPATH}/pkg/mod/github.com/project-flogo/core@$(PROJECT_FLOGO_CORE_VERSION)/support/
	cp flogo-patch/flow/instance/util2.go ${GOPATH}/pkg/mod/github.com/project-flogo/flow@$(PROJECT_FLOGO_FLOW_VERSION)/instance/
	go install github.com/TIBCOSoftware/dovetail-cli/cmd/dovetail

.PHONY: build
build:
	go build github.com/TIBCOSoftware/dovetail-cli/cmd/dovetail

.PHONY: buildtype
buildtype: install
	@$(SCRIPTS_PATH)/buildtype.sh

.PHONY: test_all
test_all: dovetail-tests hyperledger-fabric-tests corda-tests

.PHONY: dovetail-tests
dovetail-tests:
	@$(TEST_SCRIPTS_PATH)/dovetail.sh

.PHONY: hyperledger-fabric-tests
hyperledger-fabric-tests:
	@$(TEST_SCRIPTS_PATH)/hyperledger-fabric.sh

.PHONY: corda-tests
corda-tests:
	@$(TEST_SCRIPTS_PATH)/corda.sh
