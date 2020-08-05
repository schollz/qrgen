package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mdp/qrterminal/v3"
	"rsc.io/qr"
)

var levelFlag string
var quietZoneFlag int

func getLevel(s string) qr.Level {
	switch l := strings.ToLower(s); l {
	case "l":
		return qr.L
	case "m":
		return qr.M
	case "h":
		return qr.H
	default:
		return -1
	}
}

func main() {
	levelFlag = "l"
	quietZoneFlag = 2
	level := getLevel(levelFlag)

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	cfg := qrterminal.Config{
		Level:     level,
		Writer:    os.Stdout,
		QuietZone: quietZoneFlag,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
	}

	if runtime.GOOS == "windows" {
		cfg.Writer = colorable.NewColorableStdout()
		cfg.BlackChar = qrterminal.BLACK
		cfg.WhiteChar = qrterminal.WHITE
	}

	qrterminal.GenerateWithConfig(strings.TrimSpace(string(data)), cfg)
}
