<template>
	<div class="modal-wrapper" @click="checkClickOutside">
		<div class="modal" :class="OwnerAddress === 'in escrow' ? 'escrow' : ''">
			<video v-if="ImageUrl.split('.').pop() === 'mp4'" class="media" :src="ImageUrl" autoplay="true" loop="true" mute="true" playsinline="true" />
			<img v-else class="media" :src="ImageUrl" />
			<div class="nft">
				<div class="info">
					<div class="collection">{{ Collection }}</div>
					<div class="name">{{ Name }}</div>
				</div>
				<div class="status">Owned by {{ OwnerAddress === 'in escrow' ? OwnerAddress : OwnerAddress.slice(0, 10) + '...' + OwnerAddress.slice(-4) }}</div>
			</div>
			<div class="title">Offers</div>
			<div class="offer-table">
				<div class="table-header">
					<div class="table-cell">Loan</div>
					<div class="table-cell">Interest</div>
					<div class="table-cell">Payback</div>
					<div class="table-cell">Duration</div>
					<div v-if="loggedAddress === OwnerAddress && Array.isArray(offers) && offers.length > 0" class="table-cell"></div>
				</div>
				<div v-if="Array.isArray(offers) && offers.length > 0" class="table-rows">
					<div v-for="offer in offers" v-bind:key="offer.offerStartAt" class="table-row">
						<div class="table-cell">{{ offer.amount / 1000000 }} {{ offer.denom.toLocaleUpperCase().split('U').pop() }}</div>
						<div class="table-cell">{{ offer.interest / 1000000 }}</div>
						<div class="table-cell">{{ offer.paybackAmount / 1000000 }} {{ offer.denom.toLocaleUpperCase().split('U').pop() }}</div>
						<div class="table-cell">{{ offer.paybackDuration }} Days</div>
						<div v-if="loggedAddress === OwnerAddress" class="table-cell">
							<button @click="acceptOffer(offer)">{{ isAcceptLoading ? 'Loading...' : 'Accept' }}</button>
						</div>
					</div>
				</div>
				<div v-else class="table-rows"><div class="table-row">No Offer</div></div>
			</div>
			<div class="loan">
				<div class="input-group">
					<div class="label">Loan Value</div>
					<div class="loan-input">
						<input type="number" v-model="loanValue" />
						<select v-model="loanAsset">
							<option>ATOM</option>
							<option>AKT</option>
							<option>CRO</option>
							<option>DVPN</option>
							<option>IRIS</option>
							<option>XPRT</option>
						</select>
					</div>
				</div>
				<div class="input-group">
					<div class="label">Interest</div>
					<div class="interest-input">
						<input type="number" v-model="interest" />
						<span>%</span>
					</div>
				</div>
			</div>
			<div class="payback">
				<div class="input-group">
					<div class="label">Payback Value</div>
					<div>
						<b>{{ paybackValue }} {{ loanAsset }}</b>
					</div>
				</div>
				<div class="input-group">
					<div class="label">Payback Duration</div>
					<div><b>14 Days</b></div>
				</div>
			</div>
			<div class="cta" :class="isValidated ? 'active' : ''" @click="submit">{{ isLoading ? 'Loading...' : 'Make an Offer' }}</div>
		</div>
	</div>
</template>

<style scoped>
body {
	overflow: hidden;
}

.modal-wrapper {
	width: 100vw;
	height: 100%;
	overflow: scroll;
	position: fixed;
	top: 0;
	left: 0;
	display: flex;
	flex-direction: column;
	align-items: center;
	background: rgba(0, 0, 0, 0.7);
}

.modal {
	position: relative;
	margin-top: 60px;
	margin-bottom: 60px;
	background: rgba(0, 0, 0, 0.8);
	border: 1px solid white;
	border-radius: 10px;
	max-width: 420px;
	width: 100%;
	padding: 40px;
}

.modal.escrow {
	opacity: 0.3;
}

.media {
	border-radius: 10px;
	width: 100%;
}

.nft {
	margin-top: 12px;
	display: flex;
	justify-content: space-between;
}

.collection {
	font-size: 16px;
	color: #aaa;
}

.name {
	margin-top: 4px;
	font-size: 20px;
	color: white;
}

.status {
	font-size: 16px;
	color: white;
}

.title {
	margin-top: 24px;
	font-size: 18px;
	color: white;
}

.offer-table {
	margin-top: 6px;
	margin-bottom: 28px;
	border: 1px solid white;
	border-radius: 5px;
}

.table-header {
	display: flex;
	padding: 8px 8px 4px;
	color: rgba(255, 255, 255, 0.7);
	font-size: 14px;
}

.table-row {
	display: flex;
	justify-content: center;
	padding: 4px 8px 8px;
	border-radius: 5px;
}

.table-row:hover {
	cursor: pointer;
	background: rgba(0, 0, 0, 0.5);
}

.table-cell {
	width: 100%;
	min-width: 20px;
	display: flex;
	align-items: center;
	justify-content: flex-end;
}

.table-cell > button {
	background: transparent;
	border: 1px solid white;
	border-radius: 3px;
	color: white;
	padding: 2px 4px;
}

.table-cell > button:hover {
	background: white;
	color: black;
	cursor: pointer;
}

.loan {
	border-top: 1px solid rgba(255, 255, 255, 0.5);
	padding-top: 16px;
}

.payback {
	margin-top: 28px;
}

.input-group {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-top: 8px;
	font-size: 18px;
	height: 24px;
}

input {
	text-align: right;
	background: transparent;
	border: 1px solid white;
	padding: 2px;
	border-radius: 5px;
	color: white;
	font-size: 16px;
	width: 100px;
}

input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
	-webkit-appearance: none;
	margin: 0;
}

input[type='number'] {
	-moz-appearance: textfield; /* Firefox */
	outline: none;
}

select {
	padding: 2px;
	padding-left: 0;
	background: transparent;
	border: 1px solid white;
	border-radius: 5px;
	color: white;
	font-size: 16px;
}

.loan-input {
	display: flex;
}

.loan-input > input {
	border-top-right-radius: 0;
	border-bottom-right-radius: 0;
	border-right: unset;
}

.loan-input > select {
	background: transparent;
	border: 1px solid white;
	border-left: unset;
	border-top-left-radius: 0;
	border-bottom-left-radius: 0;
}

.interest-input > input {
	margin-right: 4px;
	width: 40px;
}

.cta {
	margin-top: 36px;
	padding: 16px;
	color: #333333;
	font-size: 20px;
	font-weight: bold;
	text-align: center;
	border: 1px solid rgba(255, 255, 255, 0.5);
	background: rgba(255, 255, 255, 0.5);
	border-radius: 10px;
}

.cta:hover {
	cursor: not-allowed;
}

.cta.active {
	background: rgba(255, 255, 255, 1);
	border: 1px solid rgba(255, 255, 255, 1);
}

.cta.active:hover {
	cursor: pointer;
}
</style>

<script>
export default {
	name: 'NftModal',
	props: ['Collection', 'Id', 'ImageUrl', 'Name', 'NftCreatorAddress', 'OwnerAddress'],
	data() {
		return {
			isLoading: false,
			isAcceptLoading: false,
			loanValue: 0,
			loanAsset: 'ATOM',
			interest: 0
		}
	},
	computed: {
		isValidated() {
			return !!this.loggedAddress
		},
		paybackValue() {
			return parseFloat((this.loanValue * this.interest) / 100 + parseFloat(this.loanValue || 0) || 0)
		},
		loggedAddress() {
			return this.$store.getters['common/wallet/address']
		},
		offers() {
			return (
				this.$store.getters['sapienscosmos.ibb.ibb/getNftOfferList']({
					params: {
						id: this.Id
					}
				})?.OfferList ?? []
			)
		}
	},
	async created() {
		await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryNftOfferList', {
			options: { all: true },
			params: {
				id: this.Id
			}
		})
	},
	methods: {
		checkClickOutside(e) {
			if (e.target === e.currentTarget) {
				this.$emit('click-outside')
			}
		},
		async acceptOffer(offer) {
			this.isAcceptLoading = true
			console.log(offer)
			const loggedAddress = this.$store.getters['common/wallet/address']
			const asset = offer.denom.toLocaleUpperCase().split('U').pop()
			const value = {
				creator: loggedAddress,
				blockHeight: 0,
				asset: asset,
				amount: offer.amount,
				denom: offer.denom
			}

			await this.$store.dispatch(`sapienscosmos.ibb.ibb/sendMsgCreateBorrow`, {
				value,
				fee: []
			})
			await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryPoolLoad', {
				options: { all: true },
				params: {}
			})
			await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryUserLoad', {
				options: { all: true },
				params: {
					id: this.$store.getters['common/wallet/address']
				}
			})
			await this.$store.dispatch('sapienscosmos.ibb.ibb/sendMsgChooseOffer', {
				value: {
					nftId: offer.nftId
				},
				fee: []
			})
			this.$emit('click-outside')
		},
		async submit() {
			if (!this.isValidated || !this.loggedAddress) {
				return null
			}
			this.isLoading = true

			const value = {
				creator: this.loggedAddress,
				denom: `u${this.loanAsset.toLocaleLowerCase()}`,
				amount: this.loanValue * 1000000,
				interest: this.interest * 1000000,
				paybackAmount: this.paybackValue * 1000000,
				paybackDuration: 14,
				nftId: parseInt(this.Id)
			}

			await this.$store.dispatch(`sapienscosmos.ibb.ibb/sendMsgCreateOffer`, {
				value,
				fee: []
			})

			this.$emit('click-outside')
		}
	}
}
</script>
