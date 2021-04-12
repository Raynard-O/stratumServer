package db

type Request struct {
	ID        string `json:"id" db:"id"`
	RequestedAt string `json:"requested_at" db:"requested_at"`

	CreatedOn string `json:"created_on" db:"created_on"`
}
