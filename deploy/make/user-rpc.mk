VERSION=latest
SERVER_NAME=user
SERVER_TYPE=rpc

DOCKER_REPO_TEST=crpi-j6xgprtkk3jhxlja.us-west-1.personal.cr.aliyuncs.com/chat-qy/${SERVER_NAME}-${SERVER_TYPE}-dev
VERSION_TEST=${VERSION}
APP_NAME_TEST=chat-${SERVER_NAME}-${SERVER_TYPE}-test
DOCKER_FILE_TEST=./deploy/dockerfile/Dockerfile_${SERVER_NAME}-${SERVER_TYPE}_dev

build-test:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${SERVER_NAME}-${SERVER_TYPE} ./apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go
	docker build . -f ${DOCKER_FILE_TEST} --no-cache -t ${APP_NAME_TEST}
	@echo  'build OK'

tag-test:
	@echo  'create tag ${VERSION_TEST}'
	docker tag ${APP_NAME_TEST} ${DOCKER_REPO_TEST}:${VERSION_TEST}
	@echo  'tag OK'

publish-test:
	#docker login --username=tom179 crpi-j6xgprtkk3jhxlja.us-west-1.personal.cr.aliyuncs.com # 为什么不登录也可以？
	@echo 'pushing ${VERSION_TEST} to ${DOCKER_REPO_TEST}'
	docker push ${DOCKER_REPO_TEST}:${VERSION_TEST}
	@echo  'publish OK'

release-test: build-test tag-test publish-test