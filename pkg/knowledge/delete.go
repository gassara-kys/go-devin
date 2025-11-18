package knowledge

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func (s *Service) Delete(ctx context.Context, noteID string) error {
	if noteID == "" {
		return fmt.Errorf("noteID is required")
	}
	path := fmt.Sprintf("/knowledge/%s", url.PathEscape(noteID))
	return s.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}
