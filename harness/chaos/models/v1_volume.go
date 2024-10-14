// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Volume v1 volume
//
// swagger:model v1.Volume
type V1Volume struct {

	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	AwsElasticBlockStore struct {
		V1AWSElasticBlockStoreVolumeSource
	} `json:"awsElasticBlockStore,omitempty"`

	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// +optional
	AzureDisk struct {
		V1AzureDiskVolumeSource
	} `json:"azureDisk,omitempty"`

	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// +optional
	AzureFile struct {
		V1AzureFileVolumeSource
	} `json:"azureFile,omitempty"`

	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	// +optional
	Cephfs struct {
		V1CephFSVolumeSource
	} `json:"cephfs,omitempty"`

	// Cinder represents a cinder volume attached and mounted on kubelets host machine.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	Cinder struct {
		V1CinderVolumeSource
	} `json:"cinder,omitempty"`

	// ConfigMap represents a configMap that should populate this volume
	// +optional
	ConfigMap struct {
		V1ConfigMapVolumeSource
	} `json:"configMap,omitempty"`

	// CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).
	// +optional
	Csi struct {
		V1CSIVolumeSource
	} `json:"csi,omitempty"`

	// DownwardAPI represents downward API about the pod that should populate this volume
	// +optional
	DownwardAPI struct {
		V1DownwardAPIVolumeSource
	} `json:"downwardAPI,omitempty"`

	// EmptyDir represents a temporary directory that shares a pod's lifetime.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	// +optional
	EmptyDir struct {
		V1EmptyDirVolumeSource
	} `json:"emptyDir,omitempty"`

	// Ephemeral represents a volume that is handled by a cluster storage driver.
	// The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,
	// and deleted when the pod is removed.
	//
	// Use this if:
	// a) the volume is only needed while the pod runs,
	// b) features of normal volumes like restoring from snapshot or capacity
	//    tracking are needed,
	// c) the storage driver is specified through a storage class, and
	// d) the storage driver supports dynamic volume provisioning through
	//    a PersistentVolumeClaim (see EphemeralVolumeSource for more
	//    information on the connection between this volume type
	//    and PersistentVolumeClaim).
	//
	// Use PersistentVolumeClaim or one of the vendor-specific
	// APIs for volumes that persist for longer than the lifecycle
	// of an individual pod.
	//
	// Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to
	// be used that way - see the documentation of the driver for
	// more information.
	//
	// A pod can use both types of ephemeral volumes and
	// persistent volumes at the same time.
	//
	// This is a beta feature and only available when the GenericEphemeralVolume
	// feature gate is enabled.
	//
	// +optional
	Ephemeral struct {
		V1EphemeralVolumeSource
	} `json:"ephemeral,omitempty"`

	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// +optional
	Fc struct {
		V1FCVolumeSource
	} `json:"fc,omitempty"`

	// FlexVolume represents a generic volume resource that is
	// provisioned/attached using an exec based plugin.
	// +optional
	FlexVolume struct {
		V1FlexVolumeSource
	} `json:"flexVolume,omitempty"`

	// Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running
	// +optional
	Flocker struct {
		V1FlockerVolumeSource
	} `json:"flocker,omitempty"`

	// GCEPersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	GcePersistentDisk struct {
		V1GCEPersistentDiskVolumeSource
	} `json:"gcePersistentDisk,omitempty"`

	// GitRepo represents a git repository at a particular revision.
	// DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an
	// EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir
	// into the Pod's container.
	// +optional
	GitRepo struct {
		V1GitRepoVolumeSource
	} `json:"gitRepo,omitempty"`

	// Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md
	// +optional
	Glusterfs struct {
		V1GlusterfsVolumeSource
	} `json:"glusterfs,omitempty"`

	// HostPath represents a pre-existing file or directory on the host
	// machine that is directly exposed to the container. This is generally
	// used for system agents or other privileged things that are allowed
	// to see the host machine. Most containers will NOT need this.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// ---
	// TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not
	// mount host directories as read/write.
	// +optional
	HostPath struct {
		V1HostPathVolumeSource
	} `json:"hostPath,omitempty"`

	// ISCSI represents an ISCSI Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://examples.k8s.io/volumes/iscsi/README.md
	// +optional
	Iscsi struct {
		V1ISCSIVolumeSource
	} `json:"iscsi,omitempty"`

	// Volume's name.
	// Must be a DNS_LABEL and unique within the pod.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// NFS represents an NFS mount on the host that shares a pod's lifetime
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	Nfs struct {
		V1NFSVolumeSource
	} `json:"nfs,omitempty"`

	// PersistentVolumeClaimVolumeSource represents a reference to a
	// PersistentVolumeClaim in the same namespace.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	// +optional
	PersistentVolumeClaim struct {
		V1PersistentVolumeClaimVolumeSource
	} `json:"persistentVolumeClaim,omitempty"`

	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk struct {
		V1PhotonPersistentDiskVolumeSource
	} `json:"photonPersistentDisk,omitempty"`

	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	// +optional
	PortworxVolume struct {
		V1PortworxVolumeSource
	} `json:"portworxVolume,omitempty"`

	// Items for all in one resources secrets, configmaps, and downward API
	Projected struct {
		V1ProjectedVolumeSource
	} `json:"projected,omitempty"`

	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	// +optional
	Quobyte struct {
		V1QuobyteVolumeSource
	} `json:"quobyte,omitempty"`

	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime.
	// More info: https://examples.k8s.io/volumes/rbd/README.md
	// +optional
	Rbd struct {
		V1RBDVolumeSource
	} `json:"rbd,omitempty"`

	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// +optional
	ScaleIO struct {
		V1ScaleIOVolumeSource
	} `json:"scaleIO,omitempty"`

	// Secret represents a secret that should populate this volume.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	// +optional
	Secret struct {
		V1SecretVolumeSource
	} `json:"secret,omitempty"`

	// StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
	// +optional
	Storageos struct {
		V1StorageOSVolumeSource
	} `json:"storageos,omitempty"`

	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	// +optional
	VsphereVolume struct {
		V1VsphereVirtualDiskVolumeSource
	} `json:"vsphereVolume,omitempty"`
}

