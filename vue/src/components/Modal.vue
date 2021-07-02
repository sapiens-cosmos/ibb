<template>
	<div class="modal-wrapper" @click="checkClickOutside">
		<div class="modal">
			<div class="asset">
				<img class="asset-icon" src="@/assets/images/icons/atom.png" />
				<div class="asset-name">{{ Asset }}</div>
			</div>
			<div v-if="type === 'Deposit' || type === 'Withdraw'" class="modes">
				<div @click="changeType('Deposit')" class="mode" :class="type === 'Deposit' ? 'active-blue' : ''">Deposit</div>
				<div @click="changeType('Withdraw')" class="mode" :class="type === 'Withdraw' ? 'active-orange' : ''">Withdraw</div>
			</div>
			<div v-else-if="type === 'Borrow' || type === 'Repay'" class="modes">
				<div @click="changeType('Borrow')" class="mode" :class="type === 'Borrow' ? 'active-orange' : ''">Borrow</div>
				<div @click="changeType('Repay')" class="mode" :class="type === 'Repay' ? 'active-blue' : ''">Repay</div>
			</div>
			<div class="body">
				<div class="amount">
					<div class="title">
						<div class="property">{{ type }} Amount</div>
						<div v-if="type === 'Deposit'" class="value">
							Wallet Balance: {{ Number.isNaN(parseFloat(AssetBalance)) ? 0 : parseFloat(AssetBalance) / 1000000 }} {{ Asset }}
						</div>
						<div v-if="type === 'Withdraw'" class="value">
							Deposit Balance: {{ Number.isNaN(parseFloat(AssetDeposit)) ? 0 : parseFloat(AssetDeposit) / 1000000 }} {{ Asset }}
						</div>
						<div v-if="type === 'Borrow'" class="value">Borrow Limit: {{ currentBollowLimitAssetAmount.toFixed(6) }} {{ Asset }}</div>
						<div v-if="type === 'Repay'" class="value">
							Repay Balance : {{ Number.isNaN(parseFloat(AssetBorrow)) ? 0 : parseFloat(AssetBorrow) / 1000000 }} + {{ BorrowAccrued / 1000000 }} {{ Asset }}
						</div>
					</div>
					<div class="content">
						<div class="input-wrapper">
							<div class="input">
								<div class="balance">
									<input ref="balance" type="number" v-model="balance" />
									<div class="denom">{{ Asset }}</div>
								</div>
								<div class="dollar">${{ Number.isNaN(parseFloat(balance)) ? 0 : (parseFloat(balance) * AssetPrice) / 1000000 }}</div>
							</div>
							<div class="max-button" @click="makeBalanceMax">Max</div>
						</div>
					</div>
					<div class="content">
						<div class="slider">
							<input ref="range" type="range" min="0" max="100" v-model="balanceRange" />
							<output ref="indicator" class="range-indicator" :class="type === 'Deposit' || type === 'Repay' ? 'blue' : 'orange'"></output>
							<div class="slider-percentages">
								<div class="slider-percentage">0%</div>
								<div class="slider-percentage">25%</div>
								<div class="slider-percentage">50%</div>
								<div class="slider-percentage">75%</div>
								<div class="slider-percentage">100%</div>
							</div>
						</div>
					</div>
				</div>
				<div class="stats">
					<div class="title">Stats</div>
					<div class="content">
						<div class="property">APY</div>
						<div class="value">{{ type === 'Deposit' || type === 'Withdraw' ? DepositApy / 10000 : BorrowApy / 10000 }}%</div>
					</div>
					<div class="content">
						<div class="property">Balance</div>
						<div class="value">
							{{ type === 'Deposit' || type === 'Withdraw' ? (AssetDeposit ? AssetDeposit / 1000000 : 0) : AssetBorrow ? AssetBorrow / 1000000 : 0 }}
							{{ Asset }}
						</div>
					</div>
				</div>
				<div v-if="type === 'Deposit'" class="collateral">
					<div class="title">Collateral</div>
					<div class="content">
						<div class="property">Collateral Factor</div>
						<div class="value">{{ CollatoralFactor }}%</div>
					</div>
					<div class="content">
						<div class="property">Used as Collateral</div>
						<div class="value">{{ Collatoral ? 'Yes' : 'No' }}</div>
					</div>
				</div>
				<div v-if="type !== 'Deposit' || (type === 'Deposit' && Collatoral)" class="borrow-limit">
					<div class="title">Borrow Limit</div>
					<div class="content">
						<div class="property">Your Limit</div>
						<div class="value">${{ parseFloat(currentBollowLimit).toFixed(2) }} -> ${{ nextBollowLimit.toFixed(2) }}</div>
					</div>
					<div class="content">
						<div class="property">Limit Used</div>
						<div class="value">
							{{ (((parseFloat(borrowLimitAsCollateral) - parseFloat(currentBollowLimit)) / parseFloat(borrowLimitAsCollateral)) * 100).toFixed(2) }}% ->
							{{ (((parseFloat(borrowLimitAsCollateral) - parseFloat(nextBollowLimit)) / parseFloat(borrowLimitAsCollateral)) * 100).toFixed(2) }}%
						</div>
					</div>
				</div>
				<div class="cta" :class="isValidated ? 'active' : ''" @click="submit">{{ isLoading ? 'Loading...' : type }}</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.modal-wrapper {
	width: 100vw;
	min-height: 100%;
	position: absolute;
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
	max-width: 500px;
	width: 100%;
}

.asset {
	position: absolute;
	top: -40px;
	left: 50%;
	transform: translateX(-50%);
}

.asset-icon {
	width: 80px;
	height: 80px;
}

.asset-name {
	position: absolute;
	top: 70px;
	transform: translateX(-50%);
	background: black;
	left: 50%;
	padding: 8px 8px 6px;
	font-size: 16px;
	font-weight: bold;
	text-align: center;
	border: 1px solid white;
	border-radius: 5px;
}

.modes {
	display: flex;
	justify-content: space-between;
}

.mode {
	width: 100%;
	padding: 22px 20px 20px;
	font-size: 28px;
	font-weight: bold;
	text-align: center;
	color: #ccc;
	border-bottom: 1px solid #ccc;
}

.mode.active-blue {
	color: rgba(9, 67, 121, 1);
	border-bottom: 1px solid rgba(9, 67, 121, 1);
}

.mode.active-orange {
	color: rgba(255, 139, 0, 1);
	border-bottom: 1px solid rgba(255, 139, 0, 1);
}

.mode:hover {
	cursor: pointer;
}

.body {
	padding: 0 36px 36px;
}

.input-wrapper {
	border: 1px solid white;
	border-radius: 10px;
	padding: 16px;
	width: 100%;
	display: flex;
	align-items: center;
	position: relative;
}

.input {
	width: 100%;
}

input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
	-webkit-appearance: none;
	margin: 0;
}

input[type='number'] {
	-moz-appearance: textfield; /* Firefox */
}

.range-indicator {
	display: none;
	color: white;
	font-size: 14px;
	padding: 5px 12px;
	position: absolute;
	border-radius: 4px;
	left: 50%;
	transform: translateX(-50%);
}

.range-indicator.blue {
	background: rgba(9, 67, 121, 1);
}

.range-indicator.orange {
	background: rgba(255, 139, 0, 1);
}

.balance {
	display: flex;
	align-items: center;
}

.balance > input {
	width: 100%;
	height: 100%;
	margin-right: 4px;
	margin-bottom: 1px;
	background: transparent;
	border: unset;
	color: white;
	font-size: 18px;
	text-align: right;
}

.balance > input:focus-visible {
	outline: unset;
}

.denom {
	font-size: 19px;
}

