package license

import "log"

var data = []string{
	"CERTIFIED PUBLIC ACCOUNTANT",
	"AGRICULTURAL ENGINEER",
	"ARCHITECT",
	"INTERIOR DESIGNER",
	"LANDSCAPE ARCHITECT",
	"CHEMICAL ENGINEER",
	"CHEMIST",
	"CIVIL ENGINEER",
	"CUSTOMS BROKER",
	"DENTIST",
	"REAL ESTATE BROKER",
	"INTERIOR DESIGNER",
	"NUTRITIONIST - DIETITIAN",
	"PROFESSIONAL ELECTRICAL ENGINEER",
	"REGISTERED ELECTRICAL ENGINEER",
	"MASTER ELECTRICIAN",
	"ELECTRONICS AND COMMUNICATIONS ENGINEER",
	"FORESTER",
	"GEODETIC ENGINEER",
	"JUNIOR GEODETIC ENGINEER",
	"GEOLOGIST",
	"MASTER MARINER",
	"CHIEF MATE",
	"SECOND MATE",
	"THIRD MATE",
	"MAJOR PATRON",
	"MINOR PATRON",
	"CHIEF MARINE ENGINEER",
	"SECOND MARINE ENGINEER",
	"THIRD MARINE ENGINEER",
	"FOURTH MARINE ENGINEER",
	"MOTOR ENGINEER",
	"METALLURGICAL ENGINEER",
	"MASTER PLUMBER",
	"MECHANICAL ENGINEER",
	"CERTIFIED PLANT MECHANIC",
	"MEDICAL TECHNOLOGIST",
	"MEDICAL LABORATORY TECHNICIAN",
	"PHYSICIAN",
	"MIDWIFE",
	"MINING ENGINEER",
	"NAVAL ARCHITECT",
	"NURSE",
	"OPTOMETRIST",
	"PHARMACIST",
	"PHYSICAL THERAPIST",
	"OCCUPATIONAL THERAPIST",
	"SANITARY ENGINEER",
	"SUGAR TECHNOLOGIST",
	"SOCIAL WORKER",
	"VETERINARIAN",
	"AERONAUTICAL ENGINEER",
	"CRIMINOLOGIST",
	"ENVIRONMENTAL PLANNER",
	"RADIOLOGIC TECHNOLOGIST",
	"X - RAY TECHNOLOGIST",
	"LIBRARIAN",
	"PROFESSIONAL TEACHER",
	"DRIVER - CODE 1 (MOTORCYCLES/MOTORIZED TRICYCLES)",
	"DRIVER - CODE 2 (VEHICLE UP TO 4500 KGS G V W )",
	"DRIVER - CODE 3 (VEHICLE ABOVE 4500 KGS G V W )",
	"DRIVER - CODE 4 (VEHICLE AUTOMATIC CLUTCH UP TO 4500 G V W )",
	"DRIVER - CODE 5 (VEHICLE AUTOMATIC CLUTCH ABOVE 4500 G V W )",
	"DRIVER - CODE 6 (ARTICULATED VEHICLE ABOVE 1600 KGS G V W AND BELOW)",
	"DRIVER - CODE 7 (ARTICULATED VEHICLE 1601 UP TO 4500 KGS G V W )",
	"DRIVER - CODE 8 (ARTICULATED VEHICLE 4501 KGS AND ABOVE G V W)",
	"DRIVER - CODE 9 (DISABLED)",
	"LAWYER",
	"AIR TRAFFIC CONTROL LICENSE",
	"AIR TRANSPORT RATING/FLIGHT",
	"AIR - CONDITIONING/REFRIGERATION",
	"AIRCRAFT MECHANIC LICENSE",
	"AIRFRAME AND POWERPLANT MECHANIC LICENSE",
	"AIRWAYS COMMUNICATOR LICENSE",
	"AUTO REPAIR FOREMAN",
	"AUTOMOTIVE MECHANIC (HEAVY DUTY)",
	"AUTOMOTIVE MECHANIC (LIGHT DUTY)",
	"CONDUCTOR/CONDUCTRESS",
	"DOCKMAN",
	"ELECTRIC POWER LINEMAN",
	"ELECTRICAL EQUIPMENT OPERATOR",
	"ELECTRICIAN (AUTOMOTIVE)",
	"ELECTRICIAN (BUILDING WIRING)",
	"ELECTRONICS EQUIPMENT OPERATOR",
	"HEAVY EQUIPMENT OPERATOR",
	"JUSTICE OF THE COURT OF APPEALS",
	"JUSTICE OF THE SANDIGANBAYAN",
	"MC # 11, S. 1996 (PRESSMAN)",
	"MC # 11, S. 1996 (TYPESETTER)",
	"MACHINIST (2ND CLASS)",
	"PILOT LICENSE",
	"PRINTING MACHINE OPERATOR",
	"SECURITY GUARD LICENSE",
	"TCB RADIO LICENSE",
	"TRAIN DRIVER",
	"RADIOLOGY TECHNICIAN",
	"AIRCRAFT SPECIALIST",
	"AIRLINE TRASPORT PILOT",
	"COMMERCIAL HELICOPTER PILOT",
	"COMMERCIAL PILOT",
	"FLIGHT DISPATCHER",
	"FLIGHT ENGINEER",
	"FLIGHT INSTRUCTOR",
	"GROUND INSTRUCTOR",
	"PRIVATE HELICOPTER PILOT",
	"PRIVATE PILOT",
	"INSURANCE COMMISSION LICENSE",
	"RADIO ELECTRONIC OFFICER (NTC)",
	"FIRST CLASS RADIO TELEGRAPH OPERATOR (NTC)",
	"FIRST CLASS RADIO TELEPHONE OPERATOR (NTC)",
	"FIRST CLASS RADIO TELEPHONE/TELEGRAPH OPERATOR",
	"INSURANCE ADVISER",
	"MARINE ENGINEER",
}

func Seeder() {
	licns, err := All()

	if err != nil {
		panic(err)
	}

	if len(licns) == 0 {
		for _, value := range data {
			var licn License
			licn.Name = value
			_, err := licn.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("License seeded")
	}
}