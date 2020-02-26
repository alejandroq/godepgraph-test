.PHONY: dependencytree
dependencytree:
	godepgraph -s cmd/main.go | dot -Tpng -o graph.png

.PHONY: start
start:
	go run cmd/main.go