# Docker Cache Cleaner

## What I do?
```shell
systemctl stop $SERVICE_NAME
docker system prune -f
docker volume rm -f $(docker volume list -q)
systemctl start $SERVICE_NAME
```

## Why Go?
Because Golang is simple and effective tool for system administrator. :grin:

