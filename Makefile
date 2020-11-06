test:
	go test ./... -v
bench:
	go test ./... -run=match_nothing -bench . -count 5 -benchmem | tee benchmark.txt
