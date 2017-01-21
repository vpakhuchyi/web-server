package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

//Request is a struct for received JSON
type Request struct {
	Site       []string `json:"site"`
	SearchText string   `json:"searchText"`
}

//Response is a struct for transmited JSON
type Response struct {
	FoundAtSite string `json:"foundAtSite"`
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
		result, err := searchForArgOnSites(reqjson.SearchText, reqjson.Site)

		if err != nil {
			c.Error(err)
		} else {
			respjson.FoundAtSite = reqjson.Site[result]
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

//searchForArgOnSites receive a string "text" for search and a []string "sites" for searching place;
//func returns int that represent element from the "sites" where the "text" was found;
//it returns -1 if the "text" wasn't found.

func searchForArgOnSites(text string, sites []string) (r int, err error) {
	fmt.Println(sites)
	retext := regexp.MustCompile(text)
	rehttp := regexp.MustCompile(`^https?://`)
	var tmp string
	for i, val := range sites {
		if val != "" {
			if rehttp.MatchString(val) {
				tmp = val
			}
			tmp = "http://" + val

			if checkConnectionToURL(tmp) != nil {
				err = errors.New("HTTP Code 400 Invailid request")
				continue
			}

			if retext.Find(getContentFromURL(tmp)) != nil {
				r = i
			}
		}

	}
	r = -1
	err = errors.New("HTTP Code 204 No Content")
	return r, err
}
