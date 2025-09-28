package aws

import "fmt"

type InvalidAWSConfigError struct {
	cause error
}

func (e *InvalidAWSConfigError) Error() string {
	return fmt.Sprintf("error loading AWS default config: %w", e.cause)
}
