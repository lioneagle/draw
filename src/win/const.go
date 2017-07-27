package win

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

// gdi

// Stock Logical Objects
const (
	WHITE_BRUSH         = 0
	LTGRAY_BRUSH        = 1
	GRAY_BRUSH          = 2
	DKGRAY_BRUSH        = 3
	BLACK_BRUSH         = 4
	NULL_BRUSH          = 5
	HOLLOW_BRUSH        = NULL_BRUSH
	WHITE_PEN           = 6
	BLACK_PEN           = 7
	NULL_PEN            = 8
	OEM_FIXED_FONT      = 10
	ANSI_FIXED_FONT     = 11
	ANSI_VAR_FONT       = 12
	SYSTEM_FONT         = 13
	DEVICE_DEFAULT_FONT = 14
	DEFAULT_PALETTE     = 15
	SYSTEM_FIXED_FONT   = 16
	DEFAULT_GUI_FONT    = 17
	DC_BRUSH            = 18
	DC_PEN              = 19
)

// gdiplus
const (
	Ok                        GpStatus = 0
	GenericError              GpStatus = 1
	InvalidParameter          GpStatus = 2
	OutOfMemory               GpStatus = 3
	ObjectBusy                GpStatus = 4
	InsufficientBuffer        GpStatus = 5
	NotImplemented            GpStatus = 6
	Win32Error                GpStatus = 7
	WrongState                GpStatus = 8
	Aborted                   GpStatus = 9
	FileNotFound              GpStatus = 10
	ValueOverflow             GpStatus = 11
	AccessDenied              GpStatus = 12
	UnknownImageFormat        GpStatus = 13
	FontFamilyNotFound        GpStatus = 14
	FontStyleNotFound         GpStatus = 15
	NotTrueTypeFont           GpStatus = 16
	UnsupportedGdiplusVersion GpStatus = 17
	GdiplusNotInitialized     GpStatus = 18
	PropertyNotFound          GpStatus = 19
	PropertyNotSupported      GpStatus = 20
	ProfileNotFound           GpStatus = 21
)

const (
	BrushTypeSolidColor     GpBrushType = 0
	BrushTypeHatchFill      GpBrushType = 1
	BrushTypeTextureFill    GpBrushType = 2
	BrushTypePathGradient   GpBrushType = 3
	BrushTypeLinearGradient GpBrushType = 4
)

const (
	WrapModeTile       GpWrapMode = 0
	WrapModeTileFlipX  GpWrapMode = 1
	WrapModeTileFlipY  GpWrapMode = 2
	WrapModeTileFlipXY GpWrapMode = 3
	WrapModeClamp      GpWrapMode = 4
)

const (
	TextRenderingHintSystemDefault            TextRenderingHint = 0
	TextRenderingHintSingleBitPerPixelGridFit TextRenderingHint = 1
	TextRenderingHintSingleBitPerPixel        TextRenderingHint = 2
	TextRenderingHintAntiAliasGridFit         TextRenderingHint = 3
	TextRenderingHintAntiAlias                TextRenderingHint = 4
	TextRenderingHintClearTypeGridFit         TextRenderingHint = 5
)

const (
	QualityModeInvalid QualityMode = -1
	QualityModeDefault QualityMode = 0
	QualityModeLow     QualityMode = 1
	QualityModeHigh    QualityMode = 2
)

const (
	SmoothingModeInvalid      SmoothingMode = -1 //SmoothingMode(QualityModeInvalid)
	SmoothingModeDefault      SmoothingMode = 0  // SmoothingMode(QualityModeDefault)
	SmoothingModeHighSpeed    SmoothingMode = 1  // SmoothingMode(QualityModeLow)
	SmoothingModeHighQuality  SmoothingMode = 2  // SmoothingMode(QualityModeHigh)
	SmoothingModeNone         SmoothingMode = 3  // SmoothingMode(QualityModeHigh + 1)
	SmoothingModeAntiAlias8x4 SmoothingMode = 4  // SmoothingMode(QualityModeHigh + 2)
	SmoothingModeAntiAlias    SmoothingMode = 4  // SmoothingModeAntiAlias8x4
	SmoothingModeAntiAlias8x8 SmoothingMode = 5  // SmoothingModeAntiAlias + 1
)

const (
	StringAlignmentNear   StringAlignment = 0
	StringAlignmentCenter StringAlignment = 1
	StringAlignmentFar    StringAlignment = 2
)

const (
	StringFormatFlagsDirectionRightToLeft  StringFormatFlags = 0x00000001
	StringFormatFlagsDirectionVertical     StringFormatFlags = 0x00000002
	StringFormatFlagsNoFitBlackBox         StringFormatFlags = 0x00000004
	StringFormatFlagsDisplayFormatControl  StringFormatFlags = 0x00000020
	StringFormatFlagsNoFontFallback        StringFormatFlags = 0x00000400
	StringFormatFlagsMeasureTrailingSpaces StringFormatFlags = 0x00000800
	StringFormatFlagsNoWrap                StringFormatFlags = 0x00001000
	StringFormatFlagsLineLimit             StringFormatFlags = 0x00002000
	StringFormatFlagsNoClip                StringFormatFlags = 0x00004000
)

const (
	LANG_NEUTRAL LANGID = 0
)

const (
	UnitWorld      GpUnit = 0
	UnitDisplay    GpUnit = 1
	UnitPixel      GpUnit = 2
	UnitPoint      GpUnit = 3
	UnitInch       GpUnit = 4
	UnitDocument   GpUnit = 5
	UnitMillimeter GpUnit = 6
)
