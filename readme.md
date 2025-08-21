
# 🇮🇷 Persian License Plate Lookup

A simple Golang application to lookup Iranian car license plates.
This project provides a clean JSON dataset of Iranian plates (`plates.json`) and exposes it in two ways:

---

1. **REST API** – query plates via HTTP.
2. **Terminal UI (TUI)** – query plates directly from the command line.



## 🚀 Usage (TUI & API)

---

### 1. Run as API
Start the HTTP server:

```
go run ./cmd/api/api.go
```
Then query a plate (example):

```
curl http://localhost:4000/p?w=46&y=س
```


Response:

```
{
  "province": "گیلان",
  "city": "رشت"
}
```
---

### 2. Run as Terminal App (TUI)
```
go run ./cmd/tui/tui.go
```

Example usage inside terminal:

Enter province code: 11
Enter letter: ب
```
➡ Province: گیلان
➡ City: رشت
```

---

📦 Dataset
	•	plates.json contains the normalized Persian plate data.
	•	Data is structured for O(1) lookup using nested maps:
```
{
  "11": {
    "ب": {
      "province": "گیلان",
      "city": "رشت"
    }
  }
}
```
