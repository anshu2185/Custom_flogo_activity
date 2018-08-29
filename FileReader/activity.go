package FileReader 


import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"bufio"
	"os"
)

// log is the default package logger
var log = logger.GetLogger("activity-akash file reader")

const (
	filename   = "filename"
	result    =  "result"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval

	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
 	var out string
	for fileScanner.Scan() {
		out = fileScanner.Text()
		
	}
	context.SetOutput(result, out)
	return true, nil
}

