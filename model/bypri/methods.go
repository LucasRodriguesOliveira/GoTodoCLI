package byprimodel

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}

	if s[i].Priority == s[j].Priority {
		return s[i].Pos() < s[j].Pos()
	}

	return s[i].Priority < s[j].Priority
}
