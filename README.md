# TODO
## backend
* migration
* multi stage build

## frontend
* workflowでfmtとlint
* multi stage build
* crud実装

## other
* compose.prod.yaml
* readme.mdの整理

# How To Setup
```
cp .env.template .env
docker compose up -d
```

# Build
```
cp .env.template .env
docker compose -f compose.prod.yaml up -d
```
