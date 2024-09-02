package core

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/agaUHO/aga/system"
	"github.com/radovskyb/watcher"
)

// WatcherModulesPlugins -- watcher the plugins and modules folder
func WatcherModulesPlugins() {
	w := watcher.New()
	w.SetMaxEvents(1)
	// Only notify rename and move events.
	// w.FilterOps(watcher.Create, watcher.Rename, watcher.Write, watcher.Move, watcher.Remove, watcher.Chmod)
	w.FilterOps(watcher.Create, watcher.Rename, watcher.Move, watcher.Remove)
	go func() {
		for {
			select {
			case event := <-w.Event:
				if !strings.HasSuffix(event.Path, ".swp") {
					fmt.Println("A change was detected " + event.Path + ", you must restart the system")
					// SendNotification(system.AdminUserId, "aga", "Un cambio a sido detectado, el sistema debe reiniciarse.")
				}
			case err := <-w.Error:
				fmt.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch these plugins for changes.
	if err := w.Add(system.Path + "/plugins"); err != nil {
		log.Fatalln(err)
	}
	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
