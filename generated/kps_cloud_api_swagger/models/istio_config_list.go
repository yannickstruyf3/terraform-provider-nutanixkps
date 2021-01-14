// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IstioConfigList IstioConfigList istioConfigList
//
// This type is used for returning a response of IstioConfigList
// swagger:model IstioConfigList
type IstioConfigList struct {

	// adapters
	Adapters IstioAdapters `json:"adapters,omitempty"`

	// attribute manifests
	AttributeManifests AttributeManifests `json:"attributeManifests,omitempty"`

	// authorization policies
	AuthorizationPolicies AuthorizationPolicies `json:"authorizationPolicies,omitempty"`

	// cluster rbac configs
	ClusterRbacConfigs ClusterRbacConfigs `json:"clusterRbacConfigs,omitempty"`

	// destination rules
	DestinationRules *DestinationRules `json:"destinationRules,omitempty"`

	// envoy filters
	EnvoyFilters EnvoyFilters `json:"envoyFilters,omitempty"`

	// gateways
	Gateways Gateways `json:"gateways,omitempty"`

	// handlers
	Handlers IstioHandlers `json:"handlers,omitempty"`

	// http Api spec bindings
	HTTPAPISpecBindings HTTPAPISpecBindings `json:"httpApiSpecBindings,omitempty"`

	// http Api specs
	HTTPAPISpecs HTTPAPISpecs `json:"httpApiSpecs,omitempty"`

	// instances
	Instances IstioInstances `json:"instances,omitempty"`

	// istio validations
	IstioValidations IstioValidations `json:"istioValidations,omitempty"`

	// mesh policies
	MeshPolicies MeshPolicies `json:"meshPolicies,omitempty"`

	// namespace
	// Required: true
	Namespace *Namespace `json:"namespace"`

	// peer authentications
	PeerAuthentications PeerAuthentications `json:"peerAuthentications,omitempty"`

	// policies
	Policies Policies `json:"policies,omitempty"`

	// quota spec bindings
	QuotaSpecBindings QuotaSpecBindings `json:"quotaSpecBindings,omitempty"`

	// quota specs
	QuotaSpecs QuotaSpecs `json:"quotaSpecs,omitempty"`

	// rbac configs
	RbacConfigs RbacConfigs `json:"rbacConfigs,omitempty"`

	// request authentications
	RequestAuthentications RequestAuthentications `json:"requestAuthentications,omitempty"`

	// rules
	Rules IstioRules `json:"rules,omitempty"`

	// service entries
	ServiceEntries ServiceEntries `json:"serviceEntries,omitempty"`

	// service mesh policies
	ServiceMeshPolicies ServiceMeshPolicies `json:"serviceMeshPolicies,omitempty"`

	// service mesh rbac configs
	ServiceMeshRbacConfigs ServiceMeshRbacConfigs `json:"serviceMeshRbacConfigs,omitempty"`

	// service role bindings
	ServiceRoleBindings ServiceRoleBindings `json:"serviceRoleBindings,omitempty"`

	// service roles
	ServiceRoles ServiceRoles `json:"serviceRoles,omitempty"`

	// sidecars
	Sidecars Sidecars `json:"sidecars,omitempty"`

	// templates
	Templates IstioTemplates `json:"templates,omitempty"`

	// virtual services
	VirtualServices *VirtualServices `json:"virtualServices,omitempty"`

	// workload entries
	WorkloadEntries WorkloadEntries `json:"workloadEntries,omitempty"`
}

// Validate validates this istio config list
func (m *IstioConfigList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdapters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributeManifests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterRbacConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvoyFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandlers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPAPISpecBindings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPAPISpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeshPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerAuthentications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotaSpecBindings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotaSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRbacConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestAuthentications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceMeshPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceMeshRbacConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRoleBindings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSidecars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IstioConfigList) validateAdapters(formats strfmt.Registry) error {

	if swag.IsZero(m.Adapters) { // not required
		return nil
	}

	if err := m.Adapters.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("adapters")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateAttributeManifests(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributeManifests) { // not required
		return nil
	}

	if err := m.AttributeManifests.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributeManifests")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateAuthorizationPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthorizationPolicies) { // not required
		return nil
	}

	if err := m.AuthorizationPolicies.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authorizationPolicies")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateClusterRbacConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterRbacConfigs) { // not required
		return nil
	}

	if err := m.ClusterRbacConfigs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("clusterRbacConfigs")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateDestinationRules(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationRules) { // not required
		return nil
	}

	if m.DestinationRules != nil {
		if err := m.DestinationRules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationRules")
			}
			return err
		}
	}

	return nil
}

func (m *IstioConfigList) validateEnvoyFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvoyFilters) { // not required
		return nil
	}

	if err := m.EnvoyFilters.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("envoyFilters")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateGateways(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateways) { // not required
		return nil
	}

	if err := m.Gateways.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gateways")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateHandlers(formats strfmt.Registry) error {

	if swag.IsZero(m.Handlers) { // not required
		return nil
	}

	if err := m.Handlers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("handlers")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateHTTPAPISpecBindings(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPAPISpecBindings) { // not required
		return nil
	}

	if err := m.HTTPAPISpecBindings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("httpApiSpecBindings")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateHTTPAPISpecs(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPAPISpecs) { // not required
		return nil
	}

	if err := m.HTTPAPISpecs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("httpApiSpecs")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.Instances) { // not required
		return nil
	}

	if err := m.Instances.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("instances")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateMeshPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.MeshPolicies) { // not required
		return nil
	}

	if err := m.MeshPolicies.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("meshPolicies")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	if m.Namespace != nil {
		if err := m.Namespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

func (m *IstioConfigList) validatePeerAuthentications(formats strfmt.Registry) error {

	if swag.IsZero(m.PeerAuthentications) { // not required
		return nil
	}

	if err := m.PeerAuthentications.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("peerAuthentications")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validatePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	if err := m.Policies.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policies")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateQuotaSpecBindings(formats strfmt.Registry) error {

	if swag.IsZero(m.QuotaSpecBindings) { // not required
		return nil
	}

	if err := m.QuotaSpecBindings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quotaSpecBindings")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateQuotaSpecs(formats strfmt.Registry) error {

	if swag.IsZero(m.QuotaSpecs) { // not required
		return nil
	}

	if err := m.QuotaSpecs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("quotaSpecs")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateRbacConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.RbacConfigs) { // not required
		return nil
	}

	if err := m.RbacConfigs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rbacConfigs")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateRequestAuthentications(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestAuthentications) { // not required
		return nil
	}

	if err := m.RequestAuthentications.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requestAuthentications")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	if err := m.Rules.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rules")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateServiceEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEntries) { // not required
		return nil
	}

	if err := m.ServiceEntries.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceEntries")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateServiceMeshPolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceMeshPolicies) { // not required
		return nil
	}

	if err := m.ServiceMeshPolicies.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceMeshPolicies")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateServiceMeshRbacConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceMeshRbacConfigs) { // not required
		return nil
	}

	if err := m.ServiceMeshRbacConfigs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceMeshRbacConfigs")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateServiceRoleBindings(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceRoleBindings) { // not required
		return nil
	}

	if err := m.ServiceRoleBindings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceRoleBindings")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateServiceRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceRoles) { // not required
		return nil
	}

	if err := m.ServiceRoles.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceRoles")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateSidecars(formats strfmt.Registry) error {

	if swag.IsZero(m.Sidecars) { // not required
		return nil
	}

	if err := m.Sidecars.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sidecars")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	if err := m.Templates.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("templates")
		}
		return err
	}

	return nil
}

func (m *IstioConfigList) validateVirtualServices(formats strfmt.Registry) error {

	if swag.IsZero(m.VirtualServices) { // not required
		return nil
	}

	if m.VirtualServices != nil {
		if err := m.VirtualServices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualServices")
			}
			return err
		}
	}

	return nil
}

func (m *IstioConfigList) validateWorkloadEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkloadEntries) { // not required
		return nil
	}

	if err := m.WorkloadEntries.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("workloadEntries")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IstioConfigList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IstioConfigList) UnmarshalBinary(b []byte) error {
	var res IstioConfigList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
