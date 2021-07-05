Interstellar provides **borrowing** services for any asset on a chain that supports the **IBC protocol** - both fungible tokens as well as non-fungible tokens are supported.

For fungible tokens, the protocol uses a **single-asset pool logic:**

- A depositor can deposit an asset to the pool, earning interest
- A borrower can borrow that asset from the pool, paying interest to the depositors
- A liquidator can liquidate loans once the LTV ratio has been reached

For NFTs, an **auction-based system** is used:

- The owner of an NFT can put the NFT up for a "collateral auction", in which potential lenders can make loan bids
- Once a loan bid has been accepted by the NFT owner, the NFT owner receives the collateral and puts the NFT into escrow
- Upon payback of the loan (plus interest) the NFT is released back to the original owner. If the loan is not payed back within the specified timeframe, the NFT is sent to the collateral lender

## Fungible Token Depositing & Borrowing

<img width="542" alt="image" src="https://user-images.githubusercontent.com/631758/124327920-a2b7eb00-dbc3-11eb-856b-25e670f14e6b.png">

<img width="538" alt="image" src="https://user-images.githubusercontent.com/631758/124327960-aea3ad00-dbc3-11eb-8961-59e5ad8b57ea.png">

## Non-Fungible Token Borrowing

<img width="530" alt="image" src="https://user-images.githubusercontent.com/631758/124327990-bfecb980-dbc3-11eb-895b-d039b0f3f80e.png">

## Technical Details & Current Implementation Status

The blockchain was implemented using the Cosmos SDK, while the frontend was developed using Vue. Both make use of starport.

For fungible tokens the current implementation supports:

- Depositing assets into an IBC pool to earn interest
- Borrowing from a pool (by paying interest)
- A dynamically adjusting interest rate based on the pool borrow ratio
- Paying back loans
- Claiming interest
- Reclaiming the deposit
- Liquidation by CLI command (since liquidation will mostly be done by bots anyway no UI is necessary)

For non-fungible tokens the current implementation supports:

- Depositing a NFT on the IBB chain into an escrow contract and thereby opening the bidding process
- Accepting incoming bids - containing the loan value, interest rate and payback date - from multiple wallets
- Accepting a bid and thus receiving the bidders loan
- Repaying the loan (plus interest) and thus receiving the NFT back from escrow
- Liquidation of the NFT by transferring it from the escrow to the loaner in case the original NFT owner has not payed back the loan

<img width="1264" alt="Screenshot 2021-07-03 at 06 39 19" src="https://user-images.githubusercontent.com/631758/124331909-304b0900-dbcb-11eb-8953-c1ddd32c5ef0.png">

## Outlook & Next Steps

The next major steps for Interstellar are:

- Adding a custom token to incentivise protocol usage.
    - Users will be able to stake the token and pay fees at a reduced rate.
    - Users will be able to use the custom token to provide collateral for NFTs
- IBC NFTs will be supported once the broader Cosmos ecosystem has reached consensus on a standard
