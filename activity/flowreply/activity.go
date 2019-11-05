package flowreply

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"

)

func init() {
	_ = activity.Register(&Activity{}, New)
}

type Input struct {
	Reply map[string]interface{} `md:"reply,required"` // Set of mappings to execute when the activity runs
}

var activityMd = activity.ToMetadata(&Input{})

func New(ctx activity.InitContext) (activity.Activity, error) {


	act := &Activity{}


	return act, nil
}

// Activity is an Activity that is used to reply/return via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type Activity struct {

}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"reply": i.Reply,
	}
}

func (o *Input) FromMap(values map[string]interface{}) error {
	var err error
	o.Reply, err = coerce.ToObject(values["reply"])
	if err != nil {
		return err
	}
	return nil
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	actionCtx := ctx.ActivityHost()

	input := &Input{}

	err = ctx.GetInputObject(input)
	if err != nil {
		actionCtx.Reply(nil, err)
	}

	actionCtx.Reply(input.Reply, nil)

	return true, nil
}