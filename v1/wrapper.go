package wrapper

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"strings"
	"strconv"
)

func get_display_number() ( result string ) {
	command := exec.Command( "/bin/bash" , "-c" , "/usr/local/bin/getDisplayNumber"  )
	out, err := command.Output()
	if err != nil {
		fmt.Sprintf( "%s\n" , err )
	}
	result = string( out[:] )
	return
}

var APPEND_DISPLAY bool = false
func exec_process( bash_command string , arguments ...string ) ( result string ) {
	result = "failed"
	command := exec.Command( bash_command , arguments... )
	if APPEND_DISPLAY == true {
		display_number := get_display_number()
		command.Env = append( os.Environ() , fmt.Sprintf( "DISPLAY=%s" , display_number ) )
	}
	out , err := command.Output()
	if err != nil {
		fmt.Println( bash_command )
		fmt.Println( arguments )
		fmt.Println( err )
	} else {
		result = string( out[:] )
	}
	return
}

type Wrapper struct {
	Window struct {
		Id int
		Name string
		Geometry struct {
			X int
			Y int
			Center struct {
				X int
				Y int
			}
		}
	}
	Monitors struct {
		Primary struct {
			Name string
			X int
			Y int
		}
		Secondary struct {
			Name string
			X int
			Y int
		}
	}
}

func ( xdo *Wrapper ) GetMonitors() {
	info := exec_process( "/bin/bash" , "-c" , "xrandr --prop | grep connected" )
	lines := strings.Split( info , "\n" )
	for _ , line := range lines {
		words := strings.Split( line , " " )
		if len( words ) < 3 { continue }
		switch words[ 2 ] {
			case "Primary":
				Name := words[ 0 ]
				size := words[ 3 ]
				size_components := strings.Split( size , "x" )
				X , _ := strconv.Atoi( size_components[ 0 ] )
				Y , _ := strconv.Atoi( strings.Split( size_components[ 1 ] , "+" )[ 0 ] )
				xdo.Monitors.Primary.Name = Name
				xdo.Monitors.Primary.X = X
				xdo.Monitors.Primary.Y = Y
			case "Secondary":
				Name := words[ 0 ]
				size := words[ 3 ]
				size_components := strings.Split( size , "x" )
				X , _ := strconv.Atoi( size_components[ 0 ] )
				Y , _ := strconv.Atoi( strings.Split( size_components[ 1 ] , "+" )[ 0 ] )
				xdo.Monitors.Secondary.Name = Name
				xdo.Monitors.Secondary.X = X
				xdo.Monitors.Secondary.Y = Y
		}
	}
}

func ( xdo *Wrapper ) AttachClass( options ...int ) {
	number_of_tries := 20
	sleep_milliseconds := 1000
	if len( options ) > 0 {
		number_of_tries = options[0]
	}
	if len( options ) > 1 {
		sleep_milliseconds = options[1]
	}
	duration , _ := time.ParseDuration( strconv.Itoa( sleep_milliseconds ) + "ms" )
	for i := 0; i < number_of_tries; i++ {
		//cmd := fmt.Sprintf( "xdotool search --desktop 0 --name '%s'" , xdo.Window.Name )
		// Spotify Changes the Title of the XWindow to be the "Now Playing" Track Information
		// http://manpages.ubuntu.com/manpages/trusty/man1/xdotool.1.html#window%20stack
		// http://manpages.ubuntu.com/manpages/trusty/man3/XQueryTree.3.html
		cmd := fmt.Sprintf( "xdotool search --class '%s'" , xdo.Window.Name )
		info := exec_process( "/bin/bash" , "-c" , cmd )
		if info == "failed" {
			APPEND_DISPLAY = !APPEND_DISPLAY
			info = exec_process( "/bin/bash" , "-c" , cmd )
		}
		lines := strings.Split( info , "\n" )
		lines = lines[ 0 : ( len( lines ) - 1 ) ]
		window_id , error := strconv.Atoi( lines[ ( len( lines ) - 1 ) ] )
		if error != nil {
			time.Sleep( duration )
		} else {
			xdo.Window.Id = window_id
			return
		}
	}
}

// func ( xdo *Wrapper ) AttachClass( options ...int ) {
// 	number_of_tries := 20
// 	sleep_milliseconds := 1000
// 	if len( options ) > 0 {
// 		number_of_tries = options[0]
// 	}
// 	if len( options ) > 1 {
// 		sleep_milliseconds = options[1]
// 	}
// 	duration , _ := time.ParseDuration( strconv.Itoa( sleep_milliseconds ) + "ms" )
// 	for i := 0; i < number_of_tries; i++ {
// 		//cmd := fmt.Sprintf( "xdotool search --desktop 0 --name '%s'" , xdo.Window.Name )
// 		// Spotify Changes the Title of the XWindow to be the "Now Playing" Track Information
// 		// http://manpages.ubuntu.com/manpages/trusty/man1/xdotool.1.html#window%20stack
// 		// http://manpages.ubuntu.com/manpages/trusty/man3/XQueryTree.3.html
// 		cmd := fmt.Sprintf( "xdotool search --class '%s'" , xdo.Window.Name )
// 		info := exec_process( "/bin/bash" , "-c" , cmd )
// 		if info == "failed" {
// 			APPEND_DISPLAY = !APPEND_DISPLAY
// 			info = exec_process( "/bin/bash" , "-c" , cmd )
// 		}
// 		lines := strings.Split( info , "\n" )
// 		window_id , error := strconv.Atoi( lines[ 1 ] )
// 		if error != nil {
// 			time.Sleep( duration )
// 		} else {
// 			xdo.Window.Id = window_id
// 			return
// 		}
// 	}
// }

func ( xdo *Wrapper ) Attach( options ...int ) {
	number_of_tries := 20
	sleep_milliseconds := 1000
	if len( options ) > 0 {
		number_of_tries = options[0]
	}
	if len( options ) > 1 {
		sleep_milliseconds = options[1]
	}
	duration , _ := time.ParseDuration( strconv.Itoa( sleep_milliseconds ) + "ms")
	for i := 0; i < number_of_tries; i++ {
		//cmd := fmt.Sprintf( "xdotool search --desktop 0 --name '%s'" , xdo.Window.Name )

		// Spotify Changes the Title of the XWindow to be the "Now Playing" Track Information
		// http://manpages.ubuntu.com/manpages/trusty/man1/xdotool.1.html#window%20stack
		// http://manpages.ubuntu.com/manpages/trusty/man3/XQueryTree.3.html

		cmd := fmt.Sprintf( "xdotool search --name '%s'" , xdo.Window.Name )
		info := exec_process( "/bin/bash" , "-c" , cmd )
		if info == "failed" {
			APPEND_DISPLAY = !APPEND_DISPLAY
			info = exec_process( "/bin/bash" , "-c" , cmd )
		}
		lines := strings.Split( info , "\n" )
		window_id , error := strconv.Atoi( lines[ 0 ] )
		if error != nil {
			time.Sleep( duration )
		} else {
			xdo.Window.Id = window_id
			return
		}
	}
}

func ( xdo *Wrapper ) Activate() {
	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool windowactivate %d" , xdo.Window.Id ) )
}

func ( xdo *Wrapper ) Focus() {
	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool windowfocus %d" , xdo.Window.Id ) )
}

func ( xdo *Wrapper ) Refocus() {
	xdo.Activate()
	sleep_duration , _ := time.ParseDuration( "300ms" )
	time.Sleep( sleep_duration )
	xdo.Focus()
}

func ( xdo *Wrapper ) GetGeometry() {
	xdo.Refocus()
	info := exec_process( "/bin/bash" , "-c" , "xdotool getactivewindow getwindowgeometry" )
	lines := strings.Split( info , "\n" )
	geometry_components := strings.Split( strings.Split( lines[ 2 ] , "Geometry: " )[ 1 ] , "x" )
	xdo.Window.Geometry.X , _ = strconv.Atoi( geometry_components[0] )
	xdo.Window.Geometry.Y , _ = strconv.Atoi( geometry_components[1] )
	xdo.Window.Geometry.Center.X = ( xdo.Window.Geometry.X / 2 )
	xdo.Window.Geometry.Center.Y = ( xdo.Window.Geometry.Y / 2 )
}

func ( xdo *Wrapper ) UnMaximize() {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "wmctrl -rf %d -b remove,maximized_ver,maximized_horz" , xdo.Window.Id ) )
}

func ( xdo *Wrapper ) Maximize() {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , "xdotool key F11" )
}

func ( xdo *Wrapper ) FullScreen() {
	xdo.Maximize()
}

func ( xdo *Wrapper ) MoveMouse( X int , Y int ) {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool mousemove %d %d" , X , Y ) )
}

func ( xdo *Wrapper ) LeftClick() {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , "xdotool click 1" )
}

func ( xdo *Wrapper ) RightClick() {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , "xdotool click 2" )
}

func ( xdo *Wrapper ) DoubleClick() {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , "xdotool click --repeat 2 --delay 200 1" )
}

func ( xdo *Wrapper ) CenterMouse() {
	xdo.Refocus()
	xdo.MoveMouse( xdo.Window.Geometry.Center.X , xdo.Window.Geometry.Center.Y )
}

func ( xdo *Wrapper ) PressKey( key string ) {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool key '%s'" , key )  )
}

func ( xdo *Wrapper ) SetWindowSize( x int , y int ) {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool windowsize %d %d %d" , xdo.Window.Id , x , y ) )
}

func ( xdo *Wrapper ) MoveWindow( x int , y int ) {
	xdo.Refocus()
	exec_process( "/bin/bash" , "-c" , fmt.Sprintf( "xdotool windowmove %d %d %d" , xdo.Window.Id , x , y ) )
}