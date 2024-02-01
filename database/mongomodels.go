package database

// bson normally takes the lower case of the fields
// omiempty gives the value null if the field is empty
// omiempty didn't work with strings => just an empty string
type BlogPost struct {
	Author           string
	Title            string
	Description      string
	Content          string
	Tags             []string `bson:"tags, omiempty"`
	Likes            int
	NumberOfComments int `bson:"_comments, omiempty"`
	Comments         []string
}
