<template>
	<div class="container">
		<div class="user">
			<UserDeposit @click-asset="openModal" />
			<UserBorrow @click-asset="openModal" />
		</div>
		<div class="pools">
			<DepositPools @click-asset="openModal" :pools="pools" />
			<BorrowPools @click-asset="openModal" :pools="pools" />
		</div>
		<Modal v-if="isModalOpen" @click-outside="closeModal" />
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
			isModalOpen: false
		}
	},
	async created() {
		await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryPoolAll', {
			options: { subscribe: true, all: true },
			params: {}
		})
	},
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		},
		pools() {
			return (
				this.$store.getters['sapienscosmos.ibb.ibb/getPoolAll']({
					params: {}
				})?.Pool ?? []
			)
		}
	},
	methods: {
		openModal() {
			this.isModalOpen = true
		},
		closeModal() {
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
