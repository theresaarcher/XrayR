// Package mydispatcher Package dispatcher implement the rate limiter and the online device counter
package mydispatcher

//go:generate go run github.com/xtls/xray-core/common/errors/errorgen

// Type returns the feature type token for the custom dispatcher feature.
// It intentionally differs from routing.DispatcherType() to avoid replacing
// the core dispatcher and causing type assertion panics in xray-core.
func Type() interface{} {
	// Consumers should use server.GetFeature(mydispatcher.Type()) to access it.
	return (*DefaultDispatcher)(nil)
}
