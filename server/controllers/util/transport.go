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
    curFlat := t.Is
    agFlat := against.Is

    for _, curItem := range curFlat {
        for _, agFlat := range agFlat {
            if curItem == agFlat {
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
    curFlat := t.Is
    agFlat := against.Is

    // Simple length check first
    if len(curFlat) != len(agFlat) {
        return false
    }

    // Build map with all curFlat keys
    var curFlatMap map[string]bool
    for _, item := range curFlat {
        curFlatMap[item] = false
    }

    // Loop through each agFlat item and check curFlatMap
    for _, agItem := range agFlat {
        // Check if item exists in curFlatMap
        if _, ok := curFlatMap[agItem]; ok {
            curFlatMap[agItem] = true
        } else {
            // If item doesn't exist in map then they aren't equal
            return false
        }
    }

    // Check curFlatMap and make sure all true
    for _, item := range curFlatMap {
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
        Is: []string,
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
