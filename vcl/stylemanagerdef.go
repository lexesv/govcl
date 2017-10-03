package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TStyleManager struct {
}

func NewStyleManager() *TStyleManager {
	return new(TStyleManager)
}

func (s *TStyleManager) Initialize() {
	StyleManager_Initialize()
}

// StyleManager_UnInitialize
func (s *TStyleManager) UnInitialize() {
	StyleManager_UnInitialize()
}

// StyleManager_IsValidStyle
func (s *TStyleManager) IsValidStyle(filename string) bool {
	return StyleManager_IsValidStyle(filename)
}

// StyleManager_LoadFromFile
func (s *TStyleManager) LoadFromFile(filename string) uintptr {
	return StyleManager_LoadFromFile(filename)
}

// StyleManager_CheckSysClassName
func (s *TStyleManager) CheckSysClassName(className string) bool {
	return StyleManager_CheckSysClassName(className)
}

// StyleManager_TrySetStyle
func (s *TStyleManager) TrySetStyle(name string, showErrorDialog bool) bool {
	return StyleManager_TrySetStyle(name, showErrorDialog)
}

// StyleManager_SetStyle1
func (s *TStyleManager) SetStyle(handle uintptr) {
	StyleManager_SetStyle1(handle)
}

// StyleManager_SetStyle2
func (s *TStyleManager) SetStyle2(name string) {
	StyleManager_SetStyle2(name)
}

// StyleManager_TryLoadFromResource
func (s *TStyleManager) TryLoadFromResource(instance uintptr, resName, resType string, handle *uintptr) bool {
	return StyleManager_TryLoadFromResource(instance, resName, resType, handle)
}

func (s *TStyleManager) SetStyleFromFileName(filename string) {
	StyleManager.SetStyle(StyleManager.LoadFromFile(filename))
}
