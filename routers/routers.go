package routers

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vpakhuchyi/web-server/models"
)

//GETJSONHandler is a GET handler for "/searchText"
func GETJSONHandler(c *gin.Context) {
	t, _ := template.ParseFiles("templates/default.tmpl")
	t.Execute(c.Writer, nil)
}

//POSTJSONHandler is a POST handler for "/searchText";
//it checks incoming JSON and sends a result of "searchForArgsOnEachSite" func as JSON response.
func POSTJSONHandler(c *gin.Context) {
	var reqjson models.Request
	if c.Bind(&reqjson) == nil && len(reqjson.Site) > 0 && len(reqjson.SearchText) > 0 {
		var respjson models.Response
		result, err := SearchForArgOnSites(reqjson.SearchText, reqjson.Site)
		code := 400
		if err != nil {
			c.Error(err)

			if err.Error() == "no content" {
				code = 204
			}
			c.Writer.WriteString("HTTP code " + strconv.Itoa(code) + " " + err.Error())
		} else {
			respjson.FoundAtSite = result
			c.JSON(http.StatusOK, respjson)
		}

	}
}

func getContentFromURL(url string) []byte {
	if url != "" {
		res, _ := http.Get(url)
		c, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		return c
	}
	return nil
}

func checkConnectionToURL(url string) error {
	_, err := http.Get(url)
	return err
}

//SearchForArgOnSites receive a string "text" for search and a []string "sites" for searching place;
//func returns string and error;
//if text was found it returnes string with URL where it was found and nil error;
//func returns empty string and not-nil error when text wasn't found.
func SearchForArgOnSites(text string, sites []string) (r string, err error) {
	if text == "" {
		err = errors.New("invailid request")
		return r, err
	}
	retext := regexp.MustCompile(text)

	for _, val := range sites {

		if checkConnectionToURL(val) != nil {
			err = errors.New("invailid request")
			continue
		}

		if retext.Find(getContentFromURL(val)) != nil {
			r = val
			err = nil
			break
		} else {
			err = errors.New("no content")
		}

	}

	return r, err
}
