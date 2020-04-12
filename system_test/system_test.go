package system

import (
	"testing"

	"github.com/caevv/fare-estimation/service"
)

func TestService(t *testing.T) {
	service.Start("paths.csv")
}
