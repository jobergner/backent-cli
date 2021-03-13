echo `

` >> benchmark_results.txt;
date >> benchmark_results.txt;
go test -bench=. -benchmem >> benchmark_results.txt

