<template>
	<div class="user-deposit">
		<div class="deposit-status">
			<div class="deposit-balance">
				<div class="title">Deposit Balance</div>
				<div class="value">$400.12</div>
			</div>
			<div class="net-apy">
				<div class="title">Net APY</div>
				<div class="value">32.59%</div>
			</div>
		</div>
		<div class="asset-table">
			<div class="table-header">
				<div class="table-cell"><span>Asset</span></div>
				<div class="table-cell"><span>APY</span></div>
				<div class="table-cell"><span>Balance</span></div>
				<div class="table-cell"><span>Collateral</span></div>
				<div class="table-cell"><span>Interest</span></div>
			</div>
			<div class="table-rows">
				<div class="table-row" @click="clickAsset">
					<div class="table-cell"><span>ATOM</span></div>
					<div class="table-cell"><span>32.59%</span></div>
					<div class="table-cell"><span>34.55</span></div>
					<div class="table-cell"><input type="checkbox" /></div>
					<div class="table-cell"><button>Claim</button></div>
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

.table-rows {
	font-weight: bold;
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

.table-cell > span {
	line-height: 14px;
}
</style>

<script>
export default {
	name: 'UserDeposit',
	methods: {
		clickAsset() {
			console.log(
				this.$store.getters['sapienscosmos.ibb.ibb/getUserLoad']({
					params: {
						id: 'cosmos1f5zznu7gp8fx8nftdlg243sw3kx66wl2ma3d46'
					}
				})
			)
			this.$emit('click-asset')
		}
	},
	props: {
		toggleModal: Function
	}
}
</script>
