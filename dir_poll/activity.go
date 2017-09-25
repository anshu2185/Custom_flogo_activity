package dir_poll

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os"
	"fmt"
	"strconv"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-dir-poller")

const (
	ivDirectoryName   = "directory_name"
	ivAddToFlow = "addToFlow"

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
	addToFlow, _ := toBool(context.GetInput(ivAddToFlow))

	flag := 0

	if _, err := os.Stat(dirName); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			//fmt.Println("File does not exist")
			//msg := "File does not exist"
			flag = 0
		} else {
			// other error
		}
	} else {
		//fmt.Println("Found the file to be polled")
		//msg := "Found the file to be polled"
		flag = 1
	}

	if addToFlow && ( flag == 1 ) {
		context.SetOutput(ovMessage, "Found the file to be polled")
	} else {
		context.SetOutput(ovMessage, "File not found")
	}


	return true, nil
}


func toBool(val interface{}) (bool, error) {
	
		b, ok := val.(bool)
		if !ok {
			s, ok := val.(string)
	
			if !ok {
				return false, fmt.Errorf("unable to convert to boolean")
			}
	
			var err error
			b, err = strconv.ParseBool(s)
	
			if err != nil {
				return false, err
			}
		}
	
		return b, nil
	}