package message

import "errors"

var ErrInvalidInput = errors.New("mismatch hash")
var ErrInvalidHash = errors.New("argon2id: hash is not in the correct format")
var ErrIncompatibleVariant = errors.New("argon2id: incompatible variant of argon2")
var ErrIncompatibleVersion = errors.New("argon2id: incompatible version of argon2")
var ErrInvalidEmailOrPasword = errors.New("wrong password or email")
