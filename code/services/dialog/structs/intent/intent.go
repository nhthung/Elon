package intent

type Intent struct {
	Name string
	TrainPhrases []string
	Responses [][]string
	FollowupIntents []string
	Contexts []string
}