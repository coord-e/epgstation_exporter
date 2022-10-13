// Copyright 2022 coord_e
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  	 http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package epgstation

type ScheduleChannelItem struct {
	ID                 int64  `json:"id"`
	ServiceID          int    `json:"serviceId"`
	NetworkID          int    `json:"networkId"`
	Name               string `json:"name"`
	RemoteControlKeyID *int   `json:"remoteControlKeyId"`
	HasLogoData        bool   `json:"hasLogoData"`
	ChannelType        string `json:"channelType"`
}

type ScheduleProgramItem struct {
	ID                 int64              `json:"id"`
	ChannelID          int64              `json:"channelId"`
	StartAt            int64              `json:"startAt"`
	EndAt              int64              `json:"endAt"`
	IsFree             bool               `json:"isFree"`
	Name               string             `json:"name"`
	Description        *string            `json:"description"`
	Extended           *string            `json:"extended"`
	RawExtended        *map[string]string `json:"rawExtended"`
	Genre1             *int               `json:"genre1"`
	SubGenre1          *int               `json:"subGenre1"`
	Genre2             *int               `json:"genre2"`
	SubGenre2          *int               `json:"subGenre2"`
	Genre3             *int               `json:"genre3"`
	SubGenre3          *int               `json:"subGenre3"`
	VideoType          *string            `json:"videoType"`
	VideoResolution    *string            `json:"videoResolution"`
	VideoStreamContent *int               `json:"videoStreamContent"`
	VideoComponentType *int               `json:"videoComponentType"`
	AudioSamplingRate  *int               `json:"audioSamplingRate"`
	AudioComponentType *int               `json:"audioComponentType"`
}

type Schedule struct {
	Channel  ScheduleChannelItem   `json:"channel"`
	Programs []ScheduleProgramItem `json:"programs"`
}

type ChannelItem struct {
	ID                 int64  `json:"id"`
	ServiceID          int    `json:"serviceId"`
	NetworkID          int    `json:"networkId"`
	Name               string `json:"name"`
	HalfWidthName      string `json:"halfWidthName"`
	RemoteControlKeyID *int   `json:"remoteControlKeyId"`
	HasLogoData        bool   `json:"hasLogoData"`
	ChannelType        string `json:"channelType"`
	Channel            string `json:"channel"`
}

type StorageItem struct {
	Name      string `json:"name"`
	Available int64  `json:"available"`
	Used      int64  `json:"used"`
	Total     int64  `json:"total"`
}

type StorageInfo struct {
	Items []StorageItem `json:"items"`
}

type StreamInfoItem struct {
	StreamID    int                `json:"streamId"`
	Type        string             `json:"type"`
	Mode        int                `json:"mode"`
	IsEnable    bool               `json:"isEnable"`
	ChannelID   int64              `json:"channelId"`
	Name        string             `json:"name"`
	StartAt     int64              `json:"startAt"`
	EndAt       int64              `json:"endAt"`
	Description *string            `json:"description"`
	Extended    *string            `json:"extended"`
	RawExtended *map[string]string `json:"rawExtended"`
	VideoFileID *int               `json:"videoFileId"`
	RecordedID  *int               `json:"recordedId"`
}

type StreamInfo struct {
	Items []StreamInfoItem `json:"items"`
}

type VideoFile struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
}

type DropLogFile struct {
	ID            int64 `json:"id"`
	ErrorCnt      int64 `json:"errorCnt"`
	DropCnt       int64 `json:"dropCnt"`
	ScramblingCnt int64 `json:"scramblingCnt"`
}

type RecordedTag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type RecordedItem struct {
	ID                 int64              `json:"id"`
	RuleID             *int64             `json:"ruleId"`
	ProgramID          *int64             `json:"programId"`
	ChannelID          int64              `json:"channelId"`
	StartAt            int64              `json:"startAt"`
	EndAt              int64              `json:"endAt"`
	Name               string             `json:"name"`
	Description        *string            `json:"description"`
	Extended           *string            `json:"extended"`
	RawExtended        *map[string]string `json:"rawExtended"`
	Genre1             *int               `json:"genre1"`
	SubGenre1          *int               `json:"subGenre1"`
	Genre2             *int               `json:"genre2"`
	SubGenre2          *int               `json:"subGenre2"`
	Genre3             *int               `json:"genre3"`
	SubGenre3          *int               `json:"subGenre3"`
	VideoType          *string            `json:"videoType"`
	VideoResolution    *string            `json:"videoResolution"`
	VideoStreamContent *int               `json:"videoStreamContent"`
	VideoComponentType *int               `json:"videoComponentType"`
	AudioSamplingRate  *int               `json:"audioSamplingRate"`
	AudioComponentType *int               `json:"audioComponentType"`
	IsRecording        bool               `json:"isRecording"`
	Thumbnails         *[]int64           `json:"thumbnails"`
	VideoFiles         *[]VideoFile       `json:"videoFiles"`
	DropLogFile        *DropLogFile       `json:"dropLogFile"`
	Tags               *[]RecordedTag     `json:"tags"`
	IsEncoding         bool               `json:"isEncoding"`
	IsProtected        bool               `json:"isProtected"`
}

type Records struct {
	Records []RecordedItem `json:"records"`
	Total   int64          `json:"total"`
}

type EncodeProgramItem struct {
	ID       int64        `json:"id"`
	Mode     string       `json:"mode"`
	Recorded RecordedItem `json:"recorded"`
	Percent  *float64     `json:"percent"`
	Log      *string      `json:"log"`
}

type EncodeInfo struct {
	RunningItems []EncodeProgramItem `json:"runningItems"`
	WaitItems    []EncodeProgramItem `json:"waitItems"`
}
