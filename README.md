# ipranges tool

This tool fetches the IP ranges used by GCP compute.

Google publishes the raw data here: https://cloud.google.com/compute/docs/faq#find_ip_range.

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

## Fetch the regions

```console
$ ipranges -regions
asia-east1
asia-east2
asia-northeast1
asia-northeast2
asia-northeast3
asia-south1
asia-south2
asia-southeast1
asia-southeast2
australia-southeast1
australia-southeast2
europe-central2
europe-north1
europe-southwest1
europe-west1
europe-west2
europe-west3
europe-west4
europe-west6
europe-west8
europe-west9
global
northamerica-northeast1
northamerica-northeast2
southamerica-east1
southamerica-west1
us-central1
us-central2
us-east1
us-east4
us-east5
us-east7
us-south1
us-west1
us-west2
us-west3
us-west4
```

# Fetch IPs for a specific region

```console
$ ipranges -region us-central2
35.186.0.0/17
35.186.128.0/20
35.206.32.0/19
35.220.46.0/24
35.242.46.0/24
107.167.160.0/20
108.59.88.0/21
173.255.120.0/21
```