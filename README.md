# Uncommons

> https://ninja.org/uncommons

## Root-contracts

### Whitelist developers

> Contract address: `0xe195fdde701ccfe1dba41e5802bea0d671306037`

### Withdraw and logs

> Contract address: `0xdbc45d7c3da97ce2986044006aa14e9507ca8fe3`

## Install

### Install from source code

Clone this repository and run

```
go install -v ./...
```

### Download `gunc` file

> Coming soon

## Run

This client based on go-etherum, if you used to `geth`, this `gunc` will be easy with you

```
gunc
```

### Mine

```
gunc --mine --miner.threads 4 --gcmode=archive --syncmode=full --rpc --rpccorsdomain "*" console
```
