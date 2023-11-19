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
- [ ] Add type to listings and filter capability
- [ ] Add equipment (like pickaxes and mine carts) and buildings (like a mine)
- [ ] Better locking system. Helper methods maybe?
- [ ] Need a way to save player data. Postgres or maybe just hack JSON file for now?
- [ ] More ranks? Like amatuer and apprentice and mining engineer?
- [ ] Add a mine/dig option? Use pickaxes to get minerals?
- [ ] Maybe a prospect option to find more land?
