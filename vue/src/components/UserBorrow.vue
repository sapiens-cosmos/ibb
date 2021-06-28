<template>
	<div class="user-borrow">
		<div class="borrow-status">
			<div class="borrow-balance">
				<div class="title">Borrow Balance</div>
				<div class="value">$325.52</div>
			</div>
			<div class="net-apy">
				<div class="title">Borrow Limit</div>
				<div class="value">$400.12</div>
			</div>
		</div>
		<div class="asset-table">
			<div class="table-header">
				<div class="table-cell">Asset</div>
				<div class="table-cell">APY</div>
				<div class="table-cell">Balance</div>
				<div class="table-cell">Borrow Limit</div>
			</div>
			<div class="table-rows">
				<div class="table-row" @click="clickAsset">
					<div class="table-cell">ATOM</div>
					<div class="table-cell">64.59%</div>
					<div class="table-cell">26.75</div>
					<div class="table-cell">34.55</div>
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
	name: 'Userborrow',
	methods: {
		clickAsset() {
			this.$emit('click-asset')
		}
	}
}
</script>
