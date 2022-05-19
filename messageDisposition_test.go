package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockMessageDisposition creates a MessageDisposition
func mockMessageDisposition() *MessageDisposition {
	md := NewMessageDisposition()
	md.FormatVersion = FormatVersion
	md.TestProductionCode = EnvironmentProduction
	md.MessageDuplicationCode = MessageDuplicationOriginal
	md.MessageStatusIndicator = "2"
	return md
}

// TestMockMessageDisposition validates mockMessageDisposition
func TestMockMessageDisposition(t *testing.T) {
	md := mockMessageDisposition()

	require.NoError(t, md.Validate(), "mockMessageDisposition does not validate and will break other tests")
}

// TestParseMessageDisposition parses a known MessageDisposition record string
func TestParseMessageDisposition(t *testing.T) {
	var line = "{1100}30P 2"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseMessageDisposition())

	record := r.currentFEDWireMessage.MessageDisposition
	require.Equal(t, "30", record.FormatVersion)
	require.Equal(t, "P", record.TestProductionCode)
	require.Empty(t, record.MessageDuplicationCode)
	require.Equal(t, "2", record.MessageStatusIndicator)
}

// TestWriteMessageDisposition writes a MessageDisposition record string
func TestWriteMessageDisposition(t *testing.T) {
	var line = "{1100}30P 2"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseMessageDisposition())

	record := r.currentFEDWireMessage.MessageDisposition
	require.Equal(t, line, record.String())
}

// TestMessageDispositionTagError validates a MessageDisposition tag
func TestMessageDispositionTagError(t *testing.T) {
	md := mockMessageDisposition()
	md.tag = "{9999}"

	require.EqualError(t, md.Validate(), fieldError("tag", ErrValidTagForType, md.tag).Error())
}
