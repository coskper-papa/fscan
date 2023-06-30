package main

import (
	"fmt"
	"github.com/coskper-papa/fscan/Plugins"
	"github.com/coskper-papa/fscan/common"
	"time"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	t := time.Now().Sub(start)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", t)
}
