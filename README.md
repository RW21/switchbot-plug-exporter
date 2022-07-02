# switchbot-plug-exporter

Export switchbot plug metrics.

## Usage

Run from docker.

```
docker run -p 9101:9101 

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
