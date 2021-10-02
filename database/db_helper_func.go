package database

// TODO change to be a map later so lookup is better, until then suffer
// Check if str is inside the list s
func contains(s []*string, str string) bool {
	for _, v := range s {
		if *v == str {
			return true
		}
	}

	return false
}

func (db *DB) removeElement(list []*string, item string) []*string {
	// find index of item in list
	index := -1
	for i, n := range list {
		if &item == n {
			index = i
		}
	}

	if index != -1 {
		// remove element at that spot
		list = append(list[:index], list[index+1:]...)
	}

	return list

}

func (db *DB) findCommon(list1 []*string, list2 []*string) []*string {

	var commonList []*string

	for i := 0; i < len(list1); i++ {
		if contains(list2, *list1[i]) {
			commonList = append(commonList, list1[i])
		}
	}

	return commonList
}
