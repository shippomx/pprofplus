web:
	cd front && rm -rf dist && npm run build && docker build -t pprof:web .

bin:
	export GO111MODULE=on && go build

