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

func CmdCreateOffer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createOffer [denom] [amount] [paybackAmount] [paybackDuration] [offerStartAt] [id]",
		Short: "Broadcast message createOffer",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDenom := string(args[0])
			argsAmount := string(args[1])
			argsPaybackAmount := string(args[2])
			argsPaybackDuration := string(args[3])
			argsOfferStartAt := string(args[4])
			argsId := string(args[5])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOffer(clientCtx.GetFromAddress().String(), string(argsDenom), string(argsAmount), string(argsPaybackAmount), string(argsPaybackDuration), string(argsOfferStartAt), string(argsId))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
