package util

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"regexp"
	"strings"
)

func CopyData(origin interface{}, target interface{}) error {
	ori, err := json.Marshal(origin)
	if err != nil {
		return err
	}

	return json.Unmarshal(ori, &target)
}

func IsEmpty(val interface{}) bool {
	switch v := val.(type) {
	case int:
		return v < 1
	case int64:
		return v < 1
	case string:
		return v == ""
	}
	return false
}

func ReturnNil(val interface{}) interface{} {
	if IsEmpty(val) {
		return nil
	}

	return val
}

func BuildDataHash(fields ...string) string {
	for i, v := range fields {
		fields[i] = strings.ToLower(v)
	}
	return SHA512(fields...)
}

func EncodeToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func DecodeFromBase64(strEnc string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(strEnc)
}

// Convert bytes to int64
func BytesToInt64(b []byte, numBytes int) int64 {
	// Ensure that the byte slice is not longer than numBytes
	truncatedBytes := b[:int(math.Min(float64(len(b)), float64(numBytes)))]

	// Convert bytes to int64
	var result int64
	for _, v := range truncatedBytes {
		result <<= 8
		result |= int64(v)
	}
	return result
}

func StrReplaceWithRegex(regExp, value, replace string) string {
	return regexp.MustCompile(regExp).ReplaceAllString(value, replace)
}
