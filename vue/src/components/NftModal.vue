<template>
	<div class="modal-wrapper" @click="checkClickOutside">
		<div class="modal">
			<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
			<div class="nft">
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
				<div class="status">Owned by cosmos0x2fa...ef2k</div>
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
						<b>{{ parseFloat((loanValue * interest) / 100 + parseFloat(loanValue || 0) || 0) }} {{ loanAsset }}</b>
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

.media {
	border-radius: 10px;
	width: 100%;
}

.nft {
	margin-top: 16px;
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
	font-size: 18px;
	color: white;
}

.loan {
	margin-top: 24px;
}

.payback {
	margin-top: 28px;
}

.input-group {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-top: 12px;
	font-size: 18px;
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

select {
	padding: 2px;
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
	props: [],
	data() {
		return {
			isLoading: false,
			loanValue: 0,
			loanAsset: 'ATOM',
			interest: 0
		}
	},
	computed: {
		isValidated() {
			return true
		}
	},
	methods: {
		checkClickOutside(e) {
			if (e.target === e.currentTarget) {
				this.$emit('click-outside')
			}
		},
		async submit() {
			if (!this.isValidated) {
				return null
			}
			this.isLoading = true
			this.$emit('click-outside')
		}
	}
}
</script>
