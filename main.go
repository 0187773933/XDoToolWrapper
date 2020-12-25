package main

import (
	// "fmt"
	xdotool "github.com/0187773933/XDoToolWrapper/v1"
)

func main() {
	// xdo := xdotool.Wrapper{}
	// xdo.Window.Name = "Chrome"
	// xdo.GetMonitors()
	// xdo.Attach( 3 , 500 )
	// xdo.Refocus()
	// xdo.GetGeometry()
	// fmt.Println( xdo )

	xdo := xdotool.Wrapper{}
	xdo.Window.Name = "spotify"
	xdo.GetMonitors()
	xdo.AttachClass( 3 , 500 )
	xdo.SetWindowSize( 1280 , 733 )
	xdo.MoveWindow( 300 , 212 )
}