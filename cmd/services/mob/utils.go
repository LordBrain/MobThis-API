package mob

import "math/rand"

// getRandomMobber returns a random mobber from the given slice, excluding specified mobbers if provided
func getRandomMobber(mobbers []string, exclude ...string) string {
	if len(mobbers) == 0 {
		return "" // Return an empty string if the slice is empty
	}

	// Create a map to efficiently check for excluded mobbers
	excluded := make(map[string]bool)
	for _, e := range exclude {
		excluded[e] = true
	}

	// Create a filtered slice without excluded mobbers
	filteredMobbers := make([]string, 0, len(mobbers))
	for _, mobber := range mobbers {
		if !excluded[mobber] {
			filteredMobbers = append(filteredMobbers, mobber)
		}
	}

	if len(filteredMobbers) == 0 {
		return "" // Return an empty string if the filtered slice is empty
	}

	randomIndex := rand.Intn(len(filteredMobbers))
	return filteredMobbers[randomIndex]
}
