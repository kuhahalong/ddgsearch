package ddgsearch

import (
	"testing"
)

func TestExtractVQD(t *testing.T) {
	tests := []struct {
		name     string
		html     []byte
		keywords string
		want     string
		wantErr  bool
	}{
		{
			name:     "valid vqd double quotes",
			html:     []byte(`<script>vqd="123456";</script>`),
			keywords: "test",
			want:     "123456",
			wantErr:  false,
		},
		{
			name:     "valid vqd single quotes",
			html:     []byte(`<script>vqd='789012';</script>`),
			keywords: "test",
			want:     "789012",
			wantErr:  false,
		},
		{
			name:     "valid vqd with ampersand",
			html:     []byte(`<script>vqd=345678&other=value</script>`),
			keywords: "test",
			want:     "345678",
			wantErr:  false,
		},
		{
			name:     "no vqd",
			html:     []byte(`<script>no vqd here</script>`),
			keywords: "test",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVQD(tt.html, tt.keywords)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVQD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractVQD() = %v, want %v", got, tt.want)
			}
		})
	}
}