.dollar {
	margin-top: 4px;
	font-size: 16px;
	text-align: right;
}

.max-button {
	padding-left: 16px;
	font-size: 15px;
}

.max-button:hover {
	cursor: pointer;
}

.slider {
	margin-top: 4px;
	position: relative;
	width: 100%;
}

.slider > input {
	width: 100%;
}

.slider > input::-webkit-slider-thumb {
	appearance: none;
	border-radius: 50%;
	background: blue;
	cursor: pointer;
}

.slider > input::-webkit-slider-thumb:hover {
	background: red;
}

.slider > input:active::-webkit-slider-thumb {
	background: yellow;
}

.slider > input::-moz-range-thumb {
	appearance: none;
	border-radius: 50%;
	background: blue;
	cursor: pointer;
}

.slider > input::-moz-range-thumb:hover {
	background: red;
}

.slider > input:active::-moz-range-thumb {
	background: yellow;
}

.slider > input:focus {
	background: black;
}

.slider-percentages {
	display: flex;
	justify-content: space-between;
}

.slider-percentage {
	font-size: 14px;
	font-weight: bold;
}

.title {
	margin-top: 36px;
	display: flex;
	justify-content: space-between;
	font-size: 16px;
	font-weight: bold;
	color: rgba(255, 255, 255, 0.7);
}

