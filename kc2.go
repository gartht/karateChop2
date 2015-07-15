package kc2

func Chop(searchVal int, input []int) (locationIndex int) {

	if len(input) == 0 {
		locationIndex = -1
	} else {
		locationIndex = recursiveBinSearch(0, len(input)-1, searchVal, input)
	}
	return
}

func recursiveBinSearch(lowIndex, highIndex, searchVal int, input []int) (locationIndex int) {
	currentIndex := (highIndex + lowIndex) / 2

	if lowIndex == highIndex {
		if input[lowIndex] == searchVal {
			locationIndex = lowIndex
		} else {
			locationIndex = -1
		}
		return
	}

	if highIndex-lowIndex == 1 {
		//this case happens when we have gaps in the sorted array
		if input[highIndex] == searchVal {
			locationIndex = highIndex
		} else {
			if input[lowIndex] == searchVal {
				locationIndex = lowIndex
			} else {
				locationIndex = -1
			}
		}
		return
	}

	if input[currentIndex] == searchVal {
		locationIndex = currentIndex
		return
	}

	if input[currentIndex] < searchVal {
		lowIndex = currentIndex
	} else {
		highIndex = currentIndex
	}
	locationIndex = recursiveBinSearch(lowIndex, highIndex, searchVal, input)
	return
}
