# SANDBOX

DAODISEO Sandbox Blockchain Environment

## Achilles app-chain binaries installation (achillesd)

```
go: go version go1.22.9 linux/amd64
name: achilles
```

## Prerequisites

### Install go

```
sudo apt-get update
sudo apt-get upgrade
apt install mc htop screen git gcc make
```

### Put GOPATH and GOBIN to bash_profile

### Use source ~/.profile

## Binary building

### Clone source from repo

```
git clone git@github.com:daodiseomoney/SANDBOX.git
```

### Open source directory

```
cd SANDBOX/achilles
```

### Build binary

```
make build
```

### Copy binary to usr/local/bin or to root

```
cp $(GOBIN)/achilles /usr/local/bin (or copy it to root as achillesd and use it like ./achillesd)
```

## Network launch

### Init

```bash:
achillesd init "<moniker-name>" --chain-id test-core-1
```

### Set minimum-gas-prices = "" in app.toml to minimum-gas-prices = "1uodis"

### Generate keys

```bash:
# To create new keypair - make sure you save the mnemonics!
achillesd keys add <key-name>
```

or

```
# Restore existing odin wallet with mnemonic seed phrase.
# You will be prompted to enter mnemonic seed.
achillesd keys add <key-name> --recover
```

or

```
# Add keys using ledger
achillesd keys show <key-name> --ledger
```

Check your key:

```
# Query the keystore for your public address
achillesd keys show <key-name> -a
```

### Create account to genesis

```
achillesd genesis add-genesis-account <key-name> 1000000uodis
```

### Create GenTX

```
# Create the gentx.
# Note, your gentx will be rejected if you use any amount greater than 1000000uodis.
# Make sure that all participants built their gentx files without typos.

achillesd genesis gentx <key-name> 1000000uodis \
  --pubkey=$(achillesd tendermint show-validator) \
  --chain-id=test-core-1 \
  --moniker="my-moniker" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01"
```

### Add all accounts to genesis

```
# Add account addresses of all participants before generating genesis.
# (whose Gentx files you're using to generate genesis)
achillesd genesis add-genesis-account <account-address> 1000000uodis
```

### Generate genesis

```
achillesd genesis collect-gentxs
```

### Start network

```
achillesd start
```
