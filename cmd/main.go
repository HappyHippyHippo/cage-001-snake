// Package cmd @todo doc
package main

import (
	"github.com/happyhippyhippo/cage"
)

func main() {
	// create game application
	game := (&cage.Game{}).Init()
	defer func() { _ = game.Close() }()
	/*
	   // initialize the logger
	   log := game.Logger()
	   logConsole := cage.NewLogStreamConsole()
	   logConsole.SetLevel(cage.LogInfo)
	   logConsole.AddChannel("cage")
	   _ = log.AddStream(logConsole)

	   // set the window
	   game.Window().SetWindowSize(cage.VectorI32{640, 400})
	   game.Window().SetResizeMode(cage.ResizeEnabled)

	   // register the game scenes
	   scenes := game.Scenes()
	   _ = scenes.AddScene("boot", func() cage.ITraversable { return (&boot.Scene{}).Init() })
	   _ = scenes.AddScene("present", func() cage.ITraversable { return (&present.Scene{}).Init() })
	   _ = scenes.QueueScene("boot")

	   // run the game app

	   	if e := game.Run(); e != nil {
	   		log.Broadcast(cage.LogFatal, "fatal error : %s", cage.LogCtxMsgArgs(e.Error()))
	   		return
	   	}

	   log.Broadcast(cage.LogInfo, "application terminated")
	*/
}
