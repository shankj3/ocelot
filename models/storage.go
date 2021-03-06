package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
	//pb "github.com/shankj3/ocelot/admin/models"
)

var (
	HOOKHANDLER_VALIDATION = "pre-build-validation"
	CHECKMARK              = "\u2713"
	FAILED                 = "\u2717"
)

type BuildOutput struct {
	BuildId  int64  `json:"buildId,omitempty"`
	Output   []byte `json:"output,omitempty"`
	OutputId int64  `json:"outputId,omitempty"` // generated by postgres
}

func (o *BuildOutput) Equals(n *BuildOutput) bool {
	if o.BuildId != n.BuildId || bytes.Compare(o.Output, n.Output) != 0 || o.OutputId != n.OutputId {
		return false
	}
	return true
}

func (o *BuildOutput) Validate() error {
	if o.Output == nil {
		return NewErrInvalid("no build output to store")
	}
	if o.BuildId == 0 {
		return NewErrInvalid("build id required")
	}
	return nil
}

//this is store inside of build_stage_details
type StageResult struct {
	BuildId       int64 //foreign key on build_summary table
	StageResultId int64 //generated by postgres
	Stage         string
	Status        int
	Error         string
	Messages      []string
	StartTime     time.Time
	StageDuration float64
}

func (r *StageResult) Validate() error {
	if len(r.Stage) == 0 {
		return NewErrInvalid("result stage must be set")
	}
	return nil
}

// errors for validation
type ErrInvalid struct {
	message string
}

func NewErrInvalid(message string) *ErrInvalid {
	return &ErrInvalid{
		message: message,
	}
}
func (e *ErrInvalid) Error() string {
	return e.message
}

//JsonStringArray exists because list of messages belonging to stages is stored as json in db. To get it to
//unparse from DB correctly, we have to implement custom Value() + Scan(). Check postgres.go for usage
type JsonStringArray []string

func (f JsonStringArray) Value() (driver.Value, error) {
	j, err := json.Marshal(f)
	return j, err
}

func (f *JsonStringArray) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("unable to cast source to []byte")
	}
	if err := json.Unmarshal(source, f); err != nil {
		return err
	}
	return nil
}

func NewMap() JsonStringMap {
	return make(map[string]string)
}

type JsonStringMap map[string]string

func (f JsonStringMap) Value() (driver.Value, error) {
	j, err := json.Marshal(f)
	return j, err
}

func (f JsonStringMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("unable to cast source to []byte")
	}
	if err := json.Unmarshal(source, f); err != nil {
		return err
	}
	return nil
}