package db

import (
	// "fmt"
	"os"
	"testing"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// fmt.Println(os.Getwd())
	os.Chdir("../../../")
	c, err := config.LoadConfig()
	assert.NoError(t, err)
	db, err := Init(c.DBUrl)
	assert.NotEqual(t, db, nil, "db should have some value")
	assert.NoError(t, err)
}
