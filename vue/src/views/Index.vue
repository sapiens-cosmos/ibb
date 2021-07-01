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
		<div class="nftList">
			<NftList @click-nft-card="openNftModal" />
		</div>
		<Modal v-bind="pool" v-bind:initialType="type" v-if="isModalOpen" @click-outside="closeModal" />
		<NftModal v-bind="pool" v-bind:initialType="type" v-if="isNftModalOpen" @click-outside="closeNftModal" />
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

.nftList {
	margin-top: 48px;
	padding-bottom: 48px;
}
</style>

<script>
import UserDeposit from '../components/UserDeposit'
import UserBorrow from '../components/UserBorrow'
import DepositPools from '../components/DepositPools'
import BorrowPools from '../components/BorrowPools'
import NftList from '../components/NftList'
import Modal from '../components/Modal'
import NftModal from '../components/NftModal'

export default {
	name: 'Index',
	data() {
		return {
			isModalOpen: false,
			isNftModalOpen: false,
			pool: null,
			nft: null,
			type: ''
		}
	},
	async created() {
		await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryPoolLoad', {
			options: { all: true },
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
		},
		openNftModal(nft) {
			this.nft = nft
			this.isNftModalOpen = true
		},
		closeNftModal() {
			this.nft = null
			this.isNftModalOpen = false
		}
	},
	components: {
		UserDeposit,
		UserBorrow,
		DepositPools,
		BorrowPools,
		NftList,
		Modal,
		NftModal
	}
}
</script>
