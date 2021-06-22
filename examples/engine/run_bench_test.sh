# go tool pprof profile.out
# go test  -bench=Mod -benchmem -cpuprofile profile.out -memprofile memprofile.out
echo `

` >> benchmark_results.txt;
date >> benchmark_results.txt;
go test -bench=. -benchmem >> benchmark_results.txt

