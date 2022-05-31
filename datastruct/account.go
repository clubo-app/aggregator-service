package datastruct

type AggregatedAccount struct {
	Id      string            `json:"id"`
	Profile AggregatedProfile `json:"profile"`
	Email   string            `json:"email"`
	Type    string            `json:"type,omitempty"`
	Role    string            `json:"role,omitempty"`
}
