<template>
	<div class="wallet-wrapper">
		<div v-if="loggedAddress">{{ loggedAddress }}</div>
		<button v-else @click="connectWallet">Connect Wallet</button>
		<div v-if="isSigning" class="sign-up-form">
			<div class="input-group">
				<label>Name:</label>
				<input type="text" v-model="name" />
			</div>
			<div class="input-group">
				<label>Password:</label>
				<input type="password" v-model="password" />
			</div>
			<div class="input-group">
				<label>Mnemonic:</label>
				<textarea v-model="mnemonic"></textarea>
			</div>
			<button @click="createWallet">Create Wallet</button>
		</div>
	</div>
</template>

<style scoped>
.wallet-wrapper {
	position: relative;
}

.sign-up-form {
	position: absolute;
	right: 0;
	top: 22px;
	background: white;
	padding: 16px;
	display: flex;
	flex-direction: column;
}

.input-group {
	display: flex;
}

.input-group:not(:first-child) {
	margin-top: 8px;
}

.input-group > label {
	width: 80px;
	color: black;
	margin-right: 4px;
}

.sign-up-form > button {
	margin-top: 12px;
}
</style>

<script>
export default {
	name: 'Wallet',
	data() {
		return {
			isSigning: false,
			name: '',
			password: '',
			mnemonic: '',
			loggedAddress: ''
		}
	},
	beforeCreate: function () {
		const vuexModule = ['common', 'wallet']
		for (let i = 1; i <= vuexModule.length; i++) {
			const submod = vuexModule.slice(0, i)
			if (!this.$store.hasModule(submod)) {
				console.log('Module ' + vuexModule + ' has not been registered!')
				this._depsLoaded = false
				break
			}
		}
	},
	async created() {
		const wallet = localStorage.getItem('wallet')
		if (wallet) {
			const walletInfo = JSON.parse(wallet)
			await this.$store.dispatch('common/wallet/unlockWallet', {
				name: walletInfo.name,
				password: walletInfo.password
			})
			this.loggedAddress = this.$store.getters['common/wallet/address']
			await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryUserLoad', {
				options: { all: true },
				params: {
					id: this.loggedAddress
				}
			})
		}
	},
	methods: {
		connectWallet: async function () {
			this.isSigning = true
		},
		createWallet: async function () {
			await this.$store.dispatch('common/wallet/createWalletWithMnemonic', {
				name: this.name,
				password: this.password,
				mnemonic: this.mnemonic,
				HDpath: "m/44'/118'/0'/0/", // BIP32/44 derivation path
				prefix: 'cosmos' // Address prfix for this chain
			})
			localStorage.setItem(
				'wallet',
				JSON.stringify({
					name: this.name,
					password: this.password
				})
			)
			this.isSigning = false
			this.loggedAddress = this.$store.getters['common/wallet/address']
			await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryUserLoad', {
				options: { subscribe: true, all: true },
				params: {
					id: this.loggedAddress
				}
			})
		}
	}
}
</script>
