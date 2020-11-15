test:
	go test ./... -v
bench:
	go test -run=match_nothing -bench . -count 5 -timeout 90m -benchmem | tee BenchmarkResult/benchmark.txt
benchstat:
	benchstat -csv BenchmarkResult/benchmark.txt | tee BenchmarkResult/benchstats.csv.txt
	go run cmd/parse_benchstats/main.go BenchmarkResult/benchstats.csv.txt
operation_stats:
	go test -run TestOperationStats | tee operation_stats.txt
