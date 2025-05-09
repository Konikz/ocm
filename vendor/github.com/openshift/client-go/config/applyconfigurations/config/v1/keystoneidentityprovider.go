// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// KeystoneIdentityProviderApplyConfiguration represents a declarative configuration of the KeystoneIdentityProvider type for use
// with apply.
type KeystoneIdentityProviderApplyConfiguration struct {
	OAuthRemoteConnectionInfoApplyConfiguration `json:",inline"`
	DomainName                                  *string `json:"domainName,omitempty"`
}

// KeystoneIdentityProviderApplyConfiguration constructs a declarative configuration of the KeystoneIdentityProvider type for use with
// apply.
func KeystoneIdentityProvider() *KeystoneIdentityProviderApplyConfiguration {
	return &KeystoneIdentityProviderApplyConfiguration{}
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *KeystoneIdentityProviderApplyConfiguration) WithURL(value string) *KeystoneIdentityProviderApplyConfiguration {
	b.OAuthRemoteConnectionInfoApplyConfiguration.URL = &value
	return b
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *KeystoneIdentityProviderApplyConfiguration) WithCA(value *ConfigMapNameReferenceApplyConfiguration) *KeystoneIdentityProviderApplyConfiguration {
	b.OAuthRemoteConnectionInfoApplyConfiguration.CA = value
	return b
}

// WithTLSClientCert sets the TLSClientCert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSClientCert field is set to the value of the last call.
func (b *KeystoneIdentityProviderApplyConfiguration) WithTLSClientCert(value *SecretNameReferenceApplyConfiguration) *KeystoneIdentityProviderApplyConfiguration {
	b.OAuthRemoteConnectionInfoApplyConfiguration.TLSClientCert = value
	return b
}

// WithTLSClientKey sets the TLSClientKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSClientKey field is set to the value of the last call.
func (b *KeystoneIdentityProviderApplyConfiguration) WithTLSClientKey(value *SecretNameReferenceApplyConfiguration) *KeystoneIdentityProviderApplyConfiguration {
	b.OAuthRemoteConnectionInfoApplyConfiguration.TLSClientKey = value
	return b
}

// WithDomainName sets the DomainName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DomainName field is set to the value of the last call.
func (b *KeystoneIdentityProviderApplyConfiguration) WithDomainName(value string) *KeystoneIdentityProviderApplyConfiguration {
	b.DomainName = &value
	return b
}