// Validate validates this v1 volume
func (m *V1Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsElasticBlockStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCephfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCinder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownwardAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmptyDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEphemeral(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlexVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlocker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcePersistentDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlusterfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistentVolumeClaim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotonPersistentDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortworxVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuobyte(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRbd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScaleIO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphereVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Volume) validateAwsElasticBlockStore(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsElasticBlockStore) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateAzureDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureDisk) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateAzureFile(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureFile) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateCephfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Cephfs) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateCinder(formats strfmt.Registry) error {
	if swag.IsZero(m.Cinder) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateConfigMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigMap) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateCsi(formats strfmt.Registry) error {
	if swag.IsZero(m.Csi) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateDownwardAPI(formats strfmt.Registry) error {
	if swag.IsZero(m.DownwardAPI) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateEmptyDir(formats strfmt.Registry) error {
	if swag.IsZero(m.EmptyDir) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateEphemeral(formats strfmt.Registry) error {
	if swag.IsZero(m.Ephemeral) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateFc(formats strfmt.Registry) error {
	if swag.IsZero(m.Fc) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateFlexVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.FlexVolume) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateFlocker(formats strfmt.Registry) error {
	if swag.IsZero(m.Flocker) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateGcePersistentDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.GcePersistentDisk) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateGitRepo(formats strfmt.Registry) error {
	if swag.IsZero(m.GitRepo) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateGlusterfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Glusterfs) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateHostPath(formats strfmt.Registry) error {
	if swag.IsZero(m.HostPath) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateIscsi(formats strfmt.Registry) error {
	if swag.IsZero(m.Iscsi) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateNfs(formats strfmt.Registry) error {
	if swag.IsZero(m.Nfs) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validatePersistentVolumeClaim(formats strfmt.Registry) error {
	if swag.IsZero(m.PersistentVolumeClaim) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validatePhotonPersistentDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.PhotonPersistentDisk) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validatePortworxVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.PortworxVolume) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateProjected(formats strfmt.Registry) error {
	if swag.IsZero(m.Projected) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateQuobyte(formats strfmt.Registry) error {
	if swag.IsZero(m.Quobyte) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateRbd(formats strfmt.Registry) error {
	if swag.IsZero(m.Rbd) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateScaleIO(formats strfmt.Registry) error {
	if swag.IsZero(m.ScaleIO) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateStorageos(formats strfmt.Registry) error {
	if swag.IsZero(m.Storageos) { // not required
		return nil
	}

	return nil
}

func (m *V1Volume) validateVsphereVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.VsphereVolume) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this v1 volume based on the context it is used
func (m *V1Volume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsElasticBlockStore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCephfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCinder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDownwardAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmptyDir(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEphemeral(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlexVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlocker(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcePersistentDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGitRepo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGlusterfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostPath(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePersistentVolumeClaim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhotonPersistentDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortworxVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjected(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuobyte(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRbd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScaleIO(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsphereVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Volume) contextValidateAwsElasticBlockStore(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateAzureDisk(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateAzureFile(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateCephfs(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateCinder(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateConfigMap(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateCsi(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateDownwardAPI(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateEmptyDir(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateEphemeral(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateFc(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateFlexVolume(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateFlocker(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateGcePersistentDisk(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateGitRepo(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateGlusterfs(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateHostPath(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateIscsi(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateNfs(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidatePersistentVolumeClaim(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidatePhotonPersistentDisk(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidatePortworxVolume(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateProjected(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateQuobyte(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateRbd(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateScaleIO(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateSecret(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateStorageos(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *V1Volume) contextValidateVsphereVolume(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *V1Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Volume) UnmarshalBinary(b []byte) error {
	var res V1Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
