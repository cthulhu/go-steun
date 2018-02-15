package retry

type RetryOperation struct {
	BeforeRetryCallback func(error)
	AfterRetryCallback  func(error)
	FinalError          error
}

// Builds new retryable operation
func New() *RetryOperation {
	return &RetryOperation{dummyCallback, dummyCallback, nil}
}

// Setup custom action before
func (rop *RetryOperation) BeforeRetry(before func(error)) {
	rop.BeforeRetryCallback = before
}

// Setup custom action after
func (rop *RetryOperation) AfterRetry(after func(error)) {
	rop.AfterRetryCallback = after
}

// Execute operation with retry
func (rop *RetryOperation) Do(mainFunc func() error) {
	if err := mainFunc(); err != nil {
		rop.BeforeRetryCallback(err)
		rop.FinalError = mainFunc()
		rop.AfterRetryCallback(err)
	}
}

func dummyCallback(err error) {}
