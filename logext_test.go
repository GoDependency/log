package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {

	SetOutputLevel(Ldebug)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Warnf("Warn: foo\n")
	Warn("Warn: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")

	SetOutputLevel(Linfo)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Warnf("Warn: foo\n")
	Warn("Warn: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")
}

func TestRotateLog(t *testing.T) {
	logger, err := New("/tmp/test.log", "", Ldefault)
	if err != nil {
		t.FailNow()
	}

	logger.SetOutputLevel(Ldebug)
	logger.EnableRotate()
	Std = logger

	ticker := time.Tick(5 * time.Second)

	for {

		<-ticker
		Debugf("this is test log %s", "ray")
	}

}
