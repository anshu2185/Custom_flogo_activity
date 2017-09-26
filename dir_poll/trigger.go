package dir_poll


import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"time"
	"github.com/radovskyb/watcher"
	"io/ioutil"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"fmt"
)

// log is the default package logger
var log = logger.GetLogger("trigger-tibco-mqtt")

// MqttTrigger is simple MQTT trigger
type MqttTrigger struct {
		metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config

}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MQTTFactory{metadata: md}
}

// MQTTFactory MQTT Trigger factory
type MQTTFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *MQTTFactory) New(config *trigger.Config) trigger.Trigger {
	return &MqttTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *MqttTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Init implements ext.Trigger.Init
func (t *MqttTrigger) Init(runner action.Runner) {
	t.runner = runner
}

// Start implements ext.Trigger.Start
func (t *MqttTrigger) Start() error {

	
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
				panic(err)
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
					panic(err)
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
		panic(err)
	}
	
	return nil
}
	
	//If a file is modified after initial time then fileLAstmodified > initial time 
//Hence fileLastmodified - initial time > 0
func showfiles(dirPath string, initialTime time.Time, filename string, fileLastModified time.Time) {
	
	files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        panic(err)
	}
	
    for _, f := range files {
		if f.Name() == filename && ( fileLastModified.Sub(initialTime) > 0 ) {
			fmt.Println(f.Name())
		}             
	}
}


// Stop implements ext.Trigger.Stop
func (t *MqttTrigger) Stop() error {
	//unsubscribe from topic
	
	return nil
}

