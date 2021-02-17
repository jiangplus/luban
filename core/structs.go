package core

import "time"

type FlowContext struct {
	Timeout   int64
	Envs      []string
	Params    map[string]interface{}
	TaskStates map[string]*TaskState
	Runtime string
	TaskMap map[string]map[string]bool
}

type FlowSpec struct {
	Name   string
	Author string
	Desc   string
	Timeout int64
	Labels []string
	Envs    []string
	Tasks  []TaskSpec
	Params map[string]interface{}
	TaskType string `yaml:"task_type"`
	DockerImage string `yaml:"docker_image"`
}

type RangeSpec struct {
	From int
	To int
	Step int
}

type TaskSpec struct {
	Name    string
	Command string
	Envs    []string
	Deps    []string
	Inputs  []InputSpec
	Outputs []OutputSpec
	Params map[string]interface{}
	WithItems []interface{} `yaml:"with_items"`
	WithRange RangeSpec `yaml:"with_range"`
	Namegen string
	ParentTask *TaskSpec
	TaskType string `yaml:"task_type"`
	DockerImage string `yaml:"docker_image"`
	Binds []string
}

type TaskState struct {
	Name string
	Status string
	StartTime time.Time
	EndTime time.Time
	Task *TaskSpec
}

type InputSpec struct {
	S3   string
	Path string
}

type OutputSpec struct {
	S3   string
	Path string
}
