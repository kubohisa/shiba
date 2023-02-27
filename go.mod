module pdbg.work/shiba/main

go 1.20

require pdbg.work/shiba/module/server v0.0.0-00010101000000-000000000000

require (
	golang.org/x/time v0.3.0 // indirect
	pdbg.work/shiba/module/exec v0.0.0-00010101000000-000000000000 // indirect
	pdbg.work/shiba/module/parse v0.0.0-00010101000000-000000000000 // indirect
	pdbg.work/shiba/module/setting v0.0.0-00010101000000-000000000000 // indirect
)

replace pdbg.work/shiba/module/server => ./src/module/server

replace pdbg.work/shiba/module/exec => ./src/exec

replace pdbg.work/shiba/module/setting => ./src/setting

replace pdbg.work/shiba/module/parse => ./src/module/parse
