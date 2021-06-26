<template>
	<div class="user-deposit">
		<div class="deposit-status">
			<div class="deposit-balance">
				<div class="title">Deposit Pools</div>
			</div>
		</div>
		<div class="asset-table">
			<div class="table-header">
				<div class="table-cell">Asset</div>
				<div class="table-cell">APY</div>
				<div class="table-cell">Balance</div>
			</div>
			<div class="table-rows">
				<div v-for="pool in pools" v-bind:key="pool.asset" class="table-row">
					<div class="table-cell">{{ pool.asset }}</div>
					<div class="table-cell">32.59%</div>
					<div class="table-cell">34.55</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.user-deposit {
	width: 100%;
	padding: 0 20px;
}

.deposit-status {
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
	name: 'UserDeposit'
}
</script>
