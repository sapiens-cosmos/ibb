<template>
	<div class="nft-list">
		<div class="title">NFT List</div>
		<div class="nft-table">
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
				</div>
			</div>
			<div class="card" @click="clickNftCard">
				<img class="media" src="https://avatars.githubusercontent.com/u/32129022?v=4" />
				<div class="info">
					<div class="collection">Nonce Sapiens</div>
					<div class="name">Delivan</div>
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
	font-size: 18px;
}

.nft-table {
	margin-top: 16px;
	padding: 24px 0 0 24px;
	border: 1px solid white;
	border-radius: 10px;
	display: flex;
	flex-wrap: wrap;
}

.card {
	width: 100%;
	max-width: 200px;
	margin-right: 24px;
	margin-bottom: 24px;
	border: 1px solid white;
	border-radius: 10px;
}

.card:hover {
	cursor: pointer;
}

.media {
	border-radius: 10px;
	width: 100%;
	max-width: 200px;
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
	methods: {
		clickNftCard(nft) {
			this.$emit('click-nft-card', nft)
		}
	}
}
</script>
