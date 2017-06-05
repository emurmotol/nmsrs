package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type License struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func LicenseSeeder() {
	data := []string{
		"AERONAUTICAL ENGINEER",
		"AGRICULTURAL ENGINEER",
		"AIR - CONDITIONING/REFRIGERATION",
		"AIR TRAFFIC CONTROL LICENSE",
		"AIR TRANSPORT RATING/FLIGHT",
		"AIRCRAFT MECHANIC LICENSE",
		"AIRCRAFT SPECIALIST",
		"AIRFRAME AND POWERPLANT MECHANIC LICENSE",
		"AIRLINE TRASPORT PILOT",
		"AIRWAYS COMMUNICATOR LICENSE",
		"ARCHITECT",
		"AUTO REPAIR FOREMAN",
		"AUTOMOTIVE MECHANIC (HEAVY DUTY)",
		"AUTOMOTIVE MECHANIC (LIGHT DUTY)",
		"CERTIFIED PLANT MECHANIC",
		"CERTIFIED PUBLIC ACCOUNTANT",
		"CHEMICAL ENGINEER",
		"CHEMIST",
		"CHIEF MARINE ENGINEER",
		"CHIEF MATE",
		"CIVIL ENGINEER",
		"COMMERCIAL HELICOPTER PILOT",
		"COMMERCIAL PILOT",
		"CONDUCTOR/CONDUCTRESS",
		"CRIMINOLOGIST",
		"CUSTOMS BROKER",
		"DENTIST",
		"DOCKMAN",
		"DRIVER - CODE 1 (MOTORCYCLES/MOTORIZED TRICYCLES)",
		"DRIVER - CODE 2 (VEHICLE UP TO 4500 KGS G V W )",
		"DRIVER - CODE 3 (VEHICLE ABOVE 4500 KGS G V W )",
		"DRIVER - CODE 4 (VEHICLE AUTOMATIC CLUTCH UP TO 4500 G V W )",
		"DRIVER - CODE 5 (VEHICLE AUTOMATIC CLUTCH ABOVE 4500 G V W )",
		"DRIVER - CODE 6 (ARTICULATED VEHICLE ABOVE 1600 KGS G V W AND BELOW)",
		"DRIVER - CODE 7 (ARTICULATED VEHICLE 1601 UP TO 4500 KGS G V W )",
		"DRIVER - CODE 8 (ARTICULATED VEHICLE 4501 KGS AND ABOVE G V W)",
		"DRIVER - CODE 9 (DISABLED)",
		"ELECTRIC POWER LINEMAN",
		"ELECTRICAL EQUIPMENT OPERATOR",
		"ELECTRICIAN (AUTOMOTIVE)",
		"ELECTRICIAN (BUILDING WIRING)",
		"ELECTRONICS AND COMMUNICATIONS ENGINEER",
		"ELECTRONICS EQUIPMENT OPERATOR",
		"ENVIRONMENTAL PLANNER",
		"FIRST CLASS RADIO TELEGRAPH OPERATOR (NTC)",
		"FIRST CLASS RADIO TELEPHONE OPERATOR (NTC)",
		"FIRST CLASS RADIO TELEPHONE/TELEGRAPH OPERATOR",
		"FLIGHT DISPATCHER",
		"FLIGHT ENGINEER",
		"FLIGHT INSTRUCTOR",
		"FORESTER",
		"FOURTH MARINE ENGINEER",
		"GEODETIC ENGINEER",
		"GEOLOGIST",
		"GROUND INSTRUCTOR",
		"HEAVY EQUIPMENT OPERATOR",
		"INSURANCE ADVISER",
		"INSURANCE COMMISSION LICENSE",
		"INTERIOR DESIGNER",
		"JUNIOR GEODETIC ENGINEER",
		"JUSTICE OF THE COURT OF APPEALS",
		"JUSTICE OF THE SANDIGANBAYAN",
		"LANDSCAPE ARCHITECT",
		"LAWYER",
		"LIBRARIAN",
		"MACHINIST (2ND CLASS)",
		"MAJOR PATRON",
		"MARINE ENGINEER",
		"MASTER ELECTRICIAN",
		"MASTER MARINER",
		"MASTER PLUMBER",
		"MC # 11, S. 1996 (PRESSMAN)",
		"MC # 11, S. 1996 (TYPESETTER)",
		"MECHANICAL ENGINEER",
		"MEDICAL LABORATORY TECHNICIAN",
		"MEDICAL TECHNOLOGIST",
		"METALLURGICAL ENGINEER",
		"MIDWIFE",
		"MINING ENGINEER",
		"MINOR PATRON",
		"MOTOR ENGINEER",
		"NAVAL ARCHITECT",
		"NURSE",
		"NUTRITIONIST - DIETITIAN",
		"OCCUPATIONAL THERAPIST",
		"OPTOMETRIST",
		"PHARMACIST",
		"PHYSICAL THERAPIST",
		"PHYSICIAN",
		"PILOT LICENSE",
		"PRINTING MACHINE OPERATOR",
		"PRIVATE HELICOPTER PILOT",
		"PRIVATE PILOT",
		"PROFESSIONAL ELECTRICAL ENGINEER",
		"PROFESSIONAL TEACHER",
		"RADIO ELECTRONIC OFFICER (NTC)",
		"RADIOLOGIC TECHNOLOGIST",
		"RADIOLOGY TECHNICIAN",
		"REAL ESTATE BROKER",
		"REGISTERED ELECTRICAL ENGINEER",
		"SANITARY ENGINEER",
		"SECOND MARINE ENGINEER",
		"SECOND MATE",
		"SECURITY GUARD LICENSE",
		"SOCIAL WORKER",
		"SUGAR TECHNOLOGIST",
		"TCB RADIO LICENSE",
		"THIRD MARINE ENGINEER",
		"THIRD MATE",
		"TRAIN DRIVER",
		"VETERINARIAN",
		"X - RAY TECHNOLOGIST",
	}

	for _, name := range data {
		license := License{Name: strings.ToUpper(name)}

		if _, err := license.Create(); err != nil {
			panic(err)
		}
	}
}

func (license *License) Create() (*License, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&license).Error; err != nil {
		return nil, err
	}
	return license, nil
}

func (license License) Search(q string) []License {
	db := database.Conn()
	defer db.Close()

	licenses := []License{}
	results := make(chan []License)

	go func() {
		db.Find(&licenses, "name LIKE ?", database.WrapLike(q))
		results <- licenses
	}()
	return <-results
}
