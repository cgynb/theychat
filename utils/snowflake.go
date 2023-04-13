package utils

import (
	"github.com/bwmarrin/snowflake"
)

func GenGroupId() (gid string) {
	node, _ := snowflake.NewNode(1)
	gid = node.Generate().Base64()
	return
}
