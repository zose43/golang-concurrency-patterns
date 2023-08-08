package main

import "context"

type JobId int
type JobType string

type Job struct {
	Descriptor *JobDescriptor
	Exec       ExecutionFn
	Args       interface{}
}

func NewJob(descriptor *JobDescriptor, exec ExecutionFn, args interface{}) Job {
	return Job{Descriptor: descriptor, Exec: exec, Args: args}
}

func (j Job) execute(ctx context.Context) Result {
	value, err := j.Exec(ctx, j.Args)
	if err != nil {
		return Result{Descriptor: j.Descriptor, Err: err}
	}
	return Result{Descriptor: j.Descriptor, Value: value}
}

type Result struct {
	Descriptor *JobDescriptor
	Value      interface{}
	Err        error
}

type JobDescriptor struct {
	Id   JobId
	Type JobType
	Meta map[string]interface{}
}

type ExecutionFn func(ctx context.Context, args interface{}) (interface{}, error)
