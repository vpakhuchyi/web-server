#Web-server (testing mode)

This app contains a simple web-server. 
There is html-tamplate with input form. Input forms for "Sites" can be changed.
App receives this data as JSON in folowing format:
{
   "Site":["https://google.com","https://yahoo.com"],
   "SearchText":"Google"
}

App searches in each site's body for a "SearchText". It returns JSON as response with folowing data if "SearchText" was found:

{
    "FoundAtSite":"URL of site where (in  body) "SearchText" was found"
}

If "SearchText" wasn't found it returns actual error in folowing style: HTTP Code "number" "message".
