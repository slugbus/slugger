package taps

// This function merges a new ping response with the current
// response
func mergeWithState(newPing BusMap, time float64, currentBusMap UpdatedBusMap) UpdatedBusMap {
	// Prepare the new state.
	newUpdatedBusMap := UpdatedBusMap{}

	for key, pingBus := range newPing {
		// Prepare a new UpdatedBus struct for each bus in the new ping
		newUpdatedBus := UpdatedBus{
			Bus: pingBus,
		}
		// Check to see if the bus we're looking
		// at was in our previous state, if it was
		// update it with speed and angle value
		//isBusStillRunning = currentBusMap[key] // is the key in the existing UpdatedBusMap
		if _, exists := currentBusMap[key]; exists {
			//testing
			newUpdatedBus.Speed = 100.0
			newUpdatedBus.Angle = 3000.0

			//distance := geo.Dist(stateBus.Lat, stateBus.Lon, pingBus.Lat, pingBus.Lon)
			//newUpdatedBus.Speed = geo.Speed(distance, t)
			//newUpdated.Angle = geo.Dir(stateBus.Lat, stateBus.Lon, pingBus.Lat, pingBus.Lon)
		}

		/*if stateBus, isInState := currentBusMap[key]; isInState {

			//distance := geo.Dist(stateBus.Lat, stateBus.Lon, pingBus.Lat, pingBus.Lon)
			//dataPoint.Speed = geo.Speed(distance, t)
			//dataPoint.Angle = geo.Dir(stateBus.Lat, stateBus.Lon, pingBus.Lat, pingBus.Lon)
		} */

		// Always push the new bus into the new updated bus map
		newUpdatedBusMap[key] = newUpdatedBus
	}

	return newUpdatedBusMap
}

/* // Merge update takes in two regular responses from the server
// and t (that is in milliseconds) and combines them to get speed
// and angle data.
func mergeUpdate(p, q Bus, t float64) UpdatedBusMap {
	// Make of map of strings
	// to buses
	mb := map[string]taps.Bus{}
	// Loop through first
	// ping
	for _, bus := range p {
		// Map the bus ID to the
		// bus datastructure
		mb[bus.ID] = bus
	}
	// Prepare a result
	result := UpdatedBus{}
	// Loop through the second ping
	for _, pingTwoBus := range q {
		// Make a bus with angles and speed
		d := BusDataPlusPlus{}
		// Add the buses' data to the bus++?
		d.Bus = pingTwoBus
		// Check if the current bus exists in ping one
		if pingOneBus, contains := mb[d.ID]; contains {
			// If it does, calculate its distance, speed , and angle
			distance := geo.Dist(pingOneBus.Lat, pingOneBus.Lon, pingTwoBus.Lat, pingTwoBus.Lon)
			d.Speed = geo.Speed(distance, t)
			d.Angle = geo.Dir(pingOneBus.Lat, pingOneBus.Lon, pingTwoBus.Lat, pingTwoBus.Lon)
		}
		// push the bus to the result
		result = append(result, d)
	}
	return result
}
*/
