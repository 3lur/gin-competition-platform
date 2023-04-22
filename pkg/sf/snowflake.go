package sf

import (
	"competition-backend/pkg/logger"
	"github.com/bwmarrin/snowflake"
)

var sf *snowflake.Node
var err error

func init() {
	sf, err = snowflake.NewNode(1)
	if err != nil {
		logger.ErrorString("SnowFlake", "New", err.Error())
	}
}

func GenerateID() int64 {
	return sf.Generate().Int64()
}
