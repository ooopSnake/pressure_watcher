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
          "usage": "0.0%"
        },
        {
          "id": "cpu 1",
          "freq": "1.008 GHz",
          "usage": "0.0%"
        },
        {
          "id": "cpu 2",
          "freq": "1.008 GHz",
          "usage": "1.0%"
        },
        {
          "id": "cpu 3",
          "freq": "1.008 GHz",
          "usage": "0.0%"
        }
      ],
      "cpuTemp": "37.2 °C"
    }

```


## License
WTFPL