func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for i, value := range flowerbed {
		if (value == 0) && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}

	return count >= n
}


/*
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true