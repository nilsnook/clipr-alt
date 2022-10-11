package main

import (
	"encoding/json"
	"os"
	"strconv"
)

func (c *clipr) initCurrentState() {
	if val := os.Getenv("ROFI_RETV"); val != "" {
		c.infolog.Printf("State: %s", val)
		c.state.val, _ = strconv.Atoi(val)
	}

	if info := os.Getenv("ROFI_INFO"); info != "" {
		c.infolog.Printf("Info: %s", info)
		json.Unmarshal([]byte(info), &c.state.info)
	}

	args := os.Args
	if len(args) > 1 {
		for k, v := range args {
			if k == 1 {
				c.infolog.Printf("Arg: %s", v)
				c.state.arg = v
			}
		}
	}
}
