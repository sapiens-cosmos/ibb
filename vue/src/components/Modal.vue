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
						<div v-if="type === 'Deposit'" class="value">Wallet Balance: {{ parseInt(AssetBalance) / 1000000 }} {{ Asset }}</div>
						<div v-if="type === 'Withdraw'" class="value">Deposit Balance: {{ parseInt(AssetDeposit) / 1000000 }} {{ Asset }}</div>
					</div>
					<div class="content">
						<div class="input-wrapper">
							<div class="input">
								<div class="balance">
									<input ref="balance" type="number" v-model="balance" />
									<div class="denom">{{ Asset }}</div>
								</div>
								<div class="dollar">${{ (parseFloat(Number.isNaN(parseFloat(balance)) ? 0 : balance) * AssetPrice) / 1000000 }}</div>
							</div>
							<div class="max-button">Max</div>
						</div>
					</div>
					<div class="content">
						<div class="slider">
							<input ref="range" type="range" min="0" max="100" value="0" />
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
						<div class="value">{{ type === 'Deposit' || type === 'Withdraw' ? AssetDeposit / 1000000 : AssetBorrow / 1000000 }} {{ Asset }}</div>
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
				<div v-else class="borrow-limit">
					<div class="title">Borrow Limit</div>
					<div class="content">
						<div class="property">Your Limit</div>
						<div class="value">$0 -> $402.54</div>
					</div>
					<div class="content">
						<div class="property">Limit Used</div>
						<div class="value">0% -> 0%</div>
					</div>
				</div>
				<div class="cta" @click="submit">{{ isLoading ? 'Loading...' : type }}</div>
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
	border: 1px solid white;
	background: white;
	border-radius: 10px;
}

.cta:hover {
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
		'CollatoralFactor',
		'DepositApy',
		'Liquidity'
	],
	data() {
		return {
			type: this.initialType,
			balance: '',
			isLoading: false
		}
	},
	async mounted() {
		this.$refs.balance.focus()
		this.$refs.range.addEventListener('input', () => {
			this.setRange()
		})
		this.setRange()
	},
	methods: {
		setRange() {
			const range = this.$refs.range
			const indicator = this.$refs.indicator
			const val = range.value
			const min = range.min ? range.min : 0
			const max = range.max ? range.max : 100
			const newVal = Number(((val - min) * 100) / (max - min))
			indicator.innerHTML = `${val}%`

			// Sorta magic numbers based on size of the native UI thumb
			indicator.style.left = `calc(${newVal}% + (${8 - newVal * 0.15}px))`
		},
		checkClickOutside(e) {
			if (e.target === e.currentTarget) {
				this.$emit('click-outside')
			}
		},
		async submit() {
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
