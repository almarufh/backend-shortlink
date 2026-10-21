package message

import "errors"

// HASH
var ErrInvalidInput = errors.New("mismatch hash")
var ErrInvalidHash = errors.New("argon2id: hash is not in the correct format")
var ErrIncompatibleVariant = errors.New("argon2id: incompatible variant of argon2")
var ErrIncompatibleVersion = errors.New("argon2id: incompatible version of argon2")
var ErrInvalidEmailOrPasword = errors.New("wrong password or email")

// MIDDLEWARE
var ErrTokenExpired = errors.New("session expired, please login again")

// AUTH
var ErrEmailExist = errors.New("email alredy exist")
var ErrEmailNotExist = errors.New("email alredy exist")
var ErrNoRowAffected = errors.New("no rows affected")

// Controller
var ErrInternal = errors.New("internal service error")

// LINKS
var ErrShortLinkExist = errors.New("short link exist")
var ErrShortLinkNotExist = errors.New("short link exist")
var ErrFailedQuery = errors.New("failed execute query")
var ErrReadRow = errors.New("failed read row")
