package reporter

import "github.com/tealeg/xlsx"

const (
	TOTAL_TOP        = iota //Render header in TOP position
	TOTAL_BOTTOM            //Render header in BOTTOM position
	TOTAL_TOP_BOTTOM        //Render header in TOP and BOTTOM position
)

const (
	BgColor     = "FFFFFFCC"
	HFgColor    = "FFE3DEBF"
	TFgColor    = "FFE8E4CB"
	BGWhite     = "FFFFFFFF"
	BorderColor = "FFB4AE9E"
)

var (
	Border = xlsx.Border{
		Left:        "hair",
		Right:       "hair",
		Top:         "hair",
		Bottom:      "hair",
		LeftColor:   BorderColor,
		RightColor:  BorderColor,
		TopColor:    BorderColor,
		BottomColor: BorderColor,
	}

	Fill = xlsx.Fill{
		PatternType: xlsx.Solid_Cell_Fill,
		FgColor:     HFgColor,
		BgColor:     BgColor,
	}

	TFill = xlsx.Fill{
		PatternType: xlsx.Solid_Cell_Fill,
		FgColor:     TFgColor,
		BgColor:     BgColor,
	}

	TitleStyle = &xlsx.Style{
		Font: xlsx.Font{
			Size: 12,
			Bold: true,
		},
		ApplyFont: true,
	}

	HeaderCellStyle = &xlsx.Style{
		Alignment: xlsx.Alignment{
			Horizontal: "center",
			Vertical:   "top",
		},
		ApplyAlignment: true,
		Border:         Border,
		ApplyBorder:    true,
		Fill:           Fill,
		ApplyFill:      true,
	}

	TotalCellStyle = &xlsx.Style{
		Border:      Border,
		ApplyBorder: true,
		Fill:        TFill,
		ApplyFill:   true,
	}

	CellStyle = &xlsx.Style{
		Border:      Border,
		ApplyBorder: true,
	}
)
