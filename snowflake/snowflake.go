package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cdsailing/pkg/log"
)

var Node *snowflake.Node

func init() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Error(err)
		return
	}
	Node = node
}
