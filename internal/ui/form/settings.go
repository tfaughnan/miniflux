// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package form // import "miniflux.app/v2/internal/ui/form"

import (
	"net/http"
	"strconv"

	"miniflux.app/v2/internal/locale"
	"miniflux.app/v2/internal/model"
)

// SettingsForm represents the settings form.
type SettingsForm struct {
	Username               string
	Password               string
	Confirmation           string
	Theme                  string
	Language               string
	Timezone               string
	EntryDirection         string
	EntryOrder             string
	EntriesPerPage         int
	KeyboardShortcuts      bool
	ShowReadingTime        bool
	CustomCSS              string
	EntrySwipe             bool
	GestureNav             string
	DisplayMode            string
	DefaultReadingSpeed    int
	CJKReadingSpeed        int
	DefaultHomePage        string
	CategoriesSortingOrder string
	MarkReadOnView         bool
	MediaPlaybackRate      float64
	BlockFilterEntryRules  string
	KeepFilterEntryRules   string
}

// Merge updates the fields of the given user.
func (s *SettingsForm) Merge(user *model.User) *model.User {
	user.Username = s.Username
	user.Theme = s.Theme
	user.Language = s.Language
	user.Timezone = s.Timezone
	user.EntryDirection = s.EntryDirection
	user.EntryOrder = s.EntryOrder
	user.EntriesPerPage = s.EntriesPerPage
	user.KeyboardShortcuts = s.KeyboardShortcuts
	user.ShowReadingTime = s.ShowReadingTime
	user.Stylesheet = s.CustomCSS
	user.EntrySwipe = s.EntrySwipe
	user.GestureNav = s.GestureNav
	user.DisplayMode = s.DisplayMode
	user.CJKReadingSpeed = s.CJKReadingSpeed
	user.DefaultReadingSpeed = s.DefaultReadingSpeed
	user.DefaultHomePage = s.DefaultHomePage
	user.CategoriesSortingOrder = s.CategoriesSortingOrder
	user.MarkReadOnView = s.MarkReadOnView
	user.MediaPlaybackRate = s.MediaPlaybackRate
	user.BlockFilterEntryRules = s.BlockFilterEntryRules
	user.KeepFilterEntryRules = s.KeepFilterEntryRules

	if s.Password != "" {
		user.Password = s.Password
	}

	return user
}

// Validate makes sure the form values are valid.
func (s *SettingsForm) Validate() *locale.LocalizedError {
	if s.Username == "" || s.Theme == "" || s.Language == "" || s.Timezone == "" || s.EntryDirection == "" || s.DisplayMode == "" || s.DefaultHomePage == "" {
		return locale.NewLocalizedError("error.settings_mandatory_fields")
	}

	if s.CJKReadingSpeed <= 0 || s.DefaultReadingSpeed <= 0 {
		return locale.NewLocalizedError("error.settings_reading_speed_is_positive")
	}

	if s.Confirmation == "" {
		// Firefox insists on auto-completing the password field.
		// If the confirmation field is blank, the user probably
		// didn't intend to change their password.
		s.Password = ""
	} else if s.Password != "" {
		if s.Password != s.Confirmation {
			return locale.NewLocalizedError("error.different_passwords")
		}
	}

	if s.MediaPlaybackRate < 0.25 || s.MediaPlaybackRate > 4 {
		return locale.NewLocalizedError("error.settings_media_playback_rate_range")
	}

	return nil
}

// NewSettingsForm returns a new SettingsForm.
func NewSettingsForm(r *http.Request) *SettingsForm {
	entriesPerPage, err := strconv.ParseInt(r.FormValue("entries_per_page"), 10, 0)
	if err != nil {
		entriesPerPage = 0
	}
	defaultReadingSpeed, err := strconv.ParseInt(r.FormValue("default_reading_speed"), 10, 0)
	if err != nil {
		defaultReadingSpeed = 0
	}
	cjkReadingSpeed, err := strconv.ParseInt(r.FormValue("cjk_reading_speed"), 10, 0)
	if err != nil {
		cjkReadingSpeed = 0
	}
	mediaPlaybackRate, err := strconv.ParseFloat(r.FormValue("media_playback_rate"), 64)
	if err != nil {
		mediaPlaybackRate = 1
	}
	return &SettingsForm{
		Username:               r.FormValue("username"),
		Password:               r.FormValue("password"),
		Confirmation:           r.FormValue("confirmation"),
		Theme:                  r.FormValue("theme"),
		Language:               r.FormValue("language"),
		Timezone:               r.FormValue("timezone"),
		EntryDirection:         r.FormValue("entry_direction"),
		EntryOrder:             r.FormValue("entry_order"),
		EntriesPerPage:         int(entriesPerPage),
		KeyboardShortcuts:      r.FormValue("keyboard_shortcuts") == "1",
		ShowReadingTime:        r.FormValue("show_reading_time") == "1",
		CustomCSS:              r.FormValue("custom_css"),
		EntrySwipe:             r.FormValue("entry_swipe") == "1",
		GestureNav:             r.FormValue("gesture_nav"),
		DisplayMode:            r.FormValue("display_mode"),
		DefaultReadingSpeed:    int(defaultReadingSpeed),
		CJKReadingSpeed:        int(cjkReadingSpeed),
		DefaultHomePage:        r.FormValue("default_home_page"),
		CategoriesSortingOrder: r.FormValue("categories_sorting_order"),
		MarkReadOnView:         r.FormValue("mark_read_on_view") == "1",
		MediaPlaybackRate:      mediaPlaybackRate,
		BlockFilterEntryRules:  r.FormValue("block_filter_entry_rules"),
		KeepFilterEntryRules:   r.FormValue("keep_filter_entry_rules"),
	}
}
