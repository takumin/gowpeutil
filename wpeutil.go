// +build windows

package wpeutil

import "syscall"

var (
	// Library
	libwpeutil uintptr

	// Functions
	createPageFile                     uintptr
	disableExtendedCharactersForVolume uintptr
	disableFirewall                    uintptr
	enableExtendedCharactersForVolume  uintptr
	enableFirewall                     uintptr
	initializeNetwork                  uintptr
	listKeyboardLayouts                uintptr
	reboot                             uintptr
	saveProfile                        uintptr
	setKeyboardLayout                  uintptr
	setMuiLanguage                     uintptr
	setUserLocale                      uintptr
	shutdown                           uintptr
	updateBootInfo                     uintptr
	waitForNetwork                     uintptr
	waitForRemovableStorage            uintptr
)

func MustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}

	return uintptr(lib)
}

func MustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}

func init() {
	// Library
	libwpeutil = MustLoadLibrary("wpeutil.dll")

	// Functions
	createPageFile = MustGetProcAddress(libwpeutil, "CreatePageFile")
	disableExtendedCharactersForVolume = MustGetProcAddress(libwpeutil, "DisableExtendedCharactersForVolume")
	disableFirewall = MustGetProcAddress(libwpeutil, "DisableFirewall")
	enableExtendedCharactersForVolume = MustGetProcAddress(libwpeutil, "EnableExtendedCharactersForVolume")
	enableFirewall = MustGetProcAddress(libwpeutil, "EnableFirewall")
	initializeNetwork = MustGetProcAddress(libwpeutil, "InitializeNetwork")
	listKeyboardLayouts = MustGetProcAddress(libwpeutil, "ListKeyboardLayouts")
	reboot = MustGetProcAddress(libwpeutil, "Reboot")
	saveProfile = MustGetProcAddress(libwpeutil, "SaveProfile")
	setKeyboardLayout = MustGetProcAddress(libwpeutil, "SetKeyboardLayout")
	setMuiLanguage = MustGetProcAddress(libwpeutil, "SetMuiLanguage")
	setUserLocale = MustGetProcAddress(libwpeutil, "SetUserLocale")
	shutdown = MustGetProcAddress(libwpeutil, "Shutdown")
	updateBootInfo = MustGetProcAddress(libwpeutil, "UpdateBootInfo")
	waitForNetwork = MustGetProcAddress(libwpeutil, "WaitForNetwork")
	waitForRemovableStorage = MustGetProcAddress(libwpeutil, "WaitForRemovableStorage")
}

func CreatePageFile() int32 {
	ret, _, _ := syscall.Syscall(createPageFile, 0, 0, 0, 0)

	return int32(ret)
}

func DisableExtendedCharactersForVolume() int32 {
	ret, _, _ := syscall.Syscall(disableExtendedCharactersForVolume, 0, 0, 0, 0)

	return int32(ret)
}

func DisableFirewall() int32 {
	ret, _, _ := syscall.Syscall(disableFirewall, 0, 0, 0, 0)

	return int32(ret)
}

func EnableExtendedCharactersForVolume() int32 {
	ret, _, _ := syscall.Syscall(enableExtendedCharactersForVolume, 0, 0, 0, 0)

	return int32(ret)
}

func EnableFirewall() int32 {
	ret, _, _ := syscall.Syscall(enableFirewall, 0, 0, 0, 0)

	return int32(ret)
}

func InitializeNetwork() int32 {
	ret, _, _ := syscall.Syscall(initializeNetwork, 0, 0, 0, 0)

	return int32(ret)
}

func ListKeyboardLayouts() int32 {
	ret, _, _ := syscall.Syscall(listKeyboardLayouts, 0, 0, 0, 0)

	return int32(ret)
}

func Reboot() int32 {
	ret, _, _ := syscall.Syscall(reboot, 0, 0, 0, 0)

	return int32(ret)
}

func SaveProfile() int32 {
	ret, _, _ := syscall.Syscall(saveProfile, 0, 0, 0, 0)

	return int32(ret)
}

func SetKeyboardLayout() int32 {
	ret, _, _ := syscall.Syscall(setKeyboardLayout, 0, 0, 0, 0)

	return int32(ret)
}

func SetMuiLanguage() int32 {
	ret, _, _ := syscall.Syscall(setMuiLanguage, 0, 0, 0, 0)

	return int32(ret)
}

func SetUserLocale() int32 {
	ret, _, _ := syscall.Syscall(setUserLocale, 0, 0, 0, 0)

	return int32(ret)
}

func Shutdown() int32 {
	ret, _, _ := syscall.Syscall(shutdown, 0, 0, 0, 0)

	return int32(ret)
}

func UpdateBootInfo() int32 {
	ret, _, _ := syscall.Syscall(updateBootInfo, 0, 0, 0, 0)

	return int32(ret)
}

func WaitForNetwork() int32 {
	ret, _, _ := syscall.Syscall(waitForNetwork, 0, 0, 0, 0)

	return int32(ret)
}

func WaitForRemovableStorage() int32 {
	ret, _, _ := syscall.Syscall(waitForRemovableStorage, 0, 0, 0, 0)

	return int32(ret)
}
