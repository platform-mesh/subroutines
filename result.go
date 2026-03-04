package subroutines

import "time"

type action int

const (
	actionContinue action = iota
	actionPending
	actionStopWithRequeue
	actionStop
)

// Result encodes the outcome of a subroutine invocation.
// The zero value represents a successful continue with no requeue.
type Result struct {
	action  action
	requeue time.Duration
	message string
}

// OK returns a Result that continues the chain with no requeue.
func OK() Result {
	return Result{}
}

// OKWithRequeue returns a Result that continues the chain and requeues after d.
func OKWithRequeue(d time.Duration) Result {
	return Result{action: actionContinue, requeue: d}
}

// Pending returns a Result that continues the chain, sets the condition to Unknown,
// and requeues after d. Use Pending when a subroutine is waiting on an external condition.
func Pending(d time.Duration, msg string) Result {
	return Result{action: actionPending, requeue: d, message: msg}
}

// StopWithRequeue returns a Result that stops the chain and requeues after d.
func StopWithRequeue(d time.Duration, msg string) Result {
	return Result{action: actionStopWithRequeue, requeue: d, message: msg}
}

// Stop returns a Result that stops the chain with no explicit requeue.
func Stop(msg string) Result {
	return Result{action: actionStop, message: msg}
}

// IsContinue returns true if the result allows the chain to continue (OK or OKWithRequeue).
func (r Result) IsContinue() bool {
	return r.action == actionContinue
}

// IsPending returns true if the result is Pending.
func (r Result) IsPending() bool {
	return r.action == actionPending
}

// IsStopWithRequeue returns true if the result stops the chain with a requeue.
func (r Result) IsStopWithRequeue() bool {
	return r.action == actionStopWithRequeue
}

// IsStop returns true if the result stops the chain with no requeue.
func (r Result) IsStop() bool {
	return r.action == actionStop
}

// Requeue returns the requeue duration.
func (r Result) Requeue() time.Duration {
	return r.requeue
}

// Message returns the result message.
func (r Result) Message() string {
	return r.message
}
