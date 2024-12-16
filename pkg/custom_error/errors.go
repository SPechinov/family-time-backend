package custom_error

var ErrCodeIsNotInRedis = New("codeIsNotInRedis")
var ErrCodesNotEqual = New("codesIsNotEqual")
var ErrCodeMaxAttempts = New("codeMaxAttempts")
var ErrUserExists = New("userExists")
var ErrUserNotExists = New("userNotExists")
var ErrUnknownAuthMethod = New("unknownAuthMethod")
var ErrIncorrectPassword = New("incorrectPassword")
var ErrNotAuthorized = New("notAuthorized")
