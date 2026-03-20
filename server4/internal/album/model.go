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
	{ID: "1", Title: "Thriller", Artist: "Michael Jackson", Price: 1500.99, OwnerID: "u3"},
	{ID: "2", Title: "Back in Black", Artist: "AC/DC", Price: 2500.99, OwnerID: "u3"},
	{ID: "3", Title: "The Bodyguard", Artist: "Whitney Houston", Price: 3500.99, OwnerID: "u3"},
}

func GetAll() []album {
	return albums
}
func GetOne(id string) *album {
	for i := range albums {
		if albums[i].ID == id {
			return &albums[i]
		}
	}
	return nil
}
func CreateOne(a album) album {
	albums = append(albums, a)
	return a
}
func UpdateOne(id string, updated album) (*album, bool) {
	for i := range albums {
		if albums[i].ID == id {
			updated.ID = id
			albums[i] = updated
			return &albums[i], true
		}
	}
	return nil, false
}
func UpdatePatchOne(id string, updated albumForPatch) (*album, bool) {
	for i := range albums {
		if albums[i].ID == id {
			if updated.Title != nil {
				albums[i].Title = *updated.Title
			}
			if updated.Artist != nil {
				albums[i].Artist = *updated.Artist
			}
			if updated.Price != nil {
				albums[i].Price = *updated.Price
			}
			return &albums[i], true
		}
	}
	return nil, false
}
func DeleteOne(id string) bool {
	for i := range albums {
		if albums[i].ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			return true
		}
	}
	return false
}
