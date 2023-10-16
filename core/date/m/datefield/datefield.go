package datefield

// DateField represents different date fields.
type DateField int

const (
	// ERA represents the era field.
	ERA DateField = iota
	// YEAR represents the year field.
	YEAR
	// MONTH represents the month field.
	MONTH
	// WEEK_OF_YEAR represents the week of the year field.
	WEEK_OF_YEAR
	// WEEK_OF_MONTH represents the week of the month field.
	WEEK_OF_MONTH
	// DAY_OF_MONTH represents the day of the month field.
	DAY_OF_MONTH
	// DAY_OF_YEAR represents the day of the year field.
	DAY_OF_YEAR
	// DAY_OF_WEEK represents the day of the week field.
	DAY_OF_WEEK
	// DAY_OF_WEEK_IN_MONTH represents the day of the week in the month field.
	DAY_OF_WEEK_IN_MONTH
	// AM_PM represents the AM/PM field.
	AM_PM
	// HOUR represents the hour field for 12-hour format.
	HOUR
	// HOUR_OF_DAY represents the hour of the day field for 24-hour format.
	HOUR_OF_DAY
	// MINUTE represents the minute field.
	MINUTE
	// SECOND represents the second field.
	SECOND
	// MILLISECOND represents the millisecond field.
	MILLISECOND
)

// GetValue returns the integer value associated with a DateField.
func (df DateField) GetValue() int {
	return int(df)
}

// Of converts an integer value to a DateField.
func Of(calendarPartIntValue int) DateField {
	switch calendarPartIntValue {
	case 0:
		return ERA
	case 1:
		return YEAR
	case 2:
		return MONTH
	case 3:
		return WEEK_OF_YEAR
	case 4:
		return WEEK_OF_MONTH
	case 5:
		return DAY_OF_MONTH
	case 6:
		return DAY_OF_YEAR
	case 7:
		return DAY_OF_WEEK
	case 8:
		return DAY_OF_WEEK_IN_MONTH
	case 9:
		return AM_PM
	case 10:
		return HOUR
	case 11:
		return HOUR_OF_DAY
	case 12:
		return MINUTE
	case 13:
		return SECOND
	case 14:
		return MILLISECOND
	default:
		return -1
	}
}
