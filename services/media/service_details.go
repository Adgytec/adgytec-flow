package media

// Media service is secondary service which provides the simple methods to upload media items
// It doesnt' require any permissions as media uploads are context dependant and will be handled by primary service action
// It also doesn't require its details in database
var serviceName = "media"
