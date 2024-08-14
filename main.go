package main

import (
	"fmt"
	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	"github.com/gnolang/gno/gnovm/pkg/gnoenv"
	rpcclient "github.com/gnolang/gno/tm2/pkg/bft/rpc/client"
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

	// get the address with the given name in the signer
	signerInfo, err := signer.Info()
	if err != nil {
		panic(err)
	}

	// Get address from signer
	addr := signerInfo.GetAddress()
	fmt.Println(addr)

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

	accountRes, _, err := client.QueryAccount(addr)
	if err != nil {
		panic(err)
	}

	txCfg := gnoclient.BaseTxCfg{
		GasFee:         "1000000ugnot",                // gas price
		GasWanted:      100000000,                     // gas limit
		AccountNumber:  accountRes.GetAccountNumber(), // account ID
		SequenceNumber: accountRes.GetSequence(),      // account nonce
		Memo:           "",                            // transaction memo
	}

	msgs := make([]gnoclient.MsgCall, 0, len(events))

	fmt.Printf("Adding %d events\n", len(events))

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
