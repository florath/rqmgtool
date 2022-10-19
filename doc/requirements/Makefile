all: artifacts/rqmgtool-graph1.png

artifacts/rqmgtool-graph1.png: artifacts/rqmgtool-graph1.dot
	unflatten -l 23 artifacts/rqmgtool-graph1.dot | \
		dot -Tpng -o artifacts/rqmgtool-graph1.png
