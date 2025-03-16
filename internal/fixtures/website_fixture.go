package fixtures

import (
	"JobFetcher/internal/domain"
	"log"

	"gorm.io/gorm"
)

// LoadWebsiteFixture ajoute des données de test pour les sites web
func LoadWebsiteFixture(db *gorm.DB) error {
	// Vérifier si l'entrée existe déjà
	var count int64
	db.Model(&domain.Website{}).Where("name = ?", "VIE FULL STACK").Count(&count)
	if count > 0 {
		log.Println("La fixture 'VIE FULL STACK' existe déjà.")
		return nil
	}

	// Création de la fixture
	website := domain.Website{
		Name:   "VIE FULL STACK",
		URL:    "https://civiweb-api-prd.azurewebsites.net/api/Offers/search",
		Source: "VIE",
		Method: "POST",
		Body:   `{"limit":1000,"skip":0,"query":"FULL STACK","activitySectorId":[],"missionsTypesIds":[],"missionsDurations":[],"gerographicZones":[],"countriesIds":[],"studiesLevelId":[],"companiesSizes":[],"specializationsIds":[],"entreprisesIds":[0],"missionStartDate":null}`,
	}

	// Enregistrement en base
	if err := db.Create(&website).Error; err != nil {
		return err
	}

	log.Println("Fixture 'VIE FULL STACK' ajoutée avec succès 🚀")
	return nil
}
