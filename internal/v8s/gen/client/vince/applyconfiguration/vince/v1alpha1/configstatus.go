/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ConfigStatusApplyConfiguration represents an declarative configuration of the ConfigStatus type for use
// with apply.
type ConfigStatusApplyConfiguration struct {
	Sites []string `json:"sites,omitempty"`
}

// ConfigStatusApplyConfiguration constructs an declarative configuration of the ConfigStatus type for use with
// apply.
func ConfigStatus() *ConfigStatusApplyConfiguration {
	return &ConfigStatusApplyConfiguration{}
}

// WithSites adds the given value to the Sites field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Sites field.
func (b *ConfigStatusApplyConfiguration) WithSites(values ...string) *ConfigStatusApplyConfiguration {
	for i := range values {
		b.Sites = append(b.Sites, values[i])
	}
	return b
}