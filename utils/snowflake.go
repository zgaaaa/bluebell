package utils

import (
	"bluebell/config"
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(config.Conf.SnowflakeConf.MachineId)
	if err != nil {
		fmt.Printf("snowflake.NewNode error: %v\n", err)
	}
}

func GetID() int64 {
	return node.Generate().Int64()
}

