package wargaming

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ResponseError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Value   string `json:"value"`
}

func (err ResponseError) Error() string {
	return "|" + strconv.Itoa(err.Code) + " " + err.Message + " | field: " + err.Field + " - value: " + err.Value
}

var apiErrorMessages = map[string]string{
	"APPLICATION_IS_BLOCKED":     "Application is blocked by the administration.",                                 // 407
	"AUTH_CANCEL":                "Application authorization cancelled by user.",                                  // 401
	"AUTH_ERROR":                 "Authentication error.",                                                         // 410
	"AUTH_EXPIRED":               "User authorization timed out.",                                                 // 403
	"INVALID_APPLICATION_ID":     "Invalid application_id.",                                                       // 407
	"INVALID_IP_ADDRESS":         "Invalid IP-address for the server application.",                                // 407
	"METHOD_DISABLED":            "Specified method is disabled.",                                                 // 405
	"METHOD_NOT_FOUND":           "Invalid API method.",                                                           // 404
	"NOT_ENOUGH_SEARCH_LENGTH":   "Search parameter is not long enough. Allowed value depends on type parameter.", // 407
	"REQUEST_LIMIT_EXCEEDED":     "Request limit is exceeded.",                                                    // 407
	"SEARCH_LIST_LIMIT_EXCEEDED": "Limit of specified names in search parameter exceeded ( >100 )",                // 407
	"SOURCE_NOT_AVAILABLE":       "Data source is not available.",                                                 // 504
	"SEARCH_NOT_SPECIFIED":       "Search parameter not specified with no account_id",                             // 402
}

// ApiErrorStringToString will return "" if nothing found.
func ApiErrorStringToString(error string) string {
	if val, ok := apiErrorMessages[error]; ok {
		return val
	}
	switch {
	case strings.HasPrefix(error, "INVALID_"):
		return fmt.Sprintf("Specified field value %s is not valid.", strings.ToLower(strings.TrimPrefix(error, "INVALID_")))
	case strings.HasSuffix(error, "_LIST_LIMIT_EXCEEDED"):
		return fmt.Sprintf("Limit of passed-in identifiers in the %s exceeded.", strings.ToLower(strings.TrimSuffix(error, "_LIST_LIMIT_EXCEEDED")))
	case strings.HasSuffix(error, "_NOT_SPECIFIED"):
		return fmt.Sprintf("Required field %s is not specified.", strings.ToLower(strings.TrimSuffix(error, "_NOT_SPECIFIED")))
	case strings.HasSuffix(error, "_NOT_FOUND"):
		return "Data not found."
	}
	return ""
}

var (
	InvalidResponse = errors.New("invalid response, 'status' was not ok but missing 'error'")
)

type BadStatusCode int

func (err BadStatusCode) Error() string {
	return strconv.Itoa(int(err))
}
