# **MQTT BT Monitor**

Repeatedly scan BT devices nearby from Raspberry Pi and send status to MQTT.

## **Prerequisities**
* Raspberry PI (tested on 3A+, 3B, Zero W)
* Docker Engine
* Docker Compose

## **Installation**
```
mkdir mqtt-bt-monitor
cd mqtt-bt-monitor
git clone https://github.com/vaclavhorejsi/mqtt-bt-monitor .
cp config.json.sample config.json
docker-compose build
```

## **Configuration**
Edit config.json file to your needs
```
{
    "delay": 30,
    "debug": false,
    "devices": [
        {
            "name": "my-phone",
            "mac": "XX:XX:XX:XX:XX:XX"
        }
    ],
    "mqtt": {
        "server": "mqtt.example.com",
        "port": "1883",
        "topic": "btmonitor",
        "username": "",
        "password": ""
    }
}
```

## **Run**
```
docker-compose up -d
```

## **Check logs**
If "debug" is se to true, you are able to check detailed log
```
docker-compose logs -f
```

## **Shutdown**
```
docker-compose down
```

## **Update**
```
cd mqtt-bt-monitor
git pull
docker-compose build
```
