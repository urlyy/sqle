################################## Parameter Definition And Check ##########################################
override GIT_VERSION    		= $(shell git rev-parse --abbrev-ref HEAD)${CUSTOM} $(shell git rev-parse HEAD)
override GIT_COMMIT     		= $(shell git rev-parse HEAD)
override LDFLAGS 				= -ldflags "-X 'main.version=\"${GIT_VERSION}\"'"
override RPM_BUILD_BIN  		= $(shell type -p rpmbuild 2>/dev/null)
override DOCKER         		= $(shell which docker)
override GONOSUMDB      		= actiontech.cloud
override GOOS           		= linux
override OS_VERSION 			= el7
override GOARCH         		= amd64
override RPMBUILD_TARGET		= x86_64
override RELEASE 				= qa
override CUSTOMER 				= standard
override GO_BUILD_FLAGS 		= -mod=vendor
override RPM_USER_GROUP_NAME 	= actiontech
override RPM_USER_NAME 			= actiontech-universe

## The docker registry to pull complier image, can be overwrite by: `make DOCKER_REGISTRY=10.0.0.1`
DOCKER_REGISTRY ?= 10.186.18.20

## The ftp host and user/pass used to upload rpm package, can be overwrite by: `make RELEASE_FTPD_HOST=ftp:ftp@10.186.18.21`
RELEASE_FTPD_HOST ?= admin:ftpadmin@10.186.18.20

## Dynamic Parameter
# The ftp host and user/pass used to upload rpm package, can be overwrite by: `make DOCKER_IMAGE=image:tag`
DOCKER_IMAGE  ?= docker-registry:5000/actiontech/universe-compiler-go1.14.1-centos6
DOTNET_DOCKER_IMAGE ?= docker-registry:5000/actiontech/universe-compiler-dotnetcore2.1
TEST_DOCKER_IMAGE  ?= docker-registry:5000/actiontech/universe-compiler-go1.14.1-ubuntu-with-docker

## Static Parameter, should not be overwrite
PROJECT_NAME = sqle
VERSION = 1.3.0
GOBIN = ${shell pwd}/bin
PARSER_PATH   = ${shell pwd}/vendor/github.com/pingcap/parser

default: install
######################################## Code Check ####################################################
## Static Code Analysis
vet: swagger
	GOOS=$(GOOS) GOARCH=amd64 go vet $$(GOOS=${GOOS} GOARCH=${GOARCH} go list ./...)
	GOOS=$(GOOS) GOARCH=amd64 go vet ./vendor/actiontech.cloud/...

## Unit Test
test: swagger parser
	cd $(PROJECT_NAME) && GOOS=$(GOOS) GOARCH=amd64 go test -v ./...

docker_test: pull_image
	CTN_NAME="universe_docker_test_$$RANDOM" && \
    $(DOCKER) run -d --entrypoint /sbin/init --add-host docker-registry:${DOCKER_REGISTRY}  --privileged --name $${CTN_NAME} -v $(shell pwd):/universe/sqle --rm -w /universe/sqle $(DOCKER_IMAGE) && \
    $(DOCKER) exec $${CTN_NAME} make test vet ; \
    $(DOCKER) stop $${CTN_NAME}

################################### Golang Binary Compile #############################################
docker_clean:
	$(DOCKER) run -v $(shell pwd):/universe --rm $(DOCKER_IMAGE) -c "cd /universe && make clean ${MAKEFLAGS}"

docker_install:
	$(DOCKER) run -v $(shell pwd):/universe --rm $(DOCKER_IMAGE) -c "cd /universe && make install $(MAKEFLAGS)"

install: swagger parser
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GO_BUILD_FLAGS) ${LDFLAGS} -tags $(GO_BUILD_TAGS) -o $(GOBIN)/sqled ./$(PROJECT_NAME)/cmd/sqled

swagger:
	GOARCH=amd64 go build -o ${shell pwd}/bin/swag ${shell pwd}/build/swag/main.go
	rm -rf ${shell pwd}/sqle/docs
	${shell pwd}/bin/swag init -g ./$(PROJECT_NAME)/api/app.go -o ${shell pwd}/sqle/docs

parser:
	cd build/goyacc && GOOS=${GOOS} GOARCH=amd64 GOBIN=$(GOBIN) go install
	$(GOBIN)/goyacc -o /dev/null ${PARSER_PATH}/parser.y
	$(GOBIN)/goyacc -o ${PARSER_PATH}/parser.go ${PARSER_PATH}/parser.y 2>&1 | egrep "(shift|reduce)/reduce" | awk '{print} END {if (NR > 0) {print "Find conflict in parser.y. Please check y.output for more information."; exit 1;}}'
	rm -f y.output

	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		sed -i -e 's|//line.*||' -e 's/yyEofCode/yyEOFCode/' ${PARSER_PATH}/parser.go; \
	elif [ $(ARCH) = $(MAC) ]; \
	then \
		/usr/bin/sed -i "" 's|//line.*||' ${PARSER_PATH}/parser.go; \
		/usr/bin/sed -i "" 's/yyEofCode/yyEOFCode/' ${PARSER_PATH}/parser.go; \
	fi

	@awk 'BEGIN{print "// Code generated by goyacc DO NOT EDIT."} {print $0}' ${PARSER_PATH}/parser.go > tmp_parser.go && mv tmp_parser.go ${PARSER_PATH}/parser.go;

# Clean
clean:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go clean

###################################### RPM Build #####################################################
pull_image:
	$(DOCKER) pull $(DOCKER_IMAGE)

docker_rpm: docker_rpm/sqle

docker_rpm/sqle: pull_image docker_install
	$(DOCKER) run -v $(shell pwd):/universe/sqle --rm $(DOCKER_IMAGE) -c "(mkdir -p /root/rpmbuild/SOURCES >/dev/null 2>&1);cd /root/rpmbuild/SOURCES; \
	(tar zcf ${PROJECT_NAME}.tar.gz /universe --transform 's/universe/${PROJECT_NAME}-${VERSION}_$(GIT_COMMIT)/' >/tmp/build.log 2>&1) && \
	(rpmbuild --define 'group_name $(RPM_USER_GROUP_NAME)' --define 'user_name $(RPM_USER_NAME)' \
	--define 'commit $(GIT_COMMIT)' --define 'os_version $(OS_VERSION)' \
	--target $(RPMBUILD_TARGET)  -bb --with qa /universe/sqle/build/sqled.spec >>/tmp/build.log 2>&1) && \
	(cat /root/rpmbuild/RPMS/$(RPMBUILD_TARGET)/${PROJECT_NAME}-${VERSION}_$(GIT_COMMIT)-qa.$(OS_VERSION).$(RPMBUILD_TARGET).rpm) || (cat /tmp/build.log && exit 1)" > $(PROJECT_NAME).$(CUSTOMER).$(RELEASE).$(OS_VERSION).$(RPMBUILD_TARGET).rpm

upload: upload/sqle

upload/sqle:
	curl -T $(shell pwd)/$(PROJECT_NAME).$(CUSTOMER).$(RELEASE).$(OS_VERSION).$(RPMBUILD_TARGET).rpm \
	ftp://$(RELEASE_FTPD_HOST)/actiontech-$(PROJECT_NAME)/qa/$(VERSION)/$(PROJECT_NAME)-$(VERSION).$(CUSTOMER).$(RELEASE).$(OS_VERSION).$(RPMBUILD_TARGET).rpm --ftp-create-dirs

###################################### docker start #####################################################
docker_start: fill_ui_dir docker_rpm/sqle
	cd ./docker-compose && docker-compose up -d
	./docker-compose/install.sh

docker_stop:
	cd ./docker-compose && docker-compose down

###################################### ui #####################################################
fill_ui_dir:
	# fill ui dir, it is used by rpm build.
	mkdir -p ./ui/static

.PHONY: help
help:
	$(warning ---------------------------------------------------------------------------------)
	$(warning Supported Variables And Values:)
	$(warning ---------------------------------------------------------------------------------)
	$(foreach v, $(.VARIABLES), $(if $(filter file,$(origin $(v))), $(info $(v)=$($(v)))))
