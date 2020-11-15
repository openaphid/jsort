test:
	go test ./... -v
bench:
	go test -run=match_nothing -bench . -count 5 -timeout 90m -benchmem | tee BenchmarkResult/benchmark.txt
benchstat:
	benchstat -csv BenchmarkResult/benchmark.txt | tee BenchmarkResult/benchstats.csv.txt
	go run cmd/parse_benchstats/main.go BenchmarkResult/benchstats.csv.txt
operation_stats:
	go test -run TestOperationStats | tee operation_stats.txt
go2_bench:
	go test github.com/openaphid/jsort/internal/sort_slice_go2 -run=match_nothing -bench . -count 5 -timeout 90m -benchmem | tee BenchmarkResult/Go2/benchmark.txt
go2_benchstat:
	benchstat -csv BenchmarkResult/Go2/benchmark.txt | tee BenchmarkResult/Go2/benchstats.csv.txt
	go run cmd/parse_benchstats/main.go BenchmarkResult/Go2/benchstats.csv.txt
