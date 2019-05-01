package lnwallet

import (
	"github.com/mraksoll4/btcutil"
	"github.com/mraksoll4/btcwallet/wallet/txrules"
	"github.com/mraksoll4/lnd/input"
)

// DefaultDustLimit is used to calculate the dust HTLC amount which will be
// send to other node during funding process.
func DefaultDustLimit() btcutil.Amount {
	return txrules.GetDustThreshold(input.P2WSHSize, txrules.DefaultRelayFeePerKb)
}
