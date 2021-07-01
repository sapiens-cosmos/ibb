package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func CmdCreateNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-nft [collection] [ownerAddress] [imageUrl] [name] [nftCreatorAddress]",
		Short: "Create a new nft",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsCollection, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsOwnerAddress, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsImageUrl, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsName, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsNftCreatorAddress, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argsSelectedOffer := types.Offer{}
			argsOffers := []*types.Offer{}

			msg := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), argsCollection, argsOwnerAddress, argsImageUrl, argsName, argsNftCreatorAddress, argsSelectedOffer, argsOffers)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-nft [id] [collection] [ownerAddress] [imageUrl] [name]",
		Short: "Update a nft",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsCollection, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsOwnerAddress, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsImageUrl, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsName, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNft(clientCtx.GetFromAddress().String(), id, argsCollection, argsOwnerAddress, argsImageUrl, argsName)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-nft [id]",
		Short: "Delete a nft by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteNft(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreateMockNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-all-nft",
		Short: "Create a set of mock nfts for basic project setup",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argsSelectedOffer := types.Offer{}
			argsOffers := []*types.Offer{}
			msgDogeCollection1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Dodge to the Moon Collection", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "https://images.news18.com/ibnlive/uploads/2021/02/1612757177_untitled-design.jpg", "Who let the Doge out", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgMattCollection1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Matt's Favorite Girl", clientCtx.GetFromAddress().String(), "https://thumbs.gfycat.com/GiddyEvilAmazondolphin-mobile.mp4", "Body Roll", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgMattCollection2 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Matt's Favorite Girl", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "https://thumbs.gfycat.com/DeliriousPossibleLeafcutterant-mobile.mp4", "Wink", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", clientCtx.GetFromAddress().String(), "https://secure.img1-fg.wfcdn.com/im/02238154/compr-r85/8470/84707680/pokemon-pikachu-wall-decal.jpg", "Pika Pika", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection2 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "https://i.pinimg.com/originals/38/bb/a8/38bba86077ca44eebf9faf9220f7466b.png", "Nr 151", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection3 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "https://sm.ign.com/ign_de/news/p/pokemon-go/pokemon-go-trollt-spieler-mit-shiny-karpador_zwa5.png", "I'm Useless", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)

			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgDogeCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgMattCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgMattCollection2)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection2)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection3)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
