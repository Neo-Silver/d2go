package memory

type Offset struct {
	GameData                    uintptr
	UnitTable                   uintptr
	UI                          uintptr
	Hover                       uintptr
	Expansion                   uintptr
	RosterOffset                uintptr
	PanelManagerContainerOffset uintptr
	WidgetStatesOffset          uintptr
	WaypointTableOffset         uintptr
	FPS                         uintptr
	KeyBindingsOffset           uintptr
	KeyBindingsSkillsOffset     uintptr
	QuestInfo                   uintptr
	TZ                          uintptr
	Quests                      uintptr
	Ping                        uintptr
	LegacyGraphics              uintptr
	CharData                    uintptr
	SelectedCharName            uintptr
	LastGameName                uintptr
	LastGamePassword            uintptr
}

func calculateOffsets(_ *Process) Offset {
	// UnitTable
	unitTableOffset := uintptr(0x1EAA3D0)

	// UI
	uiOffsetPtr := uintptr(0x1EBA0C2)

	// Hover
	hoverOffset := uintptr(0x1DFE090)

	// Expansion
	expOffset := uintptr(0x1DFD4E0)

	// Party members offset
	rosterOffset := uintptr(0x1EC06E0)

	// PanelManagerContainer
	panelManagerContainerOffset := uintptr(0x1E14E38)

	// WidgetStates
	WidgetStatesOffset := uintptr(0x1EE26F8)

	// Waypoints
	WaypointTableOffset := uintptr(0x1D5C440)

	// FPS
	fpsOffset := uintptr(0x1D5C414)

	// KeyBindings
	keyBindingsOffset := uintptr(0x19D55B4)

	// KeyBindings Skills
	keyBindingsSkillsOffset := uintptr(0x1DFE180)

	// QuestInfo
	questInfoOffset := uintptr(0x1EC6D58)

	// Terror Zones
	tzOffset := uintptr(0x25B4A10)

	// Ping
	pingOffset := uintptr(0x1DFD4E0)

	// LegacyGraphics
	legacyGfxOffset := uintptr(0x1EC6EFE)

	// CharData
	charDataOffset := uintptr(0x1E01790)

	// Selected Char Name
	selectedCharNameOffset := uintptr(0x1D53215)

	// Last Game Name
	lastGameNameOffset := uintptr(0x25FD370)

	// Last Game Password
	lastGamePasswordOffset := uintptr(0x25FD3C8)

	return Offset{
		UnitTable:                   unitTableOffset,
		UI:                          uiOffsetPtr,
		Hover:                       hoverOffset,
		Expansion:                   expOffset,
		RosterOffset:                rosterOffset,
		PanelManagerContainerOffset: panelManagerContainerOffset,
		WidgetStatesOffset:          WidgetStatesOffset,
		WaypointTableOffset:         WaypointTableOffset,
		FPS:                         fpsOffset,
		KeyBindingsOffset:           keyBindingsOffset,
		KeyBindingsSkillsOffset:     keyBindingsSkillsOffset,
		QuestInfo:                   questInfoOffset,
		TZ:                          tzOffset,
		Ping:                        pingOffset,
		LegacyGraphics:              legacyGfxOffset,
		CharData:                    charDataOffset,
		SelectedCharName:            selectedCharNameOffset,
		LastGameName:                lastGameNameOffset,
		LastGamePassword:            lastGamePasswordOffset,
	}
}
