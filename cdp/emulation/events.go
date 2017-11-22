package emulation

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventVirtualTimeBudgetExpired notification sent after the virtual time
// budget for the current VirtualTimePolicy has run out.
type EventVirtualTimeBudgetExpired struct{}

// EventVirtualTimeAdvanced notification sent after the virtual time has
// advanced.
type EventVirtualTimeAdvanced struct {
	VirtualTimeElapsed float64 `json:"virtualTimeElapsed"` // The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
}

// EventVirtualTimePaused notification sent after the virtual time has
// paused.
type EventVirtualTimePaused struct {
	VirtualTimeElapsed float64 `json:"virtualTimeElapsed"` // The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventEmulationVirtualTimeBudgetExpired,
	cdp.EventEmulationVirtualTimeAdvanced,
	cdp.EventEmulationVirtualTimePaused,
}
