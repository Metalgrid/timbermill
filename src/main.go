package main

import (
	"fmt"

	syslog "gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	var config *TMConfig
	config = LoadConfig("../config/timbermill.yaml")

	collectors := make(map[string]*syslog.Server)

	for _, collector := range config.Collectors {
		fmt.Printf("[INFO] Running collector %+v\n", collector)
		collectors[collector.Name] = MakeSyslogCollector(collector.Name, collector.Protocol, collector.Port)
	}

	// Gracefully cleanup the collectors
	defer func(collectors map[string]*syslog.Server) {
		for _, c := range collectors {
			fmt.Printf("[INFO] Shutting down collector %+v\n", c)
			c.Kill()
		}
	}(collectors)

	fmt.Scanln() // Enter to quit
}
