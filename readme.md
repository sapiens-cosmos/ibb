# Interstellar

**Interstellar** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

Interstellar provides **borrowing** services for any asset on a chain that supports the **IBC protocol** - both fungible tokens as well as non-fungible tokens are supported.

For fungible tokens, the protocol uses a **single-asset pool logic:**

- A depositor can deposit an asset to the pool, earning interest
- A borrower can borrow that asset from the pool, paying interest to the depositors
- A liquidator can liquidate loans once the LTV ratio has been reached

For NFTs, an **auction-based system** is used:

- The owner of an NFT can put the NFT up for a "collateral auction", in which potential lenders can make loan bids
- Once a loan bid has been accepted by the NFT owner, the NFT owner receives the collateral and puts the NFT into escrow
- Upon payback of the loan (plus interest) the NFT is released back to the original owner. If the loan is not payed back within the specified timeframe, the NFT is sent to the collateral lender

## Get started

```
starport serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

## Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)
