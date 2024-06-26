package metric

import (
	"errors"
	"math/rand"
)

type (
	Metric interface {
		Kind() string   // Kind - возвращает тип метрики.
		String() string // Stringer.
		IsCounter() bool
		IsGauge() bool
	}

	Metrics struct {
		// метрики пакета runtime.
		Memory MemoryMetrics
		// дополнительные метрики.
		// PollCount - счётчик, увеличивающийся на 1 при каждом обновлении метрики из пакета runtime.
		PollCount Counter
		// RandomValue - обновляемое произвольное значение.
		RandomValue Gauge
	}
)

var (
	// ErrorMetricNameIsNull - не указано имя метрики.
	ErrorMetricNameIsNull = errors.New("metrics: ошибка cоздания метрики, не указано име метрики")
	// ErrorMetricNotFound - метрика не найдена.
	ErrorMetricNotFound = errors.New("metrics: метрика не найдена")
)

func (m *Metrics) Poll() {
	m.PollCount += 1
	m.RandomValue = Gauge(rand.Float64())
	m.Memory.Poll()
}
