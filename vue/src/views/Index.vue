<template>
	<div class="container">
		<div class="user">
			<UserDeposit @click-asset="openModal" />
			<UserBorrow @click-asset="openModal" />
		</div>
		<div class="pools">
			<DepositPools @click-asset="openModal" />
			<BorrowPools @click-asset="openModal" />
		</div>
		<Modal v-bind="pool" v-bind:type="type" v-if="isModalOpen" @click-outside="closeModal" />
	</div>
</template>

<style scoped>
.user {
	display: flex;
}

.pools {
	margin-top: 48px;
	display: flex;
}
</style>

<script>
import UserDeposit from '../components/UserDeposit'
import UserBorrow from '../components/UserBorrow'
import DepositPools from '../components/DepositPools'
import BorrowPools from '../components/BorrowPools'
import Modal from '../components/Modal'

export default {
	name: 'Index',
	data() {
		return {
			isModalOpen: false,
			pool: null,
			type: ''
		}
	},
	async created() {
		await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryPoolLoad', {
			options: { subscribe: true, all: true },
			params: {}
		})
	},
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	},
	methods: {
		openModal(pool, type) {
			this.pool = pool
			this.type = type
			this.isModalOpen = true
		},
		closeModal() {
			this.pool = null
			this.type = ''
			this.isModalOpen = false
		}
	},
	components: {
		UserDeposit,
		UserBorrow,
		DepositPools,
		BorrowPools,
		Modal
	}
}
</script>
