package media

// Media service is a secondary service that provides simple methods to upload media items.
// It doesn't implement its own permission checks, as media uploads are context-dependent
// and authorization is expected to be handled by the primary service calling it or by middleware.
var serviceName = "media"
