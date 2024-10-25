package entities_test

import (
	"scoreplay/internal/business/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewMediaEntityWithOptions(t *testing.T) {
	assert := assert.New(t)
	id, err := uuid.NewRandom()
	assert.NoError(err)

	picture := "some image name"
	name := "some media name"

	footballTag, err := entities.NewTagEntity("football", entities.WithTagID(42))
	assert.NoError(err)

	frenchTeamTag, err := entities.NewTagEntity("french-team", entities.WithTagID(43))
	assert.NoError(err)

	tags := []*entities.TagEntity{footballTag, frenchTeamTag}

	mediaEntity, err := entities.NewMediaEntity(
		picture,
		name,
		tags,
		entities.WithMediaID(id),
	)

	assert.NoError(err)
	assert.Equal(id, mediaEntity.ID())
	assert.Equal(picture, mediaEntity.Picture())
	assert.Equal(name, mediaEntity.Name())
	assert.Equal(tags, mediaEntity.Tags())
}

func TestNewMediaEntityValidation(t *testing.T) {
	assert := assert.New(t)

	picture := "some picture name"
	name := "some media name"

	footballTag, err := entities.NewTagEntity("football", entities.WithTagID(42))
	assert.NoError(err)

	frenchTeamTag, err := entities.NewTagEntity("french-team", entities.WithTagID(43))
	assert.NoError(err)

	tags := []*entities.TagEntity{footballTag, frenchTeamTag}

	type testCase struct {
		sut           func() (*entities.MediaEntity, error)
		expectedError error
	}

	testCases := map[string]testCase{
		"empty picture data": {
			sut: func() (*entities.MediaEntity, error) {
				return entities.NewMediaEntity("", name, tags)
			},
			expectedError: entities.ErrEmptyPicture,
		},
		"empty media name": {
			sut: func() (*entities.MediaEntity, error) {
				return entities.NewMediaEntity(picture, "", tags)
			},
			expectedError: entities.ErrEmptyMediaName,
		},
		"empty tag list": {
			sut: func() (*entities.MediaEntity, error) {
				return entities.NewMediaEntity(picture, name, []*entities.TagEntity{})
			},
			expectedError: entities.ErrEmptyTagList,
		},
		"invalid tag": {
			sut: func() (*entities.MediaEntity, error) {
				return entities.NewMediaEntity(picture, name, []*entities.TagEntity{
					footballTag,
					{},
				})
			},
			expectedError: entities.ErrInvalidTag,
		},
	}

	for testName, test := range testCases {
		t.Run(testName, func(t *testing.T) {
			_, err := test.sut()
			assert.ErrorIs(err, test.expectedError)
		})
	}

}
