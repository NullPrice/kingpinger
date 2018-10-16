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
    "Statistics": {
        "PacketsRecv": 5,
        "PacketsSent": 5,
        "PacketLoss": 0,
        "IPAddr": {
            "IP": "216.58.198.238",
            "Zone": ""
        },
        "Addr": "google.com",
        "Rtts": [
            18004400,
            18003400,
            18003500,
            18003400,
            18004300
        ],
        "MinRtt": 18003400,
        "MaxRtt": 18004400,
        "AvgRtt": 18003800,
        "StdDevRtt": 451
    }
}
```