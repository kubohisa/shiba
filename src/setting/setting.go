package setting

import (
	"net/http"

	"pdbg.work/shiba/module/exec"
)

var (
	Functions = map[string]func(w http.ResponseWriter, r *http.Request, param []string){
		"hello": exec.Welcome,
	}

	ServerSecondLimit int = 1000
	MaxClients        int = 30

	CronTimerMicroseccond int = 1000000 // 1Sec.

	TimerExec bool = false
)
