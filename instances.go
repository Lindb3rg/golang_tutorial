package main

import (
	"fmt"
)

type InstanceState string

const (
	InstanceStateRunning InstanceState = "running"
	InstanceStateStopped InstanceState = "stopped"
)

type Instance struct {
	Name  string
	State InstanceState
}

func LaunchInstance(name string) (*Instance, error) {
	return &Instance{Name: name, State: InstanceStateRunning}, nil
}

func StopInstance(instance *Instance) error {
	instance.State = InstanceStateStopped
	return nil
}

func ObserveInstance(instance *Instance) {
	fmt.Printf("Instance %s is %s\n", instance.Name, instance.State)
}

func main() {
	instance, err := LaunchInstance("my-instance")
	if err != nil {
		fmt.Println(err)
		return
	}
	ObserveInstance(instance)

	err = StopInstance(instance)
	if err != nil {
		fmt.Println(err)
		return
	}
	ObserveInstance(instance)
}
