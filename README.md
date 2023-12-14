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
- [ ] Need a way to calculate net worth outside of current money (use net worth for promotions). Net worth would be value of resources, equipment, employees, land, etc.
- [ ] Add type to listings and filter capability
- [ ] Add equipment (like pickaxes and mine carts) and buildings (like a mine)
- [ ] Better locking system. Helper methods maybe?
- [ ] Need a way to save player data. Postgres or maybe just hack JSON file for now? There are some decent free embedded databases, like maybe bolt or sqlite?
- [ ] More ranks? Like amatuer and apprentice and mining engineer?
- [ ] Related to ranks, adjust salaries. Should start with a very low salary.
- [ ] Add a mine/dig option? Use pickaxes to get minerals?
- [ ] Maybe a prospect option to find more land?

# Types of stock
- Resource
- Equipment (Pickaxe, Mining Cart, Dynamite)
- Land (Claim, Mine)
- Employee (Worker, Specialist, Mining Engineer)
