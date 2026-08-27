package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type weatherAPI struct {
	endpoint string
	timeout  time.Duration
}

type Response struct {
	Hourly struct {
		Time               []string  `json:"time"`
		Temperature2m      []float64 `json:"temperature_2m"`
		RelativeHumidity2m []float64 `json:"relative_humidity_2m"`
		WindSpeed10m       []float64 `json:"wind_speed_10m"`
		Rain               []float64 `json:"rain"`
		WeatherCode        []int     `json:"weather_code"`
	} `json:"hourly"`
}

func (w weatherAPI) parseTime(str string) (time.Time, error) {
	layout := "2006-01-02T15:04"
	return time.Parse(layout, str)
}

func (w weatherAPI) request(ctx context.Context, city City, parameter parameter, units Units, days int) (Response, error) {
	ctx, cancel := context.WithTimeout(ctx, w.timeout)
	defer cancel()
	url := fmt.Sprintf(
		"%s?latitude=%.2f&longitude=%.2f%s%s&forecast_days=%d",
		w.endpoint,
		city.Lat, city.Long,
		parameter.param(), units.param(),
		days,
	)
	var r Response
	for attempt := range 3 {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return r, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "doors-tutorial/1.0")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			if ctx.Err() != nil {
				return r, ctx.Err()
			}
			log.Printf("weather request failed: attempt=%d city=%q parameter=%q days=%d err=%v", attempt+1, city.Name, parameter, days, err)
			if attempt < 2 && waitRetry(ctx, attempt) == nil {
				continue
			}
			return r, err
		}
		err = decodeWeatherResponse(res, &r)
		if err == nil {
			break
		}
		log.Printf("weather request failed: attempt=%d city=%q parameter=%q days=%d err=%v", attempt+1, city.Name, parameter, days, err)
		if attempt < 2 && waitRetry(ctx, attempt) == nil {
			continue
		}
		return r, err
	}
	for i, v := range r.Hourly.Time {
		t, err := w.parseTime(v)
		if err != nil {
			return r, err
		}
		if days < 3 {
			r.Hourly.Time[i] = t.Format("15:04") + " "
		} else {
			r.Hourly.Time[i] = t.Format("02.01")
		}
	}
	return r, nil
}

func decodeWeatherResponse(res *http.Response, out *Response) error {
	defer res.Body.Close()
	if res.StatusCode == http.StatusBadGateway ||
		res.StatusCode == http.StatusServiceUnavailable ||
		res.StatusCode == http.StatusGatewayTimeout ||
		res.StatusCode == http.StatusTooManyRequests {
		return fmt.Errorf("transient weather api status: %s", res.Status)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("weather api status: %s", res.Status)
	}
	return json.NewDecoder(res.Body).Decode(out)
}

func waitRetry(ctx context.Context, attempt int) error {
	delay := time.Duration(attempt+1) * 100 * time.Millisecond
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

type FloatSamples struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

type StringSamples struct {
	Labels []string `json:"labels"`
	Values []string `json:"values"`
}

func (w weatherAPI) Humidity(ctx context.Context, city City, days int) (FloatSamples, error) {
	r, err := w.request(ctx, city, humidity, noUnits, days)
	if err != nil {
		return FloatSamples{}, err
	}
	samples := FloatSamples{
		Labels: make([]string, len(r.Hourly.Time)),
		Values: make([]float64, len(r.Hourly.Time)),
	}
	for i := range r.Hourly.Time {
		samples.Labels[i] = r.Hourly.Time[i]
		samples.Values[i] = r.Hourly.RelativeHumidity2m[i]
	}
	return samples, nil
}

func (w weatherAPI) Temperature(ctx context.Context, city City, units Units, days int) (FloatSamples, error) {
	r, err := w.request(ctx, city, temperature, units, days)
	if err != nil {
		return FloatSamples{}, err
	}
	samples := FloatSamples{
		Labels: make([]string, len(r.Hourly.Time)),
		Values: make([]float64, len(r.Hourly.Time)),
	}
	for i := range r.Hourly.Time {
		samples.Labels[i] = r.Hourly.Time[i]
		samples.Values[i] = r.Hourly.Temperature2m[i]
	}
	return samples, nil
}

func (w weatherAPI) WindSpeed(ctx context.Context, city City, units Units, days int) (FloatSamples, error) {
	r, err := w.request(ctx, city, windSpeed, units, days)
	if err != nil {
		return FloatSamples{}, err
	}
	samples := FloatSamples{
		Labels: make([]string, len(r.Hourly.Time)),
		Values: make([]float64, len(r.Hourly.Time)),
	}
	for i := range r.Hourly.Time {
		samples.Labels[i] = r.Hourly.Time[i]
		samples.Values[i] = r.Hourly.WindSpeed10m[i]
	}
	return samples, nil
}

func (w weatherAPI) Code(ctx context.Context, city City, days int) (StringSamples, error) {
	r, err := w.request(ctx, city, weatherCode, Metric, days)
	if err != nil {
		return StringSamples{}, err
	}
	samples := StringSamples{
		Labels: make([]string, len(r.Hourly.Time)),
		Values: make([]string, len(r.Hourly.Time)),
	}
	for i := range r.Hourly.Time {
		samples.Labels[i] = r.Hourly.Time[i]
		str, ok := weatherCodeShort[r.Hourly.WeatherCode[i]]
		if !ok {
			str = "unknown"
		}
		samples.Values[i] = str
	}
	return samples, nil
}

var weatherCodeShort = map[int]string{
	0:  "Clear",
	1:  "Mainly clear",
	2:  "Partly cloudy",
	3:  "Overcast",
	45: "Fog",
	48: "Rime fog",
	51: "Drizzle light",
	53: "Drizzle mod",
	55: "Drizzle dense",
	56: "Frzg drizzle lgt",
	57: "Frzg drizzle hvy",
	61: "Rain light",
	63: "Rain mod",
	65: "Rain heavy",
	66: "Frzg rain lgt",
	67: "Frzg rain hvy",
	71: "Snow light",
	73: "Snow mod",
	75: "Snow heavy",
	77: "Snow grains",
	80: "Shower rain lgt",
	81: "Shower rain mod",
	82: "Shower rain hvy",
	85: "Snow shower lgt",
	86: "Snow shower hvy",
	95: "Thunderstorm",
	96: "Storm + small hail",
	99: "Storm + heavy hail",
}

type Units int

const (
	Metric Units = iota
	Imperial
	noUnits
)

func (u Units) String() string {
	if u == Imperial {
		return "Imperial"
	}
	if u == Metric {
		return "Metric"
	}
	return "unknown"
}

func (u Units) WindSpeed() string {
	if u == Imperial {
		return "MPH"
	}
	if u == Metric {
		return "KMH"
	}
	return "unknown"
}

func (u Units) Temperature() string {
	if u == Imperial {
		return " °F"
	}
	if u == Metric {
		return " °C"
	}
	return "unknown"
}

func (u Units) param() string {
	if u == Metric {
		return "&wind_speed_unit=kmh&temperature_unit=celsius&precipitation_unit=mm"
	}
	if u == Imperial {
		return "&wind_speed_unit=mph&temperature_unit=fahrenheit&precipitation_unit=inch"
	}
	return ""
}

type parameter string

const (
	temperature parameter = "temperature_2m"
	humidity    parameter = "relative_humidity_2m"
	windSpeed   parameter = "wind_speed_10m"
	weatherCode parameter = "weather_code"
)

func (u parameter) param() string {
	return "&hourly=" + string(u)
}
