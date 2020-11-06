test:
	go test ./... -v
bench:
	go test -run=match_nothing -bench . -count 5 -timeout 20m -benchmem | tee benchmark.txt
