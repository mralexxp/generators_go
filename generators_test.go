package generators

import (
	"fmt"
	"math/rand"
	"testing"
)

type ShortTest struct {
	name    string
	length  int
	wantErr bool // Ожидаем ли ошибку
}

func TestShortURL(t *testing.T) {
	tests := make([]ShortTest, 0)
	for i := 0; i < 100; i++ {
		// Максимальная длина генерируемой строки
		length := rand.Intn(30)
		tests = append(tests, ShortTest{
			name:    fmt.Sprintf("TESTCASE #%v", i+1),
			length:  length,
			wantErr: false,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShortURL(tt.length)
			// t.Logf("======%v======\nGOT: %v\nLenGOT: %v\nLength: %v\n================", tt.name, got, len(got), tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := tt.length
			if len(got) != want {
				t.Errorf("ShortURL() = %v, want %v", len(got), want)
			}

			// Тестирование на "0"-символы
			for _, char := range got {
				if char == 0 {
					t.FailNow()
				}
			}

		})
	}
}
