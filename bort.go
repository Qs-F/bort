package bort

import "io"

type Config struct {
	MaxBytes int
}

// IsBin returns true if the content provided by r is binary.
// Otherwise, returns false (it might be utf-8 text file).
// Using `git diff`'s judgement way, this can be implemented quickly.
// In most cases, the best Config.MaxBytes is 8000. In that cases, please use DefaultConfig.
func (c Config) IsBin(r io.Reader) (bool, error) {
	p := make([]byte, 1)
	pos := 0
	for {
		// read until reaching to the max value
		// if reached to the max value, this means this is the valid text file
		pos++
		if pos >= c.MaxBytes {
			return false, nil
		}

		// read a byte one by one
		n, err := r.Read(p)
		// if reached to the end of file, this means this is the valid text file
		if err == io.EOF && n == 0 {
			return false, nil
		}

		if err != nil {
			return false, err
		}

		// if there is a byte which is NUL, this file is binary file
		if p[0] == 0 {
			return true, nil
		}
	}
}

var DefaultConfig = Config{
	MaxBytes: 8000,
}

// IsBin returns true if the content provided by r is binary.
// Otherwise, returns false (it might be utf-8 text file).
// Using this function is encouraged. Details are written in IsBin() with Config
func IsBin(r io.Reader) (bool, error) {
	return DefaultConfig.IsBin(r)
}
