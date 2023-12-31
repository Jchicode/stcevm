#!/bin/bash

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
CHAINID="stcevmnetwork_9000-1"
MONIKER="stcevmnetworklocal"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the stcevmd instance
HOMEDIR="$HOME/.tmp-stc-network"
# to trace evm
#TRACE="--trace"
TRACE=""

# Path variables
CONFIG=$HOMEDIR/config/config.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# validate dependencies are installed
command -v jq >/dev/null 2>&1 || {
	echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
	exit 1
}

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
make install

# User prompt if an existing local node configuration is found.
if [ -d "$HOMEDIR" ]; then
	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
	echo "Overwrite the existing configuration and start a new local node? [y/n]"
	read -r overwrite
else
	overwrite="Y"
fi

# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# Remove the previous folder
	rm -rf "$HOMEDIR"

	# Set client config
  stcevmd config keyring-backend $KEYRING --home $HOMEDIR
  stcevmd config chain-id $CHAINID --home $HOMEDIR

	# If keys exist they should be deleted
	for KEY in "${KEYS[@]}"; do
		stcevmd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
	done

	# Set moniker and chain-id for STCEvm Network (Moniker can be anything, chain-id must be an integer)
	stcevmd init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

	# Change parameter token denominations to ciot
	jq '.app_state["staking"]["params"]["bond_denom"]="ciot"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["mint"]["params"]["mint_denom"]="ciot"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["crisis"]["constant_fee"]["denom"]="ciot"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="ciot"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["evm"]["params"]["evm_denom"]="ciot"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["inflation"]["params"]["mint_denom"]="ciot"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Set gas limit in genesis
	jq '.consensus_params["block"]["max_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	if [[ $1 == "pending" ]]; then
		if [[ "$OSTYPE" == "darwin"* ]]; then
			sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
			sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
			sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
			sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$CONFIG"
			sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
		else
			sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
			sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
			sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
			sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$CONFIG"
			sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
		fi
	fi

	# Allocate genesis accounts (cosmos formatted addresses)
	for KEY in "${KEYS[@]}"; do
		stcevmd add-genesis-account $KEY 10000000000000000000000000ciot --keyring-backend $KEYRING --home "$HOMEDIR"
	done

	# Sign genesis transaction
	stcevmd gentx ${KEYS[0]} 1000000000000000000000000ciot --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
	## In case you want to create multiple validators at genesis
	## 1. Back to `stcevmd keys add` step, init more keys
	## 2. Back to `stcevmd add-genesis-account` step, add balance for those
	## 3. Clone this ~/.stcevmd home directory into some others, let's say `~/.clonedRealioNetwork`
	## 4. Run `gentx` in each of those folders
	## 5. Copy the `gentx-*` folders under `~/.clonedEvmosd/config/gentx/` folders into the original `~/.stcevmd/config/gentx`

	# Collect genesis tx
	stcevmd collect-gentxs --home "$HOMEDIR"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	stcevmd validate-genesis --home "$HOMEDIR"

	if [[ $1 == "pending" ]]; then
		echo "pending mode is on, please wait for the first block committed."
	fi
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
stcevmd start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --minimum-gas-prices=0.0001ciot --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --home "$HOMEDIR"