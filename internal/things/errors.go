package things

import "errors"

var errMissingShowTarget = errors.New("Error: Must specify --id=ID or query")
var errMissingAuthToken = errors.New("Error: Must specify --auth-token=token")
var errMissingID = errors.New("Error: Must specify --id=id")
