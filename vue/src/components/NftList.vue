<template>
	<div class="nft-list">
		<div class="title">NFTs as Collateral</div>
		<div class="all-nfts">
			<div v-if="Array.isArray(userNftList) && userNftList.length > 0" class="sub-title">My NFTs</div>
			<div v-if="Array.isArray(userNftList) && userNftList.length > 0" class="my-nfts">
				<div v-for="nft in userNftList" v-bind:key="nft.id" class="card" :class="nft.OwnerAddress === 'in escrow' ? 'escrow' : ''" @click="clickNftCard(nft)">
					<div class="media">
						<video v-if="nft.ImageUrl.split('.').pop() === 'mp4'" :src="nft.ImageUrl" autoplay="true" loop="true" mute="true" playsinline="true" />
						<img v-else :src="nft.ImageUrl" />
					</div>
					<div class="info">
						<div class="collection">{{ nft.Collection }}</div>
						<div class="name">{{ nft.Name }}</div>
					</div>
				</div>
			</div>
			<div class="sub-title">NFTs on List</div>
			<div class="nfts-on-list">
				<div v-for="nft in nftList" v-bind:key="nft.id" class="card" @click="clickNftCard(nft)">
					<div class="media">
						<video v-if="nft.ImageUrl.split('.').pop() === 'mp4'" :src="nft.ImageUrl" autoplay="true" loop="true" mute="true" playsinline="true" />
						<img v-else :src="nft.ImageUrl" />
					</div>
					<div class="info">
						<div class="collection">{{ nft.Collection }}</div>
						<div class="name">{{ nft.Name }}</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.nft-list {
	padding: 0 20px;
}

.title {
	font-size: 20px;
}

.sub-title {
	font-size: 20px;
	margin-bottom: 10px;
}

.all-nfts {
	margin-top: 16px;
	padding: 24px 0 0 24px;
	border: 1px solid white;
	border-radius: 10px;
}

.my-nfts {
	display: flex;
	flex-wrap: wrap;
	margin-bottom: 8px;
}

.nfts-on-list {
	display: flex;
	flex-wrap: wrap;
}

.card {
	width: 100%;
	max-width: 200px;
	margin-right: 24px;
	margin-bottom: 24px;
	border: 1px solid rgba(255, 255, 255, 0.6);
	border-radius: 10px;
}

.card.escrow {
	opacity: 0.3;
}

.card:hover {
	cursor: pointer;
}

.media {
	padding: 12px;
	-webkit-box-align: center;
	align-items: center;
	display: flex;
	-webkit-box-pack: center;
	justify-content: center;
	max-height: 100%;
	max-width: 100%;
	height: 200px;
	overflow: hidden;
	position: relative;
	border-radius: 5px;
}

.media > * {
	object-fit: contain;
	width: auto;
	height: auto;
	border-radius: 5px;
	max-width: 100%;
	max-height: 100%;
}

.info {
	padding: 6px 12px 12px;
}

.collection {
	font-size: 13px;
	color: #aaa;
}

.name {
	margin-top: 6px;
	font-size: 16px;
	color: white;
}
</style>

<script>
export default {
	name: 'NftList',
	computed: {
		loggedAddress() {
			return this.$store.getters['common/wallet/address']
		},
		nftList() {
			const nftList =
				this.$store.getters['sapienscosmos.ibb.ibb/getNftLoad']({
					params: {
						...(this.loggedAddress && { id: this.loggedAddress })
					}
				})?.DashboardNft ?? []
			return nftList
		},
		userNftList() {
			const userNftList =
				this.$store.getters['sapienscosmos.ibb.ibb/getNftLoad']({
					params: {
						id: this.loggedAddress
					}
				})?.UserNft ?? []
			return userNftList
		}
	},
	methods: {
		clickNftCard(nft) {
			this.$emit('click-nft-card', nft)
		}
	}
}
</script>
