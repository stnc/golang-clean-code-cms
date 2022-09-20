package main

import "fmt"

/**
 * Observer interface is implemented by all observers,
 * Here we're passing the measurements to the observers
 * from the Subject when a weather measurement changes
 */
type observer interface {
	update(temp float32, humidity float32, pressure float32)
}
type subject interface {
	/**
	 * Both of these methods take an Observer as argument
	 */
	registerObserver(observer)
	deregisterObserver(observer)
	/**
	 * This method is called to notify all the observers
	 * when the Subject's state has changed
	 */
	notifyObservers()
} /**
 * DisplayElement interface just includes display() method
 * that we will call when the info needs to be displayed
 */
type displayElement interface {
	display()
}

type weatherData struct {
	/**
	 * Implementation of set in Go
	 * Instantiating in the constructor (newWeatherData)
	 */
	observers   map[observer]bool
	temperature float32
	humidity    float32
	pressure    float32
}

func newWeatherData() *weatherData {
	return &weatherData{
		observers: make(map[observer]bool),
	}
}

/**
 * WeatherData now implements the Subject interface
 */
func (w *weatherData) registerObserver(o observer) {
	/**
	 * When an observer registers, we add it in
	 * the map with value set as true
	 */
	w.observers[o] = true
}

func (w *weatherData) deregisterObserver(o observer) {
	/**
	 * When an observer deregisters, we remove it
	 * from the map after checking if the value exists
	 */
	if _, ok := w.observers[o]; ok {
		delete(w.observers, o)
	}
}

func (w *weatherData) notifyObservers() {
	/**
	 * This is where we tell all the observers about the state
	 * by calling the common update method
	 */
	for observer := range w.observers {
		observer.update(w.temperature, w.humidity, w.pressure)
	}
}

/**
 * We notify the Observers when we get updated
 * measurements from the Weather Station
 */
func (w *weatherData) measurementsChanged() {
	w.notifyObservers()
}

/**
 * Dummy method to test our display elements.
 * Setting the measurements not via a device.
 */
func (w *weatherData) setMeasurements(temp float32, humidity float32, pressure float32) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}

type currentConditionDisplay struct {
	temperature float32
	humidity    float32
	pressure    float32
}

func newCurrentConditionDisplay() *currentConditionDisplay {
	return &currentConditionDisplay{}
}

/**
 * This implements Observer and DisplayElement interfaces to get changes from WeatherData
 * and show the information based on its functionality respectively
 */
func (ccd *currentConditionDisplay) update(temp float32, humidity float32, pressure float32) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.pressure = pressure
	/**
	 * When update() is called, we save the measurements
	 * and call display() to show the required information
	 */
	ccd.display()
}

func (ccd *currentConditionDisplay) display() {
	fmt.Printf("Current conditions: Temperature:%.2f, Humidity:%.2f and Pressure:%.2f\n", ccd.temperature, ccd.humidity, ccd.pressure)
}

type statisticsDisplay struct {
	count   uint32
	avgTemp float32
	maxTemp float32
	minTemp float32
}

func newStatisticsDisplay() *statisticsDisplay {
	return &statisticsDisplay{}
}

/**
 * This implements the avg, max and min temperatures.
 * When update is called, it does the calculation and
 * call the display() method to show the required information
 */
func (sd *statisticsDisplay) update(temp float32, humidity float32, pressure float32) {
	sd.count++
	sd.avgTemp -= (sd.avgTemp - temp) / float32(sd.count)

	if sd.maxTemp < temp || sd.maxTemp == 0.0 {
		sd.maxTemp = temp
	}

	if sd.minTemp > temp || sd.minTemp == 0.0 {
		sd.minTemp = temp
	}

	sd.display()
}

func (sd *statisticsDisplay) display() {
	fmt.Printf("Avg/Max/Min Temperature:%.2f, %.2f, %.2f\n", sd.avgTemp, sd.maxTemp, sd.minTemp)
}

func main() {
	/**
	 * Creating the weather object
	 */
	weatherData := newWeatherData()

	/**
	 * Creating the displays and registering as
	 * observer to the Subject
	 */
	currentConditionDisplay := newCurrentConditionDisplay()
	weatherData.registerObserver(currentConditionDisplay)

	statisticsDisplay := newStatisticsDisplay()
	weatherData.registerObserver(statisticsDisplay)
	//forecastDisplay := newForecastDisplay()
	//weatherData.registerObserver(forecastDisplay)

	/**
	 * Simulate new weather measurements
	 */
	weatherData.setMeasurements(80, 65, 30.4)
	weatherData.setMeasurements(82, 70, 29.2)
	weatherData.setMeasurements(77, 90, 29.2)
}
