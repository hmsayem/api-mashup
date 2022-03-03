### Combines responses from two different API endpoints using Goroutines and Channels.

### Run
```bash
go run main.go
```

### Example
```bash
❯ curl -XGET  "http://localhost:8000/cars" | jq
[
  {
    "id": 1,
    "brand": "Mitsubishi",
    "model": "Montero",
    "model_year": 2002,
    "owner_first_name": "Alyosha",
    "owner_last_name": "Caldero",
    "owner_email": "acaldero0@behance.net"
  },
  {
    "id": 2,
    "brand": "Volkswagen",
    "model": "Passat",
    "model_year": 2008,
    "owner_first_name": "Danyelle",
    "owner_last_name": "Whoolehan",
    "owner_email": "dwhoolehan1@hostgator.com"
  },
  {
    "id": 3,
    "brand": "Saturn",
    "model": "L-Series",
    "model_year": 2003,
    "owner_first_name": "Raimund",
    "owner_last_name": "MacAughtrie",
    "owner_email": "rmacaughtrie2@discovery.com"
  },
  {
    "id": 4,
    "brand": "Jeep",
    "model": "Compass",
    "model_year": 2012,
    "owner_first_name": "Kathleen",
    "owner_last_name": "McMickan",
    "owner_email": "kmcmickan3@npr.org"
  },
  {
    "id": 5,
    "brand": "Mitsubishi",
    "model": "Lancer Evolution",
    "model_year": 2002,
    "owner_first_name": "Joshuah",
    "owner_last_name": "Frisdick",
    "owner_email": "jfrisdick4@yale.edu"
  }
]
```