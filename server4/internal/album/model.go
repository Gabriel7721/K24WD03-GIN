package album

type album struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Artist  string  `json:"artist"`
	Price   float64 `json:"price"`
	OwnerID string  `json:"owner_id"`
}
type albumForPatch struct {
	Title  *string  `json:"title"`
	Artist *string  `json:"artist"`
	Price  *float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Thriller", Artist: "Michael Jackson", Price: 1500.99, OwnerID: "u1"},
	{ID: "2", Title: "Back in Black", Artist: "AC/DC", Price: 2500.99, OwnerID: "u2"},
	{ID: "3", Title: "The Bodyguard", Artist: "Whitney Houston", Price: 3500.99, OwnerID: "u2"},
}
