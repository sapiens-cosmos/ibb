package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ibb queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListNft())
	cmd.AddCommand(CmdShowNft())
	cmd.AddCommand(CmdLoadNft())

	cmd.AddCommand(CmdListRepay())
	cmd.AddCommand(CmdShowRepay())

	cmd.AddCommand(CmdListWithdraw())
	cmd.AddCommand(CmdShowWithdraw())

	cmd.AddCommand(CmdListUser())
	cmd.AddCommand(CmdShowUser())
	cmd.AddCommand(CmdLoadUser())
	cmd.AddCommand(CmdUserBalance())

	cmd.AddCommand(CmdListBorrow())
	cmd.AddCommand(CmdShowBorrow())

	cmd.AddCommand(CmdListDeposit())
	cmd.AddCommand(CmdShowDeposit())

	cmd.AddCommand(CmdListPool())
	cmd.AddCommand(CmdShowPool())
	cmd.AddCommand(CmdLoadPoolRest())
	cmd.AddCommand(CmdLoadPoolGrpc())

	cmd.AddCommand(CmdListLiquidation())

	return cmd
}