.content {
	margin-top: 8px;
	display: flex;
	font-size: 18px;
	justify-content: space-between;
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
	name: 'Modal',
	props: [
		'initialType',
		'Asset',
		'AssetBalance',
		'AssetPrice',
		'AssetDeposit',
		'Collatoral',
		'AssetBorrow',
		'BorrowApy',
		'BorrowAccrued',
		'DepositEarned',
		'CollatoralFactor',
		'DepositApy',
		'Liquidity'
	],
	data() {
		return {
			type: this.initialType,
			balance: '',
			balanceRange: 0,
			isLoading: false
		}
	},
	computed: {
		assetPoolsWithUser() {
			const loggedAddress = this.$store.getters['common/wallet/address']
			if (!loggedAddress) {
				return 0
			}

			const userAssets = loggedAddress
				? this.$store.getters['sapienscosmos.ibb.ibb/getUserLoad']({
						params: {
							id: loggedAddress
						}
				  })?.LoadUserResponse ?? []
				: []
			const assetPools =
				this.$store.getters['sapienscosmos.ibb.ibb/getPoolLoad']({
					params: {}
				})?.LoadPoolResponse ?? []
			return assetPools.map((pool, index) => ({
				...pool,
				...userAssets[index]
			}))
		},
		borrowLimitAsCollateral() {
			return this.assetPoolsWithUser.reduce(
				(acc, userAsset) => (acc += ((((userAsset.AssetDeposit / 1000000) * userAsset.AssetPrice) / 1000000) * userAsset.CollatoralFactor) / 100),
				0
			)
		},
		currentBollowLimit() {
			return (
				this.borrowLimitAsCollateral -
				this.assetPoolsWithUser.reduce((acc, userAsset) => (acc += ((userAsset.AssetBorrow / 1000000) * userAsset.AssetPrice) / 1000000), 0)
			)
		},
		currentBollowLimitAssetAmount() {
			return (this.currentBollowLimit / (this.AssetPrice || 1)) * 1000000
		},
		nextBollowLimit() {
			const newLimit = (parseFloat(this.balance || 0) * (this.AssetPrice || 0)) / 1000000
			switch (this.type) {
				case 'Deposit': {
					return this.Collatoral ? this.currentBollowLimit + newLimit : this.currentBollowLimit
				}
				case 'Withdraw': {
					return this.Collatoral ? this.currentBollowLimit - newLimit : this.currentBollowLimit
				}
				case 'Borrow': {
					return this.currentBollowLimit - newLimit
				}
				case 'Repay': {
					return this.currentBollowLimit + newLimit
				}
				default:
					return this.currentBollowLimit
			}
		},
		isValidated() {
			const balance = parseFloat(this.balance || 0) * 1000000
			switch (this.type) {
				case 'Deposit': {
					const walletBalance = parseFloat(this.AssetBalance)
					return balance > 0 && walletBalance >= balance
				}
				case 'Withdraw': {
					const depositBalance = parseFloat(this.AssetDeposit)
					return balance > 0 && depositBalance >= balance
				}
				case 'Borrow': {
					const borrowLimit = parseFloat(this.currentBollowLimitAssetAmount * 1000000)
					return balance > 0 && borrowLimit >= balance
				}
				case 'Repay': {
					const borrowBalance = parseFloat(this.AssetBorrow)
					return balance > 0 && borrowBalance + this.BorrowAccrued >= balance
				}
				default:
					return false
			}
		}
	},
	async mounted() {
		this.$refs.balance.focus()
		this.$refs.balance.addEventListener('input', () => {
			this.setRangeByBalance()
			const indicator = this.$refs.indicator
			indicator.style.display = `none`
		})
		this.$refs.range.addEventListener('input', () => {
			this.setBalanceByRange()
			this.setIndicator()
		})
	},
	methods: {
		makeBalanceMax() {
			const rangingBalanceByType =
				this.type === 'Deposit'
					? this.AssetBalance
					: this.type === 'Withdraw'
					? this.AssetDeposit
					: this.type === 'Borrow'
					? this.currentBollowLimitAssetAmount * 1000000
					: this.AssetBorrow + this.BorrowAccrued
			this.balance = parseFloat(rangingBalanceByType || 0) / 1000000
			this.setRangeByBalance()
		},
		setIndicator(rangeValue) {
			const range = this.$refs.range
			const indicator = this.$refs.indicator
			const val = rangeValue ? rangeValue : range.value
			const min = range.min ? range.min : 0
			const max = range.max ? range.max : 100
			const newVal = Number(((val - min) * 100) / (max - min))
			indicator.innerHTML = `${val}%`
			indicator.style.display = `block`
			indicator.style.left = `calc(${newVal}% + (${8 - newVal * 0.15}px))`
		},
		setRangeByBalance() {
			const rangingBalanceByType =
				this.type === 'Deposit'
					? this.AssetBalance
					: this.type === 'Withdraw'
					? this.AssetDeposit
					: this.type === 'Borrow'
					? this.currentBollowLimitAssetAmount * 1000000
					: this.AssetBorrow + this.BorrowAccrued
			this.balanceRange = ((this.balance * 1000000) / parseFloat(rangingBalanceByType || 0)) * 100
		},
		setBalanceByRange() {
			const rangingBalanceByType =
				this.type === 'Deposit'
					? this.AssetBalance
					: this.type === 'Withdraw'
					? this.AssetDeposit
					: this.type === 'Borrow'
					? this.currentBollowLimitAssetAmount * 1000000
					: this.AssetBorrow + +this.BorrowAccrued
			this.balance = ((this.balanceRange / 100) * parseFloat(rangingBalanceByType || 0)) / 1000000
		},
		checkClickOutside(e) {
			if (e.target === e.currentTarget) {
				this.$emit('click-outside')
			}
		},
		async submit() {
			if (!this.isValidated) {
				return null
			}
			const loggedAddress = this.$store.getters['common/wallet/address']
			const asset = this.Asset.toLocaleLowerCase()
			const value = {
				creator: loggedAddress,
				blockHeight: 0,
				asset: this.Asset,
				amount: parseFloat(this.balance) * 1000000,
				denom: `u${asset}`
			}

			this.isLoading = true
			await this.$store.dispatch(`sapienscosmos.ibb.ibb/sendMsgCreate${this.type}`, {
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
			this.$emit('click-outside')
		},
		changeType(type) {
			this.type = type
		}
	}
}
</script>
