V?=0

KUBEPATH?=~/.kube/plugins/
NAME?='kubectl-in_toto'
deploy:
	@mkdir -vp $(KUBEPATH)
	@go build -o $(KUBEPATH)/$(NAME)
	echo plugin installed on $(KUBEPATH)/$(NAME)

test:
	go test ./...
