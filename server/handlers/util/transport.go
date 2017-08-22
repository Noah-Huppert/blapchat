package util

// Transport is a struct for holding event system transport types. To be used as an immutable enumeration. Some methods
// will not work if you change the .Name or .Is fields
type Transport struct {
    // Name is the value which should be used when referring to this transport
    Name string

    // Is holds transport names which the current transport is equivalent to. This can be used to group multiple transports
    // under one name.
    //
    // If empty the current transport is considered a "plain" transport. And only represents itself.
    Is []string
}

// Equivalent determines if a provided transport represents any of the same transports with the Is member variable. This
// method does not determine if two transports are exactly equal. From this use the .Equal() method.
func (t *Transport) Equivalent(against *Transport) bool {
    // Compare current transport's and provided transport's flattened Is array against each other
    currentIs := t.Is
    agIs := against.Is

    for _, curItem := range currentIs {
        for _, agItem := range agIs {
            if curItem == agItem {
                // If any of the plain transports that represent the current and provided transport match, they are equivalent
                return true
            }
        }
    }

    // If no similar Is items found, not equivalent
    return false
}

// Equal returns true if the provided transport, and the current transport, are exactly equal. This means they must
// represent the exact same types, or both be the same "plain" typed transport.
func (t *Transport) Equal(against *Transport) bool {
    // Compare current and provided transport's flattened forms
    curIs := t.Is
    agIs := against.Is

    // Simple length check first
    if len(curIs) != len(agIs) {
        return false
    }

    // Build map with all curIs keys
    var curMap map[string]bool
    for _, item := range curIs {
        curMap[item] = false
    }

    // Loop through each agIs item and check curMap
    for _, agItem := range agIs {
        // Check if item exists in curMap
        if _, ok := curMap[agItem]; ok {
            curMap[agItem] = true
        } else {
            // If item doesn't exist in map then they aren't equal
            return false
        }
    }

    // Check curMap and make sure all true
    for _, item := range curMap {
        if !item {
            // If item is false, wasn't found, not equal
            return false
        }
    }

    // If all checks pass, equal
    return true
}

// NewPlainTransport returns a transport who's Name is the only entry in the Is array
func NewPlainTransport(Name string) *Transport {
    return &Transport{
        Name: Name,
        Is: []string{},
    }
}

// Transports
const TRANSPORT_FCM string = "fcm"
const TRANSPORT_HTTP = "http"

var TransportFCM *Transport = NewPlainTransport(TRANSPORT_FCM)
var TransportHTTP *Transport = NewPlainTransport(TRANSPORT_HTTP)
var TransportBoth *Transport = &Transport{
    Name: "both",
    Is: []string{TRANSPORT_FCM, TRANSPORT_HTTP},
}

// IsValidTransport determines if the provided transport is one of the allowed transports. Returns false if not.
func IsValidTransport(t *Transport) bool {
    testerTransport := Transport{
        Name: "TesterTransport",
        Is: []string{TRANSPORT_FCM, TRANSPORT_HTTP},
    }

    return testerTransport.Equivalent(t)
}
