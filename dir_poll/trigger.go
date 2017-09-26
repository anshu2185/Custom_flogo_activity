package dir_poll

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"time"
	"github.com/radovskyb/watcher"
	"io/ioutil"
	"log"
	"fmt"
	
)

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct{
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata:md}
}

//New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config:config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
}

// Init implements trigger.Trigger.Init
func (t *MyTrigger) Init(runner action.Runner) {
	t.runner = runner
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	// start the trigger

	w := watcher.New()
	
	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)
	
	//  notify write, create, remove, rename events.
	w.FilterOps(watcher.Write, watcher.Create, watcher.Move, watcher.Remove, watcher.Rename)
	
	initialTime := time.Time( time.Now() )	
	
	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
				//fmt.Println("The below file was changed:")
				fmt.Println(event.Path)

				if event.Path != "-" {
					for _, f := range w.WatchedFiles() {
						//fmt.Printf("%s: %s \n", path, f.Name())
						showfiles(event.Path, initialTime, f.Name(), f.ModTime())
					}
				}

				initialTime = time.Now()
			
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch test_folder recursively for changes.

	handlers := t.config.Handlers
	for _, handler := range handlers {
				dirName := handler.Settings["dirName"].(string)
				if err := w.AddRecursive(dirName); err != nil {
					log.Fatalln(err)
				}
			
	}

	
	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s \n", path, f.Name())
	}

	fmt.Println()

	// Trigger 2 events after watcher started.
	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Create, nil)
		w.TriggerEvent(watcher.Remove, nil)
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
	
	return nil
}


//If a file is modified after initial time then fileLAstmodified > initial time 
//Hence fileLastmodified - initial time > 0
func showfiles(dirPath string, initialTime time.Time, filename string, fileLastModified time.Time) {
	
	files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        log.Fatal(err)
	}
	
    for _, f := range files {
		if f.Name() == filename && ( fileLastModified.Sub(initialTime) > 0 ) {
			fmt.Println(f.Name())
		}             
	}
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}
