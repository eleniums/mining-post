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
