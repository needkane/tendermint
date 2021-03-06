# Changelog

## 0.9.0 (March 6, 2017)

BREAKING CHANGES:

- Update ABCI to v0.4.0, where Query is now `Query(RequestQuery) ResponseQuery`, enabling precise proofs at particular heights:

```
message RequestQuery{
	bytes data = 1;
	string path = 2;
	uint64 height = 3;
	bool prove = 4; 
}

message ResponseQuery{
	CodeType          code        = 1;
	int64             index       = 2;
	bytes             key         = 3;
	bytes             value       = 4;
	bytes             proof       = 5;
	uint64            height      = 6;
	string            log         = 7;
}
```


- `BlockMeta` data type unifies its Hash and PartSetHash under a `BlockID`:

```
type BlockMeta struct {
	BlockID BlockID `json:"block_id"` // the block hash and partsethash
	Header  *Header `json:"header"`   // The block's Header
}
```

- `tendermint gen_validator` command output is now pure JSON
- `ValidatorSet` data type:
	- expose a `Proposer` field. Note this means the `Proposer` is persisted with the `State`.
	- change `.Proposer()` to `.GetProposer()`

FEATURES:

- New RPC endpoint `/commit?height=X` returns header and commit for block at height `X`
- Client API for each endpoint, including mocks for testing

IMPROVEMENTS:

- `Node` is now a `BaseService`
- Simplified starting Tendermint in-process from another application
- Better organized Makefile
- Scripts for auto-building binaries across platforms
- Docker image improved, slimmed down (using Alpine), and changed from tendermint/tmbase to tendermint/tendermint
- New repo files: `CONTRIBUTING.md`, Github `ISSUE_TEMPLATE`, `CHANGELOG.md`
- Improvements on CircleCI for managing build/test artifacts
- Handshake replay is doen through the consensus package, possibly using a mockApp
- Graceful shutdown of RPC listeners
- Tests for the PEX reactor and DialSeeds

BUG FIXES:

- Check peer.Send for failure before updating PeerState in consensus
- Fix panic in `/dial_seeds` with invalid addresses
- Fix proposer selection logic in ValidatorSet by taking the address into account in the `accumComparable`
- Fix inconcistencies with `ValidatorSet.Proposer` across restarts by persisting it in the `State`


## 0.8.0 (January 13, 2017)

BREAKING CHANGES:

- New data type `BlockID` to represent blocks:

```
type BlockID struct {
	Hash        []byte        `json:"hash"`
	PartsHeader PartSetHeader `json:"parts"`
}
```

- `Vote` data type now includes validator address and index: 

```
type Vote struct {
	ValidatorAddress []byte           `json:"validator_address"`
	ValidatorIndex   int              `json:"validator_index"`
	Height           int              `json:"height"`
	Round            int              `json:"round"`
	Type             byte             `json:"type"`
	BlockID          BlockID          `json:"block_id"` // zero if vote is nil.
	Signature        crypto.Signature `json:"signature"`
}
```

- Update TMSP to v0.3.0, where it is now called ABCI and AppendTx is DeliverTx
- Hex strings in the RPC are now "0x" prefixed


FEATURES:

- New message type on the ConsensusReactor, `Maj23Msg`, for peers to alert others they've seen a Maj23, 
in order to track and handle conflicting votes intelligently to prevent Byzantine faults from causing halts:

```
type VoteSetMaj23Message struct {
	Height  int
	Round   int
	Type    byte
	BlockID types.BlockID
}
```

- Configurable block part set size
- Validator set changes
- Optionally skip TimeoutCommit if we have all the votes
- Handshake between Tendermint and App on startup to sync latest state and ensure consistent recovery from crashes
- GRPC server for BroadcastTx endpoint

IMPROVEMENTS:

- Less verbose logging
- Better test coverage (37% -> 49%)
- Canonical SignBytes for signable types
- Write-Ahead Log for Mempool and Consensus via go-autofile 
- Better in-process testing for the consensus reactor and byzantine faults
- Better crash/restart testing for individual nodes at preset failure points, and of networks at arbitrary points
- Better abstraction over timeout mechanics

BUG FIXES:

- Fix memory leak in mempool peer
- Fix panic on POLRound=-1
- Actually set the CommitTime
- Actually send BeginBlock message
- Fix a liveness issues caused by Byzantine proposals/votes. Uses the new `Maj23Msg`.


## 0.7.4 (December 14, 2016)

FEATURES:

- Enable the Peer Exchange reactor with the `--pex` flag for more resilient gossip network (feature still in development, beware dragons)

IMPROVEMENTS:

- Remove restrictions on RPC endpoint `/dial_seeds` to enable manual network configuration

## 0.7.3 (October 20, 2016)

IMPROVEMENTS:

- Type safe FireEvent
- More WAL/replay tests
- Cleanup some docs

BUG FIXES:

- Fix deadlock in mempool for synchronous apps
- Replay handles non-empty blocks
- Fix race condition in HeightVoteSet

## 0.7.2 (September 11, 2016)

BUG FIXES:

- Set mustConnect=false so tendermint will retry connecting to the app

## 0.7.1 (September 10, 2016)

FEATURES:

- New TMSP connection for Query/Info
- New RPC endpoints:
	- `tmsp_query`
	- `tmsp_info`
- Allow application to filter peers through Query (off by default)

IMPROVEMENTS:

- TMSP connection type enforced at compile time
- All listen/client urls use a "tcp://" or "unix://" prefix

BUG FIXES:

- Save LastSignature/LastSignBytes to `priv_validator.json` for recovery
- Fix event unsubscribe
- Fix fastsync/blockchain reactor

## 0.7.0 (August 7, 2016)

BREAKING CHANGES:

- Strict SemVer starting now!
- Update to ABCI v0.2.0
- Validation types now called Commit
- NewBlock event only returns the block header


FEATURES:

- TMSP and RPC support TCP and UNIX sockets
- Addition config options including block size and consensus parameters
- New WAL mode `cswal_light`; logs only the validator's own votes
- New RPC endpoints: 
	- for starting/stopping profilers, and for updating config
	- `/broadcast_tx_commit`, returns when tx is included in a block, else an error
	- `/unsafe_flush_mempool`, empties the mempool


IMPROVEMENTS:

- Various optimizations
- Remove bad or invalidated transactions from the mempool cache (allows later duplicates)
- More elaborate testing using CircleCI including benchmarking throughput on 4 digitalocean droplets

BUG FIXES:

- Various fixes to WAL and replay logic
- Various race conditions
