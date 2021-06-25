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
				<div class="table-row">
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
	padding: 20px;
	border: 1px solid white;
	border-radius: 10px;
}

.table-header {
	display: flex;
	color: rgba(255, 255, 255, 0.7);
}

.table-rows {
	margin-top: 16px;
	font-weight: bold;
}

.table-row {
	display: flex;
}

.table-cell {
	width: 100%;
	min-width: 40px;
	text-align: right;
}

.table-cell:first-child {
	text-align: left;
}
</style>

<script>
export default {
	name: 'Userborrow'
}
</script>
