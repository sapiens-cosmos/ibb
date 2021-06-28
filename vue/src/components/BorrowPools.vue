<template>
	<div class="user-borrow">
		<div class="borrow-status">
			<div class="borrow-balance">
				<div class="title">Borrow Pools</div>
			</div>
		</div>
		<div class="asset-table">
			<div class="table-header">
				<div class="table-cell">Asset</div>
				<div class="table-cell">APY</div>
				<div class="table-cell">Liquidity</div>
			</div>
			<div v-if="Array.isArray(pools)" class="table-rows">
				<div v-for="pool in pools" v-bind:key="pool.id" class="table-row" @click="clickAsset">
					<div class="table-cell">{{ pool.Asset }}</div>
					<div class="table-cell">{{ pool.BorrowApy }}</div>
					<div class="table-cell">{{ `${pool.Liquidity} ${pool.Asset}` }}</div>
				</div>
			</div>
			<div v-else class="table-rows">
				<div class="table-row">
					<div class="table-cell">NONE</div>
					<div class="table-cell">NONE</div>
					<div class="table-cell">NONE</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.user-borrow {
	width: 100%;
	padding: 0 20px;
}

.borrow-status {
	display: flex;
	justify-content: space-between;
}

.title {
	font-size: 18px;
}

.value {
	font-size: 26px;
	font-weight: bold;
}

.asset-table {
	margin-top: 20px;
	padding: 12px;
	border: 1px solid white;
	border-radius: 10px;
}

.table-header {
	display: flex;
	padding: 8px 12px;
	color: rgba(255, 255, 255, 0.7);
	font-size: 14px;
}

.table-row {
	display: flex;
	padding: 8px 12px 7px;
	border-radius: 5px;
}

.table-row:hover {
	cursor: pointer;
	background: rgba(0, 0, 0, 0.5);
}

.table-cell {
	width: 100%;
	min-width: 40px;
	display: flex;
	align-items: center;
	justify-content: flex-end;
}

.table-cell:first-child {
	justify-content: flex-start;
}
</style>

<script>
export default {
	name: 'BorrowPools',
	methods: {
		clickAsset() {
			this.$emit('click-asset')
		}
	},
	computed: {
		pools() {
			return (
				this.$store.getters['sapienscosmos.ibb.ibb/getPoolLoad']({
					params: {}
				})?.LoadPoolResponse ?? []
			)
		}
	}
}
</script>
