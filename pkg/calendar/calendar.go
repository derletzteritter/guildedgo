package calendar

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/itschip/guildedgo/pkg/channel"
)

type Client interface {
	PerformRequest(method, url string, data any) (io.ReadCloser, error)
	Decode(body io.ReadCloser, v any) error
}

const (
	guildedApi = "https://www.guilded.gg/api/v1"
)

type CalendarEvent struct {
	ID               int    `json:"id"`
	ServerID         string `json:"serverId"`
	ChannelID        string `json:"channelId"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	Location         string `json:"location,omitempty"`
	URL              string `json:"url,omitempty"`
	Color            int    `json:"color,omitempty"`
	Repeats          bool   `json:"repeats,omitempty"`
	SeriesID         string `json:"seriesId,omitempty"`
	RoleIDs          []int  `json:"roleIds,omitempty"`
	RSVPDisabled     bool   `json:"rsvpDisabled,omitempty"`
	IsAllDay         bool   `json:"isAllDay,omitempty"`
	RSVPLimit        int    `json:"rsvpLimit,omitempty"`
	AutoFillWaitlist bool   `json:"autoFillWaitlist,omitempty"`
	StartsAt         string `json:"startsAt"`
	Duration         int    `json:"duration,omitempty"`
	IsPrivate        bool   `json:"isPrivate,omitempty"`
	channel.Mention  `json:"mentions,omitempty"`
	CreatedAt        string `json:"createdAt"`
	CreatedBy        string `json:"createdBy"`
	Cancellation     struct {
		Description string `json:"description,omitempty"`
		CreatedBy   string `json:"createdBy,omitempty"`
	} `json:"cancellation,omitempty"`
}

type CalendarEventRsvp struct {
	CalendarEventID int    `json:"calendarEventId"`
	ChannelID       string `json:"channelId"`
	ServerID        string `json:"serverId"`
	UserID          string `json:"userId"`
	Status          string `json:"status"`
	CreatedBy       string `json:"createdBy"`
	CreatedAt       string `json:"createdAt"`
	UpdatedBy       string `json:"updatedBy,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
}

type CreateParams struct {
	// Name of the event (min length 1; max length 60)
	Name string `json:"name"`

	// Description of the event (min length 1; max length 8000)
	Description string `json:"description,omitempty"`

	// Location of the event (min length 1; max length 8000)
	Location string `json:"location,omitempty"`

	// The ISO 8601 timestamp that the event starts at
	StartsAt string `json:"startsAt,omitempty"`

	// A URL to associate with the event
	URL string `json:"url,omitempty"`

	// The integer value corresponds to the decimal RGB representation for the color.
	// The color of the event when viewing in the calendar (min 0; max 16777215)
	Color int `json:"color,omitempty"`

	// Does the event last all day? If passed with duration, duration will only be applied
	// if it is an interval of minutes represented in days (e.g., duration: 2880)
	IsAllDay bool `json:"isAllDay,omitempty"`

	// When disabled, users will not be able to RSVP to the event
	RSVPDisabled bool `json:"rsvpDisabled,omitempty"`

	// The number of RSVPs to allow before waitlisting RSVPs (min 1)
	RSVPLimit int `json:"rsvpLimit,omitempty"`

	// When rsvpLimit is set, users from the waitlist will be added as space becomes available in the event
	AutoFillWaitlist bool `json:"autoFillWaitlist,omitempty"`

	// The duration of the event in minutes (min 1)
	Duration int `json:"duration,omitempty"`

	IsPrivate bool `json:"isPrivate,omitempty"`

	// The role IDs to restrict the event to (min items 1; must have unique items true)
	RoleIDs []int `json:"roleIds,omitempty"`

	RepeatInfo `json:"repeatInfo,omitempty"`
}

type RepeatInfo struct {
	// How often you want your event to repeat (important note: this will repeat for
	// the next 365 days unless custom is defined) (default once)
	//
	// string ("once", "everyDay", "everyWeek", "everyMonth", or "custom")
	Type string `json:"type"`

	// Used to control the end date of the event repeat (only used when type is custom;
	// if used with endDate, the earliest resultant date of the two will be used) (max 24)
	EndsAfterOccurences int `json:"endsAfterOccurences,omitempty"`

	// The ISO 8601 timestamp that the event ends at.
	// Used to control the end date of the event repeat (only used when type is custom;
	// if used with endsAfterOccurrences, the earliest resultant date of the two will be used)
	EndDate string `json:"endDate,omitempty"`

	// Used to control the day of the week that the event should repeat on
	// (only used when type is custom and when every.interval is week) (min items 1)
	//
	// string[] ("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", or "saturday")
	On []string `json:"on,omitempty"`
}

func CreateEvent(c Client, channelID string, params *CreateParams) (*CalendarEvent, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/events"

	var v struct {
		CalendarEvent CalendarEvent `json:"calendarEvent"`
	}

	body, err := c.PerformRequest(http.MethodPost, endpoint, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode channel response: %w", err)
	}

	return &v.CalendarEvent, nil
}

type GetQueryParams struct {
	Before string
	After  string
	Limit  int
}

func GetEvents(c Client, channelID string, params *GetQueryParams) ([]CalendarEvent, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/events"

	urlValues := url.Values{}
	if params.Before != "" {
		urlValues.Add("before", params.Before)
	}

	if params.After != "" {
		urlValues.Add("after", params.After)
	}

	if params.Limit != 0 {
		urlValues.Add("limit", strconv.Itoa(params.Limit))
	}

	endpoint += "?" + urlValues.Encode()

	var v struct {
		CalendarEvents []CalendarEvent `json:"calendarEvents"`
	}

	body, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode events response: %w", err)
	}

	return v.CalendarEvents, nil
}

func GetEvent(c Client, channelID string, calendarEventID int) (*CalendarEvent, error) {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/events/" + strconv.Itoa(calendarEventID)

	var v struct {
		CalendarEvent CalendarEvent `json:"calendarEvent"`
	}

	body, err := c.PerformRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)
	}

	err = c.Decode(body, &v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode event response: %w", err)
	}

	return &v.CalendarEvent, nil
}

func DeleteEvent(c Client, channelID, calendarEventID string) error {
	var err error
	endpoint := guildedApi + "/channels" + channelID + "/events/" + calendarEventID

	_, err = c.PerformRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	return nil
}
