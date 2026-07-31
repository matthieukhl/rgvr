package models

import (
	"fmt"
	"time"
)

type SnoozeLog struct {
	UserID              int      `json:"user_id"`
	TeamID              int      `json:"team_id"`
	LastIDOffsetSetted  int      `json:"last_id_offset_setted"`
	LimitCountSetted    int      `json:"limit_count_setted"`
	SnoozeLogListCount  int      `json:"snooze_log_list_count"`
	SnoozeLogList       []Snooze `json:"snooze_log_list"`
	TotalSnoozeLogCount int      `json:"total_snooze_log_count"`
}

type Snooze struct {
	State        string    `json:"state"`
	Name         string    `json:"name"`
	Label        string    `json:"label"`
	LabelFr      string    `json:"label_fr"`
	LabelEn      string    `json:"label_en"`
	LabelEs      string    `json:"label_es"`
	Comment      string    `json:"comment"`
	CreationDate time.Time `json:"creation_date"`
	EndDate      time.Time `json:"end_date"`
}

func (sl SnoozeLog) TableHeader() []string {
	return []string{
		"UserID",
		"TeamID",
		"LastIDOffsetSetted",
		"LimitCountSetted",
		"SnoozeLogListCount",
		"TotalSnoozeLogCount",
	}
}

func (sl SnoozeLog) TableRow() []string {
	return []string{
		fmt.Sprintf("%d", sl.UserID),
		fmt.Sprintf("%d", sl.TeamID),
		fmt.Sprintf("%d", sl.LastIDOffsetSetted),
		fmt.Sprintf("%d", sl.LimitCountSetted),
		fmt.Sprintf("%d", sl.SnoozeLogListCount),
		fmt.Sprintf("%d", sl.TotalSnoozeLogCount),
	}
}
