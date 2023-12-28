package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"radiophysiker/link_shorter/internal/config"
	"radiophysiker/link_shorter/internal/usecases/mocks"
)

func TestURLUseCase_CreateShortURL(t *testing.T) {
	type args struct {
		fullURL string
	}

	mocksRepoURL := mocks.NewURLFileRepository(t)
	mocksRepoURL.
		On("Save", mock.AnythingOfType("entity.URL")).Return(nil)
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				fullURL: "https://yandex.ru",
			},
			want:    "short_url",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := URLUseCase{
				urlRepository: mocksRepoURL,
				config: &config.Config{
					BaseURL:         "http://localhost:8080",
					ServerPort:      "localhost:8080",
					FileStoragePath: "tmp/url.json",
				},
			}
			got, err := us.CreateShortURL(tt.args.fullURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEmpty(t, got)
		})
	}
}
