## Useful Tool For Monitor CPU Stat

## GetStarted

### Step 1 clone this repo

```bash
git clone https://github.com/ooopSnake/pressure_watcher
```

### Step 2 build

```bash
./build_all.sh
```

### Step 3 execute

```bash
sudo ./pressure_watcher_arm
```

## Result

```json
 {
  "cpuInfo": [
    {
      "id": "cpu 0",
      "freq": "1.008 GHz",
      "usage": "21.6 %"
    },
    {
      "id": "cpu 1",
      "freq": "1.008 GHz",
      "usage": "21.1 %"
    },
    {
      "id": "cpu 2",
      "freq": "1.008 GHz",
      "usage": "25.5 %"
    },
    {
      "id": "cpu 3",
      "freq": "1.008 GHz",
      "usage": "23.5 %"
    }
  ],
  "cpuTemp": "39.9 °C"
}
```
## CommandFlags

```bash
Usage of ./pressure_watcher:
  -addr string
    	http addr , eg : 127.0.0.1 (default "0.0.0.0")
  -nohttp
    	disable http server
  -port string
    	http listen port , eg : 8080 (default "12345")
```

## License
WTFPL