<template>
	<div v-if="initialized" class="container">
		<Header />
		<router-view />
	</div>
</template>

<style>
html {
	height: 100vh;
}

body {
	background: rgb(2, 0, 36);
	background: linear-gradient(90deg, rgba(7, 25, 62, 1) 0%, rgba(9, 67, 121, 1) 41%, rgba(255, 139, 0, 1) 100%);
	margin: 0 auto;
	max-width: 1200px;
	width: 100%;
}

.container {
	background: transparent;
}
</style>

<script>
import './scss/app.scss'
// import '@starport/vue/lib/starport-vue.css'
import Header from './components/Header'

export default {
	components: {
		Header
	},
	data() {
		return {
			initialized: false
		}
	},
	computed: {
		hasWallet() {
			return this.$store.hasModule(['common', 'wallet'])
		}
	},
	async created() {
		await this.$store.dispatch('common/env/init')
		this.initialized = true
	},
	errorCaptured(err) {
		console.log(err)
		return false
	}
}
</script>
