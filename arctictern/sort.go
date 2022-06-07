package arctictern

type ByFileVersion []string

func (a ByFileVersion) Len() int {
	return len(a)
}

func (a ByFileVersion) Less(i, j int) bool {
	ai := GetVersion(a[i])
	aj := GetVersion(a[j])
	return ai < aj
}

func (a ByFileVersion) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
