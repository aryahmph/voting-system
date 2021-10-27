package payload

type CountVotesResponse struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Votes uint32 `json:"votes"`
}
