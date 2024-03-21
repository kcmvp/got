package got

import (
	"database/sql"
	"fmt"
	"github.com/kcmvp/got/boot"
	"github.com/kcmvp/got/internal"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samber/do/v2"
	typetostring "github.com/samber/go-type-to-string"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type BootTestSuite struct {
	suite.Suite
}

func TestBootTestSuite(t *testing.T) {
	suite.Run(t, &BootTestSuite{})
}

func (b *BootTestSuite) SetupSuite() {
	boot.InitAppWith("test")
}

func (b *BootTestSuite) TestInitialization() {
	assert.True(b.T(), strings.Contains(boot.RootDir(), "got"))
	assert.NotNil(b.T(), boot.Container())
	db := do.MustInvokeNamed[*sql.DB](internal.Container, fmt.Sprintf("%s_%s", boot.DefaultDS, typetostring.GetType[*sql.DB]()))
	assert.NotNil(b.T(), db)
	rs, err := db.Exec("select * from Product")
	assert.NoError(b.T(), err)
	cnt, _ := rs.RowsAffected()
	assert.Equal(b.T(), int64(2), cnt)
}
