# **MQTT BT Monitor**

Repeatedly scan nearby list of BT devices from Raspberry Pi and send status to MQTT.

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
    "delay": 30,                            // Delay between tests
    "debug": false,                         // Detailed log
    "devices": [                            // Device list
        {
            "name": "my-phone",             // Name showed in MQTT
            "mac": "XX:XX:XX:XX:XX:XX"      // BTMAC of tested device
        }
    ],
    "mqtt": {                               // MQTT settings
        "server": "mqtt.example.com",       // Server address (only mqtt protocol supported)
        "port": "1883",                     // Port
        "topic": "btmonitor",               // Topic where messages will be propagate
        "retained": false,                  // Retained messages
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
