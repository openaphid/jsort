test:
	go test -v
bench:
	go test -run=match_nothing -bench . -benchmem
