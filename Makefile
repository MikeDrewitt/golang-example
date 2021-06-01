build:
	go build -o api .

run: build
	./api

watch:
	ulimit -n 1000 # increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' make run
