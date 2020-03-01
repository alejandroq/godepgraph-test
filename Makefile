.PHONY: dependencytree
dependencytree:
	godepgraph -s cmd/main.go | dot -Tpng -o .github/dependencytree.png

.PHONY: start
start:
	go run cmd/main.go