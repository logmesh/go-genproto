# Logmesh Common Types

```shell
# Example path src:
export REPO_LOGMESH_APIS_SRC="$HOME/go/src/github.com/logmesh/apis"
```

### *logmesh.type.Filter*
```shell
mkdir -p ./type/filter
protoc -I $REPO_LOGMESH_APIS_SRC/logmesh/type \
  --go_out ./type/filter \
  --go_opt paths=source_relative \
  filter.proto
```

### *logmesh.type.IncomeTransaction*
```shell
mkdir -p ./type/income_transaction
protoc -I $REPO_LOGMESH_APIS_SRC/logmesh/type \
  --go_out ./type/income_transaction \
  --go_opt paths=source_relative \
  income_transaction.proto
```

### *logmesh.bot.v1*
```shell
mkdir -p ./bot/v1
protoc -I $REPO_LOGMESH_APIS_SRC/logmesh/bot/v1 -I $REPO_LOGMESH_APIS_SRC \
  --go_out ./bot/v1 \
  --go_opt paths=source_relative \
  --go-grpc_out ./bot/v1 \
  --go-grpc_opt paths=source_relative \
  bot.proto
```
