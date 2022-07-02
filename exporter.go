package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	targetDeviceId string

	metricsCurrent          *prometheus.Desc
	metricsVoltage          *prometheus.Desc
	metricsUptime           *prometheus.Desc
	metricsPowerConsumption *prometheus.Desc
	metricsPowerUp          *prometheus.Desc
}

func NewExporter(target string) *Exporter {
	defaultLabels := []string{"device_id", "device_type"}

	return &Exporter{
		targetDeviceId: target,

		metricsCurrent: prometheus.NewDesc(
			"electricity_current",
			"Current electrical current (A)",
			defaultLabels,
			prometheus.Labels{},
		),

		metricsVoltage: prometheus.NewDesc(
			"electricity_voltage",
			"Current electrical voltage (V)",
			defaultLabels,
			prometheus.Labels{},
		),

		metricsUptime: prometheus.NewDesc(
			"usage_minutes",
			"How long the device has been on for the day (minutes)",
			defaultLabels,
			prometheus.Labels{},
		),

		metricsPowerConsumption: prometheus.NewDesc(
			"power_consumption",
			"Power consumption of the plug (W/min)",
			defaultLabels,
			prometheus.Labels{},
		),

		metricsPowerUp: prometheus.NewDesc(
			"power_up",
			"Whether the device is up or down",
			defaultLabels,
			prometheus.Labels{},
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.metricsCurrent
	ch <- e.metricsVoltage
	ch <- e.metricsUptime
	ch <- e.metricsPowerConsumption
	ch <- e.metricsPowerUp
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	statusBody, err := retrievePlugStatus(e.targetDeviceId)
	if err != nil {
		log.Fatal(err)
	}
	ch <- prometheus.MustNewConstMetric(e.metricsCurrent, prometheus.GaugeValue, statusBody.ElectricCurrent, statusBody.DeviceId, statusBody.DeviceType)
	ch <- prometheus.MustNewConstMetric(e.metricsVoltage, prometheus.GaugeValue, statusBody.Voltage, statusBody.DeviceId, statusBody.DeviceType)
	ch <- prometheus.MustNewConstMetric(e.metricsUptime, prometheus.GaugeValue, statusBody.ElectricityOfDay, statusBody.DeviceId, statusBody.DeviceType)
	ch <- prometheus.MustNewConstMetric(e.metricsPowerConsumption, prometheus.GaugeValue, statusBody.Weight, statusBody.DeviceId, statusBody.DeviceType)

	if statusBody.Power == "on" {
		ch <- prometheus.MustNewConstMetric(e.metricsPowerUp, prometheus.GaugeValue, 1, statusBody.DeviceId, statusBody.DeviceType)
	} else {
		ch <- prometheus.MustNewConstMetric(e.metricsPowerUp, prometheus.GaugeValue, 0, statusBody.DeviceId, statusBody.DeviceType)
	}
}
