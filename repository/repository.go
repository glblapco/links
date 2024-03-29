//In a real application, the repository should talk to a DB.
//DB may be needed to this app, so it may be implemented in the future.

package repository

import (
	"proj/links/models"
)

var links []models.Link
var nID = 1

func GetLinks() []models.Link {
	return links
}

func AddLinks(link models.Link) int {
	link.ID = nID
	nID++
	links = append(links, link)
	return link.ID
}
