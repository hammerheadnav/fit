package fit

// ActivityFile represents the Activity FIT file type.
// Records sensor data and events from active sessions.
type ActivityFile struct {
	Activity    *ActivityMsg
	Sessions    []*SessionMsg
	Laps        []*LapMsg
	Lengths     []*LengthMsg
	Records     []*RecordMsg
	Events      []*EventMsg
	Hrvs        []*HrvMsg
	DeviceInfos []*DeviceInfoMsg
}

// DeviceFile represents the Device FIT file type.
// Describes a device's file structure and capabilities.
type DeviceFile struct {
	Softwares         []*SoftwareMsg
	Capabilities      []*CapabilitiesMsg
	FileCapabilities  []*FileCapabilitiesMsg
	MesgCapabilities  []*MesgCapabilitiesMsg
	FieldCapabilities []*FieldCapabilitiesMsg
}

// SettingsFile represents the Settings FIT file type.
// Describes a user’s parameters such as Age & Weight as well as device
// settings.
type SettingsFile struct {
	UserProfiles   []*UserProfileMsg
	HrmProfiles    []*HrmProfileMsg
	SdmProfiles    []*SdmProfileMsg
	BikeProfiles   []*BikeProfileMsg
	DeviceSettings []*DeviceSettingsMsg
}

// SportFile represents the Sport Settings FIT file type.
// Describes a user’s desired sport/zone settings.
type SportFile struct {
	ZonesTarget  *ZonesTargetMsg
	Sport        *SportMsg
	HrZones      []*HrZoneMsg
	PowerZones   []*PowerZoneMsg
	MetZones     []*MetZoneMsg
	SpeedZones   []*SpeedZoneMsg
	CadenceZones []*CadenceZoneMsg
}

// WorkoutFile represents the Workout FIT file type.
// Describes a structured activity that can be designed on a computer and
// transferred to a display device to guide a user through the activity.
type WorkoutFile struct {
	Workout      *WorkoutMsg
	WorkoutSteps []*WorkoutStepMsg
}

// CourseFile represents the Course FIT file type.
// Uses data from an activity to recreate a course.
type CourseFile struct {
	Course       *CourseMsg
	Laps         []*LapMsg
	CoursePoints []*CoursePointMsg
	Events       []*EventMsg
	Records      []*RecordMsg
}

// SchedulesFile represents the Schedules FIT file type.
// Provides scheduling of workouts and courses.
type SchedulesFile struct {
	Schedules []*ScheduleMsg
}

// WeightFile represents the Weight FIT file type.
// Records weight scale data.
type WeightFile struct {
	UserProfile  *UserProfileMsg
	WeightScales []*WeightScaleMsg
	DeviceInfos  []*DeviceInfoMsg
}

// TotalsFile represents the Totals FIT file type.
// Summarizes a user’s total activity, characterized by sport.
type TotalsFile struct {
	Totals []*TotalsMsg
}

// GoalsFile represents the Goals FIT file type.
// Describes a user’s exercise/health goals.
type GoalsFile struct {
	Goals []*GoalMsg
}

// BloodPressureFile represents the Bload Pressure FIT file type.
// Records blood pressure data.
type BloodPressureFile struct {
	UserProfile    *UserProfileMsg
	BloodPressures []*BloodPressureMsg
	DeviceInfos    []*DeviceInfoMsg
}

// MonitoringAFile represents the MonitoringA FIT file type.
// Records detailed monitoring data (i.e. logging interval < 24 Hr).
type MonitoringAFile struct {
	MonitoringInfo *MonitoringInfoMsg
	Monitorings    []*MonitoringMsg
	DeviceInfos    []*DeviceInfoMsg
}

// ActivitySummaryFile represents the Activity Summary FIT file type.
// Similar to Activity file, contains summary information only.
type ActivitySummaryFile struct {
	Activity *ActivityMsg
	Sessions []*SessionMsg
	Laps     []*LapMsg
}

// MonitoringDailyFile represents the Daily Monitoring FIT file type.
// Records daily summary monitoring data (i.e. logging interval = 24 hour).
type MonitoringDailyFile struct {
	MonitoringInfo *MonitoringInfoMsg
	Monitorings    []*MonitoringMsg
}

// MonitoringBFile represents the MonitoringB FIT file type.
// Records detailed monitoring data (i.e. logging interval < 24 Hr).
type MonitoringBFile struct {
	MonitoringInfo *MonitoringInfoMsg
	Monitorings    []*MonitoringMsg
	DeviceInfos    []*DeviceInfoMsg
}

// SegmentFile represents the Segment FIT file type.
// Describes timing data for virtual races.
type SegmentFile struct {
	SegmentId               *SegmentIdMsg
	SegmentLeaderboardEntry *SegmentLeaderboardEntryMsg
	SegmentLap              *SegmentLapMsg
	SegmentPoints           []*SegmentPointMsg
}

// SegmentListFile represents the Segment List FIT file type.
// Describes available segments.
type SegmentListFile struct {
	SegmentFiles []*SegmentFileMsg
}

// expandableMessages have some post-processing outside of raw values
type expandableMsg interface {
	expandComponents()
}

func (a *ActivityFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case ActivityMsg:
		a.Activity = &tmp
	case SessionMsg:
		a.Sessions = append(a.Sessions, &tmp)
	case LapMsg:
		a.Laps = append(a.Laps, &tmp)
	case LengthMsg:
		a.Lengths = append(a.Lengths, &tmp)
	case RecordMsg:
		a.Records = append(a.Records, &tmp)
	case EventMsg:
		a.Events = append(a.Events, &tmp)
	case HrvMsg:
		a.Hrvs = append(a.Hrvs, &tmp)
	case DeviceInfoMsg:
		a.DeviceInfos = append(a.DeviceInfos, &tmp)
	default:
	}
}

