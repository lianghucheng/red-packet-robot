package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"redpacket-sweep-robot/conf"
	"redpacket-sweep-robot/robot"

	llog "github.com/name5566/leaf/log"
)

func init() {
	logger, err := llog.New("release", conf.GetCfgGameInfo().LogPath, log.Lshortfile|log.LstdFlags)
	if err != nil {
		panic(err)
	}
	llog.Export(logger)
}

func main() {
	robot.Item = flag.Int("item", 0, "item")
	flag.Parse()

	robot.Init()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case sig := <-c:
		llog.Release("closing down (signal: %v)", sig)
		robot.Destroy()
	}
}
