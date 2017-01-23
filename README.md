#Web-server

This app contains a simple web-server. It receives a second JSON:

{
   "Site":["https://google.com","https://yahoo.com"],
   "SearchText":"Google"
}

App searches in each site's body for a "SearchText". It returns JSON as response with folowing data:

{
    "FoundAtSite":"URL of site where (in  body) "SearchText" was found"
}

If "SearchText" wasn't found it returns actual error in folowing style: HTTP Code "number" "message".
