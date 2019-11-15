package libwimark

// AuthenIsAccountable return true if authen type could be accountable
func AuthenIsAccountable(t PortalAuthenticationType) bool {
	switch t {
	case PortalAuthenticationTypeCallback, PortalAuthenticationTypeSMS,
		PortalAuthenticationTypeEmail:
		return true
	}
	return false
}
