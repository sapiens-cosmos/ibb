package cli

import (
	"strconv"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

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
			argsAmount, _ := cast.ToInt32E(args[1])
			argsPaybackAmount, _ := cast.ToInt32E(args[2])
			argsPaybackDuration, _ := cast.ToInt32E(args[3])
			argsOfferStartAt, _ := cast.ToInt64E(args[4])
			argsId, _ := strconv.ParseUint(args[5], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOffer(clientCtx.GetFromAddress().String(), string(argsDenom), argsAmount, argsPaybackAmount, argsPaybackDuration, argsOfferStartAt, argsId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
