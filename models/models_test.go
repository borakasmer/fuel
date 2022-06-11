package models

import "testing"

// ugly test, but it does its job.
func TestCityGenerateURL(t *testing.T) {
	ist := City{Name: "İstanbul"}
	kahramanmaras := City{Name: "Kahramanmaraş"}
	bartin := City{Name: "Bartın"}
	random := City{Name: "ğşöİİııÜüç"}

	istUrlWant := "https://www.petrolofisi.com.tr/akaryakit-fiyatlari/istanbul-akaryakit-fiyatlari"
	kahramanmarasUrlWant := "https://www.petrolofisi.com.tr/akaryakit-fiyatlari/kahramanmaras-akaryakit-fiyatlari"
	bartinUrlWant := "https://www.petrolofisi.com.tr/akaryakit-fiyatlari/bartin-akaryakit-fiyatlari"
	randomUrlWant := "https://www.petrolofisi.com.tr/akaryakit-fiyatlari/gsoiiiiuuc-akaryakit-fiyatlari"

	ist.setURL()
	kahramanmaras.setURL()
	bartin.setURL()
	random.setURL()

	if ist.Url != istUrlWant ||
		kahramanmaras.Url != kahramanmarasUrlWant ||
		bartin.Url != bartinUrlWant ||
		random.Url != randomUrlWant {
		t.Errorf("wrong urls: %s\n%s\n%s\n%s\n", ist.Url, kahramanmaras.Url, bartin.Url, random.Url)
	}
}
