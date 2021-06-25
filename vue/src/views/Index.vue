<template>
	<div class="container">
		<div class="user">
			<UserDeposit />
			<UserBorrow />
		</div>
		<div class="pools">
			<DepositPools />
			<BorrowPools />
		</div>
		<Modal />
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
	components: {
		UserDeposit,
		UserBorrow,
		DepositPools,
		BorrowPools,
		Modal
	},
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	}
}
</script>
