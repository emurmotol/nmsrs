package model

import "github.com/emurmotol/nmsrs/database"
import "strings"

type Certificate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CertificateSeeder() {
	data := []string{
		"(NEW) AUTOMOTIVE SERVICE TECHNICIAN (LIGHT DUTY) (COMPETENCY LEADING TO NL)",
		"(NEW) AUTOMOTIVE SERVICE TECHNICIAN (LIGHT DUTY) (NATIONAL CERTIFICATE I)",
		"(NEW) AUTOMOTIVE SERVICE TECHNICIAN (LIGHT DUTY) (NATIONAL CERTIFICATE II)",
		"(NEW) AUTOMOTIVE SERVICE TECHNICIAN (LIGHT DUTY) (NATIONAL CERTIFICATE III)",
		"(OLD OSS ELECTRIC ARC WELDER (NATIONAL CERTIFICATE I)",
		"(OLD OSS ELECTRIC ARC WELDER (NATIONAL CERTIFICATE II)",
		"(OLD OSS ELECTRIC ARC WELDER (NATIONAL CERTIFICATE III)",
		"(OLD OSS) BLACKSMITH (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) CADD OPERATOR (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) CADD OPERATOR (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) CADD OPERATOR (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) CALCULATOR TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) CNC MILLING MACHINE OPERATOR (COMPETENCY LEADINGTO NL)",
		"(OLD OSS) COREMAKER (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) DIE MAKER (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) DIE MAKER (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) DIE MAKER (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) DIE MAKER FORGING (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) DIE MAKER FORGING (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) DRAFTSMAN MECHANICAL (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) EDM (DIE SINKER) OPERATOR (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) ELECTRIC FURNACE MELTER (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) ELECTRIC FURNACE MELTER (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) ELECTRIC FURNACE MELTER (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) FOUNDRY PATTERN MAKER (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) FOUNDRY PATTERN MAKER (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) GRINDING MACHINE OPERATOR FOUNDRY (COMPETENCY LEADINGTO NC I)",
		"(OLD OSS) GTAW (TIG) WELDER (NATIONAL CERTIFICATE I)",
		"(OLD OSS) GTAW (TIG) WELDER (NATIONAL CERTIFICATE II)",
		"(OLD OSS) GTAW (TIG) WELDER (NATIONAL CERTIFICATE III)",
		"(OLD OSS) HEAT TREATMENT TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) HEAT TREATMENT TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) HEAT TREATMENT TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) LATHE - MACHINE SETTER OPERATOR (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) LATHE - MACHINE SETTER OPERATOR (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) LATHE - MACHINE SETTER OPERATOR (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) MACHINE TOOL MECHANIC (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) MACHINE TOOL MECHANIC (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) MACHINE TOOL MECHANIC (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) MACHINIST GENERAL (NATIONAL CERTIFICATE I)",
		"(OLD OSS) MACHINIST GENERAL (NATIONAL CERTIFICATE II)",
		"(OLD OSS) MACHINIST GENERAL (NATIONAL CERTIFICATE III)",
		"(OLD OSS) MECHANICAL ENGINEERING TECHNICIAN (MOTORS AND ENGINES) (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) METALLUGICAL LABORATORY TECHNICIAN (COMPETENCY LEADINGTO NC III)",
		"(OLD OSS) MILLING MACHINE SETTER OPERATOR (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) MILLING MACHINE SETTER OPERATOR (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) MILLING MACHINE SETTER OPERATOR (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) MOLD MAKER, PERMANENT (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) MOLDER, MANUAL (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) MOLDER, MANUAL (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) PLANT EQUIPMENT MECHANIC (MILLWRIGHT) (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) PLANT EQUIPMENT MECHANIC (MILLWRIGHT) (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) PLANT EQUIPMENT MECHANIC (MILLWRIGHT) (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) PLASTIC MOLD MAKER (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) PLASTIC MOLD MAKER (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) POLISHING LATHE OPERATOR (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) PROGRAMMER/ANALYST (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) SURFACE FINISHING OPERATOR (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) SURFACE FINISHING TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) SYSTEM ANALYST (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) TOOL AND DIE MAKER (COMPETENCY LEADING TO NC I)",
		"(OLD OSS) TOOL AND DIE MAKER (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) TOOL AND DIE MAKER (COMPETENCY LEADING TO NC III)",
		"(OLD OSS) TOOL MAKER (COMPETENCY LEADING TO NC II)",
		"(OLD OSS) TOOL MAKER (COMPETENCY LEADING TO NC III)",
		"AIRDUCT WORKER (COMPETENCY LEADING TO NC I)",
		"AIRDUCT WORKER (COMPETENCY LEADING TO NC II)",
		"AIRDUCT WORKER (COMPETENCY LEADING TO NC III)",
		"ALUMINUM WORKER (COMPETENCY LEADING TO NC II)",
		"ARCHITECTURAL DRAFTSMAN (COMPETENCY LEADING TO NC II)",
		"AUDIO ELECTRONICS SERVICE TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"AUTO BODY REPAIRMAN (COMPETENCY LEADING TO NC I)",
		"AUTO BODY REPAIRMAN (COMPETENCY LEADING TO NC II)",
		"AUTO PAINTER (COMPETENCY LEADING TO NC I)",
		"AUTO - BUS AIRCON MECHANIC (NATIONAL CERTIFICATE II)",
		"AUTO - BUS AIRCON MECHANIC (NATIONAL CERTIFICATE III)",
		"AUTOMOTIVE AIR - CON TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"AUTOMOTIVE AIR - CON TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"AUTOMOTIVE AIR - CON TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"AUTOMOTIVE BODY REPAIRER (NATIONAL CERTIFICATE II)",
		"AUTOMOTIVE BODY REPAIRER (NATIONAL CERTIFICATE III)",
		"AUTOMOTIVE ELECTRICIAN (NATIONAL CERTIFICATE II)",
		"AUTOMOTIVE ELECTRICIAN (NATIONAL CERTIFICATE III)",
		"AUTOMOTIVE ELECTRICIAN (NEW) (COMPETENCY LEADING TO NC I)",
		"AUTOMOTIVE ELECTRICIAN (NEW) (COMPETENCY LEADING TO NC II)",
		"AUTOMOTIVE ELECTRICIAN (NEW) (COMPETENCY LEADING TO NC III)",
		"AUTOMOTIVE MECHANIC (HEAVY DUTY) (NATIONAL CERTIFICATE I)",
		"AUTOMOTIVE MECHANIC (HEAVY DUTY) (NATIONAL CERTIFICATE II)",
		"AUTOMOTIVE MECHANIC (HEAVY DUTY) (NATIONAL CERTIFICATE III)",
		"AUTOMOTIVE MECHANIC (LIGHT DUTY) (NATIONAL CERTIFICATE I)",
		"AUTOMOTIVE MECHANIC (LIGHT DUTY) (NATIONAL CERTIFICATE II)",
		"AUTOMOTIVE MECHANIC (LIGHT DUTY) (NATIONAL CERTIFICATE III)",
		"AUTOMOTIVE PAINTER (COMPETENCY LEADING TO NC II)",
		"AUTOMOTIVE SERVICE MECHANIC (HEAVY DUTY) (NATIONAL CERTIFICATE I)",
		"AUTOMOTIVE SERVICE MECHANIC (HEAVY DUTY) (NATIONAL CERTIFICATE II)",
		"AUTOMOTIVE SERVICE MECHANIC (HEAVY DUTY) (NATIONAL CERTIFICATE III)",
		"AUTOMOTIVE SERVICE TECHNICIAN (HEAVY - DUTY) NEW (COMPETENCY LEADING TO NC I)",
		"AUTOMOTIVE SERVICE TECHNICIAN (HEAVY - DUTY) NEW (COMPETENCY LEADING TO NC II)",
		"AUTOMOTIVE SERVICE TECHNICIAN (HEAVY - DUTY) NEW (COMPETENCY LEADING TO NC III)",
		"AUTOMOTIVE SERVICE TECHNICIAN (HEAVY - DUTY) NEW (COMPETENCY LEADING TO NL)",
		"BABYSITTER (COMPETENCY LEADING TO NC II)",
		"BAKER (COMPETENCY LEADING TO NC I)",
		"BAKER (COMPETENCY LEADING TO NC II)",
		"BAMBOO FURNITURE MAKER (COMPETENCY LEADING TO NC II)",
		"BARTENDER (COMPETENCY LEADING TO NC I)",
		"BASKET TRUCK OPERATOR (NATIONAL CERTIFICATE II)",
		"BATTERY SERVICE WORKER (COMPETENCY LEADING TO NC I)",
		"BINDER GENERAL (COMPETENCY LEADING TO NC II)",
		"BINDER GENERAL (COMPETENCY LEADING TO NC III)",
		"BIOGAS CONSTRUCTOR (COMPETENCY LEADING TO NC II)",
		"BIOGAS TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"BOAT CAPTAIN (COMPETENCY LEADING TO NC II)",
		"BOAT ENGINEER (COMPETENCY LEADING TO NC I)",
		"BOATSWAIN (BOSUN) (COMPETENCY LEADING TO NC III)",
		"BOOKBINDER (COMPETENCY LEADING TO NC I)",
		"BOOKKEEPER (COMPETENCY LEADING TO NC I)",
		"BOTTOM MAKER (FOOTWEAR) (COMPETENCY LEADING TO NC II)",
		"BRANCH FIELD REPRESENTATIVE BRANCH FIELDMAN (COMPETENCY LEADING TONC III)",
		"BUILDING MAINTENANCE MAN (COMPETENCY LEADING TO NC II)",
		"BUILDING WIRING ELECTRICIAN (RESIDENTIAL/COMMERCIAL/INDUSTRIAL) ( (COMPETENCY LEADING TO NC I)",
		"BUILDING WIRING ELECTRICIAN (RESIDENTIAL/COMMERCIAL/INDUSTRIAL) ( (COMPETENCY LEADING TO NC II)",
		"BUILDING WIRING ELECTRICIAN (RESIDENTIAL/COMMERCIAL/INDUSTRIAL) ( (COMPETENCY LEADING TO NC III)",
		"BUS BODY BUILDER (NATIONAL CERTIFICATE II)",
		"BUS BODY BUILDER (NATIONAL CERTIFICATE III)",
		"BUS DRIVER (COMPETENCY LEADING TO NC III)",
		"BUTCHER (COMPETENCY LEADING TO NC I)",
		"BUTCHER (COMPETENCY LEADING TO NC II)",
		"BUTCHER (COMPETENCY LEADING TO NC III)",
		"CALCULATOR TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"CAREGIVER (NATIONAL CERTIFICATE I)",
		"CARETAKER (COMPETENCY LEADING TO NC I)",
		"CARPENTER (NATIONAL CERTIFICATE I)",
		"CARPENTER (NATIONAL CERTIFICATE II)",
		"CARPENTER (NATIONAL CERTIFICATE III)",
		"CENTRAL AIR - CONDITIONING MECHANIC (COMPETENCY LEADING TO NC II)",
		"CENTRAL AIR - CONDITIONING MECHANIC (COMPETENCY LEADING TO NC III)",
		"CENTRAL AIR - CONDITIONING TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"CENTRAL AIR - CONDITIONING TECHNICIAN (COMPETENCY LEADING TO NL)",
		"CHEF DE PARTIE (COMPETENCY LEADING TO NC III)",
		"CHICKEN CARETAKER (COMPETENCY LEADING TO NC I)",
		"CHICKEN CARETAKER (COMPETENCY LEADING TO NC II)",
		"CIVIL ENGINEERING TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"CIVIL ENGINEERING TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"CIVIL ENGINEERING TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"COIR FIBER DECORTICATOR OPERATOR (COMPETENCY LEADING TO NC I)",
		"COLD KITCHEN COOK (COMPETENCY LEADING TO NC I)",
		"COLD KITCHEN COOK (COMPETENCY LEADING TO NC III)",
		"COLOR SEPARATION TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"COLOR SEPARATION TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"COMPOSITOR LETTERPRESS (COMPETENCY LEADING TO NC II)",
		"COMPUTER AIDED DESIGN AND DRAFTING (CADD) OPERATOR (COMPETENCY LEADINGTO NC I)",
		"COMPUTER AIDED DESIGN AND DRAFTING (CADD) OPERATOR (COMPETENCY LEADINGTO NC II)",
		"COMPUTER AIDED DESIGN AND DRAFTING (CADD) OPERATOR (COMPETENCY LEADINGTO NC III)",
		"COMPUTER PROGRAMMER (COMPETENCY LEADING TO NC II)",
		"COMPUTER TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"COMPUTER TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"CONCRETE FINISHER (COMPETENCY LEADING TO NC II)",
		"CONCRETE FINISHER, GENERAL (COMPETENCY LEADING TO NC II)",
		"CONSTRUCTION GLAZIER (COMPETENCY LEADING TO NC II)",
		"CONSTRUCTION PAINTER (COMPETENCY LEADING TO NC II)",
		"CONSTRUCTION PAINTER (COMPETENCY LEADING TO NC III)",
		"CONSUMER ELECTRONICS MECHANIC (COMPETENCY LEADING TO NC I)",
		"CONSUMER ELECTRONICS MECHANIC (COMPETENCY LEADING TO NC II)",
		"CONSUMER ELECTRONICS MECHANIC (COMPETENCY LEADING TO NC III)",
		"CORN SHELLER OPERATOR (COMPETENCY LEADING TO NC I)",
		"COSMETOLOGIST (COMPETENCY LEADING TO NC I)",
		"COSMETOLOGIST (COMPETENCY LEADING TO NC III)",
		"COURT STENOGRAPHER (COMPETENCY LEADING TO NC I)",
		"CRAWLER CRANE OPERATOR (COMPETENCY LEADING TO NC II)",
		"CRAWLER CRANE OPERATOR (COMPETENCY LEADING TO NC III)",
		"CRAWLER TRACTOR (BULLDOZER) OPERATOR (COMPETENCY LEADING TO NC I)",
		"CRAWLER TRACTOR (BULLDOZER) OPERATOR (COMPETENCY LEADING TO NC II)",
		"CRAWLER TRACTOR (BULLDOZER) OPERATOR (COMPETENCY LEADING TO NC III)",
		"CRAWLER TRACTOR OPERATOR (COMPETENCY LEADING TO NC II)",
		"CRAWLER TRACTOR OPERATOR (COMPETENCY LEADING TO NC III)",
		"CURED MEAT PROCESSOR (COMPETENCY LEADING TO NC I)",
		"CURED MEAT PROCESSOR (COMPETENCY LEADING TO NC II)",
		"CURED MEAT PROCESSOR (COMPETENCY LEADING TO NC III)",
		"CURING MACHINE OPERATOR (TUBULAR FABRIC) (COMPETENCY LEADING TO NC I)",
		"DAIRY PLANT OPERATOR (COMPETENCY LEADING TO NC I)",
		"DAIRY PLANT OPERATOR (COMPETENCY LEADING TO NC II)",
		"DAIRY PLANT OPERATOR (COMPETENCY LEADING TO NC III)",
		"DATA COMMUNICATION TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"DATA ENCODER (COMPETENCY LEADING TO NC II)",
		"DEEPWELL PUMP OPERATOR (SUBMERSIBLE PUMP) (COMPETENCY LEADING TO NC II)",
		"DEHYDRATED FRUIT AND VEGETABLE PLANT OPERATOR (COMPETENCY LEADING TO NC I)",
		"DESKTOP ENCODER (COMPETENCY LEADING TO NC II)",
		"DESKTOP GRAPHIC ARTIST (COMPETENCY LEADING TO NC I)",
		"DESKTOP GRAPHIC ARTIST (COMPETENCY LEADING TO NC II)",
		"DESKTOP PAGE COMPOSITION OPERATOR (COMPETENCY LEADING TO NC II)",
		"DESKTOP PUBLISHING ENCODER (COMPETENCY LEADING TO NC II)",
		"DESKTOP SCANNER OPERATOR (COMPETENCY LEADING TO NC I)",
		"DESKTOP SCANNER OPERATOR (COMPETENCY LEADING TO NC III)",
		"DIE MAKER",
		"DIE MAKER (NEW) (COMPETENCY LEADING TO NC I)",
		"DIE MAKER (NEW) (COMPETENCY LEADING TO NC II)",
		"DIE MAKER (NEW) (COMPETENCY LEADING TO NC III)",
		"DIECUTTING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"DIESEL ENGINE MECHANIC (NATIONAL CERTIFICATE I)",
		"DIESEL ENGINE MECHANIC (NATIONAL CERTIFICATE II)",
		"DIESEL ENGINE MECHANIC (NATIONAL CERTIFICATE III)",
		"DIESEL FUEL INJECTION CALIBRATION TECHNICIAN (COMPETENCY LEADING TO NL)",
		"DIESEL FUEL INJECTION TECHNICIAN (COMPETENCY LEADING TO NL)",
		"DIESEL POWER PLANT OPERATOR (SMALL POWER PLANT) (COMPETENCY LEADING TO NC I)",
		"DIGITAL LOOP CARRIER TELECOM TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"DIGITAL PAIR GAIN TELECOMMUNICATION TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"DOCUMENTATION OFFICER (COMPETENCY LEADING TO NC I)",
		"DOMESTIC HELPER (COMPETENCY LEADING TO NC II)",
		"DOMESTIC RAC SERVICE TECHNICIAN (NATIONAL CERTIFICATE II)",
		"DRAFSTMAN GENERAL (COMPETENCY LEADING TO NC I)",
		"DRAFSTMAN GENERAL (COMPETENCY LEADING TO NC II)",
		"DRAFSTMAN GENERAL (COMPETENCY LEADING TO NC III)",
		"DRAFTSMAN, MECHANICAL (COMPETENCY LEADING TO NC II)",
		"DRESSMAKER (NATIONAL CERTIFICATE II)",
		"DRESSMAKER (NATIONAL CERTIFICATE III)",
		"DRIVING SCHOOL INSTRUCTOR (LIGHT DUTY) (COMPETENCY LEADING TO NC III)",
		"DUCK CARETAKER (COMPETENCY LEADING TO NC III)",
		"DUMP TRUCK DRIVER (COMPETENCY LEADING TO NC III)",
		"ELECTRIC MACHINE REWINDER (COMPETENCY LEADING TO NC I)",
		"ELECTRIC METER REPAIRMAN (COMPETENCY LEADING TO NC I)",
		"ELECTRIC MOTOR REPAIRMAN (COMPETENCY LEADING TO NC I)",
		"ELECTRICAL APPLIANCE SERVICE TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"ELECTRICAL CONTROL OPERATOR (HYDRO POWER PLANT) (COMPETENCY LEADING TO NC I)",
		"ELECTRICAL CONTROL OPERATOR (SUBSTATION) (COMPETENCY LEADING TO NC II)",
		"ELECTRICAL CONTROL OPERATOR (THERMAL POWER PLANT) (COMPETENCY LEADING TO NC I)",
		"ELECTRICAL CONTROL OPERATOR (THERMAL POWER PLANT) (NATIONAL CERTIFICATE I)",
		"ELECTRICAL METER REPAIRMAN (COMPETENCY LEADING TO NC I)",
		"ELECTRICAL METER REPAIRMAN (COMPETENCY LEADING TO NC III)",
		"EMBROIDERER (MULTI - HEAD MACHINE) (COMPETENCY LEADING TO NC II)",
		"ENGINE RATING - OILER/MOTORMAN (COMPETENCY LEADING TO NC II)",
		"EXTRACTIVE METALLURGY TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"F AND B STEWARD (COMPETENCY LEADING TO NC II)",
		"F AND B STEWARD (COMPETENCY LEADING TO NC III)",
		"FABRIC INSPECTOR (COMPETENCY LEADING TO NC II)",
		"FABRIC QUALITY INSPECTOR (COMPETENCY LEADING TO NC II)",
		"FABRIC QUALITY INSPECTOR (COMPETENCY LEADING TO NC III)",
		"FAMILY DRIVER (COMPETENCY LEADING TO NC II)",
		"FARM TRACTOR MECHANIC (COMPETENCY LEADING TO NC II)",
		"FARM TRACTOR OPERATOR (COMPETENCY LEADING TO NC I)",
		"FIELD CROP FARM WORKER (COMPETENCY LEADING TO NC I)",
		"FIELD EQUIPMENT OPERATOR (COMPETENCY LEADING TO NC I)",
		"FINISH CARPENTER (COMPETENCY LEADING TO NC I)",
		"FINISH CARPENTER (COMPETENCY LEADING TO NC III)",
		"FINISHING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"FINISHING MACHINE OPERATOR (COMPETENCY LEADING TO NC II)",
		"FINISHING MACHINE OPERATOR (COMPETENCY LEADING TO NC III)",
		"FIPER OPTIC CABLE SPLICER (COMPETENCY LEADING TO NC II)",
		"FISHPEN WORKER (LABORER,AIDE,TECHNICIAN) (COMPETENCY LEADING TO NC I)",
		"FISHPEN WORKER (LABORER,AIDE,TECHNICIAN) (COMPETENCY LEADING TO NC II)",
		"FISHPEN WORKER (LABORER,AIDE,TECHNICIAN) (COMPETENCY LEADING TO NC III)",
		"FISHPOND WORKER (COMPETENCY LEADING TO NC I)",
		"FISHPOND WORKER (COMPETENCY LEADING TO NC II)",
		"FISHPOND WORKER (COMPETENCY LEADING TO NC III)",
		"FLEXOGRAPHIC MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"FLUX CORED ARC WELDER (FCAW) (NATIONAL CERTIFICATE I)",
		"FLUX CORED ARC WELDER (FCAW) (NATIONAL CERTIFICATE II)",
		"FLUX CORED ARC WELDER (FCAW) (NATIONAL CERTIFICATE III)",
		"FLUX CORED ARC WELDER (FCAW) (NATIONAL LICENSE)",
		"FOIL STAMPING MACHINE OPERATOR (SHEET FED) (COMPETENCY LEADING TO NC I)",
		"FOLDING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"FOOD AND BEVERAGE SERVICE ATTENDANT (NATIONAL CERTIFICATE I)",
		"FOOD AND BEVERAGE SERVICE ATTENDANT (NATIONAL CERTIFICATE II)",
		"FOOD AND BEVERAGE SERVICE ATTENDANT (NATIONAL CERTIFICATE III)",
		"FOOTWEAR MAKER (COMPETENCY LEADING TO NC III)",
		"FOOTWEAR MAKER (NATIONAL CERTIFICATE II)",
		"FORKLIFT OPERATOR (NATIONAL CERTIFICATE I)",
		"FORKLIFT OPERATOR (NATIONAL CERTIFICATE III)",
		"FRONT OFFICE AGENT (COMPETENCY LEADING TO NC I)",
		"FRUIT AND VEGETABLES CANNING PLANT OPERATOR (COMPETENCY LEADING TO NC II)",
		"FUMIGATOR/EXTERMINATOR (COMPETENCY LEADING TO NC II)",
		"FURNITURE AND CABINET ASSEMBLER (COMPETENCY LEADING NC I)",
		"FURNITURE AND CABINET ASSEMBLER (COMPETENCY LEADING NC II)",
		"FURNITURE AND CABINET ASSEMBLER (COMPETENCY LEADING NC III)",
		"FURNITURE AND AUTOMOTIVE UPHOLSTERER (COMPETENCY LEADING TO NC I)",
		"FURNITURE AND AUTOMOTIVE UPHOLSTERER (COMPETENCY LEADING TO NC II)",
		"FURNITURE AND AUTOMOTIVE UPHOLSTERER (COMPETENCY LEADING TO NC III)",
		"FURNITURE FINISHER (COMPETENCY LEADING TO NC I)",
		"FURNITURE FINISHER (COMPETENCY LEADING TO NC II)",
		"FURNITURE FULL - SIZER/DRAFTSMAN (COMPETENCY LEADING TO NC II)",
		"FURNITURE FULL - SIZER/DRAFTSMAN (COMPETENCY LEADING TO NC III)",
		"FURNITURE WOOD CARVER (MACHINE) (COMPETENCY LEADING TO NC II)",
		"GARDENER (COMPETENCY LEADING TO NC I)",
		"GARMENTS PACKER (COMPETENCY LEADING TO NC I)",
		"GARMENTS PRESSER (COMPETENCY LEADING TO NC I)",
		"GARMENTS SAMPLE MAKER (COMPETENCY LEADING TO NC III)",
		"GARMENTS SEWER (COMPETENCY LEADING TO NC II)",
		"GARMENTS SEWER (NATIONAL CERTIFICATE I)",
		"GARMENTS TRIMMER/REVISER (COMPETENCY LEADING TO NC I)",
		"GARMENTS WASHER (COMPETENCY LEADING TO NC I)",
		"GAS METAL ARC WELDER (GMAW) (NATIONAL CERTIFICATE I)",
		"GAS METAL ARC WELDER (GMAW) (NATIONAL CERTIFICATE II)",
		"GAS METAL ARC WELDER (GMAW) (NATIONAL CERTIFICATE III)",
		"GAS METAL ARC WELDER (GMAW) (NATIONAL LICENSE)",
		"GAS TUNGSTEN ARC WELDER (GTAW) (NATIONAL CERTIFICATE I)",
		"GAS TUNGSTEN ARC WELDER (GTAW) (NATIONAL CERTIFICATE II)",
		"GAS TUNGSTEN ARC WELDER (GTAW) (NATIONAL CERTIFICATE III)",
		"GAS TUNGSTEN ARC WELDER (GTAW) (NATIONAL LICENSE)",
		"GENERAL ELECTRICIAN (COMPETENCY LEADING TO NC II)",
		"GENERAL MASON (NATIONAL CERTIFICATE I)",
		"GENERAL MASON (NATIONAL CERTIFICATE II)",
		"GENERAL MASON (NATIONAL CERTIFICATE III)",
		"GOVERNESS (COMPETENCY LEADING TO NC II)",
		"GRAIN DRYER OPERATOR (COMPETENCY LEADING TO NC I)",
		"GRINDING MACHINE OPERATOR (COMPETENCY LEADING TO NC II)",
		"GRINDING MACHINE OPERATOR (COMPETENCY LEADING TO NC III)",
		"HAIRSTYLIST (COMPETENCY LEADING TO NC I)",
		"HAND - GUIDED HIGH SPEED EMBROIDERY MACHINE OPERATOR (COMPETENCY LEADING TO NC II)",
		"HEAT TREATMENT TECHNICIAN (COMPETENCY LEADING NC I)",
		"HEAT TREATMENT TECHNICIAN (COMPETENCY LEADING NC II)",
		"HEAT TREATMENT TECHNICIAN (COMPETENCY LEADING NC III)",
		"HEAVY EQUIPMENT ELECTRICIAN (NATIONAL CERTIFICATE II)",
		"HEAVY EQUIPMENT ELECTRICIAN (NATIONAL CERTIFICATE III)",
		"HEAVY EQUIPMENT MECHANIC (NATIONAL CERTIFICATE I)",
		"HEAVY EQUIPMENT MECHANIC (NATIONAL CERTIFICATE II)",
		"HEAVY EQUIPMENT MECHANIC (NATIONAL CERTIFICATE III)",
		"HOISTER (COMPETENCY LEADING TO NC I)",
		"HOISTER (COMPETENCY LEADING TO NC III)",
		"HOT KITCHEN COOK (COMPETENCY LEADING TO NC I)",
		"HOT KITCHEN COOK (COMPETENCY LEADING TO NC II)",
		"HOT KITCHEN COOK (COMPETENCY LEADING TO NC III)",
		"HOUSEBOY (COMPETENCY LEADING TO NC I)",
		"HOUSEBUILDER (COMPETENCY LEADING TO NC II)",
		"HYDRAULIC EXCAVATOR (BACKHOE) OPERATOR (COMPETENCY LEADING TO NC I)",
		"HYDRAULIC EXCAVATOR (BACKHOE) OPERATOR (COMPETENCY LEADING TO NC III)",
		"HYDRAULIC MECHANIC (COMPETENCY LEADING TO NC II)",
		"ICEPLANT MECHANIC (COMPETENCY LEADING TO NC II)",
		"ICEPLANT MECHANIC (COMPETENCY LEADING TO NC III)",
		"ICEPLANT TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"ICEPLANT TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"ILLUSTRATOR (COMPETENCY LEADING TO NC I)",
		"ILLUSTRATOR (COMPETENCY LEADING TO NC III)",
		"INDUSTRIAL CUTTER (COMPETENCY LEADING TO NC I)",
		"INDUSTRIAL CUTTER (COMPETENCY LEADING TO NC II)",
		"INDUSTRIAL CUTTER (COMPETENCY LEADING TO NC III)",
		"INDUSTRIAL ELECTRICIAN (COMPETENCY LEADING TO NC I)",
		"INDUSTRIAL ELECTRICIAN (COMPETENCY LEADING TO NC II)",
		"INDUSTRIAL ELECTRICIAN (COMPETENCY LEADING TO NC III)",
		"INDUSTRIAL REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC I)",
		"INDUSTRIAL REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC II)",
		"INDUSTRIAL REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC III)",
		"INDUSTRIAL REFRIGERATION TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"INDUSTRIAL REFRIGERATION TECHNICIAN (COMPETENCY LEADING TO NL)",
		"INDUSTRIAL SEWING MACHINE MECHANIC (COMPETENCY LEADING TO NC I)",
		"INDUSTRIAL SEWING MACHINE MECHANIC (COMPETENCY LEADING TO NC II)",
		"INDUSTRIAL SEWING MACHINE MECHANIC (COMPETENCY LEADING TO NC III)",
		"INSTRUMENTATION REPAIRMAN (COMPETENCY LEADING TO NC I)",
		"INSULATOR (COMPETENCY LEADING TO NC I)",
		"INSULATOR (COMPETENCY LEADING TO NC II)",
		"INSULATOR (COMPETENCY LEADING TO NC III)",
		"JEEPNEY DRIVER (PUJ) (COMPETENCY LEADING TO NC II)",
		"JEWELRY MAKER (GOLD AND SILVER) (COMPETENCY LEADING TO NC I)",
		"JEWELRY MAKER (GOLD AND SILVER) (COMPETENCY LEADING TO NC II)",
		"JEWELRY MAKER (GOLD AND SILVER) (COMPETENCY LEADING TO NC III)",
		"KNITTER (FLAT MACHINE) , MANUAL (COMPETENCY LEADING TO NC II)",
		"LEATHER CUTTER",
		"LEATHERGOODS ASSEMBLER (COMPETENCY LEADING TO NC I)",
		"LEATHERGOODS ASSEMBLER (COMPETENCY LEADING TO NC II)",
		"LEATHERGOODS ASSEMBLER (COMPETENCY LEADING TO NC III)",
		"LETTERPRESS MACHINE OPERATOR (CYLINDER TYPE) (COMPETENCY LEADING TO NC I)",
		"LETTERPRESS MACHINE OPERATOR (PLATEN TYPE) (COMPETENCY LEADING TO NC II)",
		"LETTERPRESS MACHINE OPERATOR (PLATEN TYPE) (COMPETENCY LEADING TO NC III)",
		"LETTERPRESS PRESSMAN OPERATOR (COMPETENCY LEADING TO NC I)",
		"LINE LEADER (GARMENTS) (COMPETENCY LEADING TO NC III)",
		"LIVESTOCK FARMER (SWINE) (COMPETENCY LEADING TO NC II)",
		"LOCOMOTIVE DRIVER (COMPETENCY LEADING TO NL)",
		"LOCOMOTIVE MECHANIC (NATIONAL CERTIFICATE I)",
		"LOCOMOTIVE MECHANIC (NATIONAL CERTIFICATE II)",
		"LOCOMOTIVE MECHANIC (NATIONAL CERTIFICATE III)",
		"LOOM FIXER (SHUTTLE CAM) (COMPETENCY LEADING TO NC II)",
		"LUBRICATION SERVICE WORKER (COMPETENCY LEADING TO NC I)",
		"MAIN PIPE LAYER (COMPETENCY LEADING TO NCI)",
		"MAINTENANCE ELECTRICIAN (CONSTRUCTION SITE) (COMPETENCY LEADING TO NC III)",
		"MAITRE D' (COMPETENCY LEADING TONC III)",
		"MAJOR PATRON (COMPETENCY LEADINGTO NL)",
		"MARINE DIESEL ENGINE MECHANIC (COMPETENCY LEADING TO NC I)",
		"MARINE DIESEL ENGINE MECHANIC (COMPETENCY LEADING TO NC II)",
		"MARINE DIESEL ENGINE MECHANIC (COMPETENCY LEADING TO NC III)",
		"MARINE DIESELENGINE MECHANIC (NEW VERSION) (COMPETENCY LEADING TO NC III)",
		"MARINE ELECTRICIAN (COMPETENCY LEADING TO NC I)",
		"MARINE ELECTRICIAN (COMPETENCY LEADING TO NC II)",
		"MARINE ELECTRICIAN (COMPETENCY LEADING TO NC III)",
		"MARINE MACHINIST/FITTER/WELDER (NATIONAL CERTIFICATE I)",
		"MARINE MACHINIST/FITTER/WELDER (NATIONAL CERTIFICATE II)",
		"MARINE MACHINIST/FITTER/WELDER (NATIONAL CERTIFICATE III)",
		"MASSEUR/MASSEUSE (COMPETENCY LEADING TO NC I)",
		"MASTER FITTER (COMPETENCY LEADING TO NC III)",
		"MECHANIC ENGINEERING TECHNICIAN (HVAC) (COMPETENCY LEADING TO NC II)",
		"MECHANICAL FLASH DRYER OPERATOR (COMPETENCY LEADING TO NC I)",
		"MEDIATRONICS TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"METALLURGICAL TECHNICIAN, PHYSICAL (METALS AND ALLOYS) (COMPETENCY LEADING TO NC III)",
		"MOISTURE METER TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"MOLDMAKER (COMPETENCY LEADING TO NC I)",
		"MOTOR GRADER OPERATOR (COMPETENCY LEADING TO NC I)",
		"MOTOR GRADER OPERATOR (COMPETENCY LEADING TO NC II)",
		"MOTORCYCLE MECHANIC (NATIONAL CERTIFICATE I)",
		"MOTORCYCLE MECHANIC (NATIONAL CERTIFICATE II)",
		"MOTORCYCLE MECHANIC (NATIONAL CERTIFICATE III)",
		"MOTORCYCLE SERVICE MECHANIC (COMPETENCY LEADING TO NC I)",
		"MOTORCYCLE SERVICE MECHANIC (COMPETENCY LEADING TO NC II)",
		"MOTORCYCLE SERVICE MECHANIC (COMPETENCY LEADING TO NC III)",
		"MULTI - PASS RICE MILL OPERATOR (COMPETENCY LEADING TO NC I)",
		"OFFICE STENOGRAPHER (COMPETENCY LEADING TO NC I)",
		"OFFSET PRESS MACHINE OPERATOR (SHEET FED) (COMPETENCY LEADING TO NC I)",
		"OFFSET PRESS MACHINE OPERATOR (SHEET FED) (COMPETENCY LEADING TO NC II)",
		"OFFSET PRESS MACHINE OPERATOR (SHEET FED) (COMPETENCY LEADING TO NC III)",
		"OFFSET PRESS MACHINE OPERATOR (WEB - FEB) (COMPETENCY LEADING TO NC I)",
		"OFFSET PRESS MACHINE OPERATOR (WEB - FEB) (COMPETENCY LEADING TO NC II)",
		"OFFSET PRESS MACHINE OPERATOR (WEB - FEB) (COMPETENCY LEADING TO NC III)",
		"OPTICIAN (COMPETENCY LEADING TO NC I)",
		"OPTICIAN (COMPETENCY LEADING TO NC II)",
		"OPTICIAN (COMPETENCY LEADING TO NC III)",
		"OYSTER FARM WORKER (LABORER,AIDE,TECHNICIAN) (COMPETENCY LEADING TO NC I)",
		"OYSTER FARM WORKER (LABORER,AIDE,TECHNICIAN) (COMPETENCY LEADING TO NC II)",
		"OYSTER FARM WORKER (LABORER,AIDE,TECHNICIAN) (COMPETENCY LEADING TO NC III)",
		"PACKAGE TYPE AIR CONDITIONER/COMMERCIAL REFRIGERIGERATOR TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"PACKAGE TYPE AIR - CONDITIONER/COMMERCIAL REFRIGERATOR MECHANIC (COMPETENCY LEADING TO NC II)",
		"PACKAGE TYPE AIR - CONDITIONER/COMMERCIAL REFRIGERATOR MECHANIC (COMPETENCY LEADING TO NC III)",
		"PACKAGE TYPE AIR - CONDITIONER/COMMERCIAL REFRIGERATOR TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"PACKAGE TYPE AIRCONDITIONING SERVICE TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"PAINTER, STRUCTURAL STEEL (COMPETENCY LEADING TO NC II)",
		"PANTRYMAN (COMPETENCY LEADING TO NC I)",
		"PANTRYMAN (COMPETENCY LEADING TO NC III)",
		"PAPER CUTTER OPERATOR (COMPETENCY LEADING TO NC I)",
		"PARQUETRY WORKER (COMPETENCY LEADING TO NC II)",
		"PASTRY COOK (COMPETENCY LEADING TO NC I)",
		"PASTRY COOK (COMPETENCY LEADING TO NC III)",
		"PATTERN AND JIG MAKER (WOODEN FURNITURE) (COMPETENCY LEADING TO NC II)",
		"PATTERN AND JIG MAKER (WOODEN FURNITURE) (COMPETENCY LEADING TO NC III)",
		"PATTERN MAKER (COMPETENCY LEADING TO NC II)",
		"PATTERN MAKER (CAD)",
		"PATTERN MAKER (NEW CAD) (COMPETENCY LEADING TO NL)",
		"PATTERN MAKER, FOOTWEAR (NATIONAL CERTIFICATE II)",
		"PBX TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"PBX TECHNICIAN (COMPETENCY LEADING TO NC III)",
		"PC NETWORK TECHNICIAN (NATIONAL CERTIFICATE II)",
		"PCM TRANSMISSION TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"PERFECT BINDING OPERATOR (COMPETENCY LEADING TO NC I)",
		"PIPE FITTER/INSULATOR HVAC/R (COMPETENCY LEADING TO NC II)",
		"PIPEFITTER - FABRICATOR (NATIONAL CERTIFICATE I)",
		"PIPEFITTER - FABRICATOR (NATIONAL CERTIFICATE II)",
		"PIPEFITTER - FABRICATOR (NATIONAL CERTIFICATE III)",
		"PLANT MAINTENANCE MECHANIC (COMPETENCY LEADING TO NC I)",
		"PLANT MAINTENANCE MECHANIC (COMPETENCY LEADING TO NC II)",
		"PLANT MAINTENANCE MECHANIC (COMPETENCY LEADING TO NC III)",
		"PLANT MAINTENANCE MECHANIC (NEW VERSION) (COMPETENCY LEADING TO NC I)",
		"PLANT MAINTENANCE MECHANIC (NEW VERSION) (COMPETENCY LEADING TO NC II)",
		"PLANT MAINTENANCE MECHANIC (NEW VERSION) (COMPETENCY LEADING TO NC III)",
		"PLASTERER, GENERAL (COMPETENCY LEADING TO NC I)",
		"PLASTERER, GENERAL (COMPETENCY LEADING TO NC III)",
		"PLASTIC LAMINATING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"PLATEMAKER (COMPETENCY LEADING TO NC II)",
		"PLUMBER (NATIONAL CERTIFICATE I)",
		"PLUMBER (NATIONAL CERTIFICATE II)",
		"PLUMBER (NATIONAL CERTIFICATE III)",
		"PORTABLE DRILLING RIG OPERATOR (COMPETENCY LEADING TO NC I)",
		"POST HARVEST FACILITY SPECIALIST/COORDINATOR (COMPETENCY LEADING TO NC II)",
		"POWER PLANT ELECTRICIAN (COMPETENCY LEADING TO NC I)",
		"POWER PLANT EQUIPMENT OPERATOR (COMPETENCY LEADING TO NC I)",
		"POWER PLANT SUBSTATION MAINTENANCE ELECTRICIAN (COMPETENCY LEADING TO NC III)",
		"POWER SWITCHBOARD OPERATOR (COMPETENCY LEADING TO NC I)",
		"POWER TILLER OPERATOR (COMPETENCY LEADING TO NC I)",
		"POWER UTILITY TRAILER OPERATOR/DRIVER (TRUCK - TRACTOR TRAILER DRIVER (COMPETENCY LEADING TO NC II)",
		"PRESSURIZED IRRIGATION SYSTEM TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"PRINTING INK WORKER (COMPETENCY LEADING TO NC I)",
		"PRINTING MACHINERY MECHANIC (OFFSET) (COMPETENCY LEADING TO NC II)",
		"PRINTING MACHINERY MECHANIC (OFFSET) (COMPETENCY LEADING TO NC III)",
		"PRODUCT DEVELOPER (FURNITURE) (COMPETENCY LEADING TO NC II)",
		"PRODUCT DEVELOPER (FURNITURE) (COMPETENCY LEADING TO NC III)",
		"PRODUCT DEVELOPER (GARMENTS) (COMPETENCY LEADING TO NL)",
		"PRODUCTION SUPERVISOR (COMPETENCY LEADING TO NC II)",
		"PROFESSIONAL DRIVER (COMPETENCY LEADING TO NC I)",
		"PROFESSIONAL DRIVER (COMPETENCY LEADING TO NC II)",
		"PROFESSIONAL DRIVER (COMPETENCY LEADING TO NC III)",
		"PROFESSIONAL DRIVER (COMPETENCY LEADING TO NL)",
		"PROTECTIVE RELAY TECHNICIAN (NATIONAL CERTIFICATE II)",
		"QUALITY CONTROL INSPECTOR (GARMENT,SEWING LINE) (COMPETENCY LEADING TO NC II)",
		"QUALITY CONTROL INSPECTOR (GENERAL) (COMPETENCY LEADING TO NC II)",
		"RATTAN FURNITURE MAKER (COMPETENCY LEADING TO NC I)",
		"RATTAN FURNITURE MAKER (COMPETENCY LEADING TO NC II)",
		"RATTAN FURNITURE SAMPLE (PROTOTYPE) (COMPETENCY LEADING TO NC II)",
		"REBAR FIXER (COMPETENCY LEADING TO NC I)",
		"REBAR FIXER (COMPETENCY LEADING TO NC II)",
		"REBAR FIXER (COMPETENCY LEADING TO NC III)",
		"REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC I)",
		"REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC II)",
		"REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC III)",
		"REPRO - CAMERAMAN (COMPETENCY LEADING TO NC II)",
		"REPRODUCTION PHOTOGRAPHY TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"REPROGRAPHICS TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"RESERVATION OFFICER (COMPETENCY LEADING TO NC I)",
		"RICE REAPER OPERATOR (COMPETENCY LEADING TO NC I)",
		"RICE STRIPPER OPERATOR (COMPETENCY LEADING TO NC I)",
		"RICE THRESHER OPERATOR (COMPETENCY LEADING TO NC I)",
		"RIGGER (COMPETENCY LEADING TO NC II)",
		"RING SPINNER (COMPETENCY LEADING TO NC I)",
		"RING SPINNER (COMPETENCY LEADING TO NC II)",
		"ROAD ROLLER OPERATOR ( (COMPETENCY LEADING TO NC I)",
		"ROOFER (COMPETENCT LEADING TO NC III)",
		"ROOFER (COMPETENCY LEADING TO NC I)",
		"ROOFER (COMPETENCY LEADING TO NC II)",
		"ROOM ATTENDANT (COMPETENCY LEADING TO NC I)",
		"ROTOGRAVURE MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"ROUGH CARPENTER (COMPETENCY LEADING TO NC II)",
		"ROUGH TERRAIN CRANE OPERATOR (COMPETENCY LEADING TO NC I)",
		"ROVING QUALITY INSPECTOR (COMPETENCY LEADING TO NC II)",
		"SADDLE STITCHING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"SCAFFOLDS ERECTOR (COMPETENCY LEADING TO NC II)",
		"SEAFARER ABLE BODIED SEAMAN (COMPETENCY LEADING TO NC II)",
		"SEAFARER CHIEF COOK (COMPETENCY LEADING TO NC II)",
		"SEAFARER CHIEF STEWARD (COMPETENCY LEADING TO NL)",
		"SEAFARER DECK RATING (NATIONAL CERTIFICATE I)",
		"SEAFARER ENGINE RATING (NATIONAL CERTIFICATE I)",
		"SEAFARER ENGINE RATING (WIPER) (COMPETENCY LEADING TO NC I)",
		"SEAFARER PUMPMAN (COMPETENCY LEADING TO NC III)",
		"SEAFARER (DECK RATING - OS) (COMPETENCY LEADING TO NC I)",
		"SEAFARER - CATERING OR STEWARDING (MESSMAN) (COMPETENCY LEADING TO NC I)",
		"SEAFARER - OILER (NATIONAL CERTIFICATE I)",
		"SECRETARY (COMPETENCY LEADING TO NC I)",
		"SECURITY GUARD (COMPETENCY LEADING TO NC I)",
		"SECURITY GUARD (COMPETENCY LEADING TO NC II)",
		"SECURITY GUARD (COMPETENCY LEADING TO NC III)",
		"SERVICE PIPE LAYER (WATER MAINTENANCE MAN) (COMPETENCY LEADING TO NC II)",
		"SHEET - METAL WORKER (COMPETENCY LEADING TO NC I)",
		"SHEET - METAL WORKER (COMPETENCY LEADING TO NC II)",
		"SHEET - METAL WORKER (COMPETENCY LEADING TO NC III)",
		"SHIELDED METAL ARC WELDER (SMAW) (NATIONAL CERTIFICATE I)",
		"SHIELDED METAL ARC WELDER (SMAW) (NATIONAL CERTIFICATE II)",
		"SHIELDED METAL ARC WELDER (SMAW) (NATIONAL CERTIFICATE III)",
		"SHIELDED METAL ARC WELDER (SMAW) (NATIONAL LICENSE)",
		"SHOES AND BAGS REPAIRER (COMPETENCY LEADING TO NC I)",
		"SHOES AND BAGS REPAIRER (COMPETENCY LEADING TO NC II)",
		"SHOP ELECTRICIAN (COMPETENCY LEADING TO NC III)",
		"SHUTTERER (COMPETENCY LEADING TO NC I)",
		"SHUTTERER (COMPETENCY LEADING TO NC II)",
		"SHUTTERER (COMPETENCY LEADING TO NC III)",
		"SHUTTLE - CAM LOOMFIXER (COMPETENCY LEADING TO NC II)",
		"SILKSCREEN PRINTER (COMPETENCY LEADING TO NC I)",
		"SINGLE - PASS RICE MILL OPERATOR (COMPETENCY LEADING TO NC I)",
		"SLASHER TENDER (COMPETENCY LEADING TO NC II)",
		"SMALL ENGINE MECHANIC (COMPETENCY LEADING TO NC I)",
		"SMALL ENGINE MECHANIC (NEW)",
		"SMALL SCALE COCO - OIL MILL TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"SMOKED FISH PROCESSOR (COMPETENCY LEADING TO NC III)",
		"SMYTH SEWING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"SOUS CHEF (NATIONAL LICENSE)",
		"SPC MDF MAINTENANCE TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"SPECIAL MACHINE OPERATOR (COMPETENCY LEADING TO NC II)",
		"SPINDLE STRIPPING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"STAMPING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"STEAM BOILER MECHANIC (COMPETENCY LEADING TO NC I)",
		"STEAM BOILER OPERATOR (COMPETENCY LEADING TO NC I)",
		"STEAM TURBINE MECHANIC (COMPETENCY LEADING TO NC I)",
		"STEELMAN (REBAR) (COMPETENCY LEADING TO NC I)",
		"STEELMAN (REBAR) (COMPETENCY LEADING TO NC II)",
		"STEELMAN (REBAR) (COMPETENCY LEADING TO NC III)",
		"STENTER MACHINE OPERATOR",
		"STENTER MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"STONE SETTER/ENGRAVER (COMPETENCY LEADING TO NC II)",
		"STRIPPER (BLACK AND WHITE) (COMPETENCY LEADING TO NC I)",
		"STRUCTURAL STEEL WORKER (COMPETENCY LEADING TO NC I)",
		"STRUCTURAL STEEL WORKER (COMPETENCY LEADING TO NC II)",
		"STRUCTURAL STEEL WORKER (COMPETENCY LEADING TO NC III)",
		"STUFF TOY MAKER (COMPETENCY LEADING TO NC I)",
		"STUFF TOY MAKER (COMPETENCY LEADING TO NC III)",
		"STW PROJECT COORDINATOR (COMPETENCY LEADING TO NC II)",
		"STW PUMP SET TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"SUB - CON COORDINATOR (SEWING LINE) (COMPETENCY LEADING TO NC III)",
		"SUBMERGED ARC WELDER (SAW) (NATIONAL CERTIFICATE I)",
		"SUBMERGED ARC WELDER (SAW) (NATIONAL CERTIFICATE II)",
		"SUBMERGED ARC WELDER (SAW) (NATIONAL CERTIFICATE III)",
		"SUBSTATIONMAN (COMPETENCY LEADING TO NC I)",
		"SWINE FARM TECHNICIAN (COMPETENCY LEADING TO NC I)",
		"SWINE FARM TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"TAILOR (NATIONAL CERTIFICATE II)",
		"TAILOR (NATIONAL CERTIFICATE III)",
		"TAXICAB DRIVER (COMPETENCY LEADING TO NC II)",
		"TELECOM POWER TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"TICKETING OFFICER (COMPETENCY LEADING TO NC I)",
		"TILESETTER, GENERAL (COMPETENCY LEADING TO NC I)",
		"TILESETTER, GENERAL (COMPETENCY LEADING TO NC III)",
		"TINSMITH HVAC/R WORKER (COMPETENCY LEADING TO NC I)",
		"TINSMITH HVAC/R WORKER (COMPETENCY LEADING TO NC II)",
		"TINSMITH HVAC/R WORKER (COMPETENCY LEADING TO NC III)",
		"TINSMITH/INSULATOR HVAC/R WORKER (COMPETENCY LEADING NC II)",
		"TIRE SERVICE WORKER (COMPETENCY LEADING TO NC I)",
		"TOUR COORDINATOR (COMPETENCY LEADING TO NC II)",
		"TOUR GUIDE (COMPETENCY LEADING TO NC II)",
		"TOWER CRANE OPERATOR (COMPETENCY LEADING TO NC II)",
		"TOWER CRANE OPERATOR (COMPETENCY LEADING TO NC III)",
		"TRANSFORMER REPAIRMAN (COMPETENCY LEADING TO NC I)",
		"TRANSMISSION LINEMAN",
		"TRANSPORT AIR - CONDITIONING AND REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC I)",
		"TRANSPORT AIR - CONDITIONING AND REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC III)",
		"TRUCK BODY BUILDER (CLOSE VAN) (COMPETENCY LEADING TO NC II)",
		"TRUCK - MOUNTED CRANE OPERATOR (COMPETENCY LEADING TO NC II)",
		"TRUCK - MOUNTED CRANE OPERATOR (COMPETENCY LEADING TO NC III)",
		"TRUCK - TRACTOR MECHANIC (NATIONAL CERTIFICATE I)",
		"TRUCK - TRACTOR MECHANIC (NATIONAL CERTIFICATE II)",
		"TRUCK - TRACTOR/TRAILER DRIVER (COMPETENCY LEADING TO NL)",
		"TYPIST (COMPETENCY LEADING TO NC I)",
		"TYPISTS (COMPETENCY LEADING TO NC I)",
		"UNDERCHASSIS AND POWERTRAIN MECHANIC",
		"UNDERCHASSIS AND POWERTRAIN MECHANIC (COMPETENCY LEADING TO NC I)",
		"UNDERCHASSIS AND POWERTRAIN MECHANIC (COMPETENCY LEADING TO NC II)",
		"UNDERCHASSIS AND POWERTRAIN MECHANIC (COMPETENCY LEADING TO NC III)",
		"UPHOLSTERER, FURNITURE AND AUTOMOTIVE",
		"UPPER MAKER (FOOTWEAR) (COMPETENCY LEADING TO NC II)",
		"VEGETABLE FARM TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"VEHICLE EMISSION CONTROL TECHNICIAN III (COMPETENCY LEADING TO NC III)",
		"VIDEO ELECTRONICS SERVICE TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"VIDEO ELECTRONICS SVC TECHNICIAN",
		"WAREHOUSEMAN (COMPETENCY LEADING TO NC I)",
		"WAREHOUSEMAN (COMPETENCY LEADING TO NC III)",
		"WATCH AND CLOCK REPAIRMAN (COMPETENCY LEADING TO NC II)",
		"WATER METER REPAIRMAN (WATER METER MECHANIC) (COMPETENCY LEADING TO NC I)",
		"WATER PUMP OPERATOR (COMPETENCY LEADING TO NC I)",
		"WATER TREATMENT OPERATOR (GAS CHLORINATION) (COMPETENCY LEADING TO NC II)",
		"WATER TREATMENT PLANT OPERATOR (COMPETENCY LEADING TO NC II)",
		"WATERWASTE INVESTIGATOR (SENIOR WATER MAINTENANCE MAN) (COMPETENCY LEADING TO NC II)",
		"WEAVER (SHUTTLE LESS LOOM) (COMPETENCY LEADING TO NC II)",
		"WEAVER, GENERAL (COMPETENCY LEADING TO NC I)",
		"WEAVER, GENERAL (COMPETENCY LEADING TO NC III)",
		"WELL DRILLER (CABLE TOOL) PERCUSSION (COMPETENCY LEADING TO NC I)",
		"WHEEL LOADER OPERATOR (COMPETENCY LEADING TO NC I)",
		"WINDOW TYPE AIR CONDITIONING/DOMESTIC REFRIGERATION MECHANIC",
		"WINDOW TYPE AIR CONDITIONING/DOMESTIC REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC I)",
		"WINDOW TYPE AIR CONDITIONING/DOMESTIC REFRIGERATION MECHANIC (COMPETENCY LEADING TO NC II)",
		"WINDOW TYPE CONDITIONING/DOMESTIC REFRIGERATION MECHANIC (NC I AND NC II)",
		"WINDOW TYPE - AIR CONDITIONER/DOMESTIC REFRIGERATOR MECHANIC (COMPETENCY LEADING TO NC I)",
		"WINDOW TYPE - AIR CONDITIONER/DOMESTIC REFRIGERATOR MECHANIC (COMPETENCY LEADING TO NC III)",
		"WLL SUBSCRIBER RADIO TELECOMMUNICATION TECHNICIAN (COMPETENCY LEADING TO NC II)",
		"WLL TELECOM TECHNICIAN (IN - PLANT NETWORK) (COMPETENCY LEADING TO NC III)",
		"WOOD CARVER (FURNITURE) (COMPETENCY LEADING TO NC II)",
		"WOOD CARVER (HAND CARVING) (COMPETENCY LEADING TO NC II)",
		"WOOD CARVER (HAND CARVING) (COMPETENCY LEADING TO NC III)",
		"WOOD FINISHER (COMPETENCY LEADING TO NC II)",
		"WOOD FURNITURE PROTYPE (SAMPLE) (COMPETENCY LEADING TO NC II)",
		"WOOD TURNER",
		"WOOD WORKING MACHINE OPERATOR (COMPETENCY LEADING TO NC I)",
		"WOOD WORKING MACHINE OPERATOR (COMPETENCY LEADING TO NC II)",
		"WOOD WORKING MACHINE OPERATOR (COMPETENCY LEADING TO NC III)",
		"WOODEN SASH MAKER (COMPETENCY LEADING TO NC II)",
	}

	for _, name := range data {
		certificate := Certificate{Name: strings.ToUpper(name)}

		if _, err := certificate.Create(); err != nil {
			panic(err)
		}
	}
}

func (certificate *Certificate) Create() (*Certificate, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&certificate).Error; err != nil {
		return nil, err
	}
	return certificate, nil
}

func (certificate Certificate) Search(q string) []Certificate {
	db := database.Conn()
	defer db.Close()

	certificates := []Certificate{}
	results := make(chan []Certificate)

	go func() {
		db.Find(&certificates, "name LIKE ?", database.WrapLike(q))
		results <- certificates
	}()
	return <-results
}
