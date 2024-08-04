package actions_test

import (
	"link_shortener/actions"
	"link_shortener/links"
	"link_shortener/mocks"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

var ctrl *gomock.Controller
var mockAnalytics *mocks.MockAnalyticsEngine
var mockDA *mocks.MockLinkDA

var testLink = links.Link{
	Id:          "qiI5wXYJsh4=",
	ExternalURL: "http://www.google.com",
}

func setup(t *testing.T) {
	ctrl = gomock.NewController(t)
	mockAnalytics = mocks.NewMockAnalyticsEngine(ctrl)
	mockDA = mocks.NewMockLinkDA(ctrl)

	actions.InitLinkctions(mockDA, mockAnalytics)
}

func TestCreateSuccess(t *testing.T) {
	setup(t)

	mockDA.EXPECT().
		Save(gomock.Eq(testLink)).
		Return(&testLink, nil)

	mockAnalytics.EXPECT().
		Created(testLink)

	createdLink, err := actions.Create(testLink)

	assert.Equal(t, *createdLink, testLink)
	assert.NilError(t, err)
}

func TestCreateDuplicate(t *testing.T) {
	setup(t)

	mockDA.EXPECT().
		Save(gomock.Eq(testLink)).
		Return(nil, links.ErrUrlDuplicate{})

	mockDA.EXPECT().
		GetByUrl(gomock.Eq(testLink.ExternalURL)).
		Return(&testLink, nil)

	createdLink, err := actions.Create(testLink)

	assert.Equal(t, *createdLink, testLink)
	assert.NilError(t, err)
}

func TestCreateError(t *testing.T) {
	setup(t)

	mockDA.EXPECT().
		Save(gomock.Eq(testLink)).
		Return(nil, links.ErrRepositoryError{})

	createdLink, err := actions.Create(testLink)

	assert.Equal(t, createdLink, (*links.Link)(nil))
	assert.ErrorType(t, err, links.ErrRepositoryError{})
}

func TestGetSuccess(t *testing.T) {
	setup(t)

	mockDA.EXPECT().
		Get(testLink.Id).
		Return(&testLink, nil)

	mockAnalytics.EXPECT().Accessed(testLink, 1)

	returnedLink, err := actions.Get(testLink.Id)

	assert.Equal(t, *returnedLink, testLink)
	assert.NilError(t, err)
}

func TestGetError(t *testing.T) {
	setup(t)

	mockDA.EXPECT().
		Get(testLink.Id).
		Return(nil, links.ErrRepositoryError{})

	returnedLink, err := actions.Get(testLink.Id)

	assert.Equal(t, returnedLink, (*links.Link)(nil))
	assert.ErrorType(t, err, links.ErrRepositoryError{})
}

func TestValidateSuccess(t *testing.T) {
	setup(t)
	assert.Equal(t, actions.Validate(&testLink), nil)
}

func TestValidatePartialUrls(t *testing.T) {
	setup(t)
	invalidLink := testLink
	invalidLink.ExternalURL = "google.com"

	_, ok := actions.Validate(&invalidLink).(*url.Error)
	assert.Equal(t, ok, true)
}
