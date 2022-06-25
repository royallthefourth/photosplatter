package gallery

import "github.com/jfcg/sorty"

func SortPhotosDesc(p []Photo) {
	lsw := func(i, k, r, s int) bool {
		if p[i].Created.After(p[k].Created) {
			if r != s {
				p[r], p[s] = p[s], p[r]
			}
			return true
		}
		return false
	}
	sorty.Sort(len(p), lsw)
}
