package dir_poll

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os"
	"fmt"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-dir-poller")

const (
	ivDirectoryName   = "dir_name"

	ovMessage = "message"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval

	dirName, _ := context.GetInput(ivDirectoryName).(string)

	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			fmt.Println("File does not exist")
		} else {
			// other error
		}
	}

	return true, nil
}
