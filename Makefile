.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
.PHONY: tag
tag:
	git push --tags
.PHONY: run
run:
	go run main.go
.PHONY: build
build:
	docker build -f Dockerfile  -t user-api .