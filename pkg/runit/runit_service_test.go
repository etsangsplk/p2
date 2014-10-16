package runit

import (
	"runtime"
	"testing"

	. "github.com/anthonybishopric/gotcha"
	"github.com/platypus-platform/pp/pkg/util"
)

func TestRunitServicesCanBeStarted(t *testing.T) {
	sv := SV{util.From(runtime.Caller(0)).ExpandPath("fake_sv")}
	service := &Service{"/var/service/foo", "foo"}
	out, err := sv.Start(service)
	Assert(t).IsNil(err, "There should not have been an error starting the service")
	Assert(t).AreEqual(out, "start /var/service/foo\n", "Did not start service with correct arguments")
}
