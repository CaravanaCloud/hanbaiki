package cmd

import (
	"github.com/charmbracelet/log"
)

type ResourceType interface {
	List(data map[string]string) error
	Allocate() error
	Prepare() error
	Lease() error
	Claim() error
	Cleanup() error
	Dispose() error
}

func ResourceTypeByName(name string) ResourceType {
	switch name {
	case "aws_account":
		return AWSAccount{}
	}
	log.Printf("resource type '%s' not found", name)
	return nil
}

func runResource(cmd string, args []string, data map[string]string) {
	log.Info(cmd + " called")
	resourceTypeName := args[0]
	rtype := ResourceTypeByName(resourceTypeName)
	if rtype == nil {
		log.Info("Unknown resource type")
		return
	}

	switch cmd {
	case "allocate":
		rtype.Allocate()
	case "prepare":
		rtype.Prepare()
	case "lease":
		rtype.Lease()
	case "claim":
		rtype.Claim()
	case "cleanup":
		rtype.Cleanup()
	case "dispose":
		rtype.Dispose()
	case "list":
		rtype.List(data)
	default:
		log.Info("Unknown command")
	}

	log.Printf("Resource type[%v] cmd[%v] OK", resourceTypeName, cmd)
}
