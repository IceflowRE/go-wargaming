package wargaming

import (
	"errors"
	"fmt"
	"strconv"
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

// TODO: method for detailed message
var ApiErrorMessage = map[string]string{
	"AUTH_CANCEL":                 "Application authorization cancelled by user.",                                                          // 401
	"SEARCH_NOT_SPECIFIED":        "Search parameter not specified with no account_id",                                                     // 402
	"AUTH_EXPIRED":                "User authorization timed out.",                                                                         // 403
	"%FIELD%_NOT_SPECIFIED":       "Required field %FIELD% (will be replaced with the name of the field) is not specified.",                // 402
	"%FIELD%_NOT_FOUND":           "Data not found.",                                                                                       // 404
	"METHOD_NOT_FOUND":            "Invalid API method.",                                                                                   // 404
	"METHOD_DISABLED":             "Specified method is disabled.",                                                                         // 405
	"%FIELD%_LIST_LIMIT_EXCEEDED": "Limit of passed-in identifiers in the %FIELD% (will be replaced with the name of the field) exceeded.", // 407
	"APPLICATION_IS_BLOCKED":      "Application is blocked by the administration.",                                                         // 407
	"INVALID_%FIELD%":             "Specified field value %FIELD% (will be replaced with the name of the field) is not valid.",
	"INVALID_APPLICATION_ID":      "Invalid application_id.",                                                       // 407
	"INVALID_IP_ADDRESS":          "Invalid IP-address for the server application.",                                // 407
	"NOT_ENOUGH_SEARCH_LENGTH":    "Search parameter is not long enough. Allowed value depends on type parameter.", // 407
	"REQUEST_LIMIT_EXCEEDED":      "Request limit is exceeded.",                                                    // 407
	"SEARCH_LIST_LIMIT_EXCEEDED":  "Limit of specified names in search parameter exceeded ( >100 )",                // 407
	"AUTH_ERROR":                  "Authentication error.",                                                         // 410
	"SOURCE_NOT_AVAILABLE":        "Data source is not available.",                                                 // 504
}

var (
	InvalidResponse = errors.New("invalid response, 'status' was not ok but missing 'error'")
)

type InvalidParameter struct {
	Field string
	Value any
	Info  string
}

func NewInvalidParameter(field string, value any, info string) *InvalidParameter {
	return &InvalidParameter{
		Field: field,
		Value: value,
		Info:  info,
	}
}

func (err InvalidParameter) Error() string {
	suffix := ""
	if err.Info != "" {
		suffix = " (" + err.Info + ")."
	} else {
		suffix = "."
	}
	return fmt.Sprintf("Specified field (%s) value (%v) is not valid%s", err.Field, err.Value, suffix)
}

type BadStatusCode int

func (err BadStatusCode) Error() string {
	return strconv.Itoa(int(err))
}
