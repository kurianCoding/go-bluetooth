package bluez

const (

	//Device1Interface the bluez interface for Device1
	Device1Interface = "org.bluez.Device1"
	//Adapter1Interface the bluez interface for Adapter1
	Adapter1Interface = "org.bluez.Adapter1"
	//GattService1Interface the bluez interface for GattService1
	GattService1Interface = "org.bluez.GattService1"
	//GattCharacteristic1Interface the bluez interface for GattCharacteristic1
	GattCharacteristic1Interface = "org.bluez.GattCharacteristic1"
	//GattDescriptor1Interface the bluez interface for GattDescriptor1
	GattDescriptor1Interface = "org.bluez.GattDescriptor1"

	//ObjectManagerInterface the dbus object manager interface
	ObjectManagerInterface = "org.freedesktop.DBus.ObjectManager"
	//InterfacesRemoved the DBus signal member for InterfacesRemoved
	InterfacesRemoved = "org.freedesktop.DBus.ObjectManager.InterfacesRemoved"
	//InterfacesAdded the DBus signal member for InterfacesAdded
	InterfacesAdded = "org.freedesktop.DBus.ObjectManager.InterfacesAdded"

	//PropertiesInterface the DBus properties interface
	PropertiesInterface = "org.freedesktop.DBus.Properties"
	//PropertiesChanged the DBus properties interface and member
	PropertiesChanged = "org.freedesktop.DBus.Properties.PropertiesChanged"
)

// ObjectManagerIntrospectData introspect ObjectManager description
const ObjectManagerIntrospectData = `
<interface name="org.freedesktop.DBus.ObjectManager">
	<method name="GetManagedObjects">
		<arg name="objects" type="a{oa{sa{sv}}}" direction="out" />
	</method>
	<signal name="InterfacesAdded">
		<arg name="object" type="o"/>
		<arg name="interfaces" type="a{sa{sv}}"/>
	</signal>
	<signal name="InterfacesRemoved">
		<arg name="object" type="o"/>
		<arg name="interfaces" type="as"/>
	</signal>
</interface>`

//CharacteristicFlag a flag usable in GattCharacteristic1
type CharacteristicFlag string

//DescriptorFlag a flag usable in GattDescriptor1
type DescriptorFlag string

// Defines how the characteristic value can be used. See
// Core spec "Table 3.5: Characteristic Properties bit
// field", and "Table 3.8: Characteristic Extended
// Properties bit field"
const (
	FlagCharacteristicBroadcast                 CharacteristicFlag = "broadcast"
	FlagCharacteristicRead                                         = "read"
	FlagCharacteristicWriteWithoutResponse                         = "write-without-response"
	FlagCharacteristicWrite                                        = "write"
	FlagCharacteristicNotify                                       = "notify"
	FlagCharacteristicIndicate                                     = "indicate"
	FlagCharacteristicAuthenticatedSignedWrites                    = "authenticated-signed-writes"
	FlagCharacteristicReliableWrite                                = "reliable-write"
	FlagCharacteristicWritableAuxiliaries                          = "writable-auxiliaries"
	FlagCharacteristicEncryptRead                                  = "encrypt-read"
	FlagCharacteristicEncryptWrite                                 = "encrypt-write"
	FlagCharacteristicEncryptAuthenticatedRead                     = "encrypt-authenticated-read"
	FlagCharacteristicEncryptAuthenticatedWrite                    = "encrypt-authenticated-write"
	FlagCharacteristicSecureRead                                   = "secure-read"
	FlagCharacteristicSecureWrite                                  = "secure-write"
)

// Descriptor specific flags
const (
	FlagDescriptorRead                      DescriptorFlag = "read"
	FlagDescriptorWrite                                    = "write"
	FlagDescriptorEncryptRead                              = "encrypt-read"
	FlagDescriptorEncryptWrite                             = "encrypt-write"
	FlagDescriptorEncryptAuthenticatedRead                 = "encrypt-authenticated-read"
	FlagDescriptorEncryptAuthenticatedWrite                = "encrypt-authenticated-write"
	FlagDescriptorSecureRead                               = "secure-read"
	FlagDescriptorSecureWrite                              = "secure-write"
)
