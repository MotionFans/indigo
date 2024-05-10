// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package chat

// schema: chat.bsky.convo.defs

import (
	"encoding/json"
	"fmt"

	appbskytypes "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/lex/util"
)

// ConvoDefs_ConvoView is a "convoView" in the chat.bsky.convo.defs schema.
type ConvoDefs_ConvoView struct {
	Id          string                                     `json:"id" cborgen:"id"`
	LastMessage *ConvoDefs_ConvoView_LastMessage           `json:"lastMessage,omitempty" cborgen:"lastMessage,omitempty"`
	Members     []*appbskytypes.ActorDefs_ProfileViewBasic `json:"members" cborgen:"members"`
	Muted       bool                                       `json:"muted" cborgen:"muted"`
	Rev         string                                     `json:"rev" cborgen:"rev"`
	UnreadCount int64                                      `json:"unreadCount" cborgen:"unreadCount"`
}

type ConvoDefs_ConvoView_LastMessage struct {
	ConvoDefs_MessageView        *ConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ConvoDefs_DeletedMessageView
}

func (t *ConvoDefs_ConvoView_LastMessage) MarshalJSON() ([]byte, error) {
	if t.ConvoDefs_MessageView != nil {
		t.ConvoDefs_MessageView.LexiconTypeID = "chat.bsky.convo.defs#messageView"
		return json.Marshal(t.ConvoDefs_MessageView)
	}
	if t.ConvoDefs_DeletedMessageView != nil {
		t.ConvoDefs_DeletedMessageView.LexiconTypeID = "chat.bsky.convo.defs#deletedMessageView"
		return json.Marshal(t.ConvoDefs_DeletedMessageView)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *ConvoDefs_ConvoView_LastMessage) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "chat.bsky.convo.defs#messageView":
		t.ConvoDefs_MessageView = new(ConvoDefs_MessageView)
		return json.Unmarshal(b, t.ConvoDefs_MessageView)
	case "chat.bsky.convo.defs#deletedMessageView":
		t.ConvoDefs_DeletedMessageView = new(ConvoDefs_DeletedMessageView)
		return json.Unmarshal(b, t.ConvoDefs_DeletedMessageView)

	default:
		return nil
	}
}

// ConvoDefs_DeletedMessageView is a "deletedMessageView" in the chat.bsky.convo.defs schema.
//
// RECORDTYPE: ConvoDefs_DeletedMessageView
type ConvoDefs_DeletedMessageView struct {
	LexiconTypeID string                       `json:"$type,const=chat.bsky.convo.defs#deletedMessageView" cborgen:"$type,const=chat.bsky.convo.defs#deletedMessageView"`
	Id            string                       `json:"id" cborgen:"id"`
	Rev           string                       `json:"rev" cborgen:"rev"`
	Sender        *ConvoDefs_MessageViewSender `json:"sender,omitempty" cborgen:"sender,omitempty"`
	SentAt        string                       `json:"sentAt" cborgen:"sentAt"`
}

// ConvoDefs_LogBeginConvo is a "logBeginConvo" in the chat.bsky.convo.defs schema.
//
// RECORDTYPE: ConvoDefs_LogBeginConvo
type ConvoDefs_LogBeginConvo struct {
	LexiconTypeID string `json:"$type,const=chat.bsky.convo.defs#logBeginConvo" cborgen:"$type,const=chat.bsky.convo.defs#logBeginConvo"`
	ConvoId       string `json:"convoId" cborgen:"convoId"`
	Rev           string `json:"rev" cborgen:"rev"`
}

// ConvoDefs_LogCreateMessage is a "logCreateMessage" in the chat.bsky.convo.defs schema.
//
// RECORDTYPE: ConvoDefs_LogCreateMessage
type ConvoDefs_LogCreateMessage struct {
	LexiconTypeID string                              `json:"$type,const=chat.bsky.convo.defs#logCreateMessage" cborgen:"$type,const=chat.bsky.convo.defs#logCreateMessage"`
	ConvoId       string                              `json:"convoId" cborgen:"convoId"`
	Message       *ConvoDefs_LogCreateMessage_Message `json:"message" cborgen:"message"`
	Rev           string                              `json:"rev" cborgen:"rev"`
}

type ConvoDefs_LogCreateMessage_Message struct {
	ConvoDefs_MessageView        *ConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ConvoDefs_DeletedMessageView
}

func (t *ConvoDefs_LogCreateMessage_Message) MarshalJSON() ([]byte, error) {
	if t.ConvoDefs_MessageView != nil {
		t.ConvoDefs_MessageView.LexiconTypeID = "chat.bsky.convo.defs#messageView"
		return json.Marshal(t.ConvoDefs_MessageView)
	}
	if t.ConvoDefs_DeletedMessageView != nil {
		t.ConvoDefs_DeletedMessageView.LexiconTypeID = "chat.bsky.convo.defs#deletedMessageView"
		return json.Marshal(t.ConvoDefs_DeletedMessageView)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *ConvoDefs_LogCreateMessage_Message) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "chat.bsky.convo.defs#messageView":
		t.ConvoDefs_MessageView = new(ConvoDefs_MessageView)
		return json.Unmarshal(b, t.ConvoDefs_MessageView)
	case "chat.bsky.convo.defs#deletedMessageView":
		t.ConvoDefs_DeletedMessageView = new(ConvoDefs_DeletedMessageView)
		return json.Unmarshal(b, t.ConvoDefs_DeletedMessageView)

	default:
		return nil
	}
}

