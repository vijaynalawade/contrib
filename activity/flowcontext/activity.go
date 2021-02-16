package flowcontext

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/engine"
)

func init() {
	_ = activity.Register(&Activity{})
}

type Output struct {
	FlowName    string `md:"FlowName"`
	FlowID    string `md:"FlowId"`
	AppName    string `md:"AppName"`
	AppVersion    string `md:"AppVersion"`
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"FlowName":    o.FlowName,
		"FlowId":    o.FlowID,
		"AppName":    o.AppName,
		"AppVersion":    o.AppVersion,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.FlowName, err = coerce.ToString(values["FlowName"])
	if err != nil {
		return err
	}
	o.FlowID, err = coerce.ToString(values["FlowId"])
	if err != nil {
		return err
	}
	o.AppName, err = coerce.ToString(values["AppName"])
	if err != nil {
		return err
	}
	o.AppVersion, err = coerce.ToString(values["AppVersion"])
	if err != nil {
		return err
	}
	return nil
}

var activityMd = activity.ToMetadata(&Output{})

// Activity is an Activity that is used to flow context
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	output := &Output{}
	output.FlowName = ctx.ActivityHost().Name()
	output.FlowID = ctx.ActivityHost().ID()
	output.AppName = engine.GetAppName()
	output.AppVersion = engine.GetAppVersion()
	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
