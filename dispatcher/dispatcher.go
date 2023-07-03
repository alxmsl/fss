package dispatcher

import (
	"github.com/alxmsl/fss/naming"
	"io"
)

type Interface interface {
	DownloadPart(naming.Interface, string, io.Writer) error
	UploadPart(naming.Interface, []byte) error
}
