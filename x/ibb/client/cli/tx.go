package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdAcceptOffer())

	cmd.AddCommand(CmdCreateOffer())

	cmd.AddCommand(CmdMintNft())

	cmd.AddCommand(CmdCreateNft())
	cmd.AddCommand(CmdUpdateNft())
	cmd.AddCommand(CmdDeleteNft())

	cmd.AddCommand(CmdCreateTxHistory())
	cmd.AddCommand(CmdUpdateTxHistory())
	cmd.AddCommand(CmdDeleteTxHistory())

	cmd.AddCommand(CmdCreateBorrowAccrued())
	cmd.AddCommand(CmdUpdateBorrowAccrued())
	cmd.AddCommand(CmdDeleteBorrowAccrued())

	cmd.AddCommand(CmdCreateDepositEarned())
	cmd.AddCommand(CmdUpdateDepositEarned())
	cmd.AddCommand(CmdDeleteDepositEarned())

	cmd.AddCommand(CmdCreateApr())
	cmd.AddCommand(CmdUpdateApr())
	cmd.AddCommand(CmdDeleteApr())

	cmd.AddCommand(CmdCreateRepay())
	cmd.AddCommand(CmdUpdateRepay())
	cmd.AddCommand(CmdDeleteRepay())

	cmd.AddCommand(CmdCreateWithdraw())
	cmd.AddCommand(CmdUpdateWithdraw())
	cmd.AddCommand(CmdDeleteWithdraw())

	cmd.AddCommand(CmdCreateUser())
	cmd.AddCommand(CmdDeleteUser())

	cmd.AddCommand(CmdCreateBorrow())
	cmd.AddCommand(CmdUpdateBorrow())
	cmd.AddCommand(CmdDeleteBorrow())

	cmd.AddCommand(CmdCreateDeposit())
	cmd.AddCommand(CmdUpdateDeposit())
	cmd.AddCommand(CmdDeleteDeposit())

	cmd.AddCommand(CmdCreatePool())
	cmd.AddCommand(CmdDeletePool())
	cmd.AddCommand(CmdCreateAllPool())

	cmd.AddCommand(CmdCreateMockNft())
	return cmd
}
