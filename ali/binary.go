package ali




func Binary(nb int64) []int {
	var jmi []int 
	for nb > 0{
		jmi = append(jmi, int(nb % 2))
		nb /= 2	
	}
	for i, j := 0, len(jmi)-1; i < j; i, j = i+1, j-1 {
		jmi[i], jmi[j] = jmi[j], jmi[i]
	}
		return jmi
}