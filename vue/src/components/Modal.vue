<template>
	<div class="modal-wrapper">
		<div class="modal">
			<div class="asset">
				<img class="asset-icon" src="@/assets/images/icons/atom.png" />
				<div class="asset-name">ATOM</div>
			</div>
			<div class="modes">
				<div class="mode">Deposit</div>
				<div class="mode">Borrow</div>
			</div>
			<div class="body">
				<div class="amount">
					<div class="title">
						<div class="property">Deposit Amount</div>
						<div class="value">Wallet Balance: 334.45 ATOM</div>
					</div>
					<div class="content">
						<div class="input-wrapper">
							<div class="input">
								<div class="balance">
									<input type="text" />
									<div class="denom">ATOM</div>
								</div>
								<div class="dollar">$0</div>
							</div>
							<div class="max-button">Max</div>
						</div>
					</div>
					<div class="content">
						<div class="slide">
							<div class="slide-bar" />
							<div class="slide-circle" />
							<div class="slide-percentages">
								<div class="slide-percentage">0%</div>
								<div class="slide-percentage">25%</div>
								<div class="slide-percentage">50%</div>
								<div class="slide-percentage">75%</div>
								<div class="slide-percentage">100%</div>
							</div>
						</div>
					</div>
				</div>
				<div class="stats">
					<div class="title">Stats</div>
					<div class="content">
						<div class="property">APY</div>
						<div class="value">23.56%</div>
					</div>
					<div class="content">
						<div class="property">Balance</div>
						<div class="value">20 ATOM</div>
					</div>
				</div>
				<div class="collateral">
					<div class="title">Collateral</div>
					<div class="content">
						<div class="property">Collateral Factor</div>
						<div class="value">75%</div>
					</div>
					<div class="content">
						<div class="property">Used as Collateral</div>
						<div class="value">Yes</div>
					</div>
				</div>
				<div class="borrow-limit">
					<div class="title">Borrow Limit</div>
					<div class="content">
						<div class="property">Your Limit</div>
						<div class="value">$0 -> $402.54</div>
					</div>
					<div class="content">
						<div class="property">Limit Used</div>
						<div class="value">0% -> 0%</div>
					</div>
				</div>
				<div class="cta">Deposit</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.modal-wrapper {
	width: 100vw;
	height: 100vh;
	position: absolute;
	top: 0;
	left: 0;
	display: flex;
	flex-direction: column;
	align-items: center;
	background: rgba(0, 0, 0, 0.7);
}

.modal {
	position: relative;
	margin-top: 60px;
	background: rgba(0, 0, 0, 0.8);
	border: 1px solid white;
	border-radius: 10px;
	max-width: 500px;
	width: 100%;
}

.asset {
	position: absolute;
	top: -40px;
	left: 50%;
	transform: translateX(-50%);
}

.asset-icon {
	width: 80px;
	height: 80px;
}

.asset-name {
	position: absolute;
	top: 70px;
	transform: translateX(-50%);
	background: black;
	left: 50%;
	padding: 8px 8px 6px;
	font-size: 16px;
	font-weight: bold;
	text-align: center;
	border: 1px solid white;
	border-radius: 5px;
}

.modes {
	display: flex;
	justify-content: space-between;
}

.mode {
	width: 100%;
	padding: 22px 20px 20px;
	font-size: 28px;
	font-weight: bold;
	text-align: center;
	border-bottom: 1px solid white;
}

.body {
	padding: 0 36px 36px;
}

.input-wrapper {
	border: 1px solid white;
	border-radius: 10px;
	padding: 16px;
	width: 100%;
	display: flex;
	align-items: center;
}

.input {
	width: 100%;
}

.balance {
	display: flex;
	align-items: center;
}

.balance > input {
	width: 100%;
	height: 100%;
	margin-right: 4px;
	margin-bottom: 1px;
	background: transparent;
	border: unset;
	color: white;
	font-size: 18px;
	text-align: right;
}

.balance > input:focus-visible {
	outline: unset;
}

.denom {
	font-size: 18px;
}

.dollar {
	margin-top: 4px;
	font-size: 16px;
	text-align: right;
}

.max-button {
	padding-left: 16px;
}

.slide {
	position: relative;
	width: 100%;
}

.slide-bar {
	margin-top: 2px;
	background: rgba(255, 255, 255, 0.7);
	height: 6px;
	border-radius: 50px;
}

.slide-circle {
	position: absolute;
	top: 0;
	left: 0;
	width: 10px;
	height: 10px;
	background: white;
	border-radius: 50%;
}

.slide-percentages {
	margin-top: 12px;
	display: flex;
	justify-content: space-between;
}

.slide-percentage {
	font-size: 14px;
	font-weight: bold;
}

.title {
	margin-top: 36px;
	display: flex;
	justify-content: space-between;
	font-size: 16px;
	font-weight: bold;
	color: rgba(255, 255, 255, 0.7);
}

.content {
	margin-top: 8px;
	display: flex;
	font-size: 20px;
	justify-content: space-between;
}

.cta {
	margin-top: 36px;
	padding: 16px;
	color: #333333;
	font-size: 20px;
	font-weight: bold;
	text-align: center;
	border: 1px solid white;
	background: white;
	border-radius: 10px;
}
</style>

<script>
export default {
	name: 'Modal'
}
</script>
