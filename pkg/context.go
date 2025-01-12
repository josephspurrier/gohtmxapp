package pkg

import "context"

// ContextKey is the type for values stored in the request context.
type ContextKey string

var (
	// ContextHashed is the template variable for whether assets should be hashed or not.
	ContextHashed = ContextKey("hashed")

	// ContextPageURL is the template variable for the current page URL. It can be used to determine the active page.
	ContextPageURL = ContextKey("pageURL")
)

// ContextBool returns if the context key is true.
func ContextBool(ctx context.Context, key ContextKey) bool {
	// Determine if the value is set in the context. If it is, return the value.
	if val, ok := ctx.Value(key).(bool); ok {
		return val
	}

	// Else return false.
	return false
}

// ContextString returns the string from the context.
func ContextString(ctx context.Context, key ContextKey) string {
	// Determine if the value is set in the context. If it is, return the value.
	if val, ok := ctx.Value(key).(string); ok {
		return val
	}

	// Else return an empty string.
	return ""
}
