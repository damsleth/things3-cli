package things

import "errors"

var errMissingShowTarget = errors.New("Error: Must specify --id=ID or query")
var errMissingAuthToken = errors.New("Error: Missing Things auth token. Set THINGS_AUTH_TOKEN or pass --auth-token=TOKEN (Things > Settings > General > Things URLs).")
var errMissingID = errors.New("Error: Must specify --id=id")
var errMissingTitle = errors.New("Error: Must specify title")
var errMissingAreaTarget = errors.New("Error: Must specify --id=ID or area title")
var errMissingAreaTags = errors.New("Error: Must specify --tags or --add-tags")
