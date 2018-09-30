# Kingpinger Golang

## Install
This project requires [dep](https://github.com/golang/dep)

After installing dep for your system, run this command in your project root to get all dependencies:
```bash
dep ensure
```

## Running

Running the main process:
```bash
go run main.go
```

### Environment variables

These are layovers from the Python project, might not be used going forward

- `port` - the port that kingpinger will listen on
- `host` - unused, will probably be removed later
- `service` - kingping service url, not used right now

## Payloads

### request payload (subject to change)
```json
{
  "job_id": "402390-4942",
  "callback_url": "http://1e865c40.ngrok.io",
  "target": "google.com",
  "count": 5
}
```

### response payload (subject to change)
```json
{
    "job_id": "402390-4942",
    "PacketsRecv": 5,
    "PacketsSent": 5,
    "PacketLoss": 0,
    "IPAddr": {
        "IP": "216.58.210.46",
        "Zone": ""
    },
    "Addr": "google.com",
    "Rtts": [
        15001900,
        15003000,
        15003100,
        15003400,
        15002800
    ],
    "MinRtt": 15001900,
    "MaxRtt": 15003400,
    "AvgRtt": 15002840,
    "StdDevRtt": 508
}
```