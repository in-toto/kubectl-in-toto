V?=0

deploy:
	@mkdir -p ~/.kube/plugins/in-toto
	@go build -o ~/.kube/plugins/in-toto/in-toto
	@cp plugin.yaml ~/.kube/plugins/in-toto/

