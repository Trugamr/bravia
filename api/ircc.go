package api

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	irccPath = "/sony/ircc"
)

// IRCCCommand represents predefined IRCC remote control commands
type IRCCCommand string

const (
	IRCCHome                       IRCCCommand = "AAAAAQAAAAEAAABgAw=="
	IRCCReturn                     IRCCCommand = "AAAAAgAAAJcAAAAjAw=="
	IRCCNum1                       IRCCCommand = "AAAAAQAAAAEAAAAAAw=="
	IRCCNum2                       IRCCCommand = "AAAAAQAAAAEAAAABAw=="
	IRCCNum3                       IRCCCommand = "AAAAAQAAAAEAAAACAw=="
	IRCCNum4                       IRCCCommand = "AAAAAQAAAAEAAAADAw=="
	IRCCNum5                       IRCCCommand = "AAAAAQAAAAEAAAAEAw=="
	IRCCNum6                       IRCCCommand = "AAAAAQAAAAEAAAAFAw=="
	IRCCNum7                       IRCCCommand = "AAAAAQAAAAEAAAAGAw=="
	IRCCNum8                       IRCCCommand = "AAAAAQAAAAEAAAAHAw=="
	IRCCNum9                       IRCCCommand = "AAAAAQAAAAEAAAAIAw=="
	IRCCNum0                       IRCCCommand = "AAAAAQAAAAEAAAAJAw=="
	IRCCDOT                        IRCCCommand = "AAAAAgAAAJcAAAAdAw=="
	IRCCVolumeUp                   IRCCCommand = "AAAAAQAAAAEAAAASAw=="
	IRCCVolumeDown                 IRCCCommand = "AAAAAQAAAAEAAAATAw=="
	IRCCMute                       IRCCCommand = "AAAAAQAAAAEAAAAUAw=="
	IRCCTvPower                    IRCCCommand = "AAAAAQAAAAEAAAAVAw=="
	IRCCEPG                        IRCCCommand = "AAAAAgAAAKQAAABbAw=="
	IRCCConfirm                    IRCCCommand = "AAAAAQAAAAEAAABlAw=="
	IRCCChannelUp                  IRCCCommand = "AAAAAQAAAAEAAAAQAw=="
	IRCCChannelDown                IRCCCommand = "AAAAAQAAAAEAAAARAw=="
	IRCCUp                         IRCCCommand = "AAAAAQAAAAEAAAB0Aw=="
	IRCCDown                       IRCCCommand = "AAAAAQAAAAEAAAB1Aw=="
	IRCCLeft                       IRCCCommand = "AAAAAQAAAAEAAAA0Aw=="
	IRCCRight                      IRCCCommand = "AAAAAQAAAAEAAAAzAw=="
	IRCCDisplay                    IRCCCommand = "AAAAAQAAAAEAAAA6Aw=="
	IRCCSubTitle                   IRCCCommand = "AAAAAgAAAJcAAAAoAw=="
	IRCCAudio                      IRCCCommand = "AAAAAQAAAAEAAAAXAw=="
	IRCCMediaAudioTrack            IRCCCommand = "AAAAAQAAAAEAAAAXAw=="
	IRCCJump                       IRCCCommand = "AAAAAQAAAAEAAAA7Aw=="
	IRCCExit                       IRCCCommand = "AAAAAQAAAAEAAABjAw=="
	IRCCTv                         IRCCCommand = "AAAAAQAAAAEAAAAkAw=="
	IRCCInput                      IRCCCommand = "AAAAAQAAAAEAAAAlAw=="
	IRCCTvInput                    IRCCCommand = "AAAAAQAAAAEAAAAlAw=="
	IRCCRed                        IRCCCommand = "AAAAAgAAAJcAAAAlAw=="
	IRCCGreen                      IRCCCommand = "AAAAAgAAAJcAAAAmAw=="
	IRCCYellow                     IRCCCommand = "AAAAAgAAAJcAAAAnAw=="
	IRCCBlue                       IRCCCommand = "AAAAAgAAAJcAAAAkAw=="
	IRCCTeletext                   IRCCCommand = "AAAAAQAAAAEAAAA/Aw=="
	IRCCStop                       IRCCCommand = "AAAAAgAAAJcAAAAYAw=="
	IRCCRewind                     IRCCCommand = "AAAAAgAAAJcAAAAbAw=="
	IRCCForward                    IRCCCommand = "AAAAAgAAAJcAAAAcAw=="
	IRCCPrev                       IRCCCommand = "AAAAAgAAAJcAAAA8Aw=="
	IRCCNext                       IRCCCommand = "AAAAAgAAAJcAAAA9Aw=="
	IRCCPlay                       IRCCCommand = "AAAAAgAAAJcAAAAaAw=="
	IRCCRec                        IRCCCommand = "AAAAAgAAAJcAAAAgAw=="
	IRCCPause                      IRCCCommand = "AAAAAgAAAJcAAAAZAw=="
	IRCCOneTouchView               IRCCCommand = "AAAAAgAAABoAAABlAw=="
	IRCCGooglePlay                 IRCCCommand = "AAAAAgAAAMQAAABGAw=="
	IRCCNetflix                    IRCCCommand = "AAAAAgAAABoAAAB8Aw=="
	IRCCPartnerApp6                IRCCCommand = "AAAAAwAACB8AAAAFAw=="
	IRCCPartnerApp5                IRCCCommand = "AAAAAwAACB8AAAAEAw=="
	IRCCYouTube                    IRCCCommand = "AAAAAgAAAMQAAABHAw=="
	IRCCPartnerApp9                IRCCCommand = "AAAAAwAACB8AAAAIAw=="
	IRCCPartnerApp7                IRCCCommand = "AAAAAwAACB8AAAAGAw=="
	IRCCActionMenu                 IRCCCommand = "AAAAAgAAAMQAAABLAw=="
	IRCCApplicationLauncher        IRCCCommand = "AAAAAgAAAMQAAAAqAw=="
	IRCCHelp                       IRCCCommand = "AAAAAgAAAMQAAABNAw=="
	IRCCShopRemoteControlForcedDynamic IRCCCommand = "AAAAAgAAAJcAAABqAw=="
	IRCCWakeUp                     IRCCCommand = "AAAAAQAAAAEAAAAuAw=="
	IRCCPowerOff                   IRCCCommand = "AAAAAQAAAAEAAAAvAw=="
	IRCCSleep                      IRCCCommand = "AAAAAQAAAAEAAAAvAw=="
	IRCCHdmi1                      IRCCCommand = "AAAAAgAAABoAAABaAw=="
	IRCCHdmi2                      IRCCCommand = "AAAAAgAAABoAAABbAw=="
	IRCCHdmi3                      IRCCCommand = "AAAAAgAAABoAAABcAw=="
	IRCCDemoMode                   IRCCCommand = "AAAAAgAAAJcAAAB8Aw=="
)

// IRCCService handles IRCC (infrared compatible control) commands
type IRCCService service

// SendIRCCCommand sends an IRCC command to control the TV remotely
// The command parameter can be either a predefined IRCCCommand constant or a custom IRCC code string
func (s *IRCCService) SendIRCCCommand(command string) (*http.Response, error) {
	body := s.buildIRCCXML(command)

	req, err := s.client.NewRequest(http.MethodPost, irccPath, nil)
	if err != nil {
		return nil, err
	}

	// Override the body with XML
	req.Body = http.NoBody
	req.GetBody = nil
	req.ContentLength = int64(len(body))
	req.Body = &readCloser{bytes.NewReader(body)}

	// Set SOAP-specific headers
	req.Header.Set("Content-Type", "text/xml; charset=UTF-8")
	req.Header.Set("SOAPACTION", `"urn:schemas-sony-com:service:IRCC:1#X_SendIRCC"`)

	resp, err := s.client.client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	return resp, nil
}

// buildIRCCXML builds the SOAP XML body for an IRCC command
func (s *IRCCService) buildIRCCXML(code string) []byte {
	xml := fmt.Sprintf(`<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	<s:Body>
		<u:X_SendIRCC xmlns:u="urn:schemas-sony-com:service:IRCC:1">
			<IRCCCode>%s</IRCCCode>
		</u:X_SendIRCC>
	</s:Body>
</s:Envelope>`, code)
	return []byte(xml)
}

// readCloser wraps a bytes.Reader to implement io.ReadCloser
type readCloser struct {
	*bytes.Reader
}

func (rc *readCloser) Close() error {
	return nil
}
