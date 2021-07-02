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

func CmdMintNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mintNft [denomID] [tokenID] [tokenNm] [tokenURI] [tokenData]",
		Short: "Broadcast message mintNft",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDenomID := string(args[0])
			argsTokenID := string(args[1])
			argsTokenNm := string(args[2])
			argsTokenURI := string(args[3])
			argsTokenData := string(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintNft(clientCtx.GetFromAddress().String(), string(argsDenomID), string(argsTokenID), string(argsTokenNm), string(argsTokenURI), string(argsTokenData))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
