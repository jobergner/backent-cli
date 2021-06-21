echo `

` >> benchmark_results.txt;
date >> benchmark_results.txt;
go test -bench=. -benchmem -benchtime=12s >> benchmark_results.txt

