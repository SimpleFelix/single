package single

import (
	"fmt"
	"os"
)

// Lock tries to remove the lock file, if it exists.
// If the file is already open by another instance of the program,
// remove will fail and exit the program.
func (s *Single) Lock() (*os.File, error) {
	if err := os.Remove(s.Lockfile()); err != nil && !os.IsNotExist(err) {
		return nil, ErrAlreadyRunning
	}

	file, err := os.OpenFile(s.Lockfile(), os.O_EXCL|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	s.file = file

	return file, nil
}

// Unlock closes and removes the lockfile.
func (s *Single) Unlock() error {
	if err := s.file.Close(); err != nil {
		return fmt.Errorf("failed to close the lock file: %w", err)
	}

	if err := os.Remove(s.Lockfile()); err != nil {
		return fmt.Errorf("failed to remove the lock file: %w", err)
	}

	return nil
}
