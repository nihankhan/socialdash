// home_handler.go

package api

import (
	"net/http"
	"text/template"

	"github.com/nihankhan/socialdash/internal/services"
)

// HomeHandler handles the home page/dashboard
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch follower count from Facebook service
	fbCount, err := services.GetFacebookFriendCount()
	if err != nil {
		http.Error(w, "Error fetching Facebook follower count", http.StatusInternalServerError)
		return
	}

	// fmt.Println("fbCount: ", fbCount)
	/*
		//twCount, err := services.GetTwitterFollowerCount()
		if err != nil {
			http.Error(w, "Error fetching Twitter follower count", http.StatusInternalServerError)
			return
		}

		fmt.Println("tw: ", twCount)

	*/

	/*igCount, err := services.GetInstagramFollowerCount()
	if err != nil {
		http.Error(w, "Error fetching Instagram follower count", http.StatusInternalServerError)
		return
	} */

	// Render the dashboard template with follower counts
	tmpl, err := template.ParseFiles("./web/templates/dashboard.html")
	if err != nil {
		http.Error(w, "Error in HTML parsing", http.StatusInternalServerError)
		return
	}

	data := struct {
		FacebookCount int
		TwitterCount  int
		//InstagramCount int
	}{
		FacebookCount: fbCount,
		//TwitterCount:  twCount,
		//InstagramCount: igCount,
	}

	tmpl.Execute(w, data)
}