// ConvoDefs_LogDeleteMessage is a "logDeleteMessage" in the chat.bsky.convo.defs schema.
//
// RECORDTYPE: ConvoDefs_LogDeleteMessage
type ConvoDefs_LogDeleteMessage struct {
	LexiconTypeID string                              `json:"$type,const=chat.bsky.convo.defs#logDeleteMessage" cborgen:"$type,const=chat.bsky.convo.defs#logDeleteMessage"`
	ConvoId       string                              `json:"convoId" cborgen:"convoId"`
	Message       *ConvoDefs_LogDeleteMessage_Message `json:"message" cborgen:"message"`
	Rev           string                              `json:"rev" cborgen:"rev"`
}

type ConvoDefs_LogDeleteMessage_Message struct {
	ConvoDefs_MessageView        *ConvoDefs_MessageView
	ConvoDefs_DeletedMessageView *ConvoDefs_DeletedMessageView
}

func (t *ConvoDefs_LogDeleteMessage_Message) MarshalJSON() ([]byte, error) {
	if t.ConvoDefs_MessageView != nil {
		t.ConvoDefs_MessageView.LexiconTypeID = "chat.bsky.convo.defs#messageView"
		return json.Marshal(t.ConvoDefs_MessageView)
	}
	if t.ConvoDefs_DeletedMessageView != nil {
		t.ConvoDefs_DeletedMessageView.LexiconTypeID = "chat.bsky.convo.defs#deletedMessageView"
		return json.Marshal(t.ConvoDefs_DeletedMessageView)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *ConvoDefs_LogDeleteMessage_Message) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "chat.bsky.convo.defs#messageView":
		t.ConvoDefs_MessageView = new(ConvoDefs_MessageView)
		return json.Unmarshal(b, t.ConvoDefs_MessageView)
	case "chat.bsky.convo.defs#deletedMessageView":
		t.ConvoDefs_DeletedMessageView = new(ConvoDefs_DeletedMessageView)
		return json.Unmarshal(b, t.ConvoDefs_DeletedMessageView)

	default:
		return nil
	}
}

// ConvoDefs_LogLeaveConvo is a "logLeaveConvo" in the chat.bsky.convo.defs schema.
//
// RECORDTYPE: ConvoDefs_LogLeaveConvo
type ConvoDefs_LogLeaveConvo struct {
	LexiconTypeID string `json:"$type,const=chat.bsky.convo.defs#logLeaveConvo" cborgen:"$type,const=chat.bsky.convo.defs#logLeaveConvo"`
	ConvoId       string `json:"convoId" cborgen:"convoId"`
	Rev           string `json:"rev" cborgen:"rev"`
}

// ConvoDefs_Message is a "message" in the chat.bsky.convo.defs schema.
type ConvoDefs_Message struct {
	Embed *ConvoDefs_Message_Embed `json:"embed,omitempty" cborgen:"embed,omitempty"`
	// facets: Annotations of text (mentions, URLs, hashtags, etc)
	Facets []*appbskytypes.RichtextFacet `json:"facets,omitempty" cborgen:"facets,omitempty"`
	Id     *string                       `json:"id,omitempty" cborgen:"id,omitempty"`
	Text   string                        `json:"text" cborgen:"text"`
}

// ConvoDefs_MessageRef is a "messageRef" in the chat.bsky.convo.defs schema.
type ConvoDefs_MessageRef struct {
	Did       string `json:"did" cborgen:"did"`
	MessageId string `json:"messageId" cborgen:"messageId"`
}

// ConvoDefs_MessageView is a "messageView" in the chat.bsky.convo.defs schema.
//
// RECORDTYPE: ConvoDefs_MessageView
type ConvoDefs_MessageView struct {
	LexiconTypeID string                       `json:"$type,const=chat.bsky.convo.defs#messageView" cborgen:"$type,const=chat.bsky.convo.defs#messageView"`
	Embed         *ConvoDefs_MessageView_Embed `json:"embed,omitempty" cborgen:"embed,omitempty"`
	// facets: Annotations of text (mentions, URLs, hashtags, etc)
	Facets []*appbskytypes.RichtextFacet `json:"facets,omitempty" cborgen:"facets,omitempty"`
	Id     string                        `json:"id" cborgen:"id"`
	Rev    string                        `json:"rev" cborgen:"rev"`
	Sender *ConvoDefs_MessageViewSender  `json:"sender,omitempty" cborgen:"sender,omitempty"`
	SentAt string                        `json:"sentAt" cborgen:"sentAt"`
	Text   string                        `json:"text" cborgen:"text"`
}

// ConvoDefs_MessageViewSender is a "messageViewSender" in the chat.bsky.convo.defs schema.
type ConvoDefs_MessageViewSender struct {
	Did string `json:"did" cborgen:"did"`
}

type ConvoDefs_MessageView_Embed struct {
	EmbedRecord *appbskytypes.EmbedRecord
}

func (t *ConvoDefs_MessageView_Embed) MarshalJSON() ([]byte, error) {
	if t.EmbedRecord != nil {
		t.EmbedRecord.LexiconTypeID = "app.bsky.embed.record"
		return json.Marshal(t.EmbedRecord)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *ConvoDefs_MessageView_Embed) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.record":
		t.EmbedRecord = new(appbskytypes.EmbedRecord)
		return json.Unmarshal(b, t.EmbedRecord)

	default:
		return nil
	}
}

type ConvoDefs_Message_Embed struct {
	EmbedRecord *appbskytypes.EmbedRecord
}

func (t *ConvoDefs_Message_Embed) MarshalJSON() ([]byte, error) {
	if t.EmbedRecord != nil {
		t.EmbedRecord.LexiconTypeID = "app.bsky.embed.record"
		return json.Marshal(t.EmbedRecord)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *ConvoDefs_Message_Embed) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.embed.record":
		t.EmbedRecord = new(appbskytypes.EmbedRecord)
		return json.Unmarshal(b, t.EmbedRecord)

	default:
		return nil
	}
}
