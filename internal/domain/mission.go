package domain

type Mission struct {
	ID                int
	WebsiteId         int
	WebsiteSource     string
	MissionTitle      string
	MissionPostedDate string
	OrganizationName  string
	CountryName       string
	CityName          string
	MissionDuration   int
	MissionStartDate  string
	ViewCounter       int
	CandidateCounter  int
}
