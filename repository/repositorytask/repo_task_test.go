package repositorytask

import (
	"p03_kanban_board/config/configdb"
	"p03_kanban_board/entity"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

const ENV_TEST_PATH = "../../.env.test"

type RepoUserTestSuite struct {
	suite.Suite
	db             *gorm.DB
	repo           RepositoryTask
	defaultPayload entity.Task
}

func TestRepositoryTask(t *testing.T) {
	err := godotenv.Load(ENV_TEST_PATH)
	assert.NoError(t, err)
	//init db
	db, err := configdb.New()
	assert.NoError(t, err)

	db.AutoMigrate(entity.Task{})
	clearDbUser(db)

	repo := New(db)

	defaultPayload := entity.Task{
		Title:       "MANDI BARENG",
		Description: "MANDI BARENG BERSAMA SAMA",
		Status:      false,
		UserID:      2,
		CategoryID:  1,
	}

	testSuite := &RepoUserTestSuite{
		db:             db,
		repo:           repo,
		defaultPayload: defaultPayload,
	}
	suite.Run(t, testSuite)
}

func (suite *RepoUserTestSuite) Test_A_CreateTask() {
	suite.T().Run("Create Task Success", func(t *testing.T) {
		created, err := suite.repo.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotEmpty(t, created)
		suite.defaultPayload.ID = created.ID
	})
}

func clearDbUser(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(entity.User{})
}
