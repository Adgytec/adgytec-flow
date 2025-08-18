package helpers

import "context"

func GetContextValue(ctx context.Context, key ContextKey) (string, bool) {
	val, valOk := ctx.Value(key).(string)
	return val, valOk
}
