package gallery

type InMemory struct {
	P []Photo
}

func (i InMemory) Photos() []Photo {
	return i.P
}

func (i InMemory) ScanForChanges() {}
