package payload

type CountVotesResponse struct {
	ID    uint32 `json:"id"`
	Votes uint32 `json:"votes"`
}
