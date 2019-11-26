package main

import (
	"fmt"
	"strings"

	"github.com/3d0c/gmf"
)

type yalsaFrame struct {
	MediaType int32
	DTS       int64
	PTS       int64
	Keyframe  bool
	Size      int
}

var (
	srcVideoStream *gmf.Stream
	srcAudioStream *gmf.Stream
)

func openStream(source string) (*gmf.FmtCtx, error) {
	gmf.LogSetLevel(gmf.AV_LOG_INFO)

	var err error

	inputContext, err := gmf.NewInputCtx(source)
	if err != nil {
		return nil, err
	}
	srcVideoStream, err = inputContext.GetBestStream(gmf.AVMEDIA_TYPE_VIDEO)
	if err != nil {
		return nil, err
	}

	srcAudioStream, err = inputContext.GetBestStream(gmf.AVMEDIA_TYPE_AUDIO)
	if err != nil {
		return nil, err
	}

	inputContext.Dump()

	return inputContext, nil
}

func dumpStreamInfo(filename string, inputContext *gmf.FmtCtx) string {
	streamInfo := fmt.Sprintf("Input from: %s, bitrate: %d Kbps", filename, inputContext.BitRate()/1024)
	videoInfo := ""
	audioInfo := ""

	codecID2StrMap := map[int]string{
		27:    "h264",
		86018: "aac",
	}

	if srcVideoStream != nil {
		codecPar := srcVideoStream.GetCodecPar()

		codecString := codecID2StrMap[codecPar.GetCodecId()]
		if codecString == "" {
			codecString = fmt.Sprintf("%d", codecPar.GetCodecId())
		}

		videoInfo = fmt.Sprintf(" Stream Video: %s, %dX%d, %d Kbps, %.2f fps,",
			codecString,
			codecPar.GetWidth(), codecPar.GetHeight(),
			codecPar.GetBitRate()/1000,
			srcVideoStream.GetRFrameRate().AVR().Av2qd(),
		)
	}

	if srcAudioStream != nil {
		codecPar := srcAudioStream.GetCodecPar()

		codecString := codecID2StrMap[codecPar.GetCodecId()]
		if codecString == "" {
			codecString = fmt.Sprintf("%d", codecPar.GetCodecId())
		}

		audioInfo = fmt.Sprintf(" Stream Audio: %s, %d Kbps",
			codecString,
			codecPar.GetBitRate()/1000,
		)
	}
	return strings.Join([]string{streamInfo, videoInfo, audioInfo}, "\n")
}

func demuxing(inputContext *gmf.FmtCtx, streamInfoChan chan *yalsaFrame) {
	var (
		pkt *gmf.Packet
		err error
	)

	for {
		pkt, err = inputContext.GetNextPacket()
		if err != nil {
			break
		}
		streamIdx := pkt.StreamIndex()

		ist, err := inputContext.GetStream(streamIdx)
		if err != nil {
			break
		}
		frames, err := ist.CodecCtx().Decode(pkt)
		if err != nil {
			break
		}

		for i := range frames {
			frame := frames[i]

			kf := false
			if frame.KeyFrame() != 0 {
				kf = true
			}
			yf := &yalsaFrame{
				MediaType: ist.CodecCtx().Type(),
				PTS:       frame.Pts(),
				Keyframe:  kf,
				Size:      pkt.Size(),
			}
			streamInfoChan <- yf

			frames[i].Free()
		}

		if pkt != nil {
			pkt.Free()
		}

	}
}
