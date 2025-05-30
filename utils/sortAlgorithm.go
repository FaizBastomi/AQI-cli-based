package utils

func SelSortDescByIdxUdara(A *AirPolutions) {
	var i, idx, pass int

	for pass = 0; pass < len(*A)-1; pass++ {
		idx = pass
		for i = pass + 1; i < len(*A); i++ {
			if (*A)[i].IdxUdara > (*A)[idx].IdxUdara {
				idx = i
			}
		}
		(*A)[pass], (*A)[idx] = (*A)[idx], (*A)[pass]
	}
}

func SelSortDescByTime(A *AirPolutions) {
	var i, idx, pass int
	var j, lastValid int

	for j = range *A {
		if (*A)[j].AqiID != "" {
			lastValid++
		} else {
			break
		}
	}

	for pass = 0; pass < lastValid-1; pass++ {
		idx = pass
		for i = pass + 1; i < lastValid; i++ {
			if (*A)[i].Waktu.After((*A)[idx].Waktu) {
				idx = i
			}
		}
		(*A)[pass], (*A)[idx] = (*A)[idx], (*A)[pass]
	}
}

func InsSortDescByIdxUdara(A *AirPolutions) {
	var i, j int
	var key AirPolution

	for i = 1; i < len(*A); i++ {
		key = (*A)[i]
		j = i - 1

		for j >= 0 && (*A)[j].IdxUdara < key.IdxUdara {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}

func InsSortDescByTime(A *AirPolutions) {
	var i, j int
	var key AirPolution

	for i = 1; i < len(*A); i++ {
		key = (*A)[i]
		j = i - 1

		for j >= 0 && (*A)[j].Waktu.Before(key.Waktu) {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}
