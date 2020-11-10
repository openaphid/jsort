test:
	go test ./... -v
bench:
	go test -run=match_nothing -bench . -count 5 -timeout 90m -benchmem | tee benchmark.txt
benchstats:
	benchstats benchmark.txt
operation_stats:
	go test -run TestOperationStats | tee operation_stats.txt
