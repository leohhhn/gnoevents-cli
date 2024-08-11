package main

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gnovm/pkg/gnoenv"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
)

const (
	eventsRealmPath = "gno.land/r/gnoland/events"
	devRemote       = "localhost:26657"
	plRemote        = "https://rpc.gno.land:443"
)

func main() {
	// Initialize keybase from a directory
	keybase, _ := keys.NewKeyBaseFromDir(gnoenv.HomeDir())

	// Create signer
	signer := gnoclient.SignerFromKeybase{
		Keybase:  keybase,
		Account:  "main", // Name of your keypair in keybase
		Password: "",     // Password to decrypt your keypair
		ChainID:  "dev",  // id of gno.land chain
	}

	// Initialize the RPC client
	rpc, err := rpcclient.NewHTTPClient(devRemote)
	if err != nil {
		panic(err)
	}

	// Initialize the gnoclient
	client := gnoclient.Client{
		Signer:    signer,
		RPCClient: rpc,
	}

	// Convert Gno address string to `crypto.Address`
	addr, err := crypto.AddressFromBech32("g125em6arxsnj49vx35f0n0z34putv5ty3376fg5") // your Gno address
	if err != nil {
		panic(err)
	}

	accountRes, _, err := client.QueryAccount(addr)
	if err != nil {
		panic(err)
	}

	txCfg := gnoclient.BaseTxCfg{
		GasFee:         "1000000ugnot",                // gas price
		GasWanted:      1000000000,                    // gas limit
		AccountNumber:  accountRes.GetAccountNumber(), // account ID
		SequenceNumber: accountRes.GetSequence(),      // account nonce
		Memo:           "",                            // transaction memo
	}

	msgs := make([]gnoclient.MsgCall, 0, len(events))

	for _, e := range events {
		msg := gnoclient.MsgCall{
			PkgPath:  eventsRealmPath,
			FuncName: "AddEvent",
			Args:     e.ToArgSlice(),
		}

		msgs = append(msgs, msg)
	}

	_, err = client.Call(txCfg, msgs...)
	if err != nil {
		panic(err)
	}
}
