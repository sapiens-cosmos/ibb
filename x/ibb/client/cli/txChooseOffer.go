package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

var _ = strconv.Itoa(0)

func CmdChooseOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chooseOffer [nftId] [offerId]",
		Short: "Broadcast message chooseOffer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsNftId, _ := strconv.ParseInt(args[0], 10, 64)
			argsOfferId, _ := strconv.ParseInt(args[1], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgChooseOffer(clientCtx.GetFromAddress().String(), int32(argsNftId), int32(argsOfferId))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
