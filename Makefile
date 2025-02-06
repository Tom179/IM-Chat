user-rpc-dev:
	@make -f deploy/make/user-rpc.mk release-test # 需要指定哪个文件的具体执行目标

release-test:user-rpc-dev

deploy-test:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh