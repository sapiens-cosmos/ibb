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
			msgDogeCollection2 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Dodge to the Moon Collection", "cosmossdealtvlt3vm62xdz749r1dyr7rmmdxypzpsjy0", "https://i.pinimg.com/736x/4d/9c/82/4d9c826f50fd5c6916d1a28877c574a6.jpg", "Dogetronaut", "cosmosxylt3vm62x7z1dd74sdealtvyr7rmm9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgDogeCollection3 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Dodge to the Moon Collection", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "https://upload.wikimedia.org/wikipedia/en/5/5f/Original_Doge_meme.jpg", "Such WOW", "cosmosxylt3vm62x7z74sdealt1dyr7rmmdv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgDogeCollection4 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Dodge to the Moon Collection", "cosmosmmdxyltx1dyr7r7z74sdealtv3vm629rpzpsjy0", "https://napbots.com/wp-content/uploads/2021/05/dogecoin-reseaux-sociaux.jpg", "STONKS", "cosmosz74sdealt1dym62x7zpr7rmmdxyv9rplt3vsjy0", argsSelectedOffer, argsOffers)
			msgMattCollection1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Matt's K-pop Idol", clientCtx.GetFromAddress().String(), "https://thumbs.gfycat.com/CommonUglyCrossbill-mobile.mp4", "Black Pink in the House", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgMattCollection2 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Matt's K-pop Idol", "cosmosvm62x7z7lt3e4sd1dyr7rmmdxypsaltv9rpzjy0", "https://thumbs.gfycat.com/SplendidPessimisticDouglasfirbarkbeetle-mobile.mp4", "ARMY", "cosmos3vm62x7z7dy1xylttvr7rmmd4sdeal9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgMattCollection3 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Matt's K-pop Idol", "cosmos4sdealtv9t3vm62x7rpzpsjy1dyr7rmmdxylz70", "https://thumbs.gfycat.com/CloudyGregariousEsok.webp", "Always GD", "cosmosdxylt3vm67rmmsdealtv92x7z741dyrrpzpsjy0", argsSelectedOffer, argsOffers)
			msgMattCollection4 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Matt's K-pop Idol", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", "https://thumbs.gfycat.com/AdmirableMeaslyHarvestmen-mobile.mp4", "Dragonized", "cosmosvm62x7zxylt3pzps74sdealtv9r1dyr7rmmdjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", clientCtx.GetFromAddress().String(), "https://secure.img1-fg.wfcdn.com/im/02238154/compr-r85/8470/84707680/pokemon-pikachu-wall-decal.jpg", "Pika Pika", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection2 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", "cosmos3vm621dyx7z74r7rmmdxyltsdealtv9rpzpsjy0", "https://i.pinimg.com/originals/38/bb/a8/38bba86077ca44eebf9faf9220f7466b.png", "Nr 151", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection3 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", "cosmosxyl1dyr7rmmdtv9rpzlt3vm62x7z74sdeapsjy0", "https://www.pokewiki.de/images/9/96/Sugimori_006.png", "Heart Fire", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection4 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", "cosmos7rmmdxylt3vm62x1dyr7z74sdealtv9rpzpsjy0", "https://sm.ign.com/ign_de/news/p/pokemon-go/pokemon-go-trollt-spieler-mit-shiny-karpador_zwa5.png", "I'm Useless", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgPokemonCollection5 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "I wanna be the very best", "cosmos7z74sdealtv9r7rmmdxylt3vm62x1dyrpzpsjy0", "https://assets.pokemon.com/assets/cms2/img/pokedex/full/132.png", "Ditto!", "cosmos1dyr7rmmdxylt3vm62x7z74sdealtv9rpzpsjy0", argsSelectedOffer, argsOffers)
			msgArtCollection1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Avalanche", "cosmos3vm62x1dyr9r7rmmdxyltp7z74sdealtvzpsjy0", "https://mkpcdn.com/1000x/d41d8cd98f00b204e9800998ecf8427e_656731.jpg", "Bitgogh", "cosmos3vm62ltv9rpzps7z74sdea1dyr7rmxxyltmdjy0", argsSelectedOffer, argsOffers)
			msgArtCollection2 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Avalanche", "cosmosxyl1dyr7rmmdtv9rpzlt3vm62x7z74sdeapsjy0", "https://thumbor.forbes.com/thumbor/960x0/https%3A%2F%2Fblogs-images.forbes.com%2Frachelwolfson%2Ffiles%2F2019%2F01%2FFork-and-flip.jpg", "Bull vs Bear", "cosmos9rpzps7z74sdea1dyr73vm62xxyltltvrmmdjy0", argsSelectedOffer, argsOffers)
			msgArtCollection3 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Avalanche", "cosmos1dyr9r7rmmdxylt3vm62xp7z74sdealtvzpsjy0", "https://news.bitcoin.com/wp-content/uploads/2016/05/desktop_image.png", "Glory", "cosmos3ltv9rpzps7ea1dyr7rmmdjyvz74sdm62xxylt0", argsSelectedOffer, argsOffers)
			msgArtCollection4 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "Avalanche", "cosmosyr9r7rmmdxyltp7z74sdealtv3vm62x1dzpsjy0", "https://streetartnews.net/wp-content/uploads/2018/02/IMG_4630-copy.jpg", "RIP Banking System", "cosmos3vm62xxyltltv9rpzps7z74sdea1dyr7rmmdjy0", argsSelectedOffer, argsOffers)
			msgCryptoToysCollection1 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "CRYPTO CRYPTO TOYS", "cosmos4sdealtv3vm62x1dyr9r7rmmdxyltp7z7zpsjy0", "https://lh3.googleusercontent.com/ybTcSIzVF4NGPpggzpHDWd3MChsOmkLHBPmCzCEqS9xMpuapKufxonuuSYHPv6AQg58F4RgKKnmaIFtIFrxAbuOnB0i3qDDj-dOShqo=w600", "ARMY GUY", "cosmos3vm62xxyltltv9rpzps7z74sdea1dyr7rmmdjy0", argsSelectedOffer, argsOffers)
			msgCryptoToysCollection2 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "CRYPTO CRYPTO TOYS", "cosmos4sdealtv3vm62x1dyr9r7rmmdxyltp7z7zpsjy0", "https://lh3.googleusercontent.com/72r_O9twuaL0er_fbk9x7lov_JXylkIF3zhrSER4ZxBdG5eaqXQr43Xc3Y29a1hBfxAbDOzeqpg6IJnTbz-mPE8NnDjnr_Ya1inh=w600", "PILOT 33", "cosmos9rpzps7z74sd3vm62xxyltltvea1dyr7rmmdjy0", argsSelectedOffer, argsOffers)
			msgCryptoToysCollection3 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "CRYPTO CRYPTO TOYS", "cosmos4sdealtv3vm62x1dyr9r7rmmdxyltp7z7zpsjy0", "https://lh3.googleusercontent.com/nXvkPlMl0GU-Lo_eC9epIMDQI1BaAOiELGcpDRVztNN5rVYEszYRk38lagvfRYR_R5q9e6Jh-Beq_Uwqy0sgTYl41Xj6UHlF9Smo1Q=w600", "SECRET HUMAN DX", "cosmos2xxyltltv9rpzps7z74sdea3vm61dyr7rmmdjy0", argsSelectedOffer, argsOffers)
			msgCryptoToysCollection4 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "CRYPTO CRYPTO TOYS", "cosmos4sdealtv3vm62x1dyr9r7rmmdxyltp7z7zpsjy0", "https://lh3.googleusercontent.com/CR6VJ9Q1hw__qmnw79poCF7LsXtW1eIznlnuiQLc5BW03c16LZhd2NooEtGGUAydSgONK4OAO3014jWVkTZMKXrcp7oJ_RsaiQTi=w600", "OK GUYS WIN", "cosmos3vm62xxyltltv9rpzps7z74sdea1dyr7rmmdjy0", argsSelectedOffer, argsOffers)
			msgCryptoToysCollection5 := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), "CRYPTO CRYPTO TOYS", "cosmos4sdealtv3vm62x1dyr9r7rmmdxyltp7z7zpsjy0", "https://lh3.googleusercontent.com/ZKDwEDQWN1GX8XgcT23FPzf2EjOtHTcmpVXqDV81tksUAB1UzwgWn14BYRQ6OqZJEoYHdmTmc0o1q55twA0M_TwcXxZwN4dwargpq40=w600", "MOTHMAN NIGHT GUARD", "cosmos3vm62xxyltltv9rpzps7z74sdea1dyr7rmmdjy0", argsSelectedOffer, argsOffers)

			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgDogeCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgDogeCollection2)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgDogeCollection3)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgDogeCollection4)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgMattCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgMattCollection2)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgMattCollection3)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgMattCollection4)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection2)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection3)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection4)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgPokemonCollection5)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgArtCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgArtCollection2)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgArtCollection3)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgArtCollection4)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgArtCollection4)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCryptoToysCollection1)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCryptoToysCollection2)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCryptoToysCollection3)
			tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCryptoToysCollection4)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCryptoToysCollection5)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
