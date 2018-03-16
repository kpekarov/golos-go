package client

import (
	"errors"
	// RPC
	"github.com/asuleymanov/golos-go/apis/database"
	"github.com/asuleymanov/golos-go/apis/follow"
	"github.com/asuleymanov/golos-go/apis/market"
	"github.com/asuleymanov/golos-go/apis/networkbroadcast"
	"github.com/asuleymanov/golos-go/transactions"
	"github.com/asuleymanov/golos-go/transports"
	"github.com/asuleymanov/golos-go/transports/websocket"
	"github.com/asuleymanov/golos-go/types"
)

var KeyList = make(map[string]Keys)

func InitPrivateKey(username string, typeOfKey types.OpType, privKey string) {
	var keys Keys

	switch typeOfKey {
	case "active":
		keys.AKey = privKey
	case "owner":
		keys.OKey = privKey
	case "posting":
		keys.PKey = privKey
	case "memo":
		keys.MKey = privKey
	}

	InitPrivateKeys(username, keys)
}

func InitPrivateKeys(username string, keys Keys) {
	KeyList[username] = keys
}

// Client can be used to access Golos remote APIs.
// There is a public field for every Golos API available,
// e.g. Client.Database corresponds to database_api.
type Client struct {
	cc transports.CallCloser

	// Database represents database_api.
	Database *database.API

	// Follow represents follow_api.
	Follow *follow.API

	// Follow represents market_history_api.
	Market *market.API

	// NetworkBroadcast represents network_broadcast_api.
	NetworkBroadcast *networkbroadcast.API

	//Chain Id
	Chain *transactions.Chain
}

// NewClient creates a new RPC client that use the given CallCloser internally.
// Initialize only server present API. Absent API initialized as nil value.
func NewClient(url []string, chain string) (*Client, error) {
	call, err := initclient(url)
	if err != nil {
		return nil, err
	}
	client := &Client{cc: call}

	client.Database = database.NewAPI(client.cc)

	client.Follow, err = follow.NewAPI(client.cc)
	if err != nil {
		client.Follow = nil
	}

	client.Market, err = market.NewAPI(client.cc)
	if err != nil {
		client.Market = nil
	}

	client.NetworkBroadcast, err = networkbroadcast.NewAPI(client.cc)
	if err != nil {
		client.NetworkBroadcast = nil
	}

	client.Chain, err = initChainId(chain)
	if err != nil {
		client.Chain = transactions.GolosChain
	}

	return client, nil
}

// Close should be used to close the client when no longer needed.
// It simply calls Close() on the underlying CallCloser.
func (client *Client) Close() error {
	return client.cc.Close()
}

func initclient(url []string) (*websocket.Transport, error) {
	// Инициализация Websocket
	t, err := websocket.NewTransport(url)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func initChainId(str string) (*transactions.Chain, error) {
	var ChainId transactions.Chain
	// Определяем ChainId
	switch str {
	case "golos":
		ChainId = *transactions.GolosChain
	case "test":
		ChainId = *transactions.TestChain
	default:
		return nil, errors.New("Chain not found")
	}
	return &ChainId, nil
}
