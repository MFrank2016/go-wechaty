// Code generated by "stringer -type=MessageType"; DO NOT EDIT.

package schemas

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MessageTypeUnknown-0]
	_ = x[MessageTypeAttachment-1]
	_ = x[MessageTypeAudio-2]
	_ = x[MessageTypeContact-3]
	_ = x[MessageTypeChatHistory-4]
	_ = x[MessageTypeEmoticon-5]
	_ = x[MessageTypeImage-6]
	_ = x[MessageTypeText-7]
	_ = x[MessageTypeLocation-8]
	_ = x[MessageTypeMiniProgram-9]
	_ = x[MessageTypeTransfer-10]
	_ = x[MessageTypeRedEnvelope-11]
	_ = x[MessageTypeRecalled-12]
	_ = x[MessageTypeUrl-13]
	_ = x[MessageTypeVideo-14]
}

const _MessageType_name = "MessageTypeUnknownMessageTypeAttachmentMessageTypeAudioMessageTypeContactMessageTypeChatHistoryMessageTypeEmoticonMessageTypeImageMessageTypeTextMessageTypeLocationMessageTypeMiniProgramMessageTypeTransferMessageTypeRedEnvelopeMessageTypeRecalledMessageTypeUrlMessageTypeVideo"

var _MessageType_index = [...]uint16{0, 18, 39, 55, 73, 95, 114, 130, 145, 164, 186, 205, 227, 246, 260, 276}

func (i MessageType) String() string {
	if i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}
