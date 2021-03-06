package ga

import "net/url"

//WARNING: This file was generated. Do not edit.

//Social Hit Type
type Social struct {
	network      string
	action       string
	actionTarget string
}

// NewSocial creates a new Social Hit Type.
// Specifies the social network, for example Facebook or Google
// Plus.
// Specifies the social interaction action. For example on
// Google Plus when a user clicks the +1 button, the social
// action is 'plus'.
// Specifies the target of a social interaction. This value
// is typically a URL but can be any text.
func NewSocial(network string, action string, actionTarget string) *Social {
	h := &Social{
		network:      network,
		action:       action,
		actionTarget: actionTarget,
	}
	return h
}

func (h *Social) addFields(v url.Values) error {
	v.Add("sn", h.network)
	v.Add("sa", h.action)
	v.Add("st", h.actionTarget)
	return nil
}

func (h *Social) Copy() *Social {
	c := *h
	return &c
}
