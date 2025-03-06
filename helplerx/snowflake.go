package helplerx

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

type SnowflakeClient struct {
	Node *snowflake.Node
	mu   sync.Mutex
}

func NewSnowflakeClient(nodeID int64) (*SnowflakeClient, error) {
	client, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}

	return &SnowflakeClient{Node: client}, nil
}

func (sc *SnowflakeClient) GenerateSnowFlakeID() int64 {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	return sc.Node.Generate().Int64()
}
