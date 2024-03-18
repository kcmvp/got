package dal

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type Base struct {
	CreatedAt time.Time `ds:"autoUpdateTime"`
	CreatedBy string    `ds:"createdBy"`
	UpdatedAt time.Time `ds:"autoCreateTime"`
	UpdatedBy string    `ds:"updatedBy"`
}

type Product struct {
	Base
	Id          string `ds:"pk;"`
	Name        string `ds:"column=name"`
	FullName    string `ds:"ignore"`
	Grade       sql.NullInt32
	Address     sql.NullString
	ProductDate time.Time
	comment     string
}

func (p Product) Table() string {
	return "product"
}

type BuilderTestSuit struct {
	//builder *SqlBuilder
	suite.Suite
}

func TestBuilderSuite(t *testing.T) {
	suite.Run(t, &BuilderTestSuit{})
}

func (suite *BuilderTestSuit) SetupSuite() {
	//suite.builder = do.MustInvoke[*SqlBuilder](boot.Container())
}

//
//func (suite *BuilderTestSuit) TestBuild() {
//	e := Product{}
//	suite.builder.parse(Product{})
//	mapping := suite.builder.getMapper(fmt.Sprintf("%s_%s", typetostring.GetValueType(e), e.Table()))
//	assert.Equal(suite.T(), len(mapping), 7)
//}
