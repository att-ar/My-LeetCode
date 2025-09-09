build:
	hpack
	stack build

rebuild:
	python generate_category_runner.py ./src/Problems/${CATEGORY}
	make build

run:
	stack run

build-run: build run

rebuild-run: rebuild run
