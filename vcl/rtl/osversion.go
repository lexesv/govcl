package rtl

type TArchitecture uint32

const (
	ArIntelX86 TArchitecture = iota + 0
	ArIntelX64
	ArARM32
	ArARM64
)

type TPlatform uint32

const (
	PfWindows TPlatform = iota + 0
	PfMacOS
	PfiOS
	PfAndroid
	PfWinRT
	PfLinux
)

type TOSVersion struct {
	Name             string
	Build            int
	Major            int
	Minor            int
	ServicePackMajor int
	ServicePackMinor int
	Architecture     TArchitecture
	Platform         TPlatform
	ToString         string
}

var OSVersion TOSVersion

func (v *TOSVersion) Check1(AMajor int) bool {
	return v.Major >= AMajor
}

func (v *TOSVersion) Check2(AMajor, AMinor int) bool {
	return v.Major > AMajor || (v.Major == AMajor && v.Minor >= AMinor)
}

func (v *TOSVersion) Check3(AMajor, AMinor, AServicePackMajor int) bool {
	return v.Major > AMajor || (v.Major == AMajor && v.Minor > AMinor) ||
		((v.Major == AMajor && v.Minor == AMinor) && (v.ServicePackMajor >= AServicePackMajor))
}
