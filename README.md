# switchbot-plug-exporter

Export switchbot plug metrics.

## Usage

Run from docker or binaries in Github releases.

Docker images are located in https://hub.docker.com/repository/docker/rw21/switchbot-plug-exporter. `arm64` images are also available.

```
docker run -p 9101:9101 -e SWITCHBOT_TOKEN=<token> rw21/switchbot-plug-exporter:latest

```

The application will use the following environment variables:

| Environment variable | Default | Description                    |
|----------------------|---------|--------------------------------|
| `PORT`                 | 9101    | Port the exporter will run on. |
| `SWITHBOT_TOKEN`       | ""      | Switchbot api token            |

## Metrics

| Metric              | Type  | Description                                           |
|---------------------|-------|-------------------------------------------------------|
| power_consumption   | guage | Power consumption of the plug (W/min)                 |
| electricity_current | guage | Current electrical current (A)                        |
| electricity_voltage | guage | Current electrical voltage (V)                        |
| power_up            | guage | Whether the device is up or down                      |
| usage_minutes       | guage | How long the device has been on for the day (minutes) |

## Prometheus Configuration

You will need the ids of each of your plugs.

The ids can be determined using the switchbot api. 
`curl -XGET -H 'Authorization: AUTHTOKEN' 'https://api.switch-bot.com/v1.0/devices'`


```yaml
- job_name: 'plugmini'
  static_configs:
  - targets: 
    - <device ids>
    - <device ids>
  metrics_path: /scrape
  relabel_configs:
    - source_labels : [__address__]
      target_label: __param_target
    - source_labels: [__param_target]
      target_label: instance
    - target_label: __address__
      replacement: <address of exporter e.g. localhost:9101>
```

## See also

- https://github.com/fffonion/tplink-plug-exporter
- https://github.com/OpenWonderLabs/SwitchBotAPI#switchbot-meter-example