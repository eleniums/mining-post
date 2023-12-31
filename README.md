# mining-post
An API-based multiplayer game where you purchase mining equipment and various resources to become rich.

## Run
Run the service:
```
./scripts/run.sh
```

## Build
To build for Windows, Mac, and Linux:
```
./scripts/build.sh
```

To create release binaries (includes build.sh):
```
./scripts/release.sh
```

Releases are versioned with the latest git tag under the release folder (ex: release/v1.1.0-17-gfdba3eb).

## Docker
Build the latest Docker container image:
```
make docker
```

Run a container:
```
docker run --rm -p 9090:9090 miningpost
```

## Database
This is a fun hobby project that is not intended for actual production use. This server cannot, in its current state, scale out horizontally, only vertically. Since everything is stored in memory while running and no complex calculations are being made, it is very lightweight and very performant. As such, this project uses a simple embedded key-value database for storing player data.

The database chosen for this project is [bbolt](https://github.com/etcd-io/bbolt), which is a fork of the no longer maintained [boltdb](https://github.com/boltdb/bolt). Bolt is simple, fast, and written in pure Go. There might be scaling/performance concerns for very large amounts of data, but for the data being stored for this project, Bolt should be more than sufficient.

Here are some other options that were considered and are most likely just as viable:
- [BadgerDB](https://github.com/dgraph-io/badger): Fast key-value storage written in Go.
- [RocksDB](https://rocksdb.org): Fast key-value storage written in C++.
- [LevelDB](https://github.com/google/leveldb): Fast key-value storage created by Google and written in C++.
- [SQLite](https://www.sqlite.org): Small, fast, self-contained, SQL database. Cannot have more than one concurrent writer.

If this project was going to be deployed and made available to the public, a more commercial option would be used. Most likely one of the following:
- [AWS DynamoDB](https://aws.amazon.com/dynamodb): Fast, robust, and easy-to-use key-value storage.
- [AWS Aurora PostgreSQL](https://aws.amazon.com/rds/aurora): Managed PostgreSQL that is highly available and performant, but expensive.
- [CockroachDB](https://github.com/cockroachdb/cockroach): Distributed SQL database that is designed to scale.
- [Redis](https://redis.io/): Fast in-memory cache that can also persist data.

## Gameplay TODOs
- [x] Need a way to calculate net worth outside of current money (use net worth for promotions). Net worth would be value of resources, equipment, employees, land, etc.
- [x] Add type to listings and filter capability for name and type
- [x] More ranks? Like amatuer and apprentice and mining engineer?
- [x] Related to ranks, adjust salaries. Should start with a very low salary.
- [ ] Add equipment (like pickaxes and mine carts) and buildings (like a mine)
- [ ] Add a mine/dig option? Use pickaxes to get minerals?
- [ ] Maybe a prospect option to find more land?
- [ ] Add high score table that lists the top ten players by net worth
- [x] Put instructions in a markdown file and use embed to add to handler
- [x] Should not be able to buy or sell quantity 0 or negative
- [x] List market should show time remaining.
- [x] Should round values before returning to player. I saw this happen again with money somehow... Would be nice if money was formatted as a string like `"$12.30"`

## Technical TODOs
- [x] Better locking system.
- [x] Need a way to save player data. Postgres or maybe just hack JSON file for now? There are some decent free embedded databases, like maybe bolt or sqlite? Keep everything in memory and just save at the end of every update. Works for now, but won't allow scaling horizontally.
- [ ] Need something to autogenerate developer docs. Swagger? OpenAPI?
- [ ] Host developer docs on public endpoint
- [ ] Need to perf test updates with a large number of users, like 1 million users
- [x] Need to separate data models and contract models. I tried it both ways and keeping them combined, as expected, is turning into more of a pain than mapping
- [x] Add retries to client. Use Rican7 package.
- [x] Add middleware to limit request size. Is there pre-existing middleware in the chi library? Write it myself?
- [ ] Create dummy auth endpoint that generates a JWT regardless of password.
- [ ] Create authorizer middleware to properly validate JWT. Client ID should be passed in headers and/or context. Is there anything out there I can use? Or just write it myself?
- [ ] Add rate limiting for requests. Is there pre-existing middleware in the chi library? (chi library only has throttling for total number of requests, not per user) Write it myself with semaphore? (need client id in headers/context from authorizer first)
- [ ] Should I add middleware to allowlist certain content types and encodings? (go-chi has AllowContentType and AllowContentEncoding)
- [x] Enforce TLS 1.2 or greater on the server (already enforced)
- [x] Run locally using TLS with test certs (generated new test certs good for 100 years that work with 127.0.0.1. Had to do some weird stuff to add the IP SAN for 127.0.0.1 and insecureskipverify needs to be true since certs are self-signed)

Potential options for rate limiting (besides writing it myself) that can limit per user per unit of time (like no more than 60 requests per minute for player "tstark"):
- https://github.com/didip/tollbooth
- https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
- https://blog.logrocket.com/rate-limiting-go-application/
- https://github.com/go-chi/httprate
