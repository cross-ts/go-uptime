package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"time"
)

type Uptime struct {
	Now    time.Time
	Uptime time.Duration
	Load1  float64
	Load5  float64
	Load15 float64
}

func (u Uptime) String() string {
	now := u.Now.Format("15:04:05")
	uptime := int(u.Uptime.Hours() / 24)
	return fmt.Sprintf(" %s up %d days , load average: %.2f, %.2f, %.2f", now, uptime, u.Load1, u.Load5, u.Load15)
}

func uptime() (u Uptime) {
	sysinfo := unix.Sysinfo_t{}
	err := unix.Sysinfo(&sysinfo)
	if err != nil {
		panic(err)
	}
	u = Uptime{}
	u.Now = time.Now()
	u.Uptime = time.Duration(sysinfo.Uptime) * time.Second
	u.Load1 = float64(sysinfo.Loads[0]) / (1 << 16)
	u.Load5 = float64(sysinfo.Loads[1]) / (1 << 16)
	u.Load15 = float64(sysinfo.Loads[2]) / (1 << 16)
	return
}

func main() {
	fmt.Println(uptime())
}
