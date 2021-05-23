package meta

import "errors"

// Application error messages.

// Err03Activation not ready to be run.
var Err03Activation = errors.New("error activation")

// Err05Request some thing made a poorly formated request.
var Err05Request = errors.New("error request")

// Err07User a user error has occurred.
var Err07User = errors.New("error user")

// Err08Resource resource error, a resource has changed or is not
// present.
var Err08Resource = errors.New("error resource")

// Err09Record the data received was not expected or is not the correct
// data.
var Err09Record = errors.New("error record")
