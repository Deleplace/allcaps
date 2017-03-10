package allcaps

func Search(caps string) []string {
	phrases := db[caps]
	if len(phrases) > 3 {
		phrases = phrases[:3]
	}
	return phrases
}
