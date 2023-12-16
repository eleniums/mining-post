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

## TODOs
- [x] Need a way to calculate net worth outside of current money (use net worth for promotions). Net worth would be value of resources, equipment, employees, land, etc.
- [x] Add type to listings and filter capability for name and type
- [ ] Add equipment (like pickaxes and mine carts) and buildings (like a mine)
- [ ] Better locking system. Helper methods maybe?
- [ ] Need a way to save player data. Postgres or maybe just hack JSON file for now? There are some decent free embedded databases, like maybe bolt or sqlite?
- [x] More ranks? Like amatuer and apprentice and mining engineer?
- [x] Related to ranks, adjust salaries. Should start with a very low salary.
- [ ] Add a mine/dig option? Use pickaxes to get minerals?
- [ ] Maybe a prospect option to find more land?
- [ ] Put instructions in a markdown file and use embed to add to handler
- [ ] Need to perf test updates with a large number of users, like 1 million users
- [ ] Should not be able to buy or sell quantity 0 or negative
- [ ] List market should show time remaining.
- [ ] Should round values before returning to player. I saw this happen again with money somehow...

# Types of stock
- Commodity
- Equipment (Pickaxe, Mining Cart, Dynamite)
- Land (Claim, Mine)
- Employee (Worker, Specialist, Mining Engineer)
