package secrets

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Delete removes a secret by ID.
func (s *Service) Delete(ctx context.Context, secretID string) error {
	if secretID == "" {
		return fmt.Errorf("secretID is required")
	}
	path := fmt.Sprintf("/secrets/%s", url.PathEscape(secretID))
	return s.doJSON(ctx, http.MethodDelete, path, nil, nil, nil)
}
