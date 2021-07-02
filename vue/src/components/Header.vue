<template>
	<div class="header-wrapper">
		<div class="logo">
			<img src="@/assets/images/icons/logo.svg" />
			<span>Interstellar</span>
		</div>
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
	font-size: 36px;
	font-weight: bold;
	display: flex;
}

.logo > img {
	margin-top: 3px;
	width: 40px;
	height: 40px;
	margin-right: 2px;
}
</style>

<script>
import IconLogo from './IconLogo.vue'
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
