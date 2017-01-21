package controllers

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

//Request is a struct for received JSON
type Request struct {
	Site       []string `json:"Site"`
	SearchText string   `json:"SearchText"`
}

//Response is a struct for transmited JSON
type Response struct {
	FoundAtSite string `json:"FoundAtSite"`
}

//GETJSONHandler is a GET handler for "/searchText"
func GETJSONHandler(c *gin.Context) {
	t, _ := template.ParseFiles("templates/default.tmpl")
	t.Execute(c.Writer, nil)
}

//POSTJSONHandler is a POST handler for "/searchText";
//it checks incoming JSON and sends a result of "searchForArgsOnEachSite" func as JSON response.
func POSTJSONHandler(c *gin.Context) {
	var reqjson Request
	if c.Bind(&reqjson) == nil && len(reqjson.Site) > 0 && len(reqjson.SearchText) > 0 {
		var respjson Response
		result, err := SearchForArgOnSites(reqjson.SearchText, reqjson.Site)

		if err != nil {
			c.Error(err)
		} else {
			respjson.FoundAtSite = result
			c.JSON(http.StatusOK, respjson)
		}

		// switch err {
		// case -1:
		// 	c.Writer.WriteString("HTTP Code 204 No Content")
		// 	c.AbortWithError(204, errors.New("HTTP Code 204 No Content"))
		// case -2:
		// 	c.Writer.WriteString("HTTP Code 400 Invailid request")
		// 	c.AbortWithError(400, errors.New("HTTP Code 400 Invailid request"))
		// default:
		// 	respjson.FoundAtSite = reqjson.Site[result]
		// 	c.JSON(http.StatusOK, respjson)
		// }
	}
}

//getContentFromURL get Body from a URL and parse it to []byte;
//it returns nil when "url" is empty.
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
		err = errors.New("HTTP Code 400 invailid request")
		return r, err
	}
	retext := regexp.MustCompile(text)

	for _, val := range sites {

		if checkConnectionToURL(val) != nil {
			err = errors.New("HTTP Code 400 invailid request")
			continue
		}

		if retext.Find(getContentFromURL(val)) != nil {
			r = val
			err = nil
			break
		} else {
			err = errors.New("HTTP Code 204 no content")
		}

	}

	return r, err
}
