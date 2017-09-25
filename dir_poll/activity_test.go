package dir_poll

import (
	"io/ioutil"
	"testing"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("directory_name", "/home/akash/test")

	act.Eval(tc)

	//check result attr
}



func TestAddToFlow(t *testing.T) {
	
		act := NewActivity(getActivityMetadata())
		tc := test.NewTestActivityContext(getActivityMetadata())
	
		//setup attrs
		tc.SetInput("directory_name", "test message")
	
		act.Eval(tc)
	
		msg := tc.GetOutput("message")
	
		fmt.Println("Message: ", msg)
	
		if msg == nil {
			t.Fail()
		}
	}