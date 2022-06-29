# ipranges tool

Google Cloud publishes a JSON-formatted list of customer-usable global and regional external IP address ranges here: https://www.gstatic.com/ipranges/cloud.json.

This tool fetches that data and provides filtering by region prefix, and other useful options (IPv6, IPv4 etc).

## Installing the tool

```console
$ go install github.com/dhowden/ipranges/cmd/ipranges@latest
$ ipranges -h
Usage of ipranges:
  -fetch-timeout duration
        timeout after duration (default 5s)
  -ipv4
        show IPv4 addresses (default true)
  -ipv6
        show IPv6 addresses (default true)
  -region string
        prefix of region to match
  -regions
        only list regions
```

## List the regions

```console
$ ipranges -regions
asia-east1
asia-east2
asia-northeast1
asia-northeast2
...
```

## List IP ranges for specific region prefixes

For the `us-central2` region:

```console
$ ipranges -region us-central2
107.167.160.0/20
108.59.88.0/21
...
```

For all `us-central` regions:

```console
$ ipranges -region us-central
104.154.113.0/24
104.154.114.0/23
...
```

Exclude IPv6 addresses:

```console
$ ipranges -region us-central2 -ipv6=false
107.167.160.0/20
108.59.88.0/21
...
```

Exclude IPv4 addresses

```console
$ ipranges -region us-central2 -ipv4=false
2600:1900:4070::/44
```
