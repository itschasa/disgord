package disgord

// ref https://discord.com/developers/docs/resources/channel#thread-member-object
type ThreadMember struct {
	ID            Snowflake `json:"id,omitempty"`
	UserID        Snowflake `json:"user_id,omitempty"`
	JoinTimestamp Time      `json:"join_timestamp"`
	Flags         Flag      `json:"flags"`
}

// ref https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
	Archived            bool `json:"archived"`
	AutoArchiveDuration int  `json:"auto_archive_duration"`
	ArchiveTimestamp    Time `json:"archive_timestamp"`
	Locked              bool `json:"locked"`
	Invitable           bool `json:"inviteable,omitempty"`
}

type AutoArchiveDurationTime int

const (
	AutoArchiveThreadMinute AutoArchiveDurationTime = 60
	AutoArchiveThreadDay    AutoArchiveDurationTime = 1440
	// guild must be boosted to use the below auto archive durations.
	// ref: https://discord.com/developers/docs/resources/channel#start-thread-with-message-json-params
	AutoArchiveThreadThreeDays AutoArchiveDurationTime = 4320
	AutoArchiveThreadWeek      AutoArchiveDurationTime = 10080
)

// ref https://discord.com/developers/docs/resources/channel#start-thread-with-message-json-params
type CreateThread struct {
	Name                string                  `json:"name"`
	AutoArchiveDuration AutoArchiveDurationTime `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    int                     `json:"rate_limit_per_user,omitempty"`

	// Reason is a X-Audit-Log-Reason header field that will show up on the audit log for this action.
	Reason string `json:"-"`
}

// ref https://discord.com/developers/docs/resources/channel#start-thread-without-message-json-params
type CreateThreadNoMessage struct {
	Name                string                  `json:"name"`
	AutoArchiveDuration AutoArchiveDurationTime `json:"auto_archive_duration,omitempty"`
	// In API v9, type defaults to PRIVATE_THREAD in order to match the behavior when
	// thread documentation was first published. In API v10 this will be changed to be a required field, with no default.
	Type             ChannelType `json:"type,omitempty"`
	Invitable        bool        `json:"invitable,omitempty"`
	RateLimitPerUser int         `json:"rate_limit_per_user,omitempty"`

	// Reason is a X-Audit-Log-Reason header field that will show up on the audit log for this action.
	Reason string `json:"-"`
}