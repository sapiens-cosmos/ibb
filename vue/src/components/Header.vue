<template>
	<div class="header-wrapper">
		<div class="logo">IBB</div>
		<Wallet />
	</div>
</template>

<style scoped>
.header-wrapper {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 20px;
}

.logo {
	font-size: 48px;
	font-weight: bold;
}
</style>

<script>
import Wallet from './Wallet'

export default {
	name: 'Header',
	components: {
		Wallet
	},
	data() {
		return {}
	},
	computed: {
		hasWallet() {
			return this.$store.hasModule(['common', 'wallet'])
		}
	}
}
</script>
