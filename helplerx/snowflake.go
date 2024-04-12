package helplerx

import (
	"github.com/bwmarrin/snowflake"
)

func GenerateSnowFlakeID(nodeID int64) (int64, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return 0, err
	}

	return node.Generate().Int64(), nil
}
