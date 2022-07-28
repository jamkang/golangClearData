package rand_te

//利用随机数，对数组的打乱顺序

func ShuffleCard(cards []int) {
	for i := len(cards) - 1; i > 0; i-- {
		key := GetRand(i)
		tem := cards[key]
		cards[key] = cards[i]
		cards[i] = tem
	}
}