func (d *DeviceFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case SoftwareMsg:
		d.Softwares = append(d.Softwares, &tmp)
	case CapabilitiesMsg:
		d.Capabilities = append(d.Capabilities, &tmp)
	case FileCapabilitiesMsg:
		d.FileCapabilities = append(d.FileCapabilities, &tmp)
	case MesgCapabilitiesMsg:
		d.MesgCapabilities = append(d.MesgCapabilities, &tmp)
	case FieldCapabilitiesMsg:
		d.FieldCapabilities = append(d.FieldCapabilities, &tmp)
	default:
	}
}

func (s *SettingsFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case UserProfileMsg:
		s.UserProfiles = append(s.UserProfiles, &tmp)
	case HrmProfileMsg:
		s.HrmProfiles = append(s.HrmProfiles, &tmp)
	case SdmProfileMsg:
		s.SdmProfiles = append(s.SdmProfiles, &tmp)
	case BikeProfileMsg:
		s.BikeProfiles = append(s.BikeProfiles, &tmp)
	case DeviceSettingsMsg:
		s.DeviceSettings = append(s.DeviceSettings, &tmp)
	default:
	}
}

func (s *SportFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case ZonesTargetMsg:
		s.ZonesTarget = &tmp
	case SportMsg:
		s.Sport = &tmp
	case HrZoneMsg:
		s.HrZones = append(s.HrZones, &tmp)
	case PowerZoneMsg:
		s.PowerZones = append(s.PowerZones, &tmp)
	case MetZoneMsg:
		s.MetZones = append(s.MetZones, &tmp)
	case SpeedZoneMsg:
		s.SpeedZones = append(s.SpeedZones, &tmp)
	case CadenceZoneMsg:
		s.CadenceZones = append(s.CadenceZones, &tmp)
	default:
	}
}

func (w *WorkoutFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case WorkoutMsg:
		w.Workout = &tmp
	case WorkoutStepMsg:
		w.WorkoutSteps = append(w.WorkoutSteps, &tmp)
	default:
	}
}

func (c *CourseFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case CourseMsg:
		c.Course = &tmp
	case LapMsg:
		c.Laps = append(c.Laps, &tmp)
	case CoursePointMsg:
		c.CoursePoints = append(c.CoursePoints, &tmp)
	case RecordMsg:
		c.Records = append(c.Records, &tmp)
	case EventMsg:
		c.Events = append(c.Events, &tmp)
	default:
	}
}

func (s *SchedulesFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case ScheduleMsg:
		s.Schedules = append(s.Schedules, &tmp)
	default:
	}
}

func (w *WeightFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case UserProfileMsg:
		w.UserProfile = &tmp
	case WeightScaleMsg:
		w.WeightScales = append(w.WeightScales, &tmp)
	case DeviceInfoMsg:
		w.DeviceInfos = append(w.DeviceInfos, &tmp)
	default:
	}
}

func (t *TotalsFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case TotalsMsg:
		t.Totals = append(t.Totals, &tmp)
	default:
	}
}

func (g *GoalsFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case GoalMsg:
		g.Goals = append(g.Goals, &tmp)
	default:
	}
}

func (b *BloodPressureFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case UserProfileMsg:
		b.UserProfile = &tmp
	case BloodPressureMsg:
		b.BloodPressures = append(b.BloodPressures, &tmp)
	case DeviceInfoMsg:
		b.DeviceInfos = append(b.DeviceInfos, &tmp)
	default:
	}
}

func (m *MonitoringAFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case MonitoringInfoMsg:
		m.MonitoringInfo = &tmp
	case MonitoringMsg:
		m.Monitorings = append(m.Monitorings, &tmp)
	case DeviceInfoMsg:
		m.DeviceInfos = append(m.DeviceInfos, &tmp)
	default:
	}
}

func (a *ActivitySummaryFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case ActivityMsg:
		a.Activity = &tmp
	case SessionMsg:
		a.Sessions = append(a.Sessions, &tmp)
	case LapMsg:
		a.Laps = append(a.Laps, &tmp)
	default:
	}
}

func (m *MonitoringDailyFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case MonitoringInfoMsg:
		m.MonitoringInfo = &tmp
	case MonitoringMsg:
		m.Monitorings = append(m.Monitorings, &tmp)
	default:
	}
}

func (m *MonitoringBFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case MonitoringInfoMsg:
		m.MonitoringInfo = &tmp
	case MonitoringMsg:
		m.Monitorings = append(m.Monitorings, &tmp)
	case DeviceInfoMsg:
		m.DeviceInfos = append(m.DeviceInfos, &tmp)
	default:
	}
}

func (s *SegmentFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case SegmentIdMsg:
		s.SegmentId = &tmp
	case SegmentLeaderboardEntryMsg:
		s.SegmentLeaderboardEntry = &tmp
	case SegmentLapMsg:
		s.SegmentLap = &tmp
	case SegmentPointMsg:
		s.SegmentPoints = append(s.SegmentPoints, &tmp)
	default:
	}
}

func (s *SegmentListFile) add(msg interface{}) {
	switch tmp := msg.(type) {
	case SegmentFileMsg:
		s.SegmentFiles = append(s.SegmentFiles, &tmp)
	default:
	}
}
