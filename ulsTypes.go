package main

type AmateurULS struct {
}

// AM – Amateur
// create table dbo.PUBACC_AM
// (
//       record_type                char(2)              not null,
//       unique_system_identifier   numeric(9,0)         not null,
//       uls_file_num               char(14)             null,
//       ebf_number                 varchar(30)          null,
//       callsign                   char(10)             null,
//       operator_class             char(1)              null,
//       group_code                 char(1)              null,
//       region_code                tinyint              null,
//       trustee_callsign           char(10)             null,
//       trustee_indicator          char(1)              null,
//       physician_certification    char(1)              null,
//       ve_signature               char(1)              null,
//       systematic_callsign_change char(1)             null,
//       vanity_callsign_change     char(1)              null,
//       vanity_relationship        char(12)             null,
//       previous_callsign          char(10)             null,
//       previous_operator_class    char(1)              null,
//       trustee_name               varchar(50)          null
// )

// EN – Entity
// create table dbo.PUBACC_EN
// (
//       record_type               char(2)              not null,
//       unique_system_identifier  numeric(9,0)         not null,
//       uls_file_number           char(14)             null,
//       ebf_number                varchar(30)          null,
//       call_sign                 char(10)             null,
//       entity_type               char(2)              null,
//       licensee_id               char(9)              null,
//       entity_name               varchar(200)         null,
//       first_name                varchar(20)          null,
//       mi                        char(1)              null,
//       last_name                 varchar(20)          null,
//       suffix                    char(3)              null,
//       phone                     char(10)             null,
//       fax                       char(10)             null,
//       email                     varchar(50)          null,
//       street_address            varchar(60)          null,
//       city                      varchar(20)          null,
//       state                     char(2)              null,
//       zip_code                  char(9)              null,
//       po_box                    varchar(20)          null,
//       attention_line            varchar(35)          null,
//       sgin                      char(3)              null,
//       frn                       char(10)             null,
//       applicant_type_code       char(1)              null,
//       applicant_type_other      char(40)             null,
//       status_code               char(1)		     null,
//       status_date		       datetime		null,
//       lic_category_code		   char(1)		null,
//       linked_license_id	       numeric(9,0)	null,
//       linked_callsign		char(10)		null
// )

// HD - Application License/Header
// create table dbo.PUBACC_HD
// (
//       record_type               char(2)              not null,
//       unique_system_identifier  numeric(9,0)         not null,
//       uls_file_number           char(14)             null,
//       ebf_number                varchar(30)          null,
//       call_sign                 char(10)             null,
//       license_status            char(1)              null,
//       radio_service_code        char(2)              null,
//       grant_date                char(10)             null,
//       expired_date              char(10)             null,
//       cancellation_date         char(10)             null,
//       eligibility_rule_num      char(10)             null,
//       applicant_type_code_reserved       char(1)              null,
//       alien                     char(1)              null,
//       alien_government          char(1)              null,
//       alien_corporation         char(1)              null,
//       alien_officer             char(1)              null,
//       alien_control             char(1)              null,
//       revoked                   char(1)              null,
//       convicted                 char(1)              null,
//       adjudged                  char(1)              null,
//       involved_reserved         char(1)              null,
//       common_carrier            char(1)              null,
//       non_common_carrier        char(1)              null,
//       private_comm              char(1)              null,
//       fixed                     char(1)              null,
//       mobile                    char(1)              null,
//       radiolocation             char(1)              null,
//       satellite                 char(1)              null,
//       developmental_or_sta      char(1)              null,
//       interconnected_service    char(1)              null,
//       certifier_first_name      varchar(20)          null,
//       certifier_mi              char(1)              null,
//       certifier_last_name       varchar(20)          null,
//       certifier_suffix          char(3)              null,
//       certifier_title           char(40)             null,
//       gender                    char(1)              null,
//       african_american          char(1)              null,
//       native_american           char(1)              null,
//       hawaiian                  char(1)              null,
//       asian                     char(1)              null,
//       white                     char(1)              null,
//       ethnicity                 char(1)              null,
//       effective_date            char(10)             null,
//       last_action_date          char(10)             null,
//       auction_id                int                  null,
//       reg_stat_broad_serv       char(1)              null,
//       band_manager              char(1)              null,
//       type_serv_broad_serv      char(1)              null,
//       alien_ruling              char(1)              null,
//       licensee_name_change	   char(1)		        null,
//       whitespace_ind            char(1)              null,
//       additional_cert_choice    char(1)              null,
//       additional_cert_answer    char(1)              null,
//       discontinuation_ind       char(1)              null,
//       regulatory_compliance_ind char(1)              null,
//       eligibility_cert_900      char(1)              null,
//       transition_plan_cert_900    char(1)            null,
//       return_spectrum_cert_900  char(1)              null,
//       payment_cert_900        char(1)                null
// )

// HS – History
// create table dbo.PUBACC_HS
// (
//       record_type               char(2)              not null,
//       unique_system_identifier  numeric(9,0)         not null,
//       uls_file_number           char(14)             null,
//       callsign                  char(10)             null,
//       log_date                  char(10)             null,
//       code                      char(6)              null
// )

// CO - Comments
// create table dbo.PUBACC_CO
// (
//       record_type               char(2)              not null,
//       unique_system_identifier  numeric(9,0)         not null,
//       uls_file_num              char(14)             null,
//       callsign                  char(10)             null,
//       comment_date              char(10)             null,
//       description               varchar(255)         null,
//       status_code		char(1)		     null,
//       status_date		datetime             null
// )

// LA – License Attachment
// create table dbo.PUBACC_LA
//  (
//       record_type               char(2)              null ,
//       unique_system_identifier  numeric(9,0)         null ,
//       callsign                  char(10)             null ,
//       attachment_code           char(1)              Null ,
//       attachment_desc           varchar(60)          Null ,
//       attachment_date           char(10)             Null ,
//       attachment_filename       varchar(60)          Null ,
//       action_performed          char(1)              Null
// )

// SC – License Level Canned Special Conditions
// create table dbo.PUBACC_SC
// (
//       record_type               char(2)              null,
//       unique_system_identifier  numeric(9,0)         not null,
//       uls_file_number           char(14)              null,
//       ebf_number                varchar(30)           null,
//       callsign                  char(10)             null ,
//       special_condition_type    char(1)              null,
//       special_condition_code    int                  null,
// 	status_code		char(1)			null,
// 	status_date		datetime		null
// )

// SF – License Level Free Form Special Conditions
// create table dbo.PUBACC_SF
// (
//       record_type               char(2)              null ,
//       unique_system_identifier  numeric(9,0)         null ,
//       uls_file_number           char(14)              null,
//       ebf_number                varchar(30)           null,
//       callsign                  char(10)             null ,
//       lic_freeform_cond_type    char(1)              null ,
//       unique_lic_freeform_id    numeric(9,0)         null ,
//       sequence_number           int              null ,
//       lic_freeform_condition    varchar(255)         null,
// 	status_code		char(1)			null,
// 	status_date		datetime		null
// )

// AD - Application Detail
// create table dbo.PUBACC_AD
// (
//       Record_Type               	char(2)              null,
//       unique_system_identifier  	numeric(9,0)         not null,
//       ULS_File_Number           	char(14)             null,
//       EBF_Number                	varchar(30)          null,
//       Application_Purpose       	char(2)              null,
//       Application_Status       		char(1)              null,
//       Application_Fee_Exempt    	char(1)              null,
//       Regulatory_Fee_Exempt     	char(1)              null,
//       Source                    	char(1)              null,
//       requested_expiration_date_mmdd 	char(4)              null,
//       Receipt_Date              	char(10)             null,
//       Notification_Code         	char(1)              null,
//       Notification_date         	char(10)             null,
//       Expanding_Area_or_Contour 	char(1)              null,
//       Change_Type               	char(1)              null,
//       Original_Application_Purpose 	char(2)              null,
//       Requesting_A_Waiver       	char(1)              null,
//       How_Many_Waivers_Requested 	int                  null,
//       Any_Attachments           	char(1)              null,
//       Number_of_Requested_SIDs  	int                  null,
//       fee_control_num           	char(16)             null,
//       date_entered              	char(10)             null,
//       reason                    	varchar(255)         null,
//       frequency_coordination_indicat 	char(1)              null,
//       emergency_sta             	char(1)              null,
//       overall_change_type       	char(1)              null,
//       slow_growth_ind           	char(1)              null,
//       previous_waiver           	char(1)              null,
//       waiver_deferral_fee       	char(1)              null,
//       has_term_pending_ind		char(1)		     null,
//       use_of_service       		char(1)		     null
// )

// VC - Vanity Call Sign
// create table dbo.PUBACC_VC
// (
//       record_type               char(2)              null,
//       unique_system_identifier  numeric(9,0)         not null,
//       uls_file_number           char(14)             null,
//       ebf_number                varchar(30)          null,
//       request_sequence          int              null,
//       callsign_requested        char(10)             null
// )

// 	AD	Application Purpose
// 	AA	Assignment of Authorization
// 	AM	Amendment
// 	AR	DE Annual Report
// 	AU	Administrative Update
// 	CA	Cancellation of License
// 	CB	C Block Election
// 	DC	Data Correction
// 	DU	Duplicate License
// 	EX	Request for Extension of Time
// 	HA	HAC Report
// 	LC	Cancel a Lease
// 	LE	Extend Term of a Lease
// 	LL	"603T", no longer used
// 	LM	Modification of a Lease
// 	LN	New Lease
// 	LT	Transfer of Control of a Lessee
// 	LU	Administrative Update of a Lease
// 	MD	Modification
// 	NE	New
// 	NT	Required Notification
// 	RE	DE Reportable Event
// 	RL	Register Link/Location
// 	RM	Renewal/Modification
// 	RO	Renewal Only
// 	TC	Transfer of Control
// 	WD	Withdrawal of Application

// AD	Application Status
// 	1	Submitted
// 	2	Pending
// 	A	A Granted
// 	C	Consented To
// 	D	Dismissed
// 	E	Eliminate
// 	G	Granted
// 	H	History Only
// 	I	Inactive
// 	J	HAC Submitted
// 	K	Killed
// 	M	Consummated
// 	N	Granted in Part
// 	P	Pending Pack Filing
// 	Q	Accepted
// 	R	Returned
// 	S	Saved
// 	T	Terminated
// 	U	Unprocessable
// 	W	Withdrawn
// 	X	NA
// 	Y	Application has problems

// AD	Notification Code
// 	1	First Buildout/Coverage Requirement
// 	2	Second Buildout/Coverage Requirement
// 	3	Third Buildout/Coverage Requirement
// 	4	Fourth Buildout/Coverage Requirement
// 	A	All coverage requirements (for those that have neither 5 or 10)
// 	C	Consummation of transfer or assignment
// 	D	Request regular authorization for facilities operating under developmental authority
// 	G	Notification of compliance with yearly station commitments for licencees with approved extended implementation plans.
// 	H	Final notification that construction requirements have been met for referenced system with approved extended implementation plans.
// 	S	Construction Requirement
// 	T	Notification of Tribal Lands Construction
// 	P	Permissible Period of Discontinuance of Service or Operations (All Services)

// AD	Source
// 	B	Batch Filed
// 	I	Interactively Filed (external to FCC)
// 	M	Manually Keyed by FCC

// AD	Major/Minor Indicator
// 	J	Major
// 	N	Minor

// AD	Code	Description
// 	C	Geographic area license used to provide service to customers
// 	P	License is used for private business (internal) purposes or to meet the licensee's public safety communications needs

// AM 	Operator Class Code
// 	A	Advanced
// 	E	Amateur Extra
// 	G	General
// 	N	Novice
// 	P	Technician Plus
// 	T	Technician

// AN 	Antenna Type Code
// 	H	Hub
// 	P	Passive Repeater
// 	R	Final Receiver
// 	T	Transmit Antenna

// AT 	Attachment Code
// 	0      	GIS Map Files
// 	0C     	Confidential GIS Map Files
// 	10100  	Resumes/Vitae
// 	10101  	Additional Site Information
// 	10102  	Tribal/NHO Involvement
// 	10103  	Local Government Involvement
// 	10104  	Public Involvement
// 	10105  	Area of Potential Effects
// 	10106  	Historic Properties for Visual Effects
// 	10107  	Historic Properties for Direct Effects
// 	10108  	Mitigation
// 	10109  	Photographs
// 	10110  	Map Documents
// 	10111  	Confidential
// 	10112  	State-Specific Forms
// 	10113  	Cultural Resources Report
// 	10114  	Other
// 	10115  	Confidential Tribal Data
// 	10116  	Memorandum of Agreement (MOA)
// 	10117  	Environmental Assessment (EA)
// 	10118  	Photograph
// 	10119  	Map Documents
// 	10120  	Correspondence
// 	10121  	Finding of No Significant Impact (FONSI)
// 	10122  	Confidential Document
// 	10123  	Confidential Tribal Data Document
// 	10124  	Other Document
// 	10125  	Response to SHPO/THPO Request for Information
// 	10126  	GIS Map Files
// 	2      	Revised Filing Deadline Attachment
// 	74     	Rule 74.832(e) Certification
// 	A      	Application
// 	AA     	Status of Environmental Review
// 	AM     	Adaptive Modulation Certification
// 	B      	Cellular Cross Interest
// 	BJBA   	Bidding/Joint Bidding Agreement
// 	C      	Confidentiality
// 	D      	Divestiture
// 	DEA    	Designated Entity Agreement
// 	E      	603-T (Spectrum Leasing)
// 	F      	Fee Exemption
// 	G      	"BRSChannel 1,2,2A Notification"
// 	H      	Hurricane Relief
// 	I      	Indirect Ownership
// 	J      	Private Commons
// 	JVA    	Joint Venture Agreement
// 	K      	Sublease
// 	L      	Letter
// 	M      	800 MHz Band Reconfiguration
// 	N      	Ownership
// 	NBAND  	Rule 90.209(b)(6) Certification
// 	O      	Other
// 	OA     	Other Agreement
// 	P      	Pleading
// 	PAMSA  	Post-Auction Market Structure Agreement
// 	Q      	Confidential Pleading
// 	R      	Data Correction
// 	RSP    	Rural Service Provider Bidding Credit
// 	RWP    	Rural Wireless Partnership
// 	S      	1.2112(a)(6)
// 	T      	47 C.F.R. 17.7(e)(1) ASR Exemptions
// 	U      	Quiet Zone Consent
// 	V      	YC/YH attachment
// 	W      	Waiver
// 	X      	Tribal Govt. Certification
// 	Y      	Tribal Lands Waiver Request
// 	Z      	TCNS Internal Reply

// A2	PFR Status Code
// 	C	Petition for Reconsideration Period Closed
// 	D	Petition for Reconsideration Denied
// 	G	Petition for Reconsideration Granted
// 	O	Petition for Reconsideration Period Open
// 	R	Petition for Reconsideration Received

// BC 	Non-Parent Type Code
// 	A	Cable Network Entity
// 	B	Broadcast Network Entity
// 	C	Television Cable Operator
// 	L	Large Venue Owner or Operator
// 	M	Motion Picture Producer
// 	P	Professional Sound Company
// 	T	Television Producer

// BO	Buildout Code
// 	21	Construction Date 1
// 	22	Construction Requirements (date 2)
// 	31	Coverage Date 1
// 	32	Coverage Date 2
// 	33	Coverage Date 3
// 	34	Coverage Date 4
// 	35	Tribal Land Buildout Deadline

// 	"Note: The above codes are used for all three buildout tables (BO, BL, and BF)."

// BT 	Business Type
// 	M	Minority Owned Business
// 	R	Rural Telephone Company
// 	W	Woman owned Business

// CG

// CP	County
// 	State	County Name	County Active?
// 	AK	ALEUTIANS EAST                  	Y
// 	AK	ALEUTIANS WEST                  	Y
// 	AK	ANCHORAGE                       	Y
// 	AK	BETHEL                          	Y
// 	AK	BRISTOL BAY                     	Y
// 	AK	DENALI                          	Y
// 	AK	DILLINGHAM                      	Y
// 	AK	FAIRBANKS NORTH STAR            	Y
// 	AK	HAINES                          	Y
// 	AK	HOONAH-ANGOON                   	Y
// 	AK	JUNEAU                          	Y
// 	AK	KENAI PENINSULA                 	Y
// 	AK	KETCHIKAN GATEWAY               	Y
// 	AK	KODIAK ISLAND                   	Y
// 	AK	LAKE AND PENINSULA              	Y
// 	AK	MATANUSKA-SUSITNA               	Y
// 	AK	NOME                            	Y
// 	AK	NORTH SLOPE                     	Y
// 	AK	NORTHWEST ARCTIC                	Y
// 	AK	PETERSBURG                      	Y
// 	AK	PRINCE OF WALES-HYDER           	Y
// 	AK	PRINCE OF WALES-OUTER KETCHIKAN 	Y
// 	AK	SITKA                           	Y
// 	AK	SKAGWAY                         	Y
// 	AK	SKAGWAY-HOONAH-ANGOON           	Y
// 	AK	Skagway-Yakutat-Angoon          	N
// 	AK	SOUTHEAST FAIRBANKS             	Y
// 	AK	VALDEZ-CORDOVA                  	Y
// 	AK	WADE HAMPTON                    	Y
// 	AK	WRANGELL-PETERSBURG             	Y
// 	AK	YAKUTAT                         	Y
// 	AK	YUKON-KOYUKUK                   	Y
// 	AL	AUTAUGA                         	Y
// 	AL	BALDWIN                         	Y
// 	AL	BARBOUR                         	Y
// 	AL	BIBB                            	Y
// 	AL	BLOUNT                          	Y
// 	AL	BULLOCK                         	Y
// 	AL	BUTLER                          	Y
// 	AL	CALHOUN                         	Y
// 	AL	CHAMBERS                        	Y
// 	AL	CHEROKEE                        	Y
// 	AL	CHILTON                         	Y
// 	AL	CHOCTAW                         	Y
// 	AL	CLARKE                          	Y
// 	AL	CLAY                            	Y
// 	AL	CLEBURNE                        	Y
// 	AL	COFFEE                          	Y
// 	AL	COLBERT                         	Y
// 	AL	CONECUH                         	Y
// 	AL	COOSA                           	Y
// 	AL	COVINGTON                       	Y
// 	AL	CRENSHAW                        	Y
// 	AL	CULLMAN                         	Y
// 	AL	DALE                            	Y
// 	AL	DALLAS                          	Y
// 	AL	DEKALB                          	Y
// 	AL	ELMORE                          	Y
// 	AL	ESCAMBIA                        	Y
// 	AL	ETOWAH                          	Y
// 	AL	FAYETTE                         	Y
// 	AL	FRANKLIN                        	Y
// 	AL	GENEVA                          	Y
// 	AL	GREENE                          	Y
// 	AL	HALE                            	Y
// 	AL	HENRY                           	Y
// 	AL	HOUSTON                         	Y
// 	AL	JACKSON                         	Y
// 	AL	JEFFERSON                       	Y
// 	AL	LAMAR                           	Y
// 	AL	LAUDERDALE                      	Y
// 	AL	LAWRENCE                        	Y
// 	AL	LEE                             	Y
// 	AL	LIMESTONE                       	Y
// 	AL	LOWNDES                         	Y
// 	AL	MACON                           	Y
// 	AL	MADISON                         	Y
// 	AL	MARENGO                         	Y
// 	AL	MARION                          	Y
// 	AL	MARSHALL                        	Y
// 	AL	MOBILE                          	Y
// 	AL	MONROE                          	Y
// 	AL	MONTGOMERY                      	Y
// 	AL	MORGAN                          	Y
// 	AL	PERRY                           	Y
// 	AL	PICKENS                         	Y
// 	AL	PIKE                            	Y
// 	AL	RANDOLPH                        	Y
// 	AL	RUSSELL                         	Y
// 	AL	SHELBY                          	Y
// 	AL	ST. CLAIR                       	Y
// 	AL	SUMTER                          	Y
// 	AL	TALLADEGA                       	Y
// 	AL	TALLAPOOSA                      	Y
// 	AL	TUSCALOOSA                      	Y
// 	AL	WALKER                          	Y
// 	AL	WASHINGTON                      	Y
// 	AL	WILCOX                          	Y
// 	AL	WINSTON                         	Y
// 	AR	ARKANSAS                        	Y
// 	AR	ASHLEY                          	Y
// 	AR	BAXTER                          	Y
// 	AR	BENTON                          	Y
// 	AR	BOONE                           	Y
// 	AR	BRADLEY                         	Y
// 	AR	CALHOUN                         	Y
// 	AR	CARROLL                         	Y
// 	AR	CHICOT                          	Y
// 	AR	CLARK                           	Y
// 	AR	CLAY                            	Y
// 	AR	CLEBURNE                        	Y
// 	AR	CLEVELAND                       	Y
// 	AR	COLUMBIA                        	Y
// 	AR	CONWAY                          	Y
// 	AR	CRAIGHEAD                       	Y
// 	AR	CRAWFORD                        	Y
// 	AR	CRITTENDEN                      	Y
// 	AR	CROSS                           	Y
// 	AR	DALLAS                          	Y
// 	AR	DESHA                           	Y
// 	AR	DREW                            	Y
// 	AR	FAULKNER                        	Y
// 	AR	FRANKLIN                        	Y
// 	AR	FULTON                          	Y
// 	AR	GARLAND                         	Y
// 	AR	GRANT                           	Y
// 	AR	GREENE                          	Y
// 	AR	HEMPSTEAD                       	Y
// 	AR	HOT SPRING                      	Y
// 	AR	HOWARD                          	Y
// 	AR	INDEPENDENCE                    	Y
// 	AR	IZARD                           	Y
// 	AR	JACKSON                         	Y
// 	AR	JEFFERSON                       	Y
// 	AR	JOHNSON                         	Y
// 	AR	LAFAYETTE                       	Y
// 	AR	LAWRENCE                        	Y
// 	AR	LEE                             	Y
// 	AR	LINCOLN                         	Y
// 	AR	LITTLE RIVER                    	Y
// 	AR	LOGAN                           	Y
// 	AR	LONOKE                          	Y
// 	AR	MADISON                         	Y
// 	AR	MARION                          	Y
// 	AR	MILLER                          	Y
// 	AR	MISSISSIPPI                     	Y
// 	AR	MONROE                          	Y
// 	AR	MONTGOMERY                      	Y
// 	AR	NEVADA                          	Y
// 	AR	NEWTON                          	Y
// 	AR	OUACHITA                        	Y
// 	AR	PERRY                           	Y
// 	AR	PHILLIPS                        	Y
// 	AR	PIKE                            	Y
// 	AR	POINSETT                        	Y
// 	AR	POLK                            	Y
// 	AR	POPE                            	Y
// 	AR	PRAIRIE                         	Y
// 	AR	PULASKI                         	Y
// 	AR	RANDOLPH                        	Y
// 	AR	SALINE                          	Y
// 	AR	SCOTT                           	Y
// 	AR	SEARCY                          	Y
// 	AR	SEBASTIAN                       	Y
// 	AR	SEVIER                          	Y
// 	AR	SHARP                           	Y
// 	AR	ST. FRANCIS                     	Y
// 	AR	STONE                           	Y
// 	AR	UNION                           	Y
// 	AR	VAN BUREN                       	Y
// 	AR	WASHINGTON                      	Y
// 	AR	WHITE                           	Y
// 	AR	WOODRUFF                        	Y
// 	AR	YELL                            	Y
// 	AS	EASTERN                         	Y
// 	AS	MANUA                           	Y
// 	AS	ROSE ISLAND                     	Y
// 	AS	SWAINS ISLAND                   	Y
// 	AS	WESTERN                         	Y
// 	AZ	APACHE                          	Y
// 	AZ	COCHISE                         	Y
// 	AZ	COCONINO                        	Y
// 	AZ	GILA                            	Y
// 	AZ	GRAHAM                          	Y
// 	AZ	GREENLEE                        	Y
// 	AZ	LA PAZ                          	Y
// 	AZ	MARICOPA                        	Y
// 	AZ	MOHAVE                          	Y
// 	AZ	NAVAJO                          	Y
// 	AZ	PIMA                            	Y
// 	AZ	PINAL                           	Y
// 	AZ	SANTA CRUZ                      	Y
// 	AZ	YAVAPAI                         	Y
// 	AZ	YUMA                            	Y
// 	CA	ALAMEDA                         	Y
// 	CA	ALPINE                          	Y
// 	CA	AMADOR                          	Y
// 	CA	BUTTE                           	Y
// 	CA	CALAVERAS                       	Y
// 	CA	COLUSA                          	Y
// 	CA	CONTRA COSTA                    	Y
// 	CA	DEL NORTE                       	Y
// 	CA	EL DORADO                       	Y
// 	CA	FRESNO                          	Y
// 	CA	GLENN                           	Y
// 	CA	HUMBOLDT                        	Y
// 	CA	IMPERIAL                        	Y
// 	CA	INYO                            	Y
// 	CA	KERN                            	Y
// 	CA	KINGS                           	Y
// 	CA	LAKE                            	Y
// 	CA	LASSEN                          	Y
// 	CA	LOS ANGELES                     	Y
// 	CA	MADERA                          	Y
// 	CA	MARIN                           	Y
// 	CA	MARIPOSA                        	Y
// 	CA	MENDOCINO                       	Y
// 	CA	MERCED                          	Y
// 	CA	MODOC                           	Y
// 	CA	MONO                            	Y
// 	CA	MONTEREY                        	Y
// 	CA	NAPA                            	Y
// 	CA	NEVADA                          	Y
// 	CA	ORANGE                          	Y
// 	CA	PLACER                          	Y
// 	CA	PLUMAS                          	Y
// 	CA	RIVERSIDE                       	Y
// 	CA	SACRAMENTO                      	Y
// 	CA	SAN BENITO                      	Y
// 	CA	SAN BERNARDINO                  	Y
// 	CA	SAN DIEGO                       	Y
// 	CA	SAN FRANCISCO                   	Y
// 	CA	SAN JOAQUIN                     	Y
// 	CA	SAN LUIS OBISPO                 	Y
// 	CA	SAN MATEO                       	Y
// 	CA	SANTA BARBARA                   	Y
// 	CA	SANTA CLARA                     	Y
// 	CA	SANTA CRUZ                      	Y
// 	CA	SHASTA                          	Y
// 	CA	SIERRA                          	Y
// 	CA	SISKIYOU                        	Y
// 	CA	SOLANO                          	Y
// 	CA	SONOMA                          	Y
// 	CA	STANISLAUS                      	Y
// 	CA	SUTTER                          	Y
// 	CA	TEHAMA                          	Y
// 	CA	TRINITY                         	Y
// 	CA	TULARE                          	Y
// 	CA	TUOLUMNE                        	Y
// 	CA	VENTURA                         	Y
// 	CA	YOLO                            	Y
// 	CA	YUBA                            	Y
// 	CO	ADAMS                           	Y
// 	CO	ALAMOSA                         	Y
// 	CO	ARAPAHOE                        	Y
// 	CO	ARCHULETA                       	Y
// 	CO	BACA                            	Y
// 	CO	BENT                            	Y
// 	CO	BOULDER                         	Y
// 	CO	BROOMFIELD                      	Y
// 	CO	CHAFFEE                         	Y
// 	CO	CHEYENNE                        	Y
// 	CO	CLEAR CREEK                     	Y
// 	CO	CONEJOS                         	Y
// 	CO	COSTILLA                        	Y
// 	CO	CROWLEY                         	Y
// 	CO	CUSTER                          	Y
// 	CO	DELTA                           	Y
// 	CO	DENVER                          	Y
// 	CO	DOLORES                         	Y
// 	CO	DOUGLAS                         	Y
// 	CO	EAGLE                           	Y
// 	CO	EL PASO                         	Y
// 	CO	ELBERT                          	Y
// 	CO	FREMONT                         	Y
// 	CO	GARFIELD                        	Y
// 	CO	GILPIN                          	Y
// 	CO	GRAND                           	Y
// 	CO	GUNNISON                        	Y
// 	CO	HINSDALE                        	Y
// 	CO	HUERFANO                        	Y
// 	CO	JACKSON                         	Y
// 	CO	JEFFERSON                       	Y
// 	CO	KIOWA                           	Y
// 	CO	KIT CARSON                      	Y
// 	CO	LA PLATA                        	Y
// 	CO	LAKE                            	Y
// 	CO	LARIMER                         	Y
// 	CO	LAS ANIMAS                      	Y
// 	CO	LINCOLN                         	Y
// 	CO	LOGAN                           	Y
// 	CO	MESA                            	Y
// 	CO	MINERAL                         	Y
// 	CO	MOFFAT                          	Y
// 	CO	MONTEZUMA                       	Y
// 	CO	MONTROSE                        	Y
// 	CO	MORGAN                          	Y
// 	CO	OTERO                           	Y
// 	CO	OURAY                           	Y
// 	CO	PARK                            	Y
// 	CO	PHILLIPS                        	Y
// 	CO	PITKIN                          	Y
// 	CO	PROWERS                         	Y
// 	CO	PUEBLO                          	Y
// 	CO	RIO BLANCO                      	Y
// 	CO	RIO GRANDE                      	Y
// 	CO	ROUTT                           	Y
// 	CO	SAGUACHE                        	Y
// 	CO	SAN JUAN                        	Y
// 	CO	SAN MIGUEL                      	Y
// 	CO	SEDGWICK                        	Y
// 	CO	SUMMIT                          	Y
// 	CO	TELLER                          	Y
// 	CO	WASHINGTON                      	Y
// 	CO	WELD                            	Y
// 	CO	YUMA                            	Y
// 	CT	FAIRFIELD                       	Y
// 	CT	HARTFORD                        	Y
// 	CT	LITCHFIELD                      	Y
// 	CT	MIDDLESEX                       	Y
// 	CT	NEW HAVEN                       	Y
// 	CT	NEW LONDON                      	Y
// 	CT	TOLLAND                         	Y
// 	CT	WINDHAM                         	Y
// 	DC	Washington                      	Y
// 	DE	KENT                            	Y
// 	DE	NEW CASTLE                      	Y
// 	DE	SUSSEX                          	Y
// 	FL	ALACHUA                         	Y
// 	FL	BAKER                           	Y
// 	FL	BAY                             	Y
// 	FL	BRADFORD                        	Y
// 	FL	BREVARD                         	Y
// 	FL	BROWARD                         	Y
// 	FL	CALHOUN                         	Y
// 	FL	CHARLOTTE                       	Y
// 	FL	CITRUS                          	Y
// 	FL	CLAY                            	Y
// 	FL	COLLIER                         	Y
// 	FL	COLUMBIA                        	Y
// 	FL	DADE                            	N
// 	FL	DESOTO                          	Y
// 	FL	DIXIE                           	Y
// 	FL	DUVAL                           	Y
// 	FL	ESCAMBIA                        	Y
// 	FL	FLAGLER                         	Y
// 	FL	FRANKLIN                        	Y
// 	FL	GADSDEN                         	Y
// 	FL	GILCHRIST                       	Y
// 	FL	GLADES                          	Y
// 	FL	GULF                            	Y
// 	FL	HAMILTON                        	Y
// 	FL	HARDEE                          	Y
// 	FL	HENDRY                          	Y
// 	FL	HERNANDO                        	Y
// 	FL	HIGHLANDS                       	Y
// 	FL	HILLSBOROUGH                    	Y
// 	FL	HOLMES                          	Y
// 	FL	INDIAN RIVER                    	Y
// 	FL	JACKSON                         	Y
// 	FL	JEFFERSON                       	Y
// 	FL	LAFAYETTE                       	Y
// 	FL	LAKE                            	Y
// 	FL	LEE                             	Y
// 	FL	LEON                            	Y
// 	FL	LEVY                            	Y
// 	FL	LIBERTY                         	Y
// 	FL	MADISON                         	Y
// 	FL	MANATEE                         	Y
// 	FL	MARION                          	Y
// 	FL	MARTIN                          	Y
// 	FL	MIAMI-DADE                      	Y
// 	FL	MONROE                          	Y
// 	FL	NASSAU                          	Y
// 	FL	OKALOOSA                        	Y
// 	FL	OKEECHOBEE                      	Y
// 	FL	ORANGE                          	Y
// 	FL	OSCEOLA                         	Y
// 	FL	PALM BEACH                      	Y
// 	FL	PASCO                           	Y
// 	FL	PINELLAS                        	Y
// 	FL	POLK                            	Y
// 	FL	PUTNAM                          	Y
// 	FL	SANTA ROSA                      	Y
// 	FL	SARASOTA                        	Y
// 	FL	SEMINOLE                        	Y
// 	FL	ST. JOHNS                       	Y
// 	FL	ST. LUCIE                       	Y
// 	FL	SUMTER                          	Y
// 	FL	SUWANNEE                        	Y
// 	FL	TAYLOR                          	Y
// 	FL	UNION                           	Y
// 	FL	VOLUSIA                         	Y
// 	FL	WAKULLA                         	Y
// 	FL	WALTON                          	Y
// 	FL	WASHINGTON                      	Y
// 	GA	APPLING                         	Y
// 	GA	ATKINSON                        	Y
// 	GA	BACON                           	Y
// 	GA	BAKER                           	Y
// 	GA	BALDWIN                         	Y
// 	GA	BANKS                           	Y
// 	GA	BARROW                          	Y
// 	GA	BARTOW                          	Y
// 	GA	BEN HILL                        	Y
// 	GA	BERRIEN                         	Y
// 	GA	BIBB                            	Y
// 	GA	BLECKLEY                        	Y
// 	GA	BRANTLEY                        	Y
// 	GA	BROOKS                          	Y
// 	GA	BRYAN                           	Y
// 	GA	BULLOCH                         	Y
// 	GA	BURKE                           	Y
// 	GA	BUTTS                           	Y
// 	GA	CALHOUN                         	Y
// 	GA	CAMDEN                          	Y
// 	GA	CANDLER                         	Y
// 	GA	CARROLL                         	Y
// 	GA	CATOOSA                         	Y
// 	GA	CHARLTON                        	Y
// 	GA	CHATHAM                         	Y
// 	GA	CHATTAHOOCHEE                   	Y
// 	GA	CHATTOOGA                       	Y
// 	GA	CHEROKEE                        	Y
// 	GA	CLARKE                          	Y
// 	GA	CLAY                            	Y
// 	GA	CLAYTON                         	Y
// 	GA	CLINCH                          	Y
// 	GA	COBB                            	Y
// 	GA	COFFEE                          	Y
// 	GA	COLQUITT                        	Y
// 	GA	COLUMBIA                        	Y
// 	GA	COOK                            	Y
// 	GA	COWETA                          	Y
// 	GA	CRAWFORD                        	Y
// 	GA	CRISP                           	Y
// 	GA	DADE                            	Y
// 	GA	DAWSON                          	Y
// 	GA	DECATUR                         	Y
// 	GA	DEKALB                          	Y
// 	GA	DODGE                           	Y
// 	GA	DOOLY                           	Y
// 	GA	DOUGHERTY                       	Y
// 	GA	DOUGLAS                         	Y
// 	GA	EARLY                           	Y
// 	GA	ECHOLS                          	Y
// 	GA	EFFINGHAM                       	Y
// 	GA	ELBERT                          	Y
// 	GA	EMANUEL                         	Y
// 	GA	EVANS                           	Y
// 	GA	FANNIN                          	Y
// 	GA	FAYETTE                         	Y
// 	GA	FLOYD                           	Y
// 	GA	FORSYTH                         	Y
// 	GA	FRANKLIN                        	Y
// 	GA	FULTON                          	Y
// 	GA	GILMER                          	Y
// 	GA	GLASCOCK                        	Y
// 	GA	GLYNN                           	Y
// 	GA	GORDON                          	Y
// 	GA	GRADY                           	Y
// 	GA	GREENE                          	Y
// 	GA	GWINNETT                        	Y
// 	GA	HABERSHAM                       	Y
// 	GA	HALL                            	Y
// 	GA	HANCOCK                         	Y
// 	GA	HARALSON                        	Y
// 	GA	HARRIS                          	Y
// 	GA	HART                            	Y
// 	GA	HEARD                           	Y
// 	GA	HENRY                           	Y
// 	GA	HOUSTON                         	Y
// 	GA	IRWIN                           	Y
// 	GA	JACKSON                         	Y
// 	GA	JASPER                          	Y
// 	GA	JEFF DAVIS                      	Y
// 	GA	JEFFERSON                       	Y
// 	GA	JENKINS                         	Y
// 	GA	JOHNSON                         	Y
// 	GA	JONES                           	Y
// 	GA	LAMAR                           	Y
// 	GA	LANIER                          	Y
// 	GA	LAURENS                         	Y
// 	GA	LEE                             	Y
// 	GA	LIBERTY                         	Y
// 	GA	LINCOLN                         	Y
// 	GA	LONG                            	Y
// 	GA	LOWNDES                         	Y
// 	GA	LUMPKIN                         	Y
// 	GA	MACON                           	Y
// 	GA	MADISON                         	Y
// 	GA	MARION                          	Y
// 	GA	MCDUFFIE                        	Y
// 	GA	MCINTOSH                        	Y
// 	GA	MERIWETHER                      	Y
// 	GA	MILLER                          	Y
// 	GA	MITCHELL                        	Y
// 	GA	MONROE                          	Y
// 	GA	MONTGOMERY                      	Y
// 	GA	MORGAN                          	Y
// 	GA	MURRAY                          	Y
// 	GA	MUSCOGEE                        	Y
// 	GA	NEWTON                          	Y
// 	GA	OCONEE                          	Y
// 	GA	OGLETHORPE                      	Y
// 	GA	PAULDING                        	Y
// 	GA	PEACH                           	Y
// 	GA	PICKENS                         	Y
// 	GA	PIERCE                          	Y
// 	GA	PIKE                            	Y
// 	GA	POLK                            	Y
// 	GA	PULASKI                         	Y
// 	GA	PUTNAM                          	Y
// 	GA	QUITMAN                         	Y
// 	GA	RABUN                           	Y
// 	GA	RANDOLPH                        	Y
// 	GA	RICHMOND                        	Y
// 	GA	ROCKDALE                        	Y
// 	GA	SCHLEY                          	Y
// 	GA	SCREVEN                         	Y
// 	GA	SEMINOLE                        	Y
// 	GA	SPALDING                        	Y
// 	GA	STEPHENS                        	Y
// 	GA	STEWART                         	Y
// 	GA	SUMTER                          	Y
// 	GA	TALBOT                          	Y
// 	GA	TALIAFERRO                      	Y
// 	GA	TATTNALL                        	Y
// 	GA	TAYLOR                          	Y
// 	GA	TELFAIR                         	Y
// 	GA	TERRELL                         	Y
// 	GA	THOMAS                          	Y
// 	GA	TIFT                            	Y
// 	GA	TOOMBS                          	Y
// 	GA	TOWNS                           	Y
// 	GA	TREUTLEN                        	Y
// 	GA	TROUP                           	Y
// 	GA	TURNER                          	Y
// 	GA	TWIGGS                          	Y
// 	GA	UNION                           	Y
// 	GA	UPSON                           	Y
// 	GA	WALKER                          	Y
// 	GA	WALTON                          	Y
// 	GA	WARE                            	Y
// 	GA	WARREN                          	Y
// 	GA	WASHINGTON                      	Y
// 	GA	WAYNE                           	Y
// 	GA	WEBSTER                         	Y
// 	GA	WHEELER                         	Y
// 	GA	WHITE                           	Y
// 	GA	WHITFIELD                       	Y
// 	GA	WILCOX                          	Y
// 	GA	WILKES                          	Y
// 	GA	WILKINSON                       	Y
// 	GA	WORTH                           	Y
// 	GM	GULF OF MEXICO                  	Y
// 	GU	GUAM                            	Y
// 	HI	HAWAII                          	Y
// 	HI	HONOLULU                        	Y
// 	HI	KALAWAO                         	Y
// 	HI	KAUAI                           	Y
// 	HI	MAUI                            	Y
// 	IA	ADAIR                           	Y
// 	IA	ADAMS                           	Y
// 	IA	ALLAMAKEE                       	Y
// 	IA	APPANOOSE                       	Y
// 	IA	AUDUBON                         	Y
// 	IA	BENTON                          	Y
// 	IA	BLACK HAWK                      	Y
// 	IA	BOONE                           	Y
// 	IA	BREMER                          	Y
// 	IA	BUCHANAN                        	Y
// 	IA	BUENA VISTA                     	Y
// 	IA	BUTLER                          	Y
// 	IA	CALHOUN                         	Y
// 	IA	CARROLL                         	Y
// 	IA	CASS                            	Y
// 	IA	CEDAR                           	Y
// 	IA	CERRO GORDO                     	Y
// 	IA	CHEROKEE                        	Y
// 	IA	CHICKASAW                       	Y
// 	IA	CLARKE                          	Y
// 	IA	CLAY                            	Y
// 	IA	CLAYTON                         	Y
// 	IA	CLINTON                         	Y
// 	IA	CRAWFORD                        	Y
// 	IA	DALLAS                          	Y
// 	IA	DAVIS                           	Y
// 	IA	DECATUR                         	Y
// 	IA	DELAWARE                        	Y
// 	IA	DES MOINES                      	Y
// 	IA	DICKINSON                       	Y
// 	IA	DUBUQUE                         	Y
// 	IA	EMMET                           	Y
// 	IA	FAYETTE                         	Y
// 	IA	FLOYD                           	Y
// 	IA	FRANKLIN                        	Y
// 	IA	FREMONT                         	Y
// 	IA	GREENE                          	Y
// 	IA	GRUNDY                          	Y
// 	IA	GUTHRIE                         	Y
// 	IA	HAMILTON                        	Y
// 	IA	HANCOCK                         	Y
// 	IA	HARDIN                          	Y
// 	IA	HARRISON                        	Y
// 	IA	HENRY                           	Y
// 	IA	HOWARD                          	Y
// 	IA	HUMBOLDT                        	Y
// 	IA	IDA                             	Y
// 	IA	IOWA                            	Y
// 	IA	JACKSON                         	Y
// 	IA	JASPER                          	Y
// 	IA	JEFFERSON                       	Y
// 	IA	JOHNSON                         	Y
// 	IA	JONES                           	Y
// 	IA	KEOKUK                          	Y
// 	IA	KOSSUTH                         	Y
// 	IA	LEE                             	Y
// 	IA	LINN                            	Y
// 	IA	LOUISA                          	Y
// 	IA	LUCAS                           	Y
// 	IA	LYON                            	Y
// 	IA	MADISON                         	Y
// 	IA	MAHASKA                         	Y
// 	IA	MARION                          	Y
// 	IA	MARSHALL                        	Y
// 	IA	MILLS                           	Y
// 	IA	MITCHELL                        	Y
// 	IA	MONONA                          	Y
// 	IA	MONROE                          	Y
// 	IA	MONTGOMERY                      	Y
// 	IA	MUSCATINE                       	Y
// 	IA	O'BRIEN                         	Y
// 	IA	OSCEOLA                         	Y
// 	IA	PAGE                            	Y
// 	IA	PALO ALTO                       	Y
// 	IA	PLYMOUTH                        	Y
// 	IA	POCAHONTAS                      	Y
// 	IA	POLK                            	Y
// 	IA	POTTAWATTAMIE                   	Y
// 	IA	POWESHIEK                       	Y
// 	IA	RINGGOLD                        	Y
// 	IA	SAC                             	Y
// 	IA	SCOTT                           	Y
// 	IA	SHELBY                          	Y
// 	IA	SIOUX                           	Y
// 	IA	STORY                           	Y
// 	IA	TAMA                            	Y
// 	IA	TAYLOR                          	Y
// 	IA	UNION                           	Y
// 	IA	VAN BUREN                       	Y
// 	IA	WAPELLO                         	Y
// 	IA	WARREN                          	Y
// 	IA	WASHINGTON                      	Y
// 	IA	WAYNE                           	Y
// 	IA	WEBSTER                         	Y
// 	IA	WINNEBAGO                       	Y
// 	IA	WINNESHIEK                      	Y
// 	IA	WOODBURY                        	Y
// 	IA	WORTH                           	Y
// 	IA	WRIGHT                          	Y
// 	ID	ADA                             	Y
// 	ID	ADAMS                           	Y
// 	ID	BANNOCK                         	Y
// 	ID	BEAR LAKE                       	Y
// 	ID	BENEWAH                         	Y
// 	ID	BINGHAM                         	Y
// 	ID	BLAINE                          	Y
// 	ID	BOISE                           	Y
// 	ID	BONNER                          	Y
// 	ID	BONNEVILLE                      	Y
// 	ID	BOUNDARY                        	Y
// 	ID	BUTTE                           	Y
// 	ID	CAMAS                           	Y
// 	ID	CANYON                          	Y
// 	ID	CARIBOU                         	Y
// 	ID	CASSIA                          	Y
// 	ID	CLARK                           	Y
// 	ID	CLEARWATER                      	Y
// 	ID	CUSTER                          	Y
// 	ID	ELMORE                          	Y
// 	ID	FRANKLIN                        	Y
// 	ID	FREMONT                         	Y
// 	ID	GEM                             	Y
// 	ID	GOODING                         	Y
// 	ID	IDAHO                           	Y
// 	ID	JEFFERSON                       	Y
// 	ID	JEROME                          	Y
// 	ID	KOOTENAI                        	Y
// 	ID	LATAH                           	Y
// 	ID	LEMHI                           	Y
// 	ID	LEWIS                           	Y
// 	ID	LINCOLN                         	Y
// 	ID	MADISON                         	Y
// 	ID	MINIDOKA                        	Y
// 	ID	NEZ PERCE                       	Y
// 	ID	ONEIDA                          	Y
// 	ID	OWYHEE                          	Y
// 	ID	PAYETTE                         	Y
// 	ID	POWER                           	Y
// 	ID	SHOSHONE                        	Y
// 	ID	TETON                           	Y
// 	ID	TWIN FALLS                      	Y
// 	ID	VALLEY                          	Y
// 	ID	WASHINGTON                      	Y
// 	IL	ADAMS                           	Y
// 	IL	ALEXANDER                       	Y
// 	IL	BOND                            	Y
// 	IL	BOONE                           	Y
// 	IL	BROWN                           	Y
// 	IL	BUREAU                          	Y
// 	IL	CALHOUN                         	Y
// 	IL	CARROLL                         	Y
// 	IL	CASS                            	Y
// 	IL	CHAMPAIGN                       	Y
// 	IL	CHRISTIAN                       	Y
// 	IL	CLARK                           	Y
// 	IL	CLAY                            	Y
// 	IL	CLINTON                         	Y
// 	IL	COLES                           	Y
// 	IL	COOK                            	Y
// 	IL	CRAWFORD                        	Y
// 	IL	CUMBERLAND                      	Y
// 	IL	DE WITT                         	Y
// 	IL	DEKALB                          	Y
// 	IL	DOUGLAS                         	Y
// 	IL	DUPAGE                          	Y
// 	IL	EDGAR                           	Y
// 	IL	EDWARDS                         	Y
// 	IL	EFFINGHAM                       	Y
// 	IL	FAYETTE                         	Y
// 	IL	FORD                            	Y
// 	IL	FRANKLIN                        	Y
// 	IL	FULTON                          	Y
// 	IL	GALLATIN                        	Y
// 	IL	GREENE                          	Y
// 	IL	GRUNDY                          	Y
// 	IL	HAMILTON                        	Y
// 	IL	HANCOCK                         	Y
// 	IL	HARDIN                          	Y
// 	IL	HENDERSON                       	Y
// 	IL	HENRY                           	Y
// 	IL	IROQUOIS                        	Y
// 	IL	JACKSON                         	Y
// 	IL	JASPER                          	Y
// 	IL	JEFFERSON                       	Y
// 	IL	JERSEY                          	Y
// 	IL	JO DAVIESS                      	Y
// 	IL	JOHNSON                         	Y
// 	IL	KANE                            	Y
// 	IL	KANKAKEE                        	Y
// 	IL	KENDALL                         	Y
// 	IL	KNOX                            	Y
// 	IL	LA SALLE                        	Y
// 	IL	LAKE                            	Y
// 	IL	LAWRENCE                        	Y
// 	IL	LEE                             	Y
// 	IL	LIVINGSTON                      	Y
// 	IL	LOGAN                           	Y
// 	IL	MACON                           	Y
// 	IL	MACOUPIN                        	Y
// 	IL	MADISON                         	Y
// 	IL	MARION                          	Y
// 	IL	MARSHALL                        	Y
// 	IL	MASON                           	Y
// 	IL	MASSAC                          	Y
// 	IL	MCDONOUGH                       	Y
// 	IL	MCHENRY                         	Y
// 	IL	MCLEAN                          	Y
// 	IL	MENARD                          	Y
// 	IL	MERCER                          	Y
// 	IL	MONROE                          	Y
// 	IL	MONTGOMERY                      	Y
// 	IL	MORGAN                          	Y
// 	IL	MOULTRIE                        	Y
// 	IL	OGLE                            	Y
// 	IL	PEORIA                          	Y
// 	IL	PERRY                           	Y
// 	IL	PIATT                           	Y
// 	IL	PIKE                            	Y
// 	IL	POPE                            	Y
// 	IL	PULASKI                         	Y
// 	IL	PUTNAM                          	Y
// 	IL	RANDOLPH                        	Y
// 	IL	RICHLAND                        	Y
// 	IL	ROCK ISLAND                     	Y
// 	IL	SALINE                          	Y
// 	IL	SANGAMON                        	Y
// 	IL	SCHUYLER                        	Y
// 	IL	SCOTT                           	Y
// 	IL	SHELBY                          	Y
// 	IL	ST. CLAIR                       	Y
// 	IL	STARK                           	Y
// 	IL	STEPHENSON                      	Y
// 	IL	TAZEWELL                        	Y
// 	IL	UNION                           	Y
// 	IL	VERMILION                       	Y
// 	IL	WABASH                          	Y
// 	IL	WARREN                          	Y
// 	IL	WASHINGTON                      	Y
// 	IL	WAYNE                           	Y
// 	IL	WHITE                           	Y
// 	IL	WHITESIDE                       	Y
// 	IL	WILL                            	Y
// 	IL	WILLIAMSON                      	Y
// 	IL	WINNEBAGO                       	Y
// 	IL	WOODFORD                        	Y
// 	IN	ADAMS                           	Y
// 	IN	ALLEN                           	Y
// 	IN	BARTHOLOMEW                     	Y
// 	IN	BENTON                          	Y
// 	IN	BLACKFORD                       	Y
// 	IN	BOONE                           	Y
// 	IN	BROWN                           	Y
// 	IN	CARROLL                         	Y
// 	IN	CASS                            	Y
// 	IN	CLARK                           	Y
// 	IN	CLAY                            	Y
// 	IN	CLINTON                         	Y
// 	IN	CRAWFORD                        	Y
// 	IN	DAVIESS                         	Y
// 	IN	DE KALB                         	Y
// 	IN	DEARBORN                        	Y
// 	IN	DECATUR                         	Y
// 	IN	DELAWARE                        	Y
// 	IN	DUBOIS                          	Y
// 	IN	ELKHART                         	Y
// 	IN	FAYETTE                         	Y
// 	IN	FLOYD                           	Y
// 	IN	FOUNTAIN                        	Y
// 	IN	FRANKLIN                        	Y
// 	IN	FULTON                          	Y
// 	IN	GIBSON                          	Y
// 	IN	GRANT                           	Y
// 	IN	GREENE                          	Y
// 	IN	HAMILTON                        	Y
// 	IN	HANCOCK                         	Y
// 	IN	HARRISON                        	Y
// 	IN	HENDRICKS                       	Y
// 	IN	HENRY                           	Y
// 	IN	HOWARD                          	Y
// 	IN	HUNTINGTON                      	Y
// 	IN	JACKSON                         	Y
// 	IN	JASPER                          	Y
// 	IN	JAY                             	Y
// 	IN	JEFFERSON                       	Y
// 	IN	JENNINGS                        	Y
// 	IN	JOHNSON                         	Y
// 	IN	KNOX                            	Y
// 	IN	KOSCIUSKO                       	Y
// 	IN	LA PORTE                        	Y
// 	IN	LAGRANGE                        	Y
// 	IN	LAKE                            	Y
// 	IN	LAWRENCE                        	Y
// 	IN	MADISON                         	Y
// 	IN	MARION                          	Y
// 	IN	MARSHALL                        	Y
// 	IN	MARTIN                          	Y
// 	IN	MIAMI                           	Y
// 	IN	MONROE                          	Y
// 	IN	MONTGOMERY                      	Y
// 	IN	MORGAN                          	Y
// 	IN	NEWTON                          	Y
// 	IN	NOBLE                           	Y
// 	IN	OHIO                            	Y
// 	IN	ORANGE                          	Y
// 	IN	OWEN                            	Y
// 	IN	PARKE                           	Y
// 	IN	PERRY                           	Y
// 	IN	PIKE                            	Y
// 	IN	PORTER                          	Y
// 	IN	POSEY                           	Y
// 	IN	PULASKI                         	Y
// 	IN	PUTNAM                          	Y
// 	IN	RANDOLPH                        	Y
// 	IN	RIPLEY                          	Y
// 	IN	RUSH                            	Y
// 	IN	SCOTT                           	Y
// 	IN	SHELBY                          	Y
// 	IN	SPENCER                         	Y
// 	IN	ST. JOSEPH                      	Y
// 	IN	STARKE                          	Y
// 	IN	STEUBEN                         	Y
// 	IN	SULLIVAN                        	Y
// 	IN	SWITZERLAND                     	Y
// 	IN	TIPPECANOE                      	Y
// 	IN	TIPTON                          	Y
// 	IN	UNION                           	Y
// 	IN	VANDERBURGH                     	Y
// 	IN	VERMILLION                      	Y
// 	IN	VIGO                            	Y
// 	IN	WABASH                          	Y
// 	IN	WARREN                          	Y
// 	IN	WARRICK                         	Y
// 	IN	WASHINGTON                      	Y
// 	IN	WAYNE                           	Y
// 	IN	WELLS                           	Y
// 	IN	WHITE                           	Y
// 	IN	WHITLEY                         	Y
// 	KS	ALLEN                           	Y
// 	KS	ANDERSON                        	Y
// 	KS	ATCHISON                        	Y
// 	KS	BARBER                          	Y
// 	KS	BARTON                          	Y
// 	KS	BOURBON                         	Y
// 	KS	BROWN                           	Y
// 	KS	BUTLER                          	Y
// 	KS	CHASE                           	Y
// 	KS	CHAUTAUQUA                      	Y
// 	KS	CHEROKEE                        	Y
// 	KS	CHEYENNE                        	Y
// 	KS	CLARK                           	Y
// 	KS	CLAY                            	Y
// 	KS	CLOUD                           	Y
// 	KS	COFFEY                          	Y
// 	KS	COMANCHE                        	Y
// 	KS	COWLEY                          	Y
// 	KS	CRAWFORD                        	Y
// 	KS	DECATUR                         	Y
// 	KS	DICKINSON                       	Y
// 	KS	DONIPHAN                        	Y
// 	KS	DOUGLAS                         	Y
// 	KS	EDWARDS                         	Y
// 	KS	ELK                             	Y
// 	KS	ELLIS                           	Y
// 	KS	ELLSWORTH                       	Y
// 	KS	FINNEY                          	Y
// 	KS	FORD                            	Y
// 	KS	FRANKLIN                        	Y
// 	KS	GEARY                           	Y
// 	KS	GOVE                            	Y
// 	KS	GRAHAM                          	Y
// 	KS	GRANT                           	Y
// 	KS	GRAY                            	Y
// 	KS	GREELEY                         	Y
// 	KS	GREENWOOD                       	Y
// 	KS	HAMILTON                        	Y
// 	KS	HARPER                          	Y
// 	KS	HARVEY                          	Y
// 	KS	HASKELL                         	Y
// 	KS	HODGEMAN                        	Y
// 	KS	JACKSON                         	Y
// 	KS	JEFFERSON                       	Y
// 	KS	JEWELL                          	Y
// 	KS	JOHNSON                         	Y
// 	KS	KEARNY                          	Y
// 	KS	KINGMAN                         	Y
// 	KS	KIOWA                           	Y
// 	KS	LABETTE                         	Y
// 	KS	LANE                            	Y
// 	KS	LEAVENWORTH                     	Y
// 	KS	LINCOLN                         	Y
// 	KS	LINN                            	Y
// 	KS	LOGAN                           	Y
// 	KS	LYON                            	Y
// 	KS	MARION                          	Y
// 	KS	MARSHALL                        	Y
// 	KS	MCPHERSON                       	Y
// 	KS	MEADE                           	Y
// 	KS	MIAMI                           	Y
// 	KS	MITCHELL                        	Y
// 	KS	MONTGOMERY                      	Y
// 	KS	MORRIS                          	Y
// 	KS	MORTON                          	Y
// 	KS	NEMAHA                          	Y
// 	KS	NEOSHO                          	Y
// 	KS	NESS                            	Y
// 	KS	NORTON                          	Y
// 	KS	OSAGE                           	Y
// 	KS	OSBORNE                         	Y
// 	KS	OTTAWA                          	Y
// 	KS	PAWNEE                          	Y
// 	KS	PHILLIPS                        	Y
// 	KS	POTTAWATOMIE                    	Y
// 	KS	PRATT                           	Y
// 	KS	RAWLINS                         	Y
// 	KS	RENO                            	Y
// 	KS	REPUBLIC                        	Y
// 	KS	RICE                            	Y
// 	KS	RILEY                           	Y
// 	KS	ROOKS                           	Y
// 	KS	RUSH                            	Y
// 	KS	RUSSELL                         	Y
// 	KS	SALINE                          	Y
// 	KS	SCOTT                           	Y
// 	KS	SEDGWICK                        	Y
// 	KS	SEWARD                          	Y
// 	KS	SHAWNEE                         	Y
// 	KS	SHERIDAN                        	Y
// 	KS	SHERMAN                         	Y
// 	KS	SMITH                           	Y
// 	KS	STAFFORD                        	Y
// 	KS	STANTON                         	Y
// 	KS	STEVENS                         	Y
// 	KS	SUMNER                          	Y
// 	KS	THOMAS                          	Y
// 	KS	TREGO                           	Y
// 	KS	WABAUNSEE                       	Y
// 	KS	WALLACE                         	Y
// 	KS	WASHINGTON                      	Y
// 	KS	WICHITA                         	Y
// 	KS	WILSON                          	Y
// 	KS	WOODSON                         	Y
// 	KS	WYANDOTTE                       	Y
// 	KY	ADAIR                           	Y
// 	KY	ALLEN                           	Y
// 	KY	ANDERSON                        	Y
// 	KY	BALLARD                         	Y
// 	KY	BARREN                          	Y
// 	KY	BATH                            	Y
// 	KY	BELL                            	Y
// 	KY	BOONE                           	Y
// 	KY	BOURBON                         	Y
// 	KY	BOYD                            	Y
// 	KY	BOYLE                           	Y
// 	KY	BRACKEN                         	Y
// 	KY	BREATHITT                       	Y
// 	KY	BRECKINRIDGE                    	Y
// 	KY	BULLITT                         	Y
// 	KY	BUTLER                          	Y
// 	KY	CALDWELL                        	Y
// 	KY	CALLOWAY                        	Y
// 	KY	CAMPBELL                        	Y
// 	KY	CARLISLE                        	Y
// 	KY	CARROLL                         	Y
// 	KY	CARTER                          	Y
// 	KY	CASEY                           	Y
// 	KY	CHRISTIAN                       	Y
// 	KY	CLARK                           	Y
// 	KY	CLAY                            	Y
// 	KY	CLINTON                         	Y
// 	KY	CRITTENDEN                      	Y
// 	KY	CUMBERLAND                      	Y
// 	KY	DAVIESS                         	Y
// 	KY	EDMONSON                        	Y
// 	KY	ELLIOTT                         	Y
// 	KY	ESTILL                          	Y
// 	KY	FAYETTE                         	Y
// 	KY	FLEMING                         	Y
// 	KY	FLOYD                           	Y
// 	KY	FRANKLIN                        	Y
// 	KY	FULTON                          	Y
// 	KY	GALLATIN                        	Y
// 	KY	GARRARD                         	Y
// 	KY	GRANT                           	Y
// 	KY	GRAVES                          	Y
// 	KY	GRAYSON                         	Y
// 	KY	GREEN                           	Y
// 	KY	GREENUP                         	Y
// 	KY	HANCOCK                         	Y
// 	KY	HARDIN                          	Y
// 	KY	HARLAN                          	Y
// 	KY	HARRISON                        	Y
// 	KY	HART                            	Y
// 	KY	HENDERSON                       	Y
// 	KY	HENRY                           	Y
// 	KY	HICKMAN                         	Y
// 	KY	HOPKINS                         	Y
// 	KY	JACKSON                         	Y
// 	KY	JEFFERSON                       	Y
// 	KY	JESSAMINE                       	Y
// 	KY	JOHNSON                         	Y
// 	KY	KENTON                          	Y
// 	KY	KNOTT                           	Y
// 	KY	KNOX                            	Y
// 	KY	LARUE                           	Y
// 	KY	LAUREL                          	Y
// 	KY	LAWRENCE                        	Y
// 	KY	LEE                             	Y
// 	KY	LESLIE                          	Y
// 	KY	LETCHER                         	Y
// 	KY	LEWIS                           	Y
// 	KY	LINCOLN                         	Y
// 	KY	LIVINGSTON                      	Y
// 	KY	LOGAN                           	Y
// 	KY	LYON                            	Y
// 	KY	MADISON                         	Y
// 	KY	MAGOFFIN                        	Y
// 	KY	MARION                          	Y
// 	KY	MARSHALL                        	Y
// 	KY	MARTIN                          	Y
// 	KY	MASON                           	Y
// 	KY	MCCRACKEN                       	Y
// 	KY	MCCREARY                        	Y
// 	KY	MCLEAN                          	Y
// 	KY	MEADE                           	Y
// 	KY	MENIFEE                         	Y
// 	KY	MERCER                          	Y
// 	KY	METCALFE                        	Y
// 	KY	MONROE                          	Y
// 	KY	MONTGOMERY                      	Y
// 	KY	MORGAN                          	Y
// 	KY	MUHLENBERG                      	Y
// 	KY	NELSON                          	Y
// 	KY	NICHOLAS                        	Y
// 	KY	OHIO                            	Y
// 	KY	OLDHAM                          	Y
// 	KY	OWEN                            	Y
// 	KY	OWSLEY                          	Y
// 	KY	PENDLETON                       	Y
// 	KY	PERRY                           	Y
// 	KY	PIKE                            	Y
// 	KY	POWELL                          	Y
// 	KY	PULASKI                         	Y
// 	KY	ROBERTSON                       	Y
// 	KY	ROCKCASTLE                      	Y
// 	KY	ROWAN                           	Y
// 	KY	RUSSELL                         	Y
// 	KY	SCOTT                           	Y
// 	KY	SHELBY                          	Y
// 	KY	SIMPSON                         	Y
// 	KY	SPENCER                         	Y
// 	KY	TAYLOR                          	Y
// 	KY	TODD                            	Y
// 	KY	TRIGG                           	Y
// 	KY	TRIMBLE                         	Y
// 	KY	UNION                           	Y
// 	KY	WARREN                          	Y
// 	KY	WASHINGTON                      	Y
// 	KY	WAYNE                           	Y
// 	KY	WEBSTER                         	Y
// 	KY	WHITLEY                         	Y
// 	KY	WOLFE                           	Y
// 	KY	WOODFORD                        	Y
// 	LA	ACADIA                          	Y
// 	LA	ALLEN                           	Y
// 	LA	ASCENSION                       	Y
// 	LA	ASSUMPTION                      	Y
// 	LA	AVOYELLES                       	Y
// 	LA	BEAUREGARD                      	Y
// 	LA	BIENVILLE                       	Y
// 	LA	BOSSIER                         	Y
// 	LA	CADDO                           	Y
// 	LA	CALCASIEU                       	Y
// 	LA	CALDWELL                        	Y
// 	LA	CAMERON                         	Y
// 	LA	CATAHOULA                       	Y
// 	LA	CLAIBORNE                       	Y
// 	LA	CONCORDIA                       	Y
// 	LA	DE SOTO                         	Y
// 	LA	EAST BATON ROUGE                	Y
// 	LA	EAST CARROLL                    	Y
// 	LA	EAST FELICIANA                  	Y
// 	LA	EVANGELINE                      	Y
// 	LA	FRANKLIN                        	Y
// 	LA	GRANT                           	Y
// 	LA	IBERIA                          	Y
// 	LA	IBERVILLE                       	Y
// 	LA	JACKSON                         	Y
// 	LA	JEFFERSON                       	Y
// 	LA	JEFFERSON DAVIS                 	Y
// 	LA	LA SALLE                        	Y
// 	LA	LAFAYETTE                       	Y
// 	LA	LAFOURCHE                       	Y
// 	LA	LINCOLN                         	Y
// 	LA	LIVINGSTON                      	Y
// 	LA	MADISON                         	Y
// 	LA	MOREHOUSE                       	Y
// 	LA	NATCHITOCHES                    	Y
// 	LA	ORLEANS                         	Y
// 	LA	OUACHITA                        	Y
// 	LA	PLAQUEMINES                     	Y
// 	LA	POINTE COUPEE                   	Y
// 	LA	RAPIDES                         	Y
// 	LA	RED RIVER                       	Y
// 	LA	RICHLAND                        	Y
// 	LA	SABINE                          	Y
// 	LA	ST. BERNARD                     	Y
// 	LA	ST. CHARLES                     	Y
// 	LA	ST. HELENA                      	Y
// 	LA	ST. JAMES                       	Y
// 	LA	ST. JOHN THE BAPTIST            	Y
// 	LA	ST. LANDRY                      	Y
// 	LA	ST. MARTIN                      	Y
// 	LA	ST. MARY                        	Y
// 	LA	ST. TAMMANY                     	Y
// 	LA	TANGIPAHOA                      	Y
// 	LA	TENSAS                          	Y
// 	LA	TERREBONNE                      	Y
// 	LA	UNION                           	Y
// 	LA	VERMILION                       	Y
// 	LA	VERNON                          	Y
// 	LA	WASHINGTON                      	Y
// 	LA	WEBSTER                         	Y
// 	LA	WEST BATON ROUGE                	Y
// 	LA	WEST CARROLL                    	Y
// 	LA	WEST FELICIANA                  	Y
// 	LA	WINN                            	Y
// 	MA	BARNSTABLE                      	Y
// 	MA	BERKSHIRE                       	Y
// 	MA	BRISTOL                         	Y
// 	MA	DUKES                           	Y
// 	MA	ESSEX                           	Y
// 	MA	FRANKLIN                        	Y
// 	MA	HAMPDEN                         	Y
// 	MA	HAMPSHIRE                       	Y
// 	MA	MIDDLESEX                       	Y
// 	MA	NANTUCKET                       	Y
// 	MA	NORFOLK                         	Y
// 	MA	PLYMOUTH                        	Y
// 	MA	SUFFOLK                         	Y
// 	MA	WORCESTER                       	Y
// 	MD	ALLEGANY                        	Y
// 	MD	ANNE ARUNDEL                    	Y
// 	MD	BALTIMORE                       	Y
// 	MD	BALTIMORE CITY                  	Y
// 	MD	CALVERT                         	Y
// 	MD	CAROLINE                        	Y
// 	MD	CARROLL                         	Y
// 	MD	CECIL                           	Y
// 	MD	CHARLES                         	Y
// 	MD	DORCHESTER                      	Y
// 	MD	FREDERICK                       	Y
// 	MD	GARRETT                         	Y
// 	MD	HARFORD                         	Y
// 	MD	HOWARD                          	Y
// 	MD	KENT                            	Y
// 	MD	MONTGOMERY                      	Y
// 	MD	PRINCE GEORGE'S                 	Y
// 	MD	QUEEN ANNE'S                    	Y
// 	MD	SOMERSET                        	Y
// 	MD	ST. MARY'S                      	Y
// 	MD	TALBOT                          	Y
// 	MD	WASHINGTON                      	Y
// 	MD	WICOMICO                        	Y
// 	MD	WORCESTER                       	Y
// 	ME	ANDROSCOGGIN                    	Y
// 	ME	AROOSTOOK                       	Y
// 	ME	CUMBERLAND                      	Y
// 	ME	FRANKLIN                        	Y
// 	ME	HANCOCK                         	Y
// 	ME	KENNEBEC                        	Y
// 	ME	KNOX                            	Y
// 	ME	LINCOLN                         	Y
// 	ME	OXFORD                          	Y
// 	ME	PENOBSCOT                       	Y
// 	ME	PISCATAQUIS                     	Y
// 	ME	SAGADAHOC                       	Y
// 	ME	SOMERSET                        	Y
// 	ME	WALDO                           	Y
// 	ME	WASHINGTON                      	Y
// 	ME	YORK                            	Y
// 	MI	ALCONA                          	Y
// 	MI	ALGER                           	Y
// 	MI	ALLEGAN                         	Y
// 	MI	ALPENA                          	Y
// 	MI	ANTRIM                          	Y
// 	MI	ARENAC                          	Y
// 	MI	BARAGA                          	Y
// 	MI	BARRY                           	Y
// 	MI	BAY                             	Y
// 	MI	BENZIE                          	Y
// 	MI	BERRIEN                         	Y
// 	MI	BRANCH                          	Y
// 	MI	CALHOUN                         	Y
// 	MI	CASS                            	Y
// 	MI	CHARLEVOIX                      	Y
// 	MI	CHEBOYGAN                       	Y
// 	MI	CHIPPEWA                        	Y
// 	MI	CLARE                           	Y
// 	MI	CLINTON                         	Y
// 	MI	CRAWFORD                        	Y
// 	MI	DELTA                           	Y
// 	MI	DICKINSON                       	Y
// 	MI	EATON                           	Y
// 	MI	EMMET                           	Y
// 	MI	GENESEE                         	Y
// 	MI	GLADWIN                         	Y
// 	MI	GOGEBIC                         	Y
// 	MI	GRAND TRAVERSE                  	Y
// 	MI	GRATIOT                         	Y
// 	MI	HILLSDALE                       	Y
// 	MI	HOUGHTON                        	Y
// 	MI	HURON                           	Y
// 	MI	INGHAM                          	Y
// 	MI	IONIA                           	Y
// 	MI	IOSCO                           	Y
// 	MI	IRON                            	Y
// 	MI	ISABELLA                        	Y
// 	MI	JACKSON                         	Y
// 	MI	KALAMAZOO                       	Y
// 	MI	KALKASKA                        	Y
// 	MI	KENT                            	Y
// 	MI	KEWEENAW                        	Y
// 	MI	LAKE                            	Y
// 	MI	LAPEER                          	Y
// 	MI	LEELANAU                        	Y
// 	MI	LENAWEE                         	Y
// 	MI	LIVINGSTON                      	Y
// 	MI	LUCE                            	Y
// 	MI	MACKINAC                        	Y
// 	MI	MACOMB                          	Y
// 	MI	MANISTEE                        	Y
// 	MI	MARQUETTE                       	Y
// 	MI	MASON                           	Y
// 	MI	MECOSTA                         	Y
// 	MI	MENOMINEE                       	Y
// 	MI	MIDLAND                         	Y
// 	MI	MISSAUKEE                       	Y
// 	MI	MONROE                          	Y
// 	MI	MONTCALM                        	Y
// 	MI	MONTMORENCY                     	Y
// 	MI	MUSKEGON                        	Y
// 	MI	NEWAYGO                         	Y
// 	MI	OAKLAND                         	Y
// 	MI	OCEANA                          	Y
// 	MI	OGEMAW                          	Y
// 	MI	ONTONAGON                       	Y
// 	MI	OSCEOLA                         	Y
// 	MI	OSCODA                          	Y
// 	MI	OTSEGO                          	Y
// 	MI	OTTAWA                          	Y
// 	MI	PRESQUE ISLE                    	Y
// 	MI	ROSCOMMON                       	Y
// 	MI	SAGINAW                         	Y
// 	MI	SANILAC                         	Y
// 	MI	SCHOOLCRAFT                     	Y
// 	MI	SHIAWASSEE                      	Y
// 	MI	ST. CLAIR                       	Y
// 	MI	ST. JOSEPH                      	Y
// 	MI	TUSCOLA                         	Y
// 	MI	VAN BUREN                       	Y
// 	MI	WASHTENAW                       	Y
// 	MI	WAYNE                           	Y
// 	MI	WEXFORD                         	Y
// 	MN	AITKIN                          	Y
// 	MN	ANOKA                           	Y
// 	MN	BECKER                          	Y
// 	MN	BELTRAMI                        	Y
// 	MN	BENTON                          	Y
// 	MN	BIG STONE                       	Y
// 	MN	BLUE EARTH                      	Y
// 	MN	BROWN                           	Y
// 	MN	CARLTON                         	Y
// 	MN	CARVER                          	Y
// 	MN	CASS                            	Y
// 	MN	CHIPPEWA                        	Y
// 	MN	CHISAGO                         	Y
// 	MN	CLAY                            	Y
// 	MN	CLEARWATER                      	Y
// 	MN	COOK                            	Y
// 	MN	COTTONWOOD                      	Y
// 	MN	CROW WING                       	Y
// 	MN	DAKOTA                          	Y
// 	MN	DODGE                           	Y
// 	MN	DOUGLAS                         	Y
// 	MN	FARIBAULT                       	Y
// 	MN	FILLMORE                        	Y
// 	MN	FREEBORN                        	Y
// 	MN	GOODHUE                         	Y
// 	MN	GRANT                           	Y
// 	MN	HENNEPIN                        	Y
// 	MN	HOUSTON                         	Y
// 	MN	HUBBARD                         	Y
// 	MN	ISANTI                          	Y
// 	MN	ITASCA                          	Y
// 	MN	JACKSON                         	Y
// 	MN	KANABEC                         	Y
// 	MN	KANDIYOHI                       	Y
// 	MN	KITTSON                         	Y
// 	MN	KOOCHICHING                     	Y
// 	MN	LAC QUI PARLE                   	Y
// 	MN	LAKE                            	Y
// 	MN	LAKE OF THE WOODS               	Y
// 	MN	LE SUEUR                        	Y
// 	MN	LINCOLN                         	Y
// 	MN	LYON                            	Y
// 	MN	MAHNOMEN                        	Y
// 	MN	MARSHALL                        	Y
// 	MN	MARTIN                          	Y
// 	MN	MCLEOD                          	Y
// 	MN	MEEKER                          	Y
// 	MN	MILLE LACS                      	Y
// 	MN	MORRISON                        	Y
// 	MN	MOWER                           	Y
// 	MN	MURRAY                          	Y
// 	MN	NICOLLET                        	Y
// 	MN	NOBLES                          	Y
// 	MN	NORMAN                          	Y
// 	MN	OLMSTED                         	Y
// 	MN	OTTER TAIL                      	Y
// 	MN	PENNINGTON                      	Y
// 	MN	PINE                            	Y
// 	MN	PIPESTONE                       	Y
// 	MN	POLK                            	Y
// 	MN	POPE                            	Y
// 	MN	RAMSEY                          	Y
// 	MN	RED LAKE                        	Y
// 	MN	REDWOOD                         	Y
// 	MN	RENVILLE                        	Y
// 	MN	RICE                            	Y
// 	MN	ROCK                            	Y
// 	MN	ROSEAU                          	Y
// 	MN	SCOTT                           	Y
// 	MN	SHERBURNE                       	Y
// 	MN	SIBLEY                          	Y
// 	MN	ST. LOUIS                       	Y
// 	MN	STEARNS                         	Y
// 	MN	STEELE                          	Y
// 	MN	STEVENS                         	Y
// 	MN	SWIFT                           	Y
// 	MN	TODD                            	Y
// 	MN	TRAVERSE                        	Y
// 	MN	WABASHA                         	Y
// 	MN	WADENA                          	Y
// 	MN	WASECA                          	Y
// 	MN	WASHINGTON                      	Y
// 	MN	WATONWAN                        	Y
// 	MN	WILKIN                          	Y
// 	MN	WINONA                          	Y
// 	MN	WRIGHT                          	Y
// 	MN	YELLOW MEDICINE                 	Y
// 	MO	ADAIR                           	Y
// 	MO	ANDREW                          	Y
// 	MO	ATCHISON                        	Y
// 	MO	AUDRAIN                         	Y
// 	MO	BARRY                           	Y
// 	MO	BARTON                          	Y
// 	MO	BATES                           	Y
// 	MO	BENTON                          	Y
// 	MO	BOLLINGER                       	Y
// 	MO	BOONE                           	Y
// 	MO	BUCHANAN                        	Y
// 	MO	BUTLER                          	Y
// 	MO	CALDWELL                        	Y
// 	MO	CALLAWAY                        	Y
// 	MO	CAMDEN                          	Y
// 	MO	CAPE GIRARDEAU                  	Y
// 	MO	CARROLL                         	Y
// 	MO	CARTER                          	Y
// 	MO	CASS                            	Y
// 	MO	CEDAR                           	Y
// 	MO	CHARITON                        	Y
// 	MO	CHRISTIAN                       	Y
// 	MO	CLARK                           	Y
// 	MO	CLAY                            	Y
// 	MO	CLINTON                         	Y
// 	MO	COLE                            	Y
// 	MO	COOPER                          	Y
// 	MO	CRAWFORD                        	Y
// 	MO	DADE                            	Y
// 	MO	DALLAS                          	Y
// 	MO	DAVIESS                         	Y
// 	MO	DEKALB                          	Y
// 	MO	DENT                            	Y
// 	MO	DOUGLAS                         	Y
// 	MO	DUNKLIN                         	Y
// 	MO	FRANKLIN                        	Y
// 	MO	GASCONADE                       	Y
// 	MO	GENTRY                          	Y
// 	MO	GREENE                          	Y
// 	MO	GRUNDY                          	Y
// 	MO	HARRISON                        	Y
// 	MO	HENRY                           	Y
// 	MO	HICKORY                         	Y
// 	MO	HOLT                            	Y
// 	MO	HOWARD                          	Y
// 	MO	HOWELL                          	Y
// 	MO	IRON                            	Y
// 	MO	JACKSON                         	Y
// 	MO	JASPER                          	Y
// 	MO	JEFFERSON                       	Y
// 	MO	JOHNSON                         	Y
// 	MO	KNOX                            	Y
// 	MO	LACLEDE                         	Y
// 	MO	LAFAYETTE                       	Y
// 	MO	LAWRENCE                        	Y
// 	MO	LEWIS                           	Y
// 	MO	LINCOLN                         	Y
// 	MO	LINN                            	Y
// 	MO	LIVINGSTON                      	Y
// 	MO	MACON                           	Y
// 	MO	MADISON                         	Y
// 	MO	MARIES                          	Y
// 	MO	MARION                          	Y
// 	MO	MCDONALD                        	Y
// 	MO	MERCER                          	Y
// 	MO	MILLER                          	Y
// 	MO	MISSISSIPPI                     	Y
// 	MO	MONITEAU                        	Y
// 	MO	MONROE                          	Y
// 	MO	MONTGOMERY                      	Y
// 	MO	MORGAN                          	Y
// 	MO	NEW MADRID                      	Y
// 	MO	NEWTON                          	Y
// 	MO	NODAWAY                         	Y
// 	MO	OREGON                          	Y
// 	MO	OSAGE                           	Y
// 	MO	OZARK                           	Y
// 	MO	PEMISCOT                        	Y
// 	MO	PERRY                           	Y
// 	MO	PETTIS                          	Y
// 	MO	PHELPS                          	Y
// 	MO	PIKE                            	Y
// 	MO	PLATTE                          	Y
// 	MO	POLK                            	Y
// 	MO	PULASKI                         	Y
// 	MO	PUTNAM                          	Y
// 	MO	RALLS                           	Y
// 	MO	RANDOLPH                        	Y
// 	MO	RAY                             	Y
// 	MO	REYNOLDS                        	Y
// 	MO	RIPLEY                          	Y
// 	MO	SALINE                          	Y
// 	MO	SCHUYLER                        	Y
// 	MO	SCOTLAND                        	Y
// 	MO	SCOTT                           	Y
// 	MO	SHANNON                         	Y
// 	MO	SHELBY                          	Y
// 	MO	ST. CHARLES                     	Y
// 	MO	ST. CLAIR                       	Y
// 	MO	ST. FRANCOIS                    	Y
// 	MO	ST. LOUIS                       	Y
// 	MO	ST. LOUIS CITY                  	Y
// 	MO	STE. GENEVIEVE                  	Y
// 	MO	STODDARD                        	Y
// 	MO	STONE                           	Y
// 	MO	SULLIVAN                        	Y
// 	MO	TANEY                           	Y
// 	MO	TEXAS                           	Y
// 	MO	VERNON                          	Y
// 	MO	WARREN                          	Y
// 	MO	WASHINGTON                      	Y
// 	MO	WAYNE                           	Y
// 	MO	WEBSTER                         	Y
// 	MO	WORTH                           	Y
// 	MO	WRIGHT                          	Y
// 	MP	NORTHERN ISLANDS                	Y
// 	MP	ROTA                            	Y
// 	MP	SAIPAN                          	Y
// 	MP	TINIAN                          	Y
// 	MS	ADAMS                           	Y
// 	MS	ALCORN                          	Y
// 	MS	AMITE                           	Y
// 	MS	ATTALA                          	Y
// 	MS	BENTON                          	Y
// 	MS	BOLIVAR                         	Y
// 	MS	CALHOUN                         	Y
// 	MS	CARROLL                         	Y
// 	MS	CHICKASAW                       	Y
// 	MS	CHOCTAW                         	Y
// 	MS	CLAIBORNE                       	Y
// 	MS	CLARKE                          	Y
// 	MS	CLAY                            	Y
// 	MS	COAHOMA                         	Y
// 	MS	COPIAH                          	Y
// 	MS	COVINGTON                       	Y
// 	MS	DESOTO                          	Y
// 	MS	FORREST                         	Y
// 	MS	FRANKLIN                        	Y
// 	MS	GEORGE                          	Y
// 	MS	GREENE                          	Y
// 	MS	GRENADA                         	Y
// 	MS	HANCOCK                         	Y
// 	MS	HARRISON                        	Y
// 	MS	HINDS                           	Y
// 	MS	HOLMES                          	Y
// 	MS	HUMPHREYS                       	Y
// 	MS	ISSAQUENA                       	Y
// 	MS	ITAWAMBA                        	Y
// 	MS	JACKSON                         	Y
// 	MS	JASPER                          	Y
// 	MS	JEFFERSON                       	Y
// 	MS	JEFFERSON DAVIS                 	Y
// 	MS	JONES                           	Y
// 	MS	KEMPER                          	Y
// 	MS	LAFAYETTE                       	Y
// 	MS	LAMAR                           	Y
// 	MS	LAUDERDALE                      	Y
// 	MS	LAWRENCE                        	Y
// 	MS	LEAKE                           	Y
// 	MS	LEE                             	Y
// 	MS	LEFLORE                         	Y
// 	MS	LINCOLN                         	Y
// 	MS	LOWNDES                         	Y
// 	MS	MADISON                         	Y
// 	MS	MARION                          	Y
// 	MS	MARSHALL                        	Y
// 	MS	MONROE                          	Y
// 	MS	MONTGOMERY                      	Y
// 	MS	NESHOBA                         	Y
// 	MS	NEWTON                          	Y
// 	MS	NOXUBEE                         	Y
// 	MS	OKTIBBEHA                       	Y
// 	MS	PANOLA                          	Y
// 	MS	PEARL RIVER                     	Y
// 	MS	PERRY                           	Y
// 	MS	PIKE                            	Y
// 	MS	PONTOTOC                        	Y
// 	MS	PRENTISS                        	Y
// 	MS	QUITMAN                         	Y
// 	MS	RANKIN                          	Y
// 	MS	SCOTT                           	Y
// 	MS	SHARKEY                         	Y
// 	MS	SIMPSON                         	Y
// 	MS	SMITH                           	Y
// 	MS	STONE                           	Y
// 	MS	SUNFLOWER                       	Y
// 	MS	TALLAHATCHIE                    	Y
// 	MS	TATE                            	Y
// 	MS	TIPPAH                          	Y
// 	MS	TISHOMINGO                      	Y
// 	MS	TUNICA                          	Y
// 	MS	UNION                           	Y
// 	MS	WALTHALL                        	Y
// 	MS	WARREN                          	Y
// 	MS	WASHINGTON                      	Y
// 	MS	WAYNE                           	Y
// 	MS	WEBSTER                         	Y
// 	MS	WILKINSON                       	Y
// 	MS	WINSTON                         	Y
// 	MS	YALOBUSHA                       	Y
// 	MS	YAZOO                           	Y
// 	MT	BEAVERHEAD                      	Y
// 	MT	BIG HORN                        	Y
// 	MT	BLAINE                          	Y
// 	MT	BROADWATER                      	Y
// 	MT	CARBON                          	Y
// 	MT	CARTER                          	Y
// 	MT	CASCADE                         	Y
// 	MT	CHOUTEAU                        	Y
// 	MT	CUSTER                          	Y
// 	MT	DANIELS                         	Y
// 	MT	DAWSON                          	Y
// 	MT	DEER LODGE                      	Y
// 	MT	FALLON                          	Y
// 	MT	FERGUS                          	Y
// 	MT	FLATHEAD                        	Y
// 	MT	GALLATIN                        	Y
// 	MT	GARFIELD                        	Y
// 	MT	GLACIER                         	Y
// 	MT	GOLDEN VALLEY                   	Y
// 	MT	GRANITE                         	Y
// 	MT	HILL                            	Y
// 	MT	JEFFERSON                       	Y
// 	MT	JUDITH BASIN                    	Y
// 	MT	LAKE                            	Y
// 	MT	LEWIS AND CLARK                 	Y
// 	MT	LIBERTY                         	Y
// 	MT	LINCOLN                         	Y
// 	MT	MADISON                         	Y
// 	MT	MCCONE                          	Y
// 	MT	MEAGHER                         	Y
// 	MT	MINERAL                         	Y
// 	MT	MISSOULA                        	Y
// 	MT	MUSSELSHELL                     	Y
// 	MT	PARK                            	Y
// 	MT	PETROLEUM                       	Y
// 	MT	PHILLIPS                        	Y
// 	MT	PONDERA                         	Y
// 	MT	POWDER RIVER                    	Y
// 	MT	POWELL                          	Y
// 	MT	PRAIRIE                         	Y
// 	MT	RAVALLI                         	Y
// 	MT	RICHLAND                        	Y
// 	MT	ROOSEVELT                       	Y
// 	MT	ROSEBUD                         	Y
// 	MT	SANDERS                         	Y
// 	MT	SHERIDAN                        	Y
// 	MT	SILVER BOW                      	Y
// 	MT	STILLWATER                      	Y
// 	MT	SWEET GRASS                     	Y
// 	MT	TETON                           	Y
// 	MT	TOOLE                           	Y
// 	MT	TREASURE                        	Y
// 	MT	VALLEY                          	Y
// 	MT	WHEATLAND                       	Y
// 	MT	WIBAUX                          	Y
// 	MT	YELLOWSTONE                     	Y
// 	MT	YELLOWSTONE NATIONAL PARK       	N
// 	NC	ALAMANCE                        	Y
// 	NC	ALEXANDER                       	Y
// 	NC	ALLEGHANY                       	Y
// 	NC	ANSON                           	Y
// 	NC	ASHE                            	Y
// 	NC	AVERY                           	Y
// 	NC	BEAUFORT                        	Y
// 	NC	BERTIE                          	Y
// 	NC	BLADEN                          	Y
// 	NC	BRUNSWICK                       	Y
// 	NC	BUNCOMBE                        	Y
// 	NC	BURKE                           	Y
// 	NC	CABARRUS                        	Y
// 	NC	CALDWELL                        	Y
// 	NC	CAMDEN                          	Y
// 	NC	CARTERET                        	Y
// 	NC	CASWELL                         	Y
// 	NC	CATAWBA                         	Y
// 	NC	CHATHAM                         	Y
// 	NC	CHEROKEE                        	Y
// 	NC	CHOWAN                          	Y
// 	NC	CLAY                            	Y
// 	NC	CLEVELAND                       	Y
// 	NC	COLUMBUS                        	Y
// 	NC	CRAVEN                          	Y
// 	NC	CUMBERLAND                      	Y
// 	NC	CURRITUCK                       	Y
// 	NC	DARE                            	Y
// 	NC	DAVIDSON                        	Y
// 	NC	DAVIE                           	Y
// 	NC	DUPLIN                          	Y
// 	NC	DURHAM                          	Y
// 	NC	EDGECOMBE                       	Y
// 	NC	FORSYTH                         	Y
// 	NC	FRANKLIN                        	Y
// 	NC	GASTON                          	Y
// 	NC	GATES                           	Y
// 	NC	GRAHAM                          	Y
// 	NC	GRANVILLE                       	Y
// 	NC	GREENE                          	Y
// 	NC	GUILFORD                        	Y
// 	NC	HALIFAX                         	Y
// 	NC	HARNETT                         	Y
// 	NC	HAYWOOD                         	Y
// 	NC	HENDERSON                       	Y
// 	NC	HERTFORD                        	Y
// 	NC	HOKE                            	Y
// 	NC	HYDE                            	Y
// 	NC	IREDELL                         	Y
// 	NC	JACKSON                         	Y
// 	NC	JOHNSTON                        	Y
// 	NC	JONES                           	Y
// 	NC	LEE                             	Y
// 	NC	LENOIR                          	Y
// 	NC	LINCOLN                         	Y
// 	NC	MACON                           	Y
// 	NC	MADISON                         	Y
// 	NC	MARTIN                          	Y
// 	NC	MCDOWELL                        	Y
// 	NC	MECKLENBURG                     	Y
// 	NC	MITCHELL                        	Y
// 	NC	MONTGOMERY                      	Y
// 	NC	MOORE                           	Y
// 	NC	NASH                            	Y
// 	NC	NEW HANOVER                     	Y
// 	NC	NORTHAMPTON                     	Y
// 	NC	ONSLOW                          	Y
// 	NC	ORANGE                          	Y
// 	NC	PAMLICO                         	Y
// 	NC	PASQUOTANK                      	Y
// 	NC	PENDER                          	Y
// 	NC	PERQUIMANS                      	Y
// 	NC	PERSON                          	Y
// 	NC	PITT                            	Y
// 	NC	POLK                            	Y
// 	NC	RANDOLPH                        	Y
// 	NC	RICHMOND                        	Y
// 	NC	ROBESON                         	Y
// 	NC	ROCKINGHAM                      	Y
// 	NC	ROWAN                           	Y
// 	NC	RUTHERFORD                      	Y
// 	NC	SAMPSON                         	Y
// 	NC	SCOTLAND                        	Y
// 	NC	STANLY                          	Y
// 	NC	STOKES                          	Y
// 	NC	SURRY                           	Y
// 	NC	SWAIN                           	Y
// 	NC	TRANSYLVANIA                    	Y
// 	NC	TYRRELL                         	Y
// 	NC	UNION                           	Y
// 	NC	VANCE                           	Y
// 	NC	WAKE                            	Y
// 	NC	WARREN                          	Y
// 	NC	WASHINGTON                      	Y
// 	NC	WATAUGA                         	Y
// 	NC	WAYNE                           	Y
// 	NC	WILKES                          	Y
// 	NC	WILSON                          	Y
// 	NC	YADKIN                          	Y
// 	NC	YANCEY                          	Y
// 	ND	ADAMS                           	Y
// 	ND	BARNES                          	Y
// 	ND	BENSON                          	Y
// 	ND	BILLINGS                        	Y
// 	ND	BOTTINEAU                       	Y
// 	ND	BOWMAN                          	Y
// 	ND	BURKE                           	Y
// 	ND	BURLEIGH                        	Y
// 	ND	CASS                            	Y
// 	ND	CAVALIER                        	Y
// 	ND	DICKEY                          	Y
// 	ND	DIVIDE                          	Y
// 	ND	DUNN                            	Y
// 	ND	EDDY                            	Y
// 	ND	EMMONS                          	Y
// 	ND	FOSTER                          	Y
// 	ND	GOLDEN VALLEY                   	Y
// 	ND	GRAND FORKS                     	Y
// 	ND	GRANT                           	Y
// 	ND	GRIGGS                          	Y
// 	ND	HETTINGER                       	Y
// 	ND	KIDDER                          	Y
// 	ND	LAMOURE                         	Y
// 	ND	LOGAN                           	Y
// 	ND	MCHENRY                         	Y
// 	ND	MCINTOSH                        	Y
// 	ND	MCKENZIE                        	Y
// 	ND	MCLEAN                          	Y
// 	ND	MERCER                          	Y
// 	ND	MORTON                          	Y
// 	ND	MOUNTRAIL                       	Y
// 	ND	NELSON                          	Y
// 	ND	OLIVER                          	Y
// 	ND	PEMBINA                         	Y
// 	ND	PIERCE                          	Y
// 	ND	RAMSEY                          	Y
// 	ND	RANSOM                          	Y
// 	ND	RENVILLE                        	Y
// 	ND	RICHLAND                        	Y
// 	ND	ROLETTE                         	Y
// 	ND	SARGENT                         	Y
// 	ND	SHERIDAN                        	Y
// 	ND	SIOUX                           	Y
// 	ND	SLOPE                           	Y
// 	ND	STARK                           	Y
// 	ND	STEELE                          	Y
// 	ND	STUTSMAN                        	Y
// 	ND	TOWNER                          	Y
// 	ND	TRAILL                          	Y
// 	ND	WALSH                           	Y
// 	ND	WARD                            	Y
// 	ND	WELLS                           	Y
// 	ND	WILLIAMS                        	Y
// 	NE	ADAMS                           	Y
// 	NE	ANTELOPE                        	Y
// 	NE	ARTHUR                          	Y
// 	NE	BANNER                          	Y
// 	NE	BLAINE                          	Y
// 	NE	BOONE                           	Y
// 	NE	BOX BUTTE                       	Y
// 	NE	BOYD                            	Y
// 	NE	BROWN                           	Y
// 	NE	BUFFALO                         	Y
// 	NE	BURT                            	Y
// 	NE	BUTLER                          	Y
// 	NE	CASS                            	Y
// 	NE	CEDAR                           	Y
// 	NE	CHASE                           	Y
// 	NE	CHERRY                          	Y
// 	NE	CHEYENNE                        	Y
// 	NE	CLAY                            	Y
// 	NE	COLFAX                          	Y
// 	NE	CUMING                          	Y
// 	NE	CUSTER                          	Y
// 	NE	DAKOTA                          	Y
// 	NE	DAWES                           	Y
// 	NE	DAWSON                          	Y
// 	NE	DEUEL                           	Y
// 	NE	DIXON                           	Y
// 	NE	DODGE                           	Y
// 	NE	DOUGLAS                         	Y
// 	NE	DUNDY                           	Y
// 	NE	FILLMORE                        	Y
// 	NE	FRANKLIN                        	Y
// 	NE	FRONTIER                        	Y
// 	NE	FURNAS                          	Y
// 	NE	GAGE                            	Y
// 	NE	GARDEN                          	Y
// 	NE	GARFIELD                        	Y
// 	NE	GOSPER                          	Y
// 	NE	GRANT                           	Y
// 	NE	GREELEY                         	Y
// 	NE	HALL                            	Y
// 	NE	HAMILTON                        	Y
// 	NE	HARLAN                          	Y
// 	NE	HAYES                           	Y
// 	NE	HITCHCOCK                       	Y
// 	NE	HOLT                            	Y
// 	NE	HOOKER                          	Y
// 	NE	HOWARD                          	Y
// 	NE	JEFFERSON                       	Y
// 	NE	JOHNSON                         	Y
// 	NE	KEARNEY                         	Y
// 	NE	KEITH                           	Y
// 	NE	KEYA PAHA                       	Y
// 	NE	KIMBALL                         	Y
// 	NE	KNOX                            	Y
// 	NE	LANCASTER                       	Y
// 	NE	LINCOLN                         	Y
// 	NE	LOGAN                           	Y
// 	NE	LOUP                            	Y
// 	NE	MADISON                         	Y
// 	NE	MCPHERSON                       	Y
// 	NE	MERRICK                         	Y
// 	NE	MORRILL                         	Y
// 	NE	NANCE                           	Y
// 	NE	NEMAHA                          	Y
// 	NE	NUCKOLLS                        	Y
// 	NE	OTOE                            	Y
// 	NE	PAWNEE                          	Y
// 	NE	PERKINS                         	Y
// 	NE	PHELPS                          	Y
// 	NE	PIERCE                          	Y
// 	NE	PLATTE                          	Y
// 	NE	POLK                            	Y
// 	NE	RED WILLOW                      	Y
// 	NE	RICHARDSON                      	Y
// 	NE	ROCK                            	Y
// 	NE	SALINE                          	Y
// 	NE	SARPY                           	Y
// 	NE	SAUNDERS                        	Y
// 	NE	SCOTTS BLUFF                    	Y
// 	NE	SEWARD                          	Y
// 	NE	SHERIDAN                        	Y
// 	NE	SHERMAN                         	Y
// 	NE	SIOUX                           	Y
// 	NE	STANTON                         	Y
// 	NE	THAYER                          	Y
// 	NE	THOMAS                          	Y
// 	NE	THURSTON                        	Y
// 	NE	VALLEY                          	Y
// 	NE	WASHINGTON                      	Y
// 	NE	WAYNE                           	Y
// 	NE	WEBSTER                         	Y
// 	NE	WHEELER                         	Y
// 	NE	YORK                            	Y
// 	NH	BELKNAP                         	Y
// 	NH	CARROLL                         	Y
// 	NH	CHESHIRE                        	Y
// 	NH	COOS                            	Y
// 	NH	GRAFTON                         	Y
// 	NH	HILLSBOROUGH                    	Y
// 	NH	MERRIMACK                       	Y
// 	NH	ROCKINGHAM                      	Y
// 	NH	STRAFFORD                       	Y
// 	NH	SULLIVAN                        	Y
// 	NJ	ATLANTIC                        	Y
// 	NJ	BERGEN                          	Y
// 	NJ	BURLINGTON                      	Y
// 	NJ	CAMDEN                          	Y
// 	NJ	CAPE MAY                        	Y
// 	NJ	CUMBERLAND                      	Y
// 	NJ	ESSEX                           	Y
// 	NJ	GLOUCESTER                      	Y
// 	NJ	HUDSON                          	Y
// 	NJ	HUNTERDON                       	Y
// 	NJ	MERCER                          	Y
// 	NJ	MIDDLESEX                       	Y
// 	NJ	MONMOUTH                        	Y
// 	NJ	MORRIS                          	Y
// 	NJ	OCEAN                           	Y
// 	NJ	PASSAIC                         	Y
// 	NJ	SALEM                           	Y
// 	NJ	SOMERSET                        	Y
// 	NJ	SUSSEX                          	Y
// 	NJ	UNION                           	Y
// 	NJ	WARREN                          	Y
// 	NM	BERNALILLO                      	Y
// 	NM	CATRON                          	Y
// 	NM	CHAVES                          	Y
// 	NM	CIBOLA                          	Y
// 	NM	COLFAX                          	Y
// 	NM	CURRY                           	Y
// 	NM	DE BACA                         	Y
// 	NM	DONA ANA                        	Y
// 	NM	EDDY                            	Y
// 	NM	GRANT                           	Y
// 	NM	GUADALUPE                       	Y
// 	NM	HARDING                         	Y
// 	NM	HIDALGO                         	Y
// 	NM	LEA                             	Y
// 	NM	LINCOLN                         	Y
// 	NM	LOS ALAMOS                      	Y
// 	NM	LUNA                            	Y
// 	NM	MCKINLEY                        	Y
// 	NM	MORA                            	Y
// 	NM	OTERO                           	Y
// 	NM	QUAY                            	Y
// 	NM	RIO ARRIBA                      	Y
// 	NM	ROOSEVELT                       	Y
// 	NM	SAN JUAN                        	Y
// 	NM	SAN MIGUEL                      	Y
// 	NM	SANDOVAL                        	Y
// 	NM	SANTA FE                        	Y
// 	NM	SIERRA                          	Y
// 	NM	SOCORRO                         	Y
// 	NM	TAOS                            	Y
// 	NM	TORRANCE                        	Y
// 	NM	UNION                           	Y
// 	NM	VALENCIA                        	Y
// 	NV	CARSON CITY                     	Y
// 	NV	CHURCHILL                       	Y
// 	NV	CLARK                           	Y
// 	NV	DOUGLAS                         	Y
// 	NV	ELKO                            	Y
// 	NV	ESMERALDA                       	Y
// 	NV	EUREKA                          	Y
// 	NV	HUMBOLDT                        	Y
// 	NV	LANDER                          	Y
// 	NV	LINCOLN                         	Y
// 	NV	LYON                            	Y
// 	NV	MINERAL                         	Y
// 	NV	NYE                             	Y
// 	NV	PERSHING                        	Y
// 	NV	STOREY                          	Y
// 	NV	WASHOE                          	Y
// 	NV	WHITE PINE                      	Y
// 	NY	ALBANY                          	Y
// 	NY	ALLEGANY                        	Y
// 	NY	BRONX                           	Y
// 	NY	BROOME                          	Y
// 	NY	CATTARAUGUS                     	Y
// 	NY	CAYUGA                          	Y
// 	NY	CHAUTAUQUA                      	Y
// 	NY	CHEMUNG                         	Y
// 	NY	CHENANGO                        	Y
// 	NY	CLINTON                         	Y
// 	NY	COLUMBIA                        	Y
// 	NY	CORTLAND                        	Y
// 	NY	DELAWARE                        	Y
// 	NY	DUTCHESS                        	Y
// 	NY	ERIE                            	Y
// 	NY	ESSEX                           	Y
// 	NY	FRANKLIN                        	Y
// 	NY	FULTON                          	Y
// 	NY	GENESEE                         	Y
// 	NY	GREENE                          	Y
// 	NY	HAMILTON                        	Y
// 	NY	HERKIMER                        	Y
// 	NY	JEFFERSON                       	Y
// 	NY	KINGS                           	Y
// 	NY	LEWIS                           	Y
// 	NY	LIVINGSTON                      	Y
// 	NY	MADISON                         	Y
// 	NY	MONROE                          	Y
// 	NY	MONTGOMERY                      	Y
// 	NY	NASSAU                          	Y
// 	NY	NEW YORK                        	Y
// 	NY	NIAGARA                         	Y
// 	NY	ONEIDA                          	Y
// 	NY	ONONDAGA                        	Y
// 	NY	ONTARIO                         	Y
// 	NY	ORANGE                          	Y
// 	NY	ORLEANS                         	Y
// 	NY	OSWEGO                          	Y
// 	NY	OTSEGO                          	Y
// 	NY	PUTNAM                          	Y
// 	NY	QUEENS                          	Y
// 	NY	RENSSELAER                      	Y
// 	NY	RICHMOND                        	Y
// 	NY	ROCKLAND                        	Y
// 	NY	SARATOGA                        	Y
// 	NY	SCHENECTADY                     	Y
// 	NY	SCHOHARIE                       	Y
// 	NY	SCHUYLER                        	Y
// 	NY	SENECA                          	Y
// 	NY	ST. LAWRENCE                    	Y
// 	NY	STEUBEN                         	Y
// 	NY	SUFFOLK                         	Y
// 	NY	SULLIVAN                        	Y
// 	NY	TIOGA                           	Y
// 	NY	TOMPKINS                        	Y
// 	NY	ULSTER                          	Y
// 	NY	WARREN                          	Y
// 	NY	WASHINGTON                      	Y
// 	NY	WAYNE                           	Y
// 	NY	WESTCHESTER                     	Y
// 	NY	WYOMING                         	Y
// 	NY	YATES                           	Y
// 	OH	ADAMS                           	Y
// 	OH	ALLEN                           	Y
// 	OH	ASHLAND                         	Y
// 	OH	ASHTABULA                       	Y
// 	OH	ATHENS                          	Y
// 	OH	AUGLAIZE                        	Y
// 	OH	BELMONT                         	Y
// 	OH	BROWN                           	Y
// 	OH	BUTLER                          	Y
// 	OH	CARROLL                         	Y
// 	OH	CHAMPAIGN                       	Y
// 	OH	CLARK                           	Y
// 	OH	CLERMONT                        	Y
// 	OH	CLINTON                         	Y
// 	OH	COLUMBIANA                      	Y
// 	OH	COSHOCTON                       	Y
// 	OH	CRAWFORD                        	Y
// 	OH	CUYAHOGA                        	Y
// 	OH	DARKE                           	Y
// 	OH	DEFIANCE                        	Y
// 	OH	DELAWARE                        	Y
// 	OH	ERIE                            	Y
// 	OH	FAIRFIELD                       	Y
// 	OH	FAYETTE                         	Y
// 	OH	FRANKLIN                        	Y
// 	OH	FULTON                          	Y
// 	OH	GALLIA                          	Y
// 	OH	GEAUGA                          	Y
// 	OH	GREENE                          	Y
// 	OH	GUERNSEY                        	Y
// 	OH	HAMILTON                        	Y
// 	OH	HANCOCK                         	Y
// 	OH	HARDIN                          	Y
// 	OH	HARRISON                        	Y
// 	OH	HENRY                           	Y
// 	OH	HIGHLAND                        	Y
// 	OH	HOCKING                         	Y
// 	OH	HOLMES                          	Y
// 	OH	HURON                           	Y
// 	OH	JACKSON                         	Y
// 	OH	JEFFERSON                       	Y
// 	OH	KNOX                            	Y
// 	OH	LAKE                            	Y
// 	OH	LAWRENCE                        	Y
// 	OH	LICKING                         	Y
// 	OH	LOGAN                           	Y
// 	OH	LORAIN                          	Y
// 	OH	LUCAS                           	Y
// 	OH	MADISON                         	Y
// 	OH	MAHONING                        	Y
// 	OH	MARION                          	Y
// 	OH	MEDINA                          	Y
// 	OH	MEIGS                           	Y
// 	OH	MERCER                          	Y
// 	OH	MIAMI                           	Y
// 	OH	MONROE                          	Y
// 	OH	MONTGOMERY                      	Y
// 	OH	MORGAN                          	Y
// 	OH	MORROW                          	Y
// 	OH	MUSKINGUM                       	Y
// 	OH	NOBLE                           	Y
// 	OH	OTTAWA                          	Y
// 	OH	PAULDING                        	Y
// 	OH	PERRY                           	Y
// 	OH	PICKAWAY                        	Y
// 	OH	PIKE                            	Y
// 	OH	PORTAGE                         	Y
// 	OH	PREBLE                          	Y
// 	OH	PUTNAM                          	Y
// 	OH	RICHLAND                        	Y
// 	OH	ROSS                            	Y
// 	OH	SANDUSKY                        	Y
// 	OH	SCIOTO                          	Y
// 	OH	SENECA                          	Y
// 	OH	SHELBY                          	Y
// 	OH	STARK                           	Y
// 	OH	SUMMIT                          	Y
// 	OH	TRUMBULL                        	Y
// 	OH	TUSCARAWAS                      	Y
// 	OH	UNION                           	Y
// 	OH	VAN WERT                        	Y
// 	OH	VINTON                          	Y
// 	OH	WARREN                          	Y
// 	OH	WASHINGTON                      	Y
// 	OH	WAYNE                           	Y
// 	OH	WILLIAMS                        	Y
// 	OH	WOOD                            	Y
// 	OH	WYANDOT                         	Y
// 	OK	ADAIR                           	Y
// 	OK	ALFALFA                         	Y
// 	OK	ATOKA                           	Y
// 	OK	BEAVER                          	Y
// 	OK	BECKHAM                         	Y
// 	OK	BLAINE                          	Y
// 	OK	BRYAN                           	Y
// 	OK	CADDO                           	Y
// 	OK	CANADIAN                        	Y
// 	OK	CARTER                          	Y
// 	OK	CHEROKEE                        	Y
// 	OK	CHOCTAW                         	Y
// 	OK	CIMARRON                        	Y
// 	OK	CLEVELAND                       	Y
// 	OK	COAL                            	Y
// 	OK	COMANCHE                        	Y
// 	OK	COTTON                          	Y
// 	OK	CRAIG                           	Y
// 	OK	CREEK                           	Y
// 	OK	CUSTER                          	Y
// 	OK	DELAWARE                        	Y
// 	OK	DEWEY                           	Y
// 	OK	ELLIS                           	Y
// 	OK	GARFIELD                        	Y
// 	OK	GARVIN                          	Y
// 	OK	GRADY                           	Y
// 	OK	GRANT                           	Y
// 	OK	GREER                           	Y
// 	OK	HARMON                          	Y
// 	OK	HARPER                          	Y
// 	OK	HASKELL                         	Y
// 	OK	HUGHES                          	Y
// 	OK	JACKSON                         	Y
// 	OK	JEFFERSON                       	Y
// 	OK	JOHNSTON                        	Y
// 	OK	KAY                             	Y
// 	OK	KINGFISHER                      	Y
// 	OK	KIOWA                           	Y
// 	OK	LATIMER                         	Y
// 	OK	LE FLORE                        	Y
// 	OK	LINCOLN                         	Y
// 	OK	LOGAN                           	Y
// 	OK	LOVE                            	Y
// 	OK	MAJOR                           	Y
// 	OK	MARSHALL                        	Y
// 	OK	MAYES                           	Y
// 	OK	MCCLAIN                         	Y
// 	OK	MCCURTAIN                       	Y
// 	OK	MCINTOSH                        	Y
// 	OK	MURRAY                          	Y
// 	OK	MUSKOGEE                        	Y
// 	OK	NOBLE                           	Y
// 	OK	NOWATA                          	Y
// 	OK	OKFUSKEE                        	Y
// 	OK	OKLAHOMA                        	Y
// 	OK	OKMULGEE                        	Y
// 	OK	OSAGE                           	Y
// 	OK	OTTAWA                          	Y
// 	OK	PAWNEE                          	Y
// 	OK	PAYNE                           	Y
// 	OK	PITTSBURG                       	Y
// 	OK	PONTOTOC                        	Y
// 	OK	POTTAWATOMIE                    	Y
// 	OK	PUSHMATAHA                      	Y
// 	OK	ROGER MILLS                     	Y
// 	OK	ROGERS                          	Y
// 	OK	SEMINOLE                        	Y
// 	OK	SEQUOYAH                        	Y
// 	OK	STEPHENS                        	Y
// 	OK	TEXAS                           	Y
// 	OK	TILLMAN                         	Y
// 	OK	TULSA                           	Y
// 	OK	WAGONER                         	Y
// 	OK	WASHINGTON                      	Y
// 	OK	WASHITA                         	Y
// 	OK	WOODS                           	Y
// 	OK	WOODWARD                        	Y
// 	OR	BAKER                           	Y
// 	OR	BENTON                          	Y
// 	OR	CLACKAMAS                       	Y
// 	OR	CLATSOP                         	Y
// 	OR	COLUMBIA                        	Y
// 	OR	COOS                            	Y
// 	OR	CROOK                           	Y
// 	OR	CURRY                           	Y
// 	OR	DESCHUTES                       	Y
// 	OR	DOUGLAS                         	Y
// 	OR	GILLIAM                         	Y
// 	OR	GRANT                           	Y
// 	OR	HARNEY                          	Y
// 	OR	HOOD RIVER                      	Y
// 	OR	JACKSON                         	Y
// 	OR	JEFFERSON                       	Y
// 	OR	JOSEPHINE                       	Y
// 	OR	KLAMATH                         	Y
// 	OR	LAKE                            	Y
// 	OR	LANE                            	Y
// 	OR	LINCOLN                         	Y
// 	OR	LINN                            	Y
// 	OR	MALHEUR                         	Y
// 	OR	MARION                          	Y
// 	OR	MORROW                          	Y
// 	OR	MULTNOMAH                       	Y
// 	OR	POLK                            	Y
// 	OR	SHERMAN                         	Y
// 	OR	TILLAMOOK                       	Y
// 	OR	UMATILLA                        	Y
// 	OR	UNION                           	Y
// 	OR	WALLOWA                         	Y
// 	OR	WASCO                           	Y
// 	OR	WASHINGTON                      	Y
// 	OR	WHEELER                         	Y
// 	OR	YAMHILL                         	Y
// 	PA	ADAMS                           	Y
// 	PA	ALLEGHENY                       	Y
// 	PA	ARMSTRONG                       	Y
// 	PA	BEAVER                          	Y
// 	PA	BEDFORD                         	Y
// 	PA	BERKS                           	Y
// 	PA	BLAIR                           	Y
// 	PA	BRADFORD                        	Y
// 	PA	BUCKS                           	Y
// 	PA	BUTLER                          	Y
// 	PA	CAMBRIA                         	Y
// 	PA	CAMERON                         	Y
// 	PA	CARBON                          	Y
// 	PA	CENTRE                          	Y
// 	PA	CHESTER                         	Y
// 	PA	CLARION                         	Y
// 	PA	CLEARFIELD                      	Y
// 	PA	CLINTON                         	Y
// 	PA	COLUMBIA                        	Y
// 	PA	CRAWFORD                        	Y
// 	PA	CUMBERLAND                      	Y
// 	PA	DAUPHIN                         	Y
// 	PA	DELAWARE                        	Y
// 	PA	ELK                             	Y
// 	PA	ERIE                            	Y
// 	PA	FAYETTE                         	Y
// 	PA	FOREST                          	Y
// 	PA	FRANKLIN                        	Y
// 	PA	FULTON                          	Y
// 	PA	GREENE                          	Y
// 	PA	HUNTINGDON                      	Y
// 	PA	INDIANA                         	Y
// 	PA	JEFFERSON                       	Y
// 	PA	JUNIATA                         	Y
// 	PA	LACKAWANNA                      	Y
// 	PA	LANCASTER                       	Y
// 	PA	LAWRENCE                        	Y
// 	PA	LEBANON                         	Y
// 	PA	LEHIGH                          	Y
// 	PA	LUZERNE                         	Y
// 	PA	LYCOMING                        	Y
// 	PA	MCKEAN                          	Y
// 	PA	MERCER                          	Y
// 	PA	MIFFLIN                         	Y
// 	PA	MONROE                          	Y
// 	PA	MONTGOMERY                      	Y
// 	PA	MONTOUR                         	Y
// 	PA	NORTHAMPTON                     	Y
// 	PA	NORTHUMBERLAND                  	Y
// 	PA	PERRY                           	Y
// 	PA	PHILADELPHIA                    	Y
// 	PA	PIKE                            	Y
// 	PA	POTTER                          	Y
// 	PA	SCHUYLKILL                      	Y
// 	PA	SNYDER                          	Y
// 	PA	SOMERSET                        	Y
// 	PA	SULLIVAN                        	Y
// 	PA	SUSQUEHANNA                     	Y
// 	PA	TIOGA                           	Y
// 	PA	UNION                           	Y
// 	PA	VENANGO                         	Y
// 	PA	WARREN                          	Y
// 	PA	WASHINGTON                      	Y
// 	PA	WAYNE                           	Y
// 	PA	WESTMORELAND                    	Y
// 	PA	WYOMING                         	Y
// 	PA	YORK                            	Y
// 	PR	ADJUNTAS                        	Y
// 	PR	AGUADA                          	Y
// 	PR	AGUADILLA                       	Y
// 	PR	AGUAS BUENAS                    	Y
// 	PR	AIBONITO                        	Y
// 	PR	ANASCO                          	Y
// 	PR	ARECIBO                         	Y
// 	PR	ARROYO                          	Y
// 	PR	BARCELONETA                     	Y
// 	PR	BARRANQUITAS                    	Y
// 	PR	BAYAMON                         	Y
// 	PR	CABO ROJO                       	Y
// 	PR	CAGUAS                          	Y
// 	PR	CAMUY                           	Y
// 	PR	CANOVANAS                       	Y
// 	PR	CAROLINA                        	Y
// 	PR	CATANO                          	Y
// 	PR	CAYEY                           	Y
// 	PR	CEIBA                           	Y
// 	PR	CIALES                          	Y
// 	PR	CIDRA                           	Y
// 	PR	COAMO                           	Y
// 	PR	COMERIO                         	Y
// 	PR	COROZAL                         	Y
// 	PR	CULEBRA                         	Y
// 	PR	DORADO                          	Y
// 	PR	FAJARDO                         	Y
// 	PR	FLORIDA                         	Y
// 	PR	GUANICA                         	Y
// 	PR	GUAYAMA                         	Y
// 	PR	GUAYANILLA                      	Y
// 	PR	GUAYNABO                        	Y
// 	PR	GURABO                          	Y
// 	PR	HATILLO                         	Y
// 	PR	HORMIGUEROS                     	Y
// 	PR	HUMACAO                         	Y
// 	PR	ISABELA                         	Y
// 	PR	JAYUYA                          	Y
// 	PR	JUANA DIAZ                      	Y
// 	PR	JUNCOS                          	Y
// 	PR	LAJAS                           	Y
// 	PR	LARES                           	Y
// 	PR	LAS MARIAS                      	Y
// 	PR	LAS PIEDRAS                     	Y
// 	PR	LOIZA                           	Y
// 	PR	LUQUILLO                        	Y
// 	PR	MANATI                          	Y
// 	PR	MARICAO                         	Y
// 	PR	MAUNABO                         	Y
// 	PR	MAYAGUEZ                        	Y
// 	PR	MOCA                            	Y
// 	PR	MOROVIS                         	Y
// 	PR	NAGUABO                         	Y
// 	PR	NARANJITO                       	Y
// 	PR	OROCOVIS                        	Y
// 	PR	PATILLAS                        	Y
// 	PR	PENUELAS                        	Y
// 	PR	PONCE                           	Y
// 	PR	QUEBRADILLAS                    	Y
// 	PR	RINCON                          	Y
// 	PR	RIO GRANDE                      	Y
// 	PR	SABANA GRANDE                   	Y
// 	PR	SALINAS                         	Y
// 	PR	SAN GERMAN                      	Y
// 	PR	SAN JUAN                        	Y
// 	PR	SAN LORENZO                     	Y
// 	PR	SAN SEBASTIAN                   	Y
// 	PR	SANTA ISABEL                    	Y
// 	PR	TOA ALTA                        	Y
// 	PR	TOA BAJA                        	Y
// 	PR	TRUJILLO ALTO                   	Y
// 	PR	UTUADO                          	Y
// 	PR	VEGA ALTA                       	Y
// 	PR	VEGA BAJA                       	Y
// 	PR	VIEQUES                         	Y
// 	PR	VILLALBA                        	Y
// 	PR	YABUCOA                         	Y
// 	PR	YAUCO                           	Y
// 	RI	BRISTOL                         	Y
// 	RI	KENT                            	Y
// 	RI	NEWPORT                         	Y
// 	RI	PROVIDENCE                      	Y
// 	RI	WASHINGTON                      	Y
// 	SC	ABBEVILLE                       	Y
// 	SC	AIKEN                           	Y
// 	SC	ALLENDALE                       	Y
// 	SC	ANDERSON                        	Y
// 	SC	BAMBERG                         	Y
// 	SC	BARNWELL                        	Y
// 	SC	BEAUFORT                        	Y
// 	SC	BERKELEY                        	Y
// 	SC	CALHOUN                         	Y
// 	SC	CHARLESTON                      	Y
// 	SC	CHEROKEE                        	Y
// 	SC	CHESTER                         	Y
// 	SC	CHESTERFIELD                    	Y
// 	SC	CLARENDON                       	Y
// 	SC	COLLETON                        	Y
// 	SC	DARLINGTON                      	Y
// 	SC	DILLON                          	Y
// 	SC	DORCHESTER                      	Y
// 	SC	EDGEFIELD                       	Y
// 	SC	FAIRFIELD                       	Y
// 	SC	FLORENCE                        	Y
// 	SC	GEORGETOWN                      	Y
// 	SC	GREENVILLE                      	Y
// 	SC	GREENWOOD                       	Y
// 	SC	HAMPTON                         	Y
// 	SC	HORRY                           	Y
// 	SC	JASPER                          	Y
// 	SC	KERSHAW                         	Y
// 	SC	LANCASTER                       	Y
// 	SC	LAURENS                         	Y
// 	SC	LEE                             	Y
// 	SC	LEXINGTON                       	Y
// 	SC	MARION                          	Y
// 	SC	MARLBORO                        	Y
// 	SC	MCCORMICK                       	Y
// 	SC	NEWBERRY                        	Y
// 	SC	OCONEE                          	Y
// 	SC	ORANGEBURG                      	Y
// 	SC	PICKENS                         	Y
// 	SC	RICHLAND                        	Y
// 	SC	SALUDA                          	Y
// 	SC	SPARTANBURG                     	Y
// 	SC	SUMTER                          	Y
// 	SC	UNION                           	Y
// 	SC	WILLIAMSBURG                    	Y
// 	SC	YORK                            	Y
// 	SD	AURORA                          	Y
// 	SD	BEADLE                          	Y
// 	SD	BENNETT                         	Y
// 	SD	BON HOMME                       	Y
// 	SD	BROOKINGS                       	Y
// 	SD	BROWN                           	Y
// 	SD	BRULE                           	Y
// 	SD	BUFFALO                         	Y
// 	SD	BUTTE                           	Y
// 	SD	CAMPBELL                        	Y
// 	SD	CHARLES MIX                     	Y
// 	SD	CLARK                           	Y
// 	SD	CLAY                            	Y
// 	SD	CODINGTON                       	Y
// 	SD	CORSON                          	Y
// 	SD	CUSTER                          	Y
// 	SD	DAVISON                         	Y
// 	SD	DAY                             	Y
// 	SD	DEUEL                           	Y
// 	SD	DEWEY                           	Y
// 	SD	DOUGLAS                         	Y
// 	SD	EDMUNDS                         	Y
// 	SD	FALL RIVER                      	Y
// 	SD	FAULK                           	Y
// 	SD	GRANT                           	Y
// 	SD	GREGORY                         	Y
// 	SD	HAAKON                          	Y
// 	SD	HAMLIN                          	Y
// 	SD	HAND                            	Y
// 	SD	HANSON                          	Y
// 	SD	HARDING                         	Y
// 	SD	HUGHES                          	Y
// 	SD	HUTCHINSON                      	Y
// 	SD	HYDE                            	Y
// 	SD	JACKSON                         	Y
// 	SD	JERAULD                         	Y
// 	SD	JONES                           	Y
// 	SD	KINGSBURY                       	Y
// 	SD	LAKE                            	Y
// 	SD	LAWRENCE                        	Y
// 	SD	LINCOLN                         	Y
// 	SD	LYMAN                           	Y
// 	SD	MARSHALL                        	Y
// 	SD	MCCOOK                          	Y
// 	SD	MCPHERSON                       	Y
// 	SD	MEADE                           	Y
// 	SD	MELLETTE                        	Y
// 	SD	MINER                           	Y
// 	SD	MINNEHAHA                       	Y
// 	SD	MOODY                           	Y
// 	SD	PENNINGTON                      	Y
// 	SD	PERKINS                         	Y
// 	SD	POTTER                          	Y
// 	SD	ROBERTS                         	Y
// 	SD	SANBORN                         	Y
// 	SD	SHANNON                         	Y
// 	SD	SPINK                           	Y
// 	SD	STANLEY                         	Y
// 	SD	SULLY                           	Y
// 	SD	TODD                            	Y
// 	SD	TRIPP                           	Y
// 	SD	TURNER                          	Y
// 	SD	UNION                           	Y
// 	SD	WALWORTH                        	Y
// 	SD	YANKTON                         	Y
// 	SD	ZIEBACH                         	Y
// 	TN	ANDERSON                        	Y
// 	TN	BEDFORD                         	Y
// 	TN	BENTON                          	Y
// 	TN	BLEDSOE                         	Y
// 	TN	BLOUNT                          	Y
// 	TN	BRADLEY                         	Y
// 	TN	CAMPBELL                        	Y
// 	TN	CANNON                          	Y
// 	TN	CARROLL                         	Y
// 	TN	CARTER                          	Y
// 	TN	CHEATHAM                        	Y
// 	TN	CHESTER                         	Y
// 	TN	CLAIBORNE                       	Y
// 	TN	CLAY                            	Y
// 	TN	COCKE                           	Y
// 	TN	COFFEE                          	Y
// 	TN	CROCKETT                        	Y
// 	TN	CUMBERLAND                      	Y
// 	TN	DAVIDSON                        	Y
// 	TN	DECATUR                         	Y
// 	TN	DEKALB                          	Y
// 	TN	DICKSON                         	Y
// 	TN	DYER                            	Y
// 	TN	FAYETTE                         	Y
// 	TN	FENTRESS                        	Y
// 	TN	FRANKLIN                        	Y
// 	TN	GIBSON                          	Y
// 	TN	GILES                           	Y
// 	TN	GRAINGER                        	Y
// 	TN	GREENE                          	Y
// 	TN	GRUNDY                          	Y
// 	TN	HAMBLEN                         	Y
// 	TN	HAMILTON                        	Y
// 	TN	HANCOCK                         	Y
// 	TN	HARDEMAN                        	Y
// 	TN	HARDIN                          	Y
// 	TN	HAWKINS                         	Y
// 	TN	HAYWOOD                         	Y
// 	TN	HENDERSON                       	Y
// 	TN	HENRY                           	Y
// 	TN	HICKMAN                         	Y
// 	TN	HOUSTON                         	Y
// 	TN	HUMPHREYS                       	Y
// 	TN	JACKSON                         	Y
// 	TN	JEFFERSON                       	Y
// 	TN	JOHNSON                         	Y
// 	TN	KNOX                            	Y
// 	TN	LAKE                            	Y
// 	TN	LAUDERDALE                      	Y
// 	TN	LAWRENCE                        	Y
// 	TN	LEWIS                           	Y
// 	TN	LINCOLN                         	Y
// 	TN	LOUDON                          	Y
// 	TN	MACON                           	Y
// 	TN	MADISON                         	Y
// 	TN	MARION                          	Y
// 	TN	MARSHALL                        	Y
// 	TN	MAURY                           	Y
// 	TN	MCMINN                          	Y
// 	TN	MCNAIRY                         	Y
// 	TN	MEIGS                           	Y
// 	TN	MONROE                          	Y
// 	TN	MONTGOMERY                      	Y
// 	TN	MOORE                           	Y
// 	TN	MORGAN                          	Y
// 	TN	OBION                           	Y
// 	TN	OVERTON                         	Y
// 	TN	PERRY                           	Y
// 	TN	PICKETT                         	Y
// 	TN	POLK                            	Y
// 	TN	PUTNAM                          	Y
// 	TN	RHEA                            	Y
// 	TN	ROANE                           	Y
// 	TN	ROBERTSON                       	Y
// 	TN	RUTHERFORD                      	Y
// 	TN	SCOTT                           	Y
// 	TN	SEQUATCHIE                      	Y
// 	TN	SEVIER                          	Y
// 	TN	SHELBY                          	Y
// 	TN	SMITH                           	Y
// 	TN	STEWART                         	Y
// 	TN	SULLIVAN                        	Y
// 	TN	SUMNER                          	Y
// 	TN	TIPTON                          	Y
// 	TN	TROUSDALE                       	Y
// 	TN	UNICOI                          	Y
// 	TN	UNION                           	Y
// 	TN	VAN BUREN                       	Y
// 	TN	WARREN                          	Y
// 	TN	WASHINGTON                      	Y
// 	TN	WAYNE                           	Y
// 	TN	WEAKLEY                         	Y
// 	TN	WHITE                           	Y
// 	TN	WILLIAMSON                      	Y
// 	TN	WILSON                          	Y
// 	TX	ANDERSON                        	Y
// 	TX	ANDREWS                         	Y
// 	TX	ANGELINA                        	Y
// 	TX	ARANSAS                         	Y
// 	TX	ARCHER                          	Y
// 	TX	ARMSTRONG                       	Y
// 	TX	ATASCOSA                        	Y
// 	TX	AUSTIN                          	Y
// 	TX	BAILEY                          	Y
// 	TX	BANDERA                         	Y
// 	TX	BASTROP                         	Y
// 	TX	BAYLOR                          	Y
// 	TX	BEE                             	Y
// 	TX	BELL                            	Y
// 	TX	BEXAR                           	Y
// 	TX	BLANCO                          	Y
// 	TX	BORDEN                          	Y
// 	TX	BOSQUE                          	Y
// 	TX	BOWIE                           	Y
// 	TX	BRAZORIA                        	Y
// 	TX	BRAZOS                          	Y
// 	TX	BREWSTER                        	Y
// 	TX	BRISCOE                         	Y
// 	TX	BROOKS                          	Y
// 	TX	BROWN                           	Y
// 	TX	BURLESON                        	Y
// 	TX	BURNET                          	Y
// 	TX	CALDWELL                        	Y
// 	TX	CALHOUN                         	Y
// 	TX	CALLAHAN                        	Y
// 	TX	CAMERON                         	Y
// 	TX	CAMP                            	Y
// 	TX	CARSON                          	Y
// 	TX	CASS                            	Y
// 	TX	CASTRO                          	Y
// 	TX	CHAMBERS                        	Y
// 	TX	CHEROKEE                        	Y
// 	TX	CHILDRESS                       	Y
// 	TX	CLAY                            	Y
// 	TX	COCHRAN                         	Y
// 	TX	COKE                            	Y
// 	TX	COLEMAN                         	Y
// 	TX	COLLIN                          	Y
// 	TX	COLLINGSWORTH                   	Y
// 	TX	COLORADO                        	Y
// 	TX	COMAL                           	Y
// 	TX	COMANCHE                        	Y
// 	TX	CONCHO                          	Y
// 	TX	COOKE                           	Y
// 	TX	CORYELL                         	Y
// 	TX	COTTLE                          	Y
// 	TX	CRANE                           	Y
// 	TX	CROCKETT                        	Y
// 	TX	CROSBY                          	Y
// 	TX	CULBERSON                       	Y
// 	TX	DALLAM                          	Y
// 	TX	DALLAS                          	Y
// 	TX	DAWSON                          	Y
// 	TX	DEAF SMITH                      	Y
// 	TX	DELTA                           	Y
// 	TX	DENTON                          	Y
// 	TX	DEWITT                          	Y
// 	TX	DICKENS                         	Y
// 	TX	DIMMIT                          	Y
// 	TX	DONLEY                          	Y
// 	TX	DUVAL                           	Y
// 	TX	EASTLAND                        	Y
// 	TX	ECTOR                           	Y
// 	TX	EDWARDS                         	Y
// 	TX	EL PASO                         	Y
// 	TX	ELLIS                           	Y
// 	TX	ERATH                           	Y
// 	TX	FALLS                           	Y
// 	TX	FANNIN                          	Y
// 	TX	FAYETTE                         	Y
// 	TX	FISHER                          	Y
// 	TX	FLOYD                           	Y
// 	TX	FOARD                           	Y
// 	TX	FORT BEND                       	Y
// 	TX	FRANKLIN                        	Y
// 	TX	FREESTONE                       	Y
// 	TX	FRIO                            	Y
// 	TX	GAINES                          	Y
// 	TX	GALVESTON                       	Y
// 	TX	GARZA                           	Y
// 	TX	GILLESPIE                       	Y
// 	TX	GLASSCOCK                       	Y
// 	TX	GOLIAD                          	Y
// 	TX	GONZALES                        	Y
// 	TX	GRAY                            	Y
// 	TX	GRAYSON                         	Y
// 	TX	GREGG                           	Y
// 	TX	GRIMES                          	Y
// 	TX	GUADALUPE                       	Y
// 	TX	HALE                            	Y
// 	TX	HALL                            	Y
// 	TX	HAMILTON                        	Y
// 	TX	HANSFORD                        	Y
// 	TX	HARDEMAN                        	Y
// 	TX	HARDIN                          	Y
// 	TX	HARRIS                          	Y
// 	TX	HARRISON                        	Y
// 	TX	HARTLEY                         	Y
// 	TX	HASKELL                         	Y
// 	TX	HAYS                            	Y
// 	TX	HEMPHILL                        	Y
// 	TX	HENDERSON                       	Y
// 	TX	HIDALGO                         	Y
// 	TX	HILL                            	Y
// 	TX	HOCKLEY                         	Y
// 	TX	HOOD                            	Y
// 	TX	HOPKINS                         	Y
// 	TX	HOUSTON                         	Y
// 	TX	HOWARD                          	Y
// 	TX	HUDSPETH                        	Y
// 	TX	HUNT                            	Y
// 	TX	HUTCHINSON                      	Y
// 	TX	IRION                           	Y
// 	TX	JACK                            	Y
// 	TX	JACKSON                         	Y
// 	TX	JASPER                          	Y
// 	TX	JEFF DAVIS                      	Y
// 	TX	JEFFERSON                       	Y
// 	TX	JIM HOGG                        	Y
// 	TX	JIM WELLS                       	Y
// 	TX	JOHNSON                         	Y
// 	TX	JONES                           	Y
// 	TX	KARNES                          	Y
// 	TX	KAUFMAN                         	Y
// 	TX	KENDALL                         	Y
// 	TX	KENEDY                          	Y
// 	TX	KENT                            	Y
// 	TX	KERR                            	Y
// 	TX	KIMBLE                          	Y
// 	TX	KING                            	Y
// 	TX	KINNEY                          	Y
// 	TX	KLEBERG                         	Y
// 	TX	KNOX                            	Y
// 	TX	LA SALLE                        	Y
// 	TX	LAMAR                           	Y
// 	TX	LAMB                            	Y
// 	TX	LAMPASAS                        	Y
// 	TX	LAVACA                          	Y
// 	TX	LEE                             	Y
// 	TX	LEON                            	Y
// 	TX	LIBERTY                         	Y
// 	TX	LIMESTONE                       	Y
// 	TX	LIPSCOMB                        	Y
// 	TX	LIVE OAK                        	Y
// 	TX	LLANO                           	Y
// 	TX	LOVING                          	Y
// 	TX	LUBBOCK                         	Y
// 	TX	LYNN                            	Y
// 	TX	MADISON                         	Y
// 	TX	MARION                          	Y
// 	TX	MARTIN                          	Y
// 	TX	MASON                           	Y
// 	TX	MATAGORDA                       	Y
// 	TX	MAVERICK                        	Y
// 	TX	MCCULLOCH                       	Y
// 	TX	MCLENNAN                        	Y
// 	TX	MCMULLEN                        	Y
// 	TX	MEDINA                          	Y
// 	TX	MENARD                          	Y
// 	TX	MIDLAND                         	Y
// 	TX	MILAM                           	Y
// 	TX	MILLS                           	Y
// 	TX	MITCHELL                        	Y
// 	TX	MONTAGUE                        	Y
// 	TX	MONTGOMERY                      	Y
// 	TX	MOORE                           	Y
// 	TX	MORRIS                          	Y
// 	TX	MOTLEY                          	Y
// 	TX	NACOGDOCHES                     	Y
// 	TX	NAVARRO                         	Y
// 	TX	NEWTON                          	Y
// 	TX	NOLAN                           	Y
// 	TX	NUECES                          	Y
// 	TX	OCHILTREE                       	Y
// 	TX	OLDHAM                          	Y
// 	TX	ORANGE                          	Y
// 	TX	PALO PINTO                      	Y
// 	TX	PANOLA                          	Y
// 	TX	PARKER                          	Y
// 	TX	PARMER                          	Y
// 	TX	PECOS                           	Y
// 	TX	POLK                            	Y
// 	TX	POTTER                          	Y
// 	TX	PRESIDIO                        	Y
// 	TX	RAINS                           	Y
// 	TX	RANDALL                         	Y
// 	TX	REAGAN                          	Y
// 	TX	REAL                            	Y
// 	TX	RED RIVER                       	Y
// 	TX	REEVES                          	Y
// 	TX	REFUGIO                         	Y
// 	TX	ROBERTS                         	Y
// 	TX	ROBERTSON                       	Y
// 	TX	ROCKWALL                        	Y
// 	TX	RUNNELS                         	Y
// 	TX	RUSK                            	Y
// 	TX	SABINE                          	Y
// 	TX	SAN AUGUSTINE                   	Y
// 	TX	SAN JACINTO                     	Y
// 	TX	SAN PATRICIO                    	Y
// 	TX	SAN SABA                        	Y
// 	TX	SCHLEICHER                      	Y
// 	TX	SCURRY                          	Y
// 	TX	SHACKELFORD                     	Y
// 	TX	SHELBY                          	Y
// 	TX	SHERMAN                         	Y
// 	TX	SMITH                           	Y
// 	TX	SOMERVELL                       	Y
// 	TX	STARR                           	Y
// 	TX	STEPHENS                        	Y
// 	TX	STERLING                        	Y
// 	TX	STONEWALL                       	Y
// 	TX	SUTTON                          	Y
// 	TX	SWISHER                         	Y
// 	TX	TARRANT                         	Y
// 	TX	TAYLOR                          	Y
// 	TX	TERRELL                         	Y
// 	TX	TERRY                           	Y
// 	TX	THROCKMORTON                    	Y
// 	TX	TITUS                           	Y
// 	TX	TOM GREEN                       	Y
// 	TX	TRAVIS                          	Y
// 	TX	TRINITY                         	Y
// 	TX	TYLER                           	Y
// 	TX	UPSHUR                          	Y
// 	TX	UPTON                           	Y
// 	TX	UVALDE                          	Y
// 	TX	VAL VERDE                       	Y
// 	TX	VAN ZANDT                       	Y
// 	TX	VICTORIA                        	Y
// 	TX	WALKER                          	Y
// 	TX	WALLER                          	Y
// 	TX	WARD                            	Y
// 	TX	WASHINGTON                      	Y
// 	TX	WEBB                            	Y
// 	TX	WHARTON                         	Y
// 	TX	WHEELER                         	Y
// 	TX	WICHITA                         	Y
// 	TX	WILBARGER                       	Y
// 	TX	WILLACY                         	Y
// 	TX	WILLIAMSON                      	Y
// 	TX	WILSON                          	Y
// 	TX	WINKLER                         	Y
// 	TX	WISE                            	Y
// 	TX	WOOD                            	Y
// 	TX	YOAKUM                          	Y
// 	TX	YOUNG                           	Y
// 	TX	ZAPATA                          	Y
// 	TX	ZAVALA                          	Y
// 	UM	BAKER ISLAND                    	Y
// 	UM	HOWLAND ISLAND                  	Y
// 	UM	JARVIS ISLAND                   	Y
// 	UM	JOHNSTON ISLANDS                	Y
// 	UM	KINGMAN REEF                    	Y
// 	UM	MIDWAY ISLANDS                  	Y
// 	UM	NAVASSA ISLAND                  	Y
// 	UM	PALMYRA ATOLL                   	Y
// 	UM	WAKE ISLANDS                    	Y
// 	UT	BEAVER                          	Y
// 	UT	BOX ELDER                       	Y
// 	UT	CACHE                           	Y
// 	UT	CARBON                          	Y
// 	UT	DAGGETT                         	Y
// 	UT	DAVIS                           	Y
// 	UT	DUCHESNE                        	Y
// 	UT	EMERY                           	Y
// 	UT	GARFIELD                        	Y
// 	UT	GRAND                           	Y
// 	UT	IRON                            	Y
// 	UT	JUAB                            	Y
// 	UT	KANE                            	Y
// 	UT	MILLARD                         	Y
// 	UT	MORGAN                          	Y
// 	UT	PIUTE                           	Y
// 	UT	RICH                            	Y
// 	UT	SALT LAKE                       	Y
// 	UT	SAN JUAN                        	Y
// 	UT	SANPETE                         	Y
// 	UT	SEVIER                          	Y
// 	UT	SUMMIT                          	Y
// 	UT	TOOELE                          	Y
// 	UT	UINTAH                          	Y
// 	UT	UTAH                            	Y
// 	UT	WASATCH                         	Y
// 	UT	WASHINGTON                      	Y
// 	UT	WAYNE                           	Y
// 	UT	WEBER                           	Y
// 	VA	ACCOMACK                        	Y
// 	VA	ALBEMARLE                       	Y
// 	VA	ALEXANDRIA CITY                 	Y
// 	VA	ALLEGHANY                       	Y
// 	VA	AMELIA                          	Y
// 	VA	AMHERST                         	Y
// 	VA	APPOMATTOX                      	Y
// 	VA	ARLINGTON                       	Y
// 	VA	AUGUSTA                         	Y
// 	VA	BATH                            	Y
// 	VA	BEDFORD                         	Y
// 	VA	BEDFORD CITY                    	Y
// 	VA	BLAND                           	Y
// 	VA	BOTETOURT                       	Y
// 	VA	BRISTOL CITY                    	Y
// 	VA	BRUNSWICK                       	Y
// 	VA	BUCHANAN                        	Y
// 	VA	BUCKINGHAM                      	Y
// 	VA	BUENA VISTA CITY                	Y
// 	VA	CAMPBELL                        	Y
// 	VA	CAROLINE                        	Y
// 	VA	CARROLL                         	Y
// 	VA	CHARLES CITY                    	Y
// 	VA	CHARLOTTE                       	Y
// 	VA	CHARLOTTESVILLE CITY            	Y
// 	VA	CHESAPEAKE CITY                 	Y
// 	VA	CHESTERFIELD                    	Y
// 	VA	CLARKE                          	Y
// 	VA	CLIFTON FORGE CITY              	N
// 	VA	COLONIAL HEIGHTS CITY           	Y
// 	VA	COVINGTON CITY                  	Y
// 	VA	CRAIG                           	Y
// 	VA	CULPEPER                        	Y
// 	VA	CUMBERLAND                      	Y
// 	VA	DANVILLE CITY                   	Y
// 	VA	DICKENSON                       	Y
// 	VA	DINWIDDIE                       	Y
// 	VA	EMPORIA CITY                    	Y
// 	VA	ESSEX                           	Y
// 	VA	FAIRFAX                         	Y
// 	VA	FAIRFAX CITY                    	Y
// 	VA	FALLS CHURCH CITY               	Y
// 	VA	FAUQUIER                        	Y
// 	VA	FLOYD                           	Y
// 	VA	FLUVANNA                        	Y
// 	VA	FRANKLIN                        	Y
// 	VA	FRANKLIN CITY                   	Y
// 	VA	FREDERICK                       	Y
// 	VA	FREDERICKSBURG CITY             	Y
// 	VA	GALAX CITY                      	Y
// 	VA	GILES                           	Y
// 	VA	GLOUCESTER                      	Y
// 	VA	GOOCHLAND                       	Y
// 	VA	GRAYSON                         	Y
// 	VA	GREENE                          	Y
// 	VA	GREENSVILLE                     	Y
// 	VA	HALIFAX                         	Y
// 	VA	HAMPTON CITY                    	Y
// 	VA	HANOVER                         	Y
// 	VA	HARRISONBURG CITY               	Y
// 	VA	HENRICO                         	Y
// 	VA	HENRY                           	Y
// 	VA	HIGHLAND                        	Y
// 	VA	HOPEWELL CITY                   	Y
// 	VA	ISLE OF WIGHT                   	Y
// 	VA	JAMES CITY                      	Y
// 	VA	KING AND QUEEN                  	Y
// 	VA	KING GEORGE                     	Y
// 	VA	KING WILLIAM                    	Y
// 	VA	LANCASTER                       	Y
// 	VA	LEE                             	Y
// 	VA	LEXINGTON CITY                  	Y
// 	VA	LOUDOUN                         	Y
// 	VA	LOUISA                          	Y
// 	VA	LUNENBURG                       	Y
// 	VA	LYNCHBURG CITY                  	Y
// 	VA	MADISON                         	Y
// 	VA	MANASSAS CITY                   	Y
// 	VA	MANASSAS PARK CITY              	Y
// 	VA	MARTINSVILLE CITY               	Y
// 	VA	MATHEWS                         	Y
// 	VA	MECKLENBURG                     	Y
// 	VA	MIDDLESEX                       	Y
// 	VA	MONTGOMERY                      	Y
// 	VA	NELSON                          	Y
// 	VA	NEW KENT                        	Y
// 	VA	NEWPORT NEWS CITY               	Y
// 	VA	NORFOLK CITY                    	Y
// 	VA	NORTHAMPTON                     	Y
// 	VA	NORTHUMBERLAND                  	Y
// 	VA	NORTON CITY                     	Y
// 	VA	NOTTOWAY                        	Y
// 	VA	ORANGE                          	Y
// 	VA	PAGE                            	Y
// 	VA	PATRICK                         	Y
// 	VA	PETERSBURG CITY                 	Y
// 	VA	PITTSYLVANIA                    	Y
// 	VA	POQUOSON CITY                   	Y
// 	VA	PORTSMOUTH CITY                 	Y
// 	VA	POWHATAN                        	Y
// 	VA	PRINCE EDWARD                   	Y
// 	VA	PRINCE GEORGE                   	Y
// 	VA	PRINCE WILLIAM                  	Y
// 	VA	PULASKI                         	Y
// 	VA	RADFORD CITY                    	Y
// 	VA	RAPPAHANNOCK                    	Y
// 	VA	RICHMOND                        	Y
// 	VA	RICHMOND CITY                   	Y
// 	VA	ROANOKE                         	Y
// 	VA	ROANOKE CITY                    	Y
// 	VA	ROCKBRIDGE                      	Y
// 	VA	ROCKINGHAM                      	Y
// 	VA	RUSSELL                         	Y
// 	VA	SALEM CITY                      	Y
// 	VA	SCOTT                           	Y
// 	VA	SHENANDOAH                      	Y
// 	VA	SMYTH                           	Y
// 	VA	SOUTH BOSTON CITY               	N
// 	VA	SOUTHAMPTON                     	Y
// 	VA	SPOTSYLVANIA                    	Y
// 	VA	STAFFORD                        	Y
// 	VA	STAUNTON CITY                   	Y
// 	VA	SUFFOLK CITY                    	Y
// 	VA	SURRY                           	Y
// 	VA	SUSSEX                          	Y
// 	VA	TAZEWELL                        	Y
// 	VA	VIRGINIA BEACH CITY             	Y
// 	VA	WARREN                          	Y
// 	VA	WASHINGTON                      	Y
// 	VA	WAYNESBORO CITY                 	Y
// 	VA	WESTMORELAND                    	Y
// 	VA	WILLIAMSBURG CITY               	Y
// 	VA	WINCHESTER CITY                 	Y
// 	VA	WISE                            	Y
// 	VA	WYTHE                           	Y
// 	VA	YORK                            	Y
// 	VI	ST. CROIX                       	Y
// 	VI	ST. JOHN                        	Y
// 	VI	ST. THOMAS                      	Y
// 	VT	ADDISON                         	Y
// 	VT	BENNINGTON                      	Y
// 	VT	CALEDONIA                       	Y
// 	VT	CHITTENDEN                      	Y
// 	VT	ESSEX                           	Y
// 	VT	FRANKLIN                        	Y
// 	VT	GRAND ISLE                      	Y
// 	VT	LAMOILLE                        	Y
// 	VT	ORANGE                          	Y
// 	VT	ORLEANS                         	Y
// 	VT	RUTLAND                         	Y
// 	VT	WASHINGTON                      	Y
// 	VT	WINDHAM                         	Y
// 	VT	WINDSOR                         	Y
// 	WA	ADAMS                           	Y
// 	WA	ASOTIN                          	Y
// 	WA	BENTON                          	Y
// 	WA	CHELAN                          	Y
// 	WA	CLALLAM                         	Y
// 	WA	CLARK                           	Y
// 	WA	COLUMBIA                        	Y
// 	WA	COWLITZ                         	Y
// 	WA	DOUGLAS                         	Y
// 	WA	FERRY                           	Y
// 	WA	FRANKLIN                        	Y
// 	WA	GARFIELD                        	Y
// 	WA	GRANT                           	Y
// 	WA	GRAYS HARBOR                    	Y
// 	WA	ISLAND                          	Y
// 	WA	JEFFERSON                       	Y
// 	WA	KING                            	Y
// 	WA	KITSAP                          	Y
// 	WA	KITTITAS                        	Y
// 	WA	KLICKITAT                       	Y
// 	WA	LEWIS                           	Y
// 	WA	LINCOLN                         	Y
// 	WA	MASON                           	Y
// 	WA	OKANOGAN                        	Y
// 	WA	PACIFIC                         	Y
// 	WA	PEND OREILLE                    	Y
// 	WA	PIERCE                          	Y
// 	WA	SAN JUAN                        	Y
// 	WA	SKAGIT                          	Y
// 	WA	SKAMANIA                        	Y
// 	WA	SNOHOMISH                       	Y
// 	WA	SPOKANE                         	Y
// 	WA	STEVENS                         	Y
// 	WA	THURSTON                        	Y
// 	WA	WAHKIAKUM                       	Y
// 	WA	WALLA WALLA                     	Y
// 	WA	WHATCOM                         	Y
// 	WA	WHITMAN                         	Y
// 	WA	YAKIMA                          	Y
// 	WI	ADAMS                           	Y
// 	WI	ASHLAND                         	Y
// 	WI	BARRON                          	Y
// 	WI	BAYFIELD                        	Y
// 	WI	BROWN                           	Y
// 	WI	BUFFALO                         	Y
// 	WI	BURNETT                         	Y
// 	WI	CALUMET                         	Y
// 	WI	CHIPPEWA                        	Y
// 	WI	CLARK                           	Y
// 	WI	COLUMBIA                        	Y
// 	WI	CRAWFORD                        	Y
// 	WI	DANE                            	Y
// 	WI	DODGE                           	Y
// 	WI	DOOR                            	Y
// 	WI	DOUGLAS                         	Y
// 	WI	DUNN                            	Y
// 	WI	EAU CLAIRE                      	Y
// 	WI	FLORENCE                        	Y
// 	WI	FOND DU LAC                     	Y
// 	WI	FOREST                          	Y
// 	WI	GRANT                           	Y
// 	WI	GREEN                           	Y
// 	WI	GREEN LAKE                      	Y
// 	WI	IOWA                            	Y
// 	WI	IRON                            	Y
// 	WI	JACKSON                         	Y
// 	WI	JEFFERSON                       	Y
// 	WI	JUNEAU                          	Y
// 	WI	KENOSHA                         	Y
// 	WI	KEWAUNEE                        	Y
// 	WI	LA CROSSE                       	Y
// 	WI	LAFAYETTE                       	Y
// 	WI	LANGLADE                        	Y
// 	WI	LINCOLN                         	Y
// 	WI	MANITOWOC                       	Y
// 	WI	MARATHON                        	Y
// 	WI	MARINETTE                       	Y
// 	WI	MARQUETTE                       	Y
// 	WI	MENOMINEE                       	Y
// 	WI	MILWAUKEE                       	Y
// 	WI	MONROE                          	Y
// 	WI	OCONTO                          	Y
// 	WI	ONEIDA                          	Y
// 	WI	OUTAGAMIE                       	Y
// 	WI	OZAUKEE                         	Y
// 	WI	PEPIN                           	Y
// 	WI	PIERCE                          	Y
// 	WI	POLK                            	Y
// 	WI	PORTAGE                         	Y
// 	WI	PRICE                           	Y
// 	WI	RACINE                          	Y
// 	WI	RICHLAND                        	Y
// 	WI	ROCK                            	Y
// 	WI	RUSK                            	Y
// 	WI	SAUK                            	Y
// 	WI	SAWYER                          	Y
// 	WI	SHAWANO                         	Y
// 	WI	SHEBOYGAN                       	Y
// 	WI	ST. CROIX                       	Y
// 	WI	TAYLOR                          	Y
// 	WI	TREMPEALEAU                     	Y
// 	WI	VERNON                          	Y
// 	WI	VILAS                           	Y
// 	WI	WALWORTH                        	Y
// 	WI	WASHBURN                        	Y
// 	WI	WASHINGTON                      	Y
// 	WI	WAUKESHA                        	Y
// 	WI	WAUPACA                         	Y
// 	WI	WAUSHARA                        	Y
// 	WI	WINNEBAGO                       	Y
// 	WI	WOOD                            	Y
// 	WV	BARBOUR                         	Y
// 	WV	BERKELEY                        	Y
// 	WV	BOONE                           	Y
// 	WV	BRAXTON                         	Y
// 	WV	BROOKE                          	Y
// 	WV	CABELL                          	Y
// 	WV	CALHOUN                         	Y
// 	WV	CLAY                            	Y
// 	WV	DODDRIDGE                       	Y
// 	WV	FAYETTE                         	Y
// 	WV	GILMER                          	Y
// 	WV	GRANT                           	Y
// 	WV	GREENBRIER                      	Y
// 	WV	HAMPSHIRE                       	Y
// 	WV	HANCOCK                         	Y
// 	WV	HARDY                           	Y
// 	WV	HARRISON                        	Y
// 	WV	JACKSON                         	Y
// 	WV	JEFFERSON                       	Y
// 	WV	KANAWHA                         	Y
// 	WV	LEWIS                           	Y
// 	WV	LINCOLN                         	Y
// 	WV	LOGAN                           	Y
// 	WV	MARION                          	Y
// 	WV	MARSHALL                        	Y
// 	WV	MASON                           	Y
// 	WV	MCDOWELL                        	Y
// 	WV	MERCER                          	Y
// 	WV	MINERAL                         	Y
// 	WV	MINGO                           	Y
// 	WV	MONONGALIA                      	Y
// 	WV	MONROE                          	Y
// 	WV	MORGAN                          	Y
// 	WV	NICHOLAS                        	Y
// 	WV	OHIO                            	Y
// 	WV	PENDLETON                       	Y
// 	WV	PLEASANTS                       	Y
// 	WV	POCAHONTAS                      	Y
// 	WV	PRESTON                         	Y
// 	WV	PUTNAM                          	Y
// 	WV	RALEIGH                         	Y
// 	WV	RANDOLPH                        	Y
// 	WV	RITCHIE                         	Y
// 	WV	ROANE                           	Y
// 	WV	SUMMERS                         	Y
// 	WV	TAYLOR                          	Y
// 	WV	TUCKER                          	Y
// 	WV	TYLER                           	Y
// 	WV	UPSHUR                          	Y
// 	WV	WAYNE                           	Y
// 	WV	WEBSTER                         	Y
// 	WV	WETZEL                          	Y
// 	WV	WIRT                            	Y
// 	WV	WOOD                            	Y
// 	WV	WYOMING                         	Y
// 	WY	ALBANY                          	Y
// 	WY	BIG HORN                        	Y
// 	WY	CAMPBELL                        	Y
// 	WY	CARBON                          	Y
// 	WY	CONVERSE                        	Y
// 	WY	CROOK                           	Y
// 	WY	FREMONT                         	Y
// 	WY	GOSHEN                          	Y
// 	WY	HOT SPRINGS                     	Y
// 	WY	JOHNSON                         	Y
// 	WY	LARAMIE                         	Y
// 	WY	LINCOLN                         	Y
// 	WY	NATRONA                         	Y
// 	WY	NIOBRARA                        	Y
// 	WY	PARK                            	Y
// 	WY	PLATTE                          	Y
// 	WY	SHERIDAN                        	Y
// 	WY	SUBLETTE                        	Y
// 	WY	SWEETWATER                      	Y
// 	WY	TETON                           	Y
// 	WY	UINTA                           	Y
// 	WY	WASHAKIE                        	Y
// 	WY	WESTON                          	Y

// 	Note: The above counties are used for control points and locations

// CS 	Coser Activity Type
// 	C	Cancel
// 	N	New
// 	O	On Air Test
// 	R	Request for Reconsideration

// EN	Entity Type
// 	CE	Transferee contact
// 	CL	Licensee Contact
// 	CR	Assignor or Transferor Contact
// 	CS	Lessee Contact
// 	E 	Transferee
// 	L 	Licensee or Assignee
// 	O 	Owner
// 	R 	Assignor or Transferor
// 	S 	Lessee

// EN	Applicant Type Code
// 	Code	Description		Code Active?
// 	B 	Amateur Club                  	A
// 	C 	Corporation                   	A
// 	D 	General Partnership           	A
// 	E 	Limited Partnership           	A
// 	F 	Limited Liability Partnership 	A
// 	G 	Governmental Entity           	A
// 	H 	Other                         	A
// 	I 	Individual                    	A
// 	J 	Joint Venture                 	N
// 	L 	Limited Liability Company     A
// 	M 	Military Recreation           	A
// 	O 	Consortium                    	A
// 	P 	Partnership                   	N
// 	R 	RACES                         	A
// 	T 	Trust                         	A
// 	U 	Unincorporated Association   A

// EN	Status Code
// 	Null/Blank	Active
// 	X	Termination Pending
// 	T	Terminated
// 	Note: Status Code also appears on the following record types and uses the same definitions:
// 		"MW, CG, LM, CO, AS, SC, SF, BO, CP, LS, LO, LF, OP, BL, AN, RC, FR, F2, IR, CS, FS, FF, BF, RA, EM, PC, PA,  SG, L3, L4, O2, L5, L6, A3, F3, F4, F5, F6, P2, TP"

// EN	3.7 GHz License Type
// 	Code	Description
// 	C	Combo
// 	R	Combo-R
// 	F	Final
// 	I	Interim

// FA	Operator Class Code
// 	DB	GMDSS Radio Operator/Maintainer License
// 	DM	GMDSS Radio Maintainer's License
// 	DO	GMDSS Radio Operator's License
// 	MP	Marine Radio Operator Permit
// 	PG	General Radiotelephone Operator License
// 	RG	GMDSS Restricted Radio Operator License
// 	RL	Restricted Radiotelephone Operator Permit-Limited Use
// 	RR	Restricted Radiotelephone Operator Permit
// 	T 	Radiotelegraph Operator License
// 	T1	First Class Radiotelegraph Operator's Certificate
// 	T2	Second Class Radiotelegraph Operator's Certificate
// 	T3	Third Class Radiotelegraph Operator's Certificate

// FC	Coordinator Name
// 	02         	"Electronic Technicians Assoc. International, Inc.                                    "
// 	03         	"Elkins International, Inc                                                                  "
// 	04         	ISCET
// 	06         	Sea School
// 	07         	"Sylvan Learning Systems, Inc.                                                              "
// 	08         	"RABQSA International, Inc                                                                  "
// 	09         	National Radio Examiners
// 	11         	BFT Training Unlimited Inc
// 	13         	Laser Grade Computer Testing Inc
// 	14         	Marine Technical Institute
// 	16         	CATS d.b.a. Comira
// 	17         	Mariners Learning System
// 	18         	Kennedy Point Maritime School
// 	19         	National Marine Electronics Association
// 	20         	US Captain's Training
// 	21         	Test Placeholder
// 	A          	"The Milwaukee RA Club, Inc.                                                                "
// 	AAA        	AAA Frequency Coordination
// 	AAHTO      	American Association of State Highway and Transpor
// 	AAR        	Association of American Railroads
// 	ameri      	Americas Communications LLC
// 	AMPTP      	Alliance of Motion Picture & TV Producer
// 	AMTA       	American Mobile Telecommunications Association
// 	APCO       	Associated Public Safety Com Officrs Inc
// 	ATA        	American Trucking Association Inc
// 	ATTC       	AT&T CORP
// 	ATTMb      	AT&T Mobility LLC
// 	B          	"Laurel AR Club, Inc.                                                                       "
// 	BA         	Black & Associates
// 	C          	Anchorage AR Club
// 	CBSCm      	CBS Communications Services Inc
// 	CCS        	Consolidated Spectrum Services
// 	CET        	SiteSafe
// 	COMSH      	Comsearch
// 	CSAA       	Central Station Alarm Association
// 	D          	ARRL/VEC
// 	DGLC       	Digis LLC
// 	E          	W4VEC
// 	EKW        	Edwards and Kelcey Wireless
// 	EWA        	Enterprise Wireless Alliance
// 	F          	Greater LA AR Group
// 	FCCA       	Forestry Conservation Communications Association
// 	FCCIN      	FCC Internal Filer
// 	FIT        	Forest Industries Telecommunications
// 	FreqE      	FreqEasy LLC
// 	G          	Sunnyvale VEC ARC Inc
// 	H          	"W. Carolina AR Soc/VEC, Inc.                                                               "
// 	IMSA       	International Municipal Signal Association (IMSA)
// 	IPATH      	Intelpath
// 	ITA        	"Industrial telecommunications Assoc, Inc.                                                  "
// 	ITLA       	International Taxicab and Livery Assoc
// 	ITU1       	FCC-Gettysburg
// 	ITU2       	FCC-Gettysburg
// 	ITU3       	FCC-Gettysburg
// 	J          	Golden Empire AR Society
// 	K          	Central America VEC Inc
// 	L          	W5YI-VEC
// 	M          	MO-KAN VE Coordinator
// 	MNCI       	Micronet Communications Inc.
// 	MPI        	Microwave Planning Inc.
// 	MRFAC      	Manufacturers Radio Freq Adv Comm Inc
// 	N          	Sandarc-Vec
// 	NAA        	Newspaper Association of America
// 	NEXTL      	NEXTEL
// 	NOAA       	NOAA National Weather Service
// 	OTHER      	COORDINATOR
// 	PCIA       	Personal Communications Industry Association
// 	PEGBn      	"PEG Bandwidth, LLCPEG Bandwidth, LLC                                                       "
// 	PFCC       	Petroleum Frequency Coordinate Committee
// 	place      	Future Placeholder
// 	PRTNR      	NEXTEL PARTNERS
// 	Q          	Jefferson AR Club
// 	RDC        	"Radyn, Inc.                                                                                "
// 	RDSFT      	Radiosoft
// 	SERS       	Special Emerg Rad Svc c/o PCIA/IMSA/IAFC
// 	TELFA      	Telephone Maint Freq Advisory Committee
// 	USC        	UTC Service Corporation
// 	UTC        	Utilities Telecom Council
// 	WACor      	Wireless Applications Corp.
// 	WIA        	The Wireless Infrastructure Association

// FR	Operational Altitude Code
// 	HL	High Level
// 	HO	Helicopter Operations
// 	LL	Low Level
// 	ML	Medium Level
// 	RL	Ramp Level

// FT	Frequency Type Code
// 	L	Local Control
// 	G	Ground Control
// 	E	Emergency (121.5 MHz)
// 	O	Other

// F2	Frequency Channel Block
// 	Channel Block	Freq. Assigned	Freq. Upper Band	Radio Service Code
// 	1     	2150.00000000	2156.00000000	MD
// 	1     	2150.00000000	2156.00000000	VX
// 	172   	5855.00000000	5865.00000000	IQ
// 	174   	5865.00000000	5875.00000000	IQ
// 	174   	5865.00000000	5875.00000000	QQ
// 	175   	5865.00000000	5885.00000000	IQ
// 	175   	5865.00000000	5885.00000000	QQ
// 	176   	5875.00000000	5885.00000000	IQ
// 	176   	5875.00000000	5885.00000000	QQ
// 	178   	5885.00000000	5895.00000000	IQ
// 	178   	5885.00000000	5895.00000000	QQ
// 	180   	5895.00000000	5905.00000000	IQ
// 	180   	5895.00000000	5905.00000000	QQ
// 	181   	5895.00000000	5915.00000000	IQ
// 	181   	5895.00000000	5915.00000000	QQ
// 	182   	5905.00000000	5915.00000000	IQ
// 	182   	5905.00000000	5915.00000000	QQ
// 	184   	5915.00000000	5925.00000000	IQ
// 	2     	2156.00000000	2162.00000000	MD
// 	2     	2156.00000000	2162.00000000	VX
// 	2A    	2156.00000000	2160.00000000	MD
// 	2A    	2156.00000000	2160.00000000	VX
// 	A1    	2500.00000000	2506.00000000	MD
// 	A1    	2500.00000000	2506.00000000	VX
// 	A2    	2512.00000000	2518.00000000	MD
// 	A2    	2512.00000000	2518.00000000	VX
// 	A3    	2524.00000000	2530.00000000	MD
// 	A3    	2524.00000000	2530.00000000	VX
// 	A4    	2536.00000000	2542.00000000	MD
// 	A4    	2536.00000000	2542.00000000	VX
// 	B1    	2506.00000000	2512.00000000	MD
// 	B1    	2506.00000000	2512.00000000	VX
// 	B2    	2518.00000000	2524.00000000	MD
// 	B2    	2518.00000000	2524.00000000	VX
// 	B3    	2530.00000000	2536.00000000	MD
// 	B3    	2530.00000000	2536.00000000	VX
// 	B4    	2542.00000000	2548.00000000	MD
// 	B4    	2542.00000000	2548.00000000	VX
// 	C1    	2548.00000000	2554.00000000	MD
// 	C1    	2548.00000000	2554.00000000	VX
// 	C2    	2560.00000000	2566.00000000	MD
// 	C2    	2560.00000000	2566.00000000	VX
// 	C3    	2572.00000000	2578.00000000	MD
// 	C3    	2572.00000000	2578.00000000	VX
// 	C4    	2584.00000000	2590.00000000	MD
// 	C4    	2584.00000000	2590.00000000	VX
// 	D1    	2554.00000000	2560.00000000	MD
// 	D1    	2554.00000000	2560.00000000	VX
// 	D2    	2566.00000000	2572.00000000	MD
// 	D2    	2566.00000000	2572.00000000	VX
// 	D3    	2578.00000000	2584.00000000	MD
// 	D3    	2578.00000000	2584.00000000	VX
// 	D4    	2590.00000000	2596.00000000	MD
// 	D4    	2590.00000000	2596.00000000	VX
// 	E1    	2596.00000000	2602.00000000	MD
// 	E1    	2596.00000000	2602.00000000	VX
// 	E2    	2608.00000000	2614.00000000	MD
// 	E2    	2608.00000000	2614.00000000	VX
// 	E3    	2620.00000000	2626.00000000	MD
// 	E3    	2620.00000000	2626.00000000	VX
// 	E4    	2632.00000000	2638.00000000	MD
// 	E4    	2632.00000000	2638.00000000	VX
// 	F1    	2602.00000000	2608.00000000	MD
// 	F1    	2602.00000000	2608.00000000	VX
// 	F2    	2614.00000000	2620.00000000	MD
// 	F2    	2614.00000000	2620.00000000	VX
// 	F3    	2626.00000000	2632.00000000	MD
// 	F3    	2626.00000000	2632.00000000	VX
// 	F4    	2638.00000000	2644.00000000	MD
// 	F4    	2638.00000000	2644.00000000	VX
// 	G1    	2644.00000000	2650.00000000	MD
// 	G1    	2644.00000000	2650.00000000	VX
// 	G2    	2656.00000000	2662.00000000	MD
// 	G2    	2656.00000000	2662.00000000	VX
// 	G3    	2668.00000000	2674.00000000	MD
// 	G3    	2668.00000000	2674.00000000	VX
// 	G4    	2680.00000000	2686.00000000	MD
// 	G4    	2680.00000000	2686.00000000	VX
// 	H1    	2650.00000000	2656.00000000	MD
// 	H1    	2650.00000000	2656.00000000	VX
// 	H2    	2662.00000000	2668.00000000	MD
// 	H2    	2662.00000000	2668.00000000	VX
// 	H3    	2674.00000000	2680.00000000	MD
// 	H3    	2674.00000000	2680.00000000	VX
// 	I01   	2686.00000000	2686.12500000	MD
// 	I01   	2686.00000000	2686.12500000	VX
// 	I02   	2686.12500000	2686.25000000	MD
// 	I02   	2686.12500000	2686.25000000	VX
// 	I03   	2686.25000000	2686.37500000	MD
// 	I03   	2686.25000000	2686.37500000	VX
// 	I04   	2686.37500000	2686.50000000	MD
// 	I04   	2686.37500000	2686.50000000	VX
// 	I05   	2686.50000000	2686.62500000	MD
// 	I05   	2686.50000000	2686.62500000	VX
// 	I06   	2686.62500000	2686.75000000	MD
// 	I06   	2686.62500000	2686.75000000	VX
// 	I07   	2686.75000000	2686.87500000	MD
// 	I07   	2686.75000000	2686.87500000	VX
// 	I08   	2686.87500000	2687.00000000	MD
// 	I08   	2686.87500000	2687.00000000	VX
// 	I09   	2687.00000000	2687.12500000	MD
// 	I09   	2687.00000000	2687.12500000	VX
// 	I10   	2687.12500000	2687.25000000	MD
// 	I10   	2687.12500000	2687.25000000	VX
// 	I11   	2687.25000000	2687.37500000	MD
// 	I11   	2687.25000000	2687.37500000	VX
// 	I12   	2687.37500000	2687.50000000	MD
// 	I12   	2687.37500000	2687.50000000	VX
// 	I13   	2687.50000000	2687.62500000	MD
// 	I13   	2687.50000000	2687.62500000	VX
// 	I14   	2687.62500000	2687.75000000	MD
// 	I14   	2687.62500000	2687.75000000	VX
// 	I15   	2687.75000000	2687.87500000	MD
// 	I15   	2687.75000000	2687.87500000	VX
// 	I16   	2687.87500000	2688.00000000	MD
// 	I16   	2687.87500000	2688.00000000	VX
// 	I17   	2688.00000000	2688.12500000	MD
// 	I17   	2688.00000000	2688.12500000	VX
// 	I18   	2688.12500000	2688.25000000	MD
// 	I18   	2688.12500000	2688.25000000	VX
// 	I19   	2688.25000000	2688.37500000	MD
// 	I19   	2688.25000000	2688.37500000	VX
// 	I20   	2688.37500000	2688.50000000	MD
// 	I20   	2688.37500000	2688.50000000	VX
// 	I21   	2688.50000000	2688.62500000	MD
// 	I21   	2688.50000000	2688.62500000	VX
// 	I22   	2688.62500000	2688.75000000	MD
// 	I22   	2688.62500000	2688.75000000	VX
// 	I23   	2688.75000000	2688.87500000	MD
// 	I23   	2688.75000000	2688.87500000	VX
// 	I24   	2688.87500000	2689.00000000	MD
// 	I24   	2688.87500000	2689.00000000	VX
// 	I25   	2689.00000000	2689.12500000	MD
// 	I25   	2689.00000000	2689.12500000	VX
// 	I26   	2689.12500000	2689.25000000	MD
// 	I26   	2689.12500000	2689.25000000	VX
// 	I27   	2689.25000000	2689.37500000	MD
// 	I27   	2689.25000000	2689.37500000	VX
// 	I28   	2689.37500000	2689.50000000	MD
// 	I28   	2689.37500000	2689.50000000	VX
// 	I29   	2689.50000000	2689.62500000	MD
// 	I29   	2689.50000000	2689.62500000	VX
// 	I30   	2689.62500000	2689.75000000	MD
// 	I30   	2689.62500000	2689.75000000	VX
// 	I31   	2689.75000000	2689.87500000	MD
// 	I31   	2689.75000000	2689.87500000	VX

// HD	License Status
// 	A	Active
// 	C	Canceled
// 	E	Expired
// 	L	Pending Legal Status
// 	P	Parent Station Canceled
// 	T	Terminated
// 	X	Term Pending

// HD	Developmental/STA/Demonstration
// 	D 	Developmental
// 	M	Demonstration
// 	N	Regular
// 	S	Special Temporary Authority (STA)

// HD	Operation/Performance Requirement Certification
// 	Code	Description
// 	A	Site-Based License
// 	B	Geographic license, commercial service - licensee in its initial license term with an interim performance requirement
// 	C	Geographic license, commercial service - licensee in its initial license term with no interim performance requirement
// 	D	Geographic license, commercial service - licensee in any subsequent term
// 	E	Geographic license, private systems - licensee in its initial license term with an interim performance requirement
// 	F	Geographic license, private systems - licensee in its initial license term with no interim performance requirement
// 	G	Geographic license, private systems - licensee in any subsequent term
// 	H	Partitioned or disaggregated license without a performance requirement, for the first renewal application filed after date of deployment
// 	I	Partitioned or disaggregated license without a performance requirement, for any subsequent renewal filings

// HS	History Code
// 	10MCOM      	10 MHz Geographic Overlap Review Completed
// 	172COM      	Review of Channel 172 Completed
// 	20COM       	2065/2079 Review Completed
// 	38GOFF		Offlined For 38.6 - 40 GHz Review.
// 	602COM      	Form 602 Review Completed
// 	602OFF		Offlined for Form 602 Review
// 	603TCN		Converted from a 603T application.
// 	700COM      	Offlined for 700 MHz Review Completed
// 	702COM      	700 MHz Canadian Treaty Compliance completed
// 	703COM      	700 MHz Mexican Treaty Compliance completed
// 	704COM      	Canadian/Mexican border check completed
// 	72CCOM      	72-76 MHz Coordination Completed (Channel 4/5)
// 	72OCOM      	72-76 MHz Coordination Completed (Op Fixed)
// 	800COM      	Completed 800 MHz Band Reconfiguration Review
// 	929COM      	929.0125-929.4875 /MHz review completed
// 	ACSCOM      	Associated Call Sign Review Completed
// 	ADDCOM      	Manual Address Correction Complete
// 	AFCOM       	AFTRCC Review Completed
// 	AFRCOM      	Application for Review Consideration Completed
// 	AFROFF		Offlined for Application for Review Consideration
// 	ALCOM       	Alert List Review Completed
// 	ALOCOM      	Already Occurred Review Completed
// 	ALOFF 		Offlined for Alert List Review
// 	AMACOM      	Review of Trustee Change Completed
// 	AMSCOM      	AM Station Review Completed
// 	ANATRM		Antenna Terminated Through Auto Term Process
// 	AOCCOM      	CA county coordination review completed
// 	AOCOM       	Alien Ownership Review Completed
// 	AONCOM      	Area of Operation = N or U review completed
// 	AOOFF 		Offlined for Alien Ownership Review
// 	AOPCOM      	"Coordinate review (N/Line A, E/Line C) completed"
// 	APACC 		Application Accepted
// 	APCON 		Application Consented
// 	APCRTC      	Litigation or Court Action Closed
// 	APDIS 		Application Dismissed
// 	APDIS1		Application Dismissed due to Aged Return
// 	APDIS2		Application Dismissed due to Parent Facility Expir
// 	APGRP 		Application Granted in Part
// 	APGRT 		Application Granted
// 	APOCOM      	Coser Spec Cond Review (155.475 MHz) completed
// 	APPCOM      	Change of applicant type code review completed
// 	APPOFF		Offlined change of applicant type code review
// 	APRTN 		Application Returned
// 	APRVCC      	Canceled Coser Application Review Completed
// 	APRVIC      	Canceled Iracs Application Review Completed
// 	APTR2 		Application Set to Returned Status.
// 	APWD  		Application Withdrawn
// 	ARECOM      	Area of Operation Review Completed
// 	ASCCOM      	Associated Call Sign Review Completed
// 	ASROFF		Offlined for ASR review
// 	ATCOM       	Attachment Review Completed
// 	ATOFF 		Offlined for Attachment Review
// 	ATRCOM      	Applicant Type Review Completed
// 	ATTCOM      	Attachment Received
// 	AUTCOM      	Auto Termination Review Completed
// 	AUTHPC      	Consent Authorization Printed
// 	AUTHPR      	Authorization Printed
// 	AUTHRP      	Reference Copy Authorization Printed
// 	AUTOFF		Offlined for Auto Termination Review
// 	AUTRLP      	Registration Letter Printed
// 	BARCOM      	Station transmit loc outside CDBS station complete
// 	BAROFF		Station transmit loc outside CDBS parent station.
// 	BIDCOM      	Bidding Credit data Review Completed
// 	BILCOM      	Offline for Billing Completed
// 	BILOFF		Offlined for Billing
// 	BKPCOM      	Bankruptcy Review Completed
// 	BLDCOM      	Buildout Notification Review Completed
// 	BNDCOM      	Bandwidth Review Completed.
// 	BNDOFF		Bandwidth Review (Coser/Irac Major/Minor Failed).
// 	BOCOM       	NT/EX Buildout Untimely Filing Review Complete
// 	BOOFF 		Offlined for NT/EX Buildout Untimely Filing Review
// 	BORDCM      	Border Coordinate Review Completed
// 	BORDOF		Offlined for Border Coordinate Review
// 	BQCOM       	Basic Qualification Review Completed
// 	BQOFF 		Offlined for Basic Qualification Review
// 	BRECOM      	BR/ED Review Completed
// 	CAECOM      	Review of Coser Errors Completed
// 	CAEOFF		Review Coser Errors
// 	CANCOM      	Manual Coser Review Completed
// 	CANCSE		Cancel Coser Sent to Canada
// 	CANFRC      	Forced Coser Review Completed
// 	CANOSE		On-Air Coser Sent to Canada
// 	CANRE       	Received from Canada
// 	CANRSE		Recon Coser Sent to Canada
// 	CANSC 		Coser Screening Completed
// 	CANSE 		New Coser Sent to Canada
// 	CECOM       	CE Review Completed
// 	CEOFF 		Offlined for CE Review
// 	CFBCOM      	C/F Block Review Completed
// 	CG2COM      	CGSA review completed
// 	CGSCOM      	Alt CGSA study completed
// 	CIBCOM      	CIB Action Completed
// 	CLICOF		Review pending App for removal of canceled license
// 	CLSCOM      	Closed Bidding data Review Completed
// 	CMTCOM      	T1/T2/T3 Commercial License Review Complete
// 	CNRCOM      	Canada/Narrowbanding Review completed
// 	COMCOM      	Compliance with Section 27.1203 Review Completed
// 	CONEX 		Construction Date Extended
// 	COR   		Internal Correction Applied
// 	CORADD		Internal Correction - License Added
// 	COVNT 		Required Notification of Const/Coverage
// 	CPCOM       	Review for CPOFF Completed
// 	CPLCOM      	CP Location Review Completed
// 	CPOCOM      	Review for CP OFF Air Completed
// 	CSACOM      	Four-Letter Call Sign Assignment Complete
// 	CSPCOM      	Coser Spec Cond Review (155.475 MHz) Completed
// 	CTACOM      	Cable TV/Attributable Interest Review Completed
// 	CZCCOM      	Coordination Zone Check Completed
// 	CZCOFF		Offlined for Coordination Zone Check
// 	CZRCOM      	Canadian Zone Review Complete
// 	CZROFF		Offlined for Canadian Zone Review
// 	DBCOM       	DO/DM Consolidation Review Complete
// 	DCIACO      	DCIA FRN Review Completed
// 	DEMCOM      	Demonstration Review Completed
// 	DEMOFF		Offlined for Demonstration Review
// 	DEVCOM      	Developmental Review Completed
// 	DEVOFF		Offlined for Developmental Review
// 	DFTCOM      	Review application for criteria of R&O FCC 05-144
// 	DIREV 		Dismissed Application Reversed
// 	DMCCOM      	Demonstration - Coser clearance closed
// 	DMMCOM      	Demonstration - Coser manual offline closed
// 	DMZCOM      	Demonstration - Coser clearance zone check closed
// 	DSHCOM      	Manual Domestic MMSI Number Assignment Complete
// 	DTBCOM      	Date of Birth Review Completed
// 	DTSCOM      	Cable TV/Attributable Interest Review Completed
// 	DWVCOM      	Waiver/Deferral of Fees Review Completed
// 	EDRCOM      	Expiration Date Review Completed
// 	EFCAPR		Application receipt email failed: CORES email
// 	EFCFND		FRN Disassociation email failed: CORES email
// 	EFCFNR		FRN Re-association email failed: CORES email
// 	EFCFRN		FRN Association email failed: CORES email
// 	EFRCOM      	Electronic Filing Review Completed
// 	EFROFF		Offline for Mandatory Electronic Filing Review
// 	EFUAPR		Application receipt email failed: ULS email
// 	EFUFND		FRN Disassociation email failed: ULS email
// 	EFUFNR		FRN Re-association email failed: ULS email
// 	EFUFRN		FRN Association email failed: ULS email
// 	EFURNW	Public Safety Renewal email failed: ULS email
// 	ELGCOM      	Eligibility Review Completed
// 	ENFCOM      	Enforcement Bureau Action Completed
// 	ENFOFF		Offlined for Enforcement Bureau Action
// 	ENGCOM      	Engineering Review Completed
// 	ENGOFF		Offlined for Engineering Review
// 	ERPCOM      	Radial ERP greater than 450 watts completed
// 	ESCAPR		Application receipt email sent: CORES email
// 	ESCFND		FRN Disassociation email sent: CORES email
// 	ESCFNR		FRN Re-association email sent: CORES email
// 	ESCFRN		FRN Association email sent: CORES email
// 	ESUAPR		Application receipt email sent: ULS email
// 	ESUFND		FRN Disassociation email sent: ULS email
// 	ESUFNR		FRN Re-association email sent: ULS email
// 	ESUFRN		FRN Association email sent: ULS email
// 	ESURNW	Public Safety Renewal email sent: ULS email
// 	EXCOM       	Extension of Time Review Completed
// 	EXPCOM      	Manual License Status Change Completed
// 	EXPOFF		Offlined for Manual License Status Change
// 	EXTCOM      	Extended implementation review completed
// 	FCRCOM      	Frequency Coordination Review Completed
// 	FCROFF		Offlined for Frequency Coordination Review
// 	FORCOM      	Forbearance/Already Occurred Review Completed
// 	FRATRM		Frequency Terminated Through Auto Termination Proc
// 	FRCCOM      	FRC_OP_Class Code/Endorsement review complete
// 	FRDCOM      	Forbearance Date Review Completed
// 	FRSCOM      	Invalid frequency/radio service review completed
// 	FVCOM       	Fee Validation Completed
// 	FVDCOM      	Fee Determination Completed
// 	FVDOFF		Offlined for Fee Determination
// 	FVPCNF		Payment Confirmed
// 	FVPCOM      	Payment Verification Completed
// 	FVPOFF		Offlined for Payment Verification
// 	FXCOM       	Faxed Attachment Review Completed
// 	GCOCOM      	Review of Grant/Consumm Failure Completed
// 	GEOCOM      	Geographic Overlap Review Completed
// 	GFRCOM      	Grandfathered Station Review Completed
// 	GFROFF		Offlined for Grandfathered Station Review
// 	GMRCOM      	GMRS Application Review Completed
// 	GRCOM       	Grant Fail Review Completed
// 	GROFF 		Review Application.  Grant Failed.
// 	GRREV 		Granted Application Reversed
// 	GSOCOM      	GSO Review Completed
// 	GSOOFF		Offlined for GSO Review
// 	HFRCOM      	HF Enroute/App 27 Review Completed
// 	HRGCOM      	Hearing Review Completed
// 	HRGOFF		Offlined for Hearing
// 	IFRBRE      	Received from IFRB
// 	INACOM      	Inactive Call Sign Review Complete
// 	INSCOM      	Installment Payment data Review Completed
// 	INTCOM      	Location within 72 km of border completed
// 	INTDUP		Internal Duplicate Requested
// 	INVCOM      	Involuntary/Invountary Date Review Completed
// 	IQCOM       	Registered Location Coord Review Completed
// 	IRACC2		Condition 2 Received from IRAC
// 	IRACC3		Condition 3 Received from IRAC
// 	IRACCS		Cancel IRAC Sent
// 	IRACDW		Generate Dismissed or Withdrawn IRAC
// 	IRACFS		FAA/IRAC--Sent to OET
// 	IRACRC      	Frequency offline review complete
// 	IRACRE      	Received from IRAC
// 	IRACRO		Clearance returned from Irac with objection
// 	IRACRV		Offlined for Irac STA/Demonstration Review
// 	IRACSC		IRAC Screening Completed
// 	IRACSE		IRAC--Sent to OET
// 	IRAFRC      	Forced Irac Review Completed
// 	IRAFRV		Review - Forced Irac Canceled Due to Major Mod
// 	IRAOFC		Offlined for Open Cancel AI on Inactive App Review
// 	ISHCOM      	Manual Intl MMSI Number Assignment Complete
// 	LBOCOM      	NT/EX Lower Buildout Untimely Filing Review Comple
// 	LDACOM      	Licensing Division Analysis Completed
// 	LEASGE		Lease Generated
// 	LEAUA 		Lease Administrative Update Applied
// 	LECAN 			Lease Canceled
// 	LEEXP 			Lease Status Set to Expired
// 	LEEXT 			Lease Extention
// 	LEGCOM      	Legal Review Completed
// 	LEGOFF			Offlined for Legal Review
// 	LEISS 			Lease Issued
// 	LEMOD 			Lease Modified
// 	LESLPR      	Lease Letter Sent
// 	LETACS      	Audit Cancellation Letter Sent
// 	LETATS      	Audit Termination Letter Sent
// 	LETCAS      	Cancellation Letter Sent
// 	LETCMS      	Consummation Reminder Letter Sent
// 	LETCNS      	Construction/Coverage Reminder Letter Sent
// 	LETCOS      	Coser HIA Return Letter Sent
// 	LETDCS      	Failure to Notify Dismissal Letter Sent
// 	LETDFS      	Deleted Frequency Letter Sent
// 	LETDIS      	Dismissal Letter Sent
// 	LETDLS      	Unprocessable Dismissal Letter Sent for Live Money
// 	LETDUS      	Unprocessable Dismissal Letter Sent
// 	LETERM			Lease Terminated
// 	LETGIS      	Grant in Part Letter Sent
// 	LETLTS      	License Termination Letter Sent
// 	LETNNS      	Aircraft Termination Letter Sent
// 	LETRAN			Lease Transferred
// 	LETREG			Generate Renewal Reminder Letter
// 	LETRES      	Renewal Reminder Letter Sent
// 	LETRTS      	Return Letter Sent
// 	LETRUS      	Unprocessable Return Letter Sent
// 	LETSTS      	Ship Termination Letter Sent
// 	LETTRS			License Auto Termination Letter Sent
// 	LEXCOM      	Expired License Review Complete
// 	LEXOFF			Offlined for Expired License Review
// 	LIASS 			License Assigned (Full Assignment)
// 	LIAUA 			Administrative Update Applied
// 	LICAN 			License Canceled
// 	LICARE			License Cancellation Reversed
// 	LICCNV			License Converted
// 	LIDUP 			Duplicate Authorization Generated
// 	LIEXP 			License Status Set to Expired
// 	LIISS 			License Issued
// 	LIMOD 			License Modified
// 	LIPASS			License Assigned (Partial Assignment)
// 	LIPND 			License Assigned (Partition/Disaggregation)
// 	LIPPND			License Issued from a Partial/P&D Assignment
// 	LIREN 			License Renewed
// 	LIRIN1			License reinstated-parent station was un-canceled
// 	LIRIN2			License reinstated-parent station was undismissed
// 	LITEAT			License Terminated by Auto Term Process
// 	LITERE			License Termination Reversed
// 	LITERM			License Terminated
// 	LITIN 			License TIN Added
// 	LITRAN			License Transferred
// 	LMNCOM      	Licensee Name Mismatch Review Complete
// 	LMNOFF			Offlined for Licensee Name Mismatch
// 	LMSCOM      	LMS Review Completed
// 	LNMCOM      	Licensee name change review completed
// 	LNMOFF			Licensee name changed in CDBS
// 	LOATRM			Location Terminated Through Auto Termination Proce
// 	LOCCOM      	Location Review Completed
// 	LOCOFF			Offlined for Location Review
// 	LSLCOM      	P&D of Leased Services Review Completed
// 	LTSAPR			Application Receipt Letter sent
// 	LTSFND			FRN Disassociation Letter sent
// 	LTSFNR			FRN Re-association Letter sent
// 	LTSFRN			FRN Association Letter sent
// 	MAPCOM      	603/608 Map Upload Review Completed
// 	MARCOM      	RO/RM Market application review complete
// 	MAROFF			Offlined for RO/RM Market application review
// 	MCCOM       	MC - FC Cancellation Review Completed
// 	MDONC       	"This Call Sign previously P&Ded, review channels    "
// 	MDRCOM      	Mail Delivery Review Completed
// 	MEXCOM      	Mexico Coordination Review Completed
// 	MEXOFF			Offlined for Mexico Coordination
// 	MEXRE       	Received from Mexico
// 	MJCOM       	Major Amendment Review Completed
// 	MJREV 			Offlined for Major Amendment Review
// 	MMBCOM      	MMB Referral Completed
// 	MMBOFF			Offlined for MMB Referral
// 	MMCOM       	Multiple Mod (Conflict) Review Completed
// 	MMMCOM      	Major Merger Review Completed
// 	MMMOFF			Offlined for Major Merger Review
// 	MMOFF 			Offlined for Multiple Mod (Conflict) Review
// 	MMWCOM      	Merger Review Completed
// 	MMWOFF			Offlined for Merger Review
// 	MOBCOM      	Mobile Count Review Completed
// 	MOBOFF			Offlined for Mobile Count
// 	MODCOM      	 Offlined for Change in Receive Location Closed
// 	MODFLD      	Assignor modified a callsign - Recheck frequencies
// 	MODOFF			Offlined for Change in Receive Location
// 	MRSCOM      	NT/EX Multiple Private/Commercial Review Complete
// 	MRSOFF			Offlined for NT/EX Multiple Priv/Comm Review
// 	MTSCOM      	MTS Completed
// 	MTSOFF			Offlined for MTS
// 	MULCOM      	MULTICOM Review Completed
// 	MWLCOM      	MW Link Transmit and Receive Coordinates verificat
// 	MXCOM       	MX Processing Completed
// 	MXNCOM      	Mutually Exclusive Review (no rules) Completed
// 	MXRCOM      	Mutually Exclusive Review (rules) Completed
// 	MZRCOM      	Mexican Zone Review Complete
// 	MZROFF			Offlined for Mexican Zone Review
// 	NEPCOM      	NEPA Review Completed
// 	NEPOFF			Offlined for NEPA Review
// 	NNCOM       	N-Number Conflict with Amateur Call Sign Review Co
// 	NPCOM       	Northern Pacific Review Completed
// 	NTACOM      	NT/EX Amendment Review Completed
// 	NTAOFF			Offlined for NT/EX Amendment Review
// 	NTBCOM      	NT Buildout Deadline Review completed
// 	NTECOM      	NT/EX Error Review Complete
// 	NTEOFF			Offlined for NT/EX Error Review
// 	NTGCOM      	NT Grant Hold Review Complete
// 	NTGOFF			Offlined for NT Grant Hold
// 	NTICOM      	NTIA Coordination Review Completed
// 	NWACOM      	Non-Wireless Activity Review Completed
// 	PAATRM			Path Terminated Through Auto Term Process
// 	PARREB			"Parent Station, Once Silent, Is Broadcasting Again"
// 	PARSIL			Parent Station Went Silent
// 	PCCOM       	Private/Commercial Review Completed
// 	PCECOM      	Review of Pre-Canada Errors Completed
// 	PCEOFF			Review Pre-Canada Errors
// 	PCKCOM      	Pack Review Completed
// 	PCKOFF			Offlined for Pack Review
// 	PCRCOM      	Public Coast Review Completed
// 	PFCOM       	Public Fixed Review Completed
// 	PFDCOM      	Review of PFD at Border Completed
// 	PFDOFF			Offline for PFD at Border Review
// 	PFRCL 			Petition for Reconsideration Period Closed
// 	PFRCOM      	Petition for Reconsideration Review Completed
// 	PFRRE 			Petition for Reconsideration Received
// 	PH2COM      	Review of Phase 2 application completed
// 	PLAUPR			Paperless Authorization Printed
// 	PLRRPR			Paperless Renewal Reminder Letter
// 	PMTCOM      	Installment Payment Review Completed
// 	PNACC       	Accepted for Filing PN Generated
// 	PNACT       	Action PN Generated
// 	PNLIMR      	Market Based Auto Term PN Generated
// 	PNLISR      	Site Based Auto Term PN Generated
// 	PRCOM       	Puerto Rico Review Completed
// 	PRCOM2      	Puerto Rico Review Completed
// 	PROFF 			Offlined for Notification from Puerto Rico
// 	PTDCOM      	Petition to Deny Review Completed
// 	PUBCOM      	Part 90 Public Safety Services Review Completed
// 	PWVCOM      	Grand/waiver/func integ question review completed
// 	PWVOFF			"Checked Y to grand/waiver/func integ question 8c"""
// 	QZCOM       	NRAO/NRRO Review Completed
// 	QZCOM2      	NRAO/NRRO Review Completed
// 	QZNCOM      	Quiet Zone Notification Completed
// 	QZNOFF			Waiting for Quiet Zone Notification
// 	QZOFF 			Offlined for Notification from NRAO/NRRO
// 	QZOFF2			Waiting for Notification from NRAO/NRRO
// 	QZRVC       	Review of Quiet Zone Attribute Change Completed
// 	QZRVW 			Offlined for Review of Quiet Zone Attribute Change
// 	RCDUP 			Reference Copy Duplicate Requested
// 	RCSCOM      	Receiver Call Sign Review Completed
// 	RCSOFF			Offlined for Receiver Call Sign Review
// 	RDLCOM      	Redlight Review Completed
// 	RDLOFF			Offlined for Red Light
// 	RECAM 			Amendment Received
// 	RECAU 			Administrative Update Received
// 	RECCA 			Cancellation of License Received
// 	RECCHK			Offlined For Coser Recon Check
// 	RECCOM      	Coser Recon Check Completed
// 	RECDC 			Data Correction Application Applied
// 	RECDU 			Duplicate License Request Received
// 	RECEX 			Request for Extension of Time Received
// 	RECMD 			Modification Received
// 	RECMJ 			Major Modification Received
// 	RECNE 			New Application Received
// 	RECNT 			Required Notification Received
// 	RECRL 			Registration of Location/Link Received
// 	RECRM 			Renewal/Modification Received
// 	RECRO 			Renewal Only Received
// 	RECWD 			Withdrawal of Application Received
// 	REGCOM      	Regulatory Status Review Complete
// 	RLMCOM      	Rulemaking Decision Completed
// 	RLMOFF			Offlined for Rulemaking
// 	RLOFF 			Offlined for Link/Location Review
// 	RLPR  			Registration Letter Printed
// 	RMCOM       	Renewal/Mod Review Completed
// 	RMOFF 			Offlined for Renewal/Mod Review
// 	ROBCOM      	Renewal: Buildout Deadline Review Completed.
// 	ROMCOM      	Untimely filed RO/RM Review Completed
// 	ROMOFF			Offlined for Untimely filed RO/RM Review
// 	RONCOM      	RO: Pending NT Review completed.
// 	ROXCOM      	RO: Pending EX Review completed.
// 	RSCCOM      	Radio Service Review Completed
// 	RSCOFF			Offlined for Radio Service Review
// 	RSPCOM      	Reduced Service Providers Review Completed
// 	RTPEND			Returned to Pending
// 	SABCOM      	Radial data height-power limit review completed
// 	SBCCOM      	ML Bidding Credit Eligibility Review Completed
// 	SBGCOM      	Radial data height-power limit review completed
// 	SCBCOM      	ML Closed Bidding Eligibility Review Completed
// 	SCCOM       	Southern California Review Completed
// 	SCHCOM      	Four-Letter Call Sign Review Complete
// 	SCRCOM      	Manual Domestic MMSI Number Assignment Complete
// 	SCSCOM      	Subscriber Call Sign Review Completed
// 	SDACOM      	Additional SID review completed
// 	SDNCOM      	New SID review completed
// 	SECCOM      	Offline for compliance with Section 20.9 completed
// 	SGACOM      	Review of Slow Growth - Application - Completed
// 	SGLCOM      	Review of Slow Growth - License - Completed
// 	SHXCOM      	Ship Exemption Review Complete
// 	SIPCOM      	ML Installment Payment Elig Review Completed
// 	SLPEXC      	Streamline Processing Exclusivity Review Completed
// 	SMBCOM      	SMR Buildout Review Completed
// 	SMICOM      	Review of SMR Special Conditions Completed
// 	STACCZ			STA - Coser clearance zone check was not created
// 	STACGE			STA - Coser clearance was not created
// 	STACOM      	STA Review Completed
// 	STAOFF			Offlined for STA Review
// 	STBCOM      	Bill for Emergency STA Completed
// 	STBOFF			Offlined to Bill for Emergency STA
// 	STCCOM      	STA - Coser clearance closed
// 	STLCOM      	Short Term Lease Expiration Date Review Completed
// 	STMCOM      	STA - Coser manual offline closed
// 	STZCOM      	STA - Coser clearance zone check closed
// 	TACOM       	TowAir Check Completed
// 	TBCCOM      	TL Bidding Credit Status Review Completed
// 	TCBCOM      	TL Closed Bidding Status Review Completed
// 	TCCOM       	Tower Tolerance Check Completed
// 	TCRCOM      	TLBC Coverage Requirement Completed
// 	TCREVC      	Terminated Component Review Completed
// 	TINCOM      	TIN/Name Verification Completed
// 	TINOFF			Offlined for TIN/Name Verification
// 	TIPCOM      	TL Installment Payment Status Review Completed
// 	TLBCOM      	Tribal Land Bidding Credit Review Completed
// 	TLSCOM      	short term De Facto Transfer Leases verification c
// 	TOWCOM      	Review of 47 C.F.R. 17.7(e)(1) completed
// 	TOWGCM      	Failed Grant-Time Tower Clearance Check Completed
// 	TOWOFF			Offlined for review of 47 C.F.R. 17.7(e)(1)
// 	TOWRE       	Received from Tower
// 	TOWSE 			Sent to Tower
// 	TPECOM      	Term Pending Exclusion Review Completed
// 	TPRCOM      	Term Pending Review Completed
// 	UCFCOM      	Untimely Constructed/Filed Review Completed
// 	UCFOFF			Offlined for Untimely Constructed/Filed Review
// 	UCOCOM      	Untimely Construction Review Completed
// 	UCOOFF			Offlined for Untimely Construction Review
// 	UMRCOM      	Utility Mobile Review Completed
// 	UNCOM       	UNICOM Review Completed
// 	UNECOM      	Unicom Eligibility Review Completed
// 	VANCOM      	Schedule D Eligibility A or C review completed
// 	VRCOM       	RO Vanity Review Completed
// 	WACOM       	Other Wireless Activity Review Completed
// 	WAVCOM      	Wavier Review Completed
// 	WAVOFF			Offlined for Waiver Review
// 	WDREV 			Withdrawn Application Reversed
// 	WNTCOM      	NT Waiver Removal Review Complete
// 	WRECOM      	Waiver Renewal- 30 days Review Completed
// 	WREOFF			Offlined for Renewal Waiver- 30 days Review
// 	WSCCOM      	Wavier Special Condition Review Completed
// 	XBRCOM      	Transborder Coordination Completed
// 	XBROFF			Offlined for Transborder Coordination

// IR	OFACS Status Code
// 	50 	Pending at OET
// 	52 	Record Returned by OET
// 	53 	Pending at OET
// 	54 	Pending at OET
// 	55 	Pending at OET
// 	56 	Pending at OET
// 	57 	Pending at FAA
// 	58 	Pending at NTIA
// 	59 	Pending at OET
// 	60 	Pending at NTIA
// 	61 	Pending at NTIA
// 	62 	Under Review at OET
// 	63 	Pending at NTIA
// 	64 	Under Review at OET
// 	65 	Under Review at OET
// 	66 	Action Completed by OET
// 	67 	Action Canceled by OET
// 	68 	Pending at OET
// 	146	Pending at OET
// 	149	Pending at OET

// LC	Classification of Lease
// 	ML	Spectrum Manager Lease
// 	TL	De Facto Transfer Lease

// LC	Lease Filing Type
// 	P	Partial
// 	F	Full
// 	M	Microwave Link
// 	C	PND

// LC	Lease Term
// 	L 	Long
// 	S	Short

// LO	Location Class Code
// 	C	Centerpoint
// 	L	IQ/QQ Area of Operation Desc
// 	P	Passive Repeater Location
// 	R	Receive Location
// 	T	Transmit Location

// LO	Location Type Code
// 	Type Code	Description	Fixed/Mobile
// 	6 	6.1 meter control station 	M
// 	F 	Fixed                     	F
// 	I 	Itinerant                 	M
// 	L 	Geographic Location       	M
// 	M 	Mobile                    	M
// 	P 	P35 Centerpoint           	F
// 	T 	Temporary Fixed           	M

// LO	Structure Type
// 	B      	Building
// 	BANT   	Building with Antenna on top
// 	BMAST  	Building with Mast
// 	BPIPE  	Building with Pipe
// 	BPOLE  	Building with Pole
// 	BRIDG  	Bridge
// 	BTWR   	Building with Tower
// 	GTOWER 	Guyed Structure Used for Communication Purposes
// 	LTOWER 	Lattice Tower
// 	MAST   	Mast
// 	MTOWER 	Monopole
// 	NNGTANN	Guyed Tower Array
// 	NNLTANN	Lattice Tower Array
// 	NNMTANN	Monopole Array
// 	NNTANN 	Antenna Tower Array - 1st N = # towers  2nd N = Po
// 	NTOWER 	Multiple Structures - The N is the # of structures
// 	PIPE   	Any type of Pipe
// 	POLE   	Any type of Pole
// 	RIG    	Oil or other type of Rig
// 	SIGN   	Any Type of Sign or Billboard
// 	SILO   	Any type of Silo
// 	STACK  	Smoke Stack
// 	TANK   	"Any type of Tank (Water, Gas, etc)                 "
// 	TOWER  	Free standing or Guyed Structure used for Communic
// 	TREE   	When used as a support for an antenna
// 	UPOLE  	Utility Pole/Tower used to provide service (Electr
// 	UTOWER 	Unguyed - Free Standing Tower

// LO	Type Operation Code
// 	18	18 GHz Low Power
// 	31	31 GHz Systems
// 	38	38 GHz Systems
// 	D 	Digital Electronic Message Service (DEMS)
// 	F 	Permanent Fixed Point to Point
// 	M 	Multiple Address System (MAS)
// 	T 	Temporary Fixed/Mobile

// LO	Area of Operation
// 	Code	Description	Service
// 	A	KMRA around a Fixed Location               	LP
// 	C	Countywide                                 	MW
// 	N	"Nationwide including HI, AK, US Territories"	MW
// 	O	Other - Narrative                          	MW
// 	P	KMRA around a centerpoint                  	MW
// 	R	Box/Rectangular                            	MW
// 	S	Statewide                                  	MW
// 	U	Continental US                             	MW
// 	X	LM Control Station Meeting 6.1 Meter Rule  	LP

// MH	Channel Plan Number
// 	Plan	New/Old Plan	Freq Lower	Freq Upper	Description
// 	BRS1   	N	2496.00000000	2502.00000000	New Channel Number BRS1
// 	BRS2   	N	2618.00000000	2624.00000000	New Channel Number BRS2
// 	NA1    	N	2502.00000000	2507.50000000	New Channel Number A1
// 	NA2    	N	2507.50000000	2513.00000000	New Channel Number A2
// 	NA3    	N	2513.00000000	2518.50000000	New Channel Number A3
// 	NA4    	N	2572.00000000	2578.00000000	New Channel Number A4
// 	NB1    	N	2518.50000000	2524.00000000	New Channel Number B1
// 	NB2    	N	2524.00000000	2529.50000000	New Channel Number B2
// 	NB3    	N	2529.50000000	2535.00000000	New Channel Number B3
// 	NB4    	N	2578.00000000	2584.00000000	New Channel Number B4
// 	NC1    	N	2535.00000000	2540.50000000	New Channel Number C1
// 	NC2    	N	2540.50000000	2546.00000000	New Channel Number C2
// 	NC3    	N	2546.00000000	2551.50000000	New Channel Number C3
// 	NC4    	N	2584.00000000	2590.00000000	New Channel Number C4
// 	ND1    	N	2551.50000000	2557.00000000	New Channel Number D1
// 	ND2    	N	2557.00000000	2562.50000000	New Channel Number D2
// 	ND3    	N	2562.50000000	2568.00000000	New Channel Number D3
// 	ND4    	N	2590.00000000	2596.00000000	New Channel Number D4
// 	NE1    	N	2624.00000000	2629.50000000	New Channel Number E1
// 	NE2    	N	2629.50000000	2635.00000000	New Channel Number E2
// 	NE3    	N	2635.00000000	2640.50000000	New Channel Number E3
// 	NE4    	N	2608.00000000	2614.00000000	New Channel Number E4
// 	NF1    	N	2640.50000000	2646.00000000	New Channel Number F1
// 	NF2    	N	2646.00000000	2651.50000000	New Channel Number F2
// 	NF3    	N	2651.50000000	2657.00000000	New Channel Number F3
// 	NF4    	N	2602.00000000	2608.00000000	New Channel Number F4
// 	NG1    	N	2673.50000000	2679.00000000	New Channel Number G1
// 	NG2    	N	2679.00000000	2684.50000000	New Channel Number G2
// 	NG3    	N	2684.50000000	2690.00000000	New Channel Number G3
// 	NG4    	N	2596.00000000	2602.00000000	New Channel Number G4
// 	NH1    	N	2657.00000000	2662.50000000	New Channel Number H1
// 	NH2    	N	2662.50000000	2668.00000000	New Channel Number H2
// 	NH3    	N	2668.00000000	2673.50000000	New Channel Number H3
// 	O1     	O	2150.00000000	2156.00000000	Old Channel Number 1
// 	O2     	O	2156.00000000	2162.00000000	Old Channel Number 2
// 	O2A    	O	2156.00000000	2160.00000000	Old Channel Number 2A
// 	OA1    	O	2500.00000000	2506.00000000	Old Channel Number A1
// 	OA2    	O	2512.00000000	2518.00000000	Old Channel Number A2
// 	OA3    	O	2524.00000000	2530.00000000	Old Channel Number A3
// 	OA4    	O	2536.00000000	2542.00000000	Old Channel Number A4
// 	OB1    	O	2506.00000000	2512.00000000	Old Channel Number B1
// 	OB2    	O	2518.00000000	2524.00000000	Old Channel Number B2
// 	OB3    	O	2530.00000000	2536.00000000	Old Channel Number B3
// 	OB4    	O	2542.00000000	2548.00000000	Old Channel Number B4
// 	OC1    	O	2548.00000000	2554.00000000	Old Channel Number C1
// 	OC2    	O	2560.00000000	2566.00000000	Old Channel Number C2
// 	OC3    	O	2572.00000000	2578.00000000	Old Channel Number C3
// 	OC4    	O	2584.00000000	2590.00000000	Old Channel Number C4
// 	OD1    	O	2554.00000000	2560.00000000	Old Channel Number D1
// 	OD2    	O	2566.00000000	2572.00000000	Old Channel Number D2
// 	OD3    	O	2578.00000000	2584.00000000	Old Channel Number D3
// 	OD4    	O	2590.00000000	2596.00000000	Old Channel Number D4
// 	OE1    	O	2596.00000000	2602.00000000	Old Channel Number E1
// 	OE2    	O	2608.00000000	2614.00000000	Old Channel Number E2
// 	OE3    	O	2620.00000000	2626.00000000	Old Channel Number E3
// 	OE4    	O	2632.00000000	2638.00000000	Old Channel Number E4
// 	OF1    	O	2602.00000000	2608.00000000	Old Channel Number F1
// 	OF2    	O	2614.00000000	2620.00000000	Old Channel Number F2
// 	OF3    	O	2626.00000000	2632.00000000	Old Channel Number F3
// 	OF4    	O	2638.00000000	2644.00000000	Old Channel Number F4
// 	OG1    	O	2644.00000000	2650.00000000	Old Channel Number G1
// 	OG2    	O	2656.00000000	2662.00000000	Old Channel Number G2
// 	OG3    	O	2668.00000000	2674.00000000	Old Channel Number G3
// 	OG4    	O	2680.00000000	2686.00000000	Old Channel Number G4
// 	OH1    	O	2650.00000000	2656.00000000	Old Channel Number H1
// 	OH2    	O	2662.00000000	2668.00000000	Old Channel Number H2
// 	OH3    	O	2674.00000000	2680.00000000	Old Channel Number H3
// 	OI01   	O	2686.00000000	2686.12500000	Old Channel Number I01
// 	OI02   	O	2686.12500000	2686.25000000	Old Channel Number I02
// 	OI03   	O	2686.25000000	2686.37500000	Old Channel Number I03
// 	OI04   	O	2686.37500000	2686.50000000	Old Channel Number I04
// 	OI05   	O	2686.50000000	2686.62500000	Old Channel Number I05
// 	OI06   	O	2686.62500000	2686.75000000	Old Channel Number I06
// 	OI07   	O	2686.75000000	2686.87500000	Old Channel Number I07
// 	OI08   	O	2686.87500000	2687.00000000	Old Channel Number I08
// 	OI09   	O	2687.00000000	2687.12500000	Old Channel Number I09
// 	OI10   	O	2687.12500000	2687.25000000	Old Channel Number I10
// 	OI11   	O	2687.25000000	2687.37500000	Old Channel Number I11
// 	OI12   	O	2687.37500000	2687.50000000	Old Channel Number I12
// 	OI13   	O	2687.50000000	2687.62500000	Old Channel Number I13
// 	OI14   	O	2687.62500000	2687.75000000	Old Channel Number I14
// 	OI15   	O	2687.75000000	2687.87500000	Old Channel Number I15
// 	OI16   	O	2687.87500000	2688.00000000	Old Channel Number I16
// 	OI17   	O	2688.00000000	2688.12500000	Old Channel Number I17
// 	OI18   	O	2688.12500000	2688.25000000	Old Channel Number I18
// 	OI19   	O	2688.25000000	2688.37500000	Old Channel Number I19
// 	OI20   	O	2688.37500000	2688.50000000	Old Channel Number I20
// 	OI21   	O	2688.50000000	2688.62500000	Old Channel Number I21
// 	OI22   	O	2688.62500000	2688.75000000	Old Channel Number I22
// 	OI23   	O	2688.75000000	2688.87500000	Old Channel Number I23
// 	OI24   	O	2688.87500000	2689.00000000	Old Channel Number I24
// 	OI25   	O	2689.00000000	2689.12500000	Old Channel Number I25
// 	OI26   	O	2689.12500000	2689.25000000	Old Channel Number I26
// 	OI27   	O	2689.25000000	2689.37500000	Old Channel Number I27
// 	OI28   	O	2689.37500000	2689.50000000	Old Channel Number I28
// 	OI29   	O	2689.50000000	2689.62500000	Old Channel Number I29
// 	OI30   	O	2689.62500000	2689.75000000	Old Channel Number I30
// 	OI31   	O	2689.75000000	2689.87500000	Old Channel Number I31

// MI	Facility Type Code
// 	Code	Facility Type	Description
// 	BTA    	N	BTA AUTHORIZATION
// 	CIF    	N	Commercial ITFS Station
// 	CIFB   	B	Commercial ITFS High Power Signal Booster Station
// 	CIFH   	H	Commercial ITFS Response Station Hub
// 	CIFX   	B	Commercial ITFS Low Power Signal Booster Station
// 	CIFY   	I	Commercial ITFS Downstream I Channel Station
// 	ICB    	B	"ITFS,ITFSX HIGH-POWER BOOSTER STATION                    "
// 	ICH    	H	"ITFS,ITFSX RESPONSE STATION HUB                          "
// 	ICX    	B	"ITFS,ITFSX LOW-POWER BOOSTER STATION                     "
// 	ICY    	I	"ITFS,ITFSX DOWNSTREAM I CHANNEL STATION                  "
// 	IF     	N	INSTRUCTIONAL TELEVISION FIXED SERVICE
// 	IFB    	B	ITFS HIGH-POWER BOOSTER STATION
// 	IFC    	N	INSTRUCTIONAL TELEVISION FIXED SERVICE RESPONSE RECEIVER
// 	IFD    	N	INSTRUCTIONAL TELEVISION FIXED SERVICE DEVELOPMENTAL
// 	IFH    	H	ITFS RESPONSE STATION HUB
// 	IFS    	N	INSTRUCTIONAL TELEVISION FIXED SERVICE STUDIO LINK
// 	IFX    	B	ITFS LOW-POWER BOOSTER STATION
// 	IFY    	I	ITFS DOWNSTREAM I CHANNEL STATION
// 	MD     	N	MULTIPOINT DISTRIBUTION SERVICE STATION
// 	MDB    	B	MULTIPOINT DISTRIBUTION SERVICE SIGNAL BOOSTER STATION
// 	MDCB   	B	"MDS,ITFSX HIGH-POWER BOOSTER STATION                     "
// 	MDCH   	H	"MDS,ITFSX RESPONSE STATION HUB                           "
// 	MDCX   	B	"MDS,ITFSX LOW-POWER BOOSTER STATION                      "
// 	MDCY   	I	"MDS,ITFSX DOWNSTREAM I CHANNEL STATION                   "
// 	MDH    	H	MDS RESPONSE STATION HUB
// 	MDIB   	B	"MDS,ITFS HIGH-POWER BOOSTER STATION                      "
// 	MDIH   	H	"MDS,ITFS RESPONSE STATION HUB                            "
// 	MDIX   	B	"MDS,ITFS LOW-POWER BOOSTER STATION                       "
// 	MDIY   	I	"MDS,ITFS DOWNSTREAM I CHANNEL STATION                    "
// 	MDV    	N	MULTIPOINT DISTRIBUTION SERVICE DEVELOPMENTAL STATION
// 	MDX    	B	MDS LOW-POWER BOOSTER STATION
// 	MDY    	I	MDS DOWNSTREAM I CHANNEL STATION
// 	MICB   	B	"MDS,ITFS,ITFSX HIGH-POWER BOOSTER STATION                "
// 	MICH   	H	"MDS,ITFS,ITFSX RESPONSE STATION HUB                      "
// 	MICX   	B	"MDS,ITFS,ITFSX LOW-POWER BOOSTER STATION                 "
// 	MICY   	I	"MDS,ITFS,ITFSX DOWNSTREAM I CHANNEL STATION              "

// MI	License Type Code
// 	A	Partitioned Service Area
// 	B	BTA
// 	P	Protected Service Area (56.33 km)

// MW	Class Station Code
// 	AOX                 	Operational Fixed
// 	APC                 	ALASKA PUBLIC
// 	APX                 	ALASKA PRIVATE
// 	APX2                	ALASKA PRIVATE (TEMPORARY)
// 	AVW                 	Audio Visual Warning System (1300-1350 MHz)
// 	AX                  	AERONAUTICAL FIXED
// 	AX1                 	AERONAUTICAL FIXED (MOBILE)
// 	AX2                 	AERONAUTICAL FIXED (TEMPORARY)
// 	DGP                 	DIFFERENTIAL GPS
// 	DLT                 	Aircraft Data Link Land Test
// 	DLT1                	Aircraft Data Link Land Test (Mobile)
// 	ELT                 	ELT TEST
// 	ELT1                	ELT TEST (MOBILE)
// 	FA                  	AERONAUTICAL ENROUTE
// 	FA1                 	AERONAUTICAL ENROUTE (MOBILE)
// 	FA2                 	AERONAUTICAL ENROUTE (TEMPORARY)
// 	FAA                 	AERONAUTICAL ADVISORY (UNICOM)
// 	FAA1                	AERONAUTICAL ADVISORY (UNICOM) (MOBILE)
// 	FAA2                	AERONAUTICAL ADVISORY (UNICOM) (TEMPORARY)
// 	FAB                 	AUTOMATIC WEATHER OBSERVATION
// 	FAC                 	AIRPORT CONTROL TOWER
// 	FAS                 	AVIATION SUPPORT INSTRUCTIONAL
// 	FAS1                	AVIATION SUPPORT INSTRUCTIONAL (MOBILE)
// 	FAT                 	FLIGHT TEST
// 	FAT1                	FLIGHT TEST (MOBILE)
// 	FAT3                	FLIGHT TEST (ITINERANT)
// 	FB                  	Base
// 	FB2                 	Mobile Relay
// 	FB2A                	Mobile Relay - Airport Terminal Use
// 	FB2C                	Mobile Relay - Interconnect
// 	FB2I                	Mobile Relay - Itinerant
// 	FB2J                	Mobil Relay - Temporary Interconnect
// 	FB2K                	Mobile Relay - Stand-by Interconnect
// 	FB2L                	Mobile Relay - Itinerant Interconnect
// 	FB2S                	Mobil Relay - Stand-by
// 	FB2T                	Mobile Relay - Temporary
// 	FB4                 	Community Repeater
// 	FB4A                	Community Repeater - Airport Terminal Use
// 	FB4C                	Community Repeater - Interconnect
// 	FB4I                	Community Repeater - Itinerant
// 	FB4J                	Community Repeater - Temporary Interconnect
// 	FB4K                	Community Repeater - Stand-by Interconnect
// 	FB4L                	Community Repeater - Itinerant Interconnect
// 	FB4S                	Community Repeater - Stand-by
// 	FB4T                	Community Repeater - Temporary
// 	FB6                 	Private Carrier (profit)
// 	FB6A                	Private Carrier (profit) - Airport Terminal Use
// 	FB6C                	Private Carrier (profit) - Interconnect
// 	FB6I                	Private Carrier (profit) - Itinerant
// 	FB6J                	Private Carrier (profit) - Temporary Interconn
// 	FB6K                	Private Carrier (profit) - Stand-by Interconne
// 	FB6L                	Private Carrier (profit) - Itinerant Interconn
// 	FB6S                	Private Carrier (profit) - Stand-by
// 	FB6T                	Private Carrier (profit) - Temporary
// 	FB7                 	Private Carrier (non-profit)
// 	FB7A                	Private Carrier (non-profit) - Airport Terminal Us
// 	FB7C                	Private Carrier (non-profit) - Interconnect
// 	FB7I                	Private Carrier (non-profit) - Itinerant
// 	FB7J                	Private Carrier (non-profit) - Temporary Inter
// 	FB7K                	Private Carrier (non-profit) - Stand-by Interc
// 	FB7L                	Private Carrier (non-profit) - Itinerant Inter
// 	FB7S                	Private Carrier (non-profit) - Stand-by
// 	FB7T                	Private Carrier (non-profit) - Temporary
// 	FB8                 	Centralized Trunk Relay
// 	FB8A                	Centralized Trunk Relay - Airport Terminal Use
// 	FB8C                	Centralized Trunk Relay - Interconnect
// 	FB8I                	Centralized Trunk Relay - Itinerant
// 	FB8J                	Centralized Trunk Relay - Temporary Interconnect
// 	FB8K                	Centralized Trunk Relay - Standby Interconnect
// 	FB8L                	Centralized Trunk Relay - Itinerant Interconnect
// 	FB8S                	Centralized Trunk Relay - Standby
// 	FB8T                	Centralized Trunk Relay - Temporary
// 	FBA                 	Airport Terminal Use
// 	FBAT                	Small Base - Temporary
// 	FBBS                	Base
// 	FBC                 	Base - Interconnect
// 	FBCT                	FBCT
// 	FBGS                	Ground
// 	FBI                 	Base - Itinerant
// 	FBJ                 	Base - Temporary Interconnect
// 	FBK                 	Base - Stand-by Interconnect
// 	FBL                 	Base - Itinerant Interconnect
// 	FBS                 	Base - Stand-by
// 	FBSI                	Air-ground Signaling
// 	FBST                	Standby
// 	FBT                 	Base - Temporary
// 	FC                  	PUBLIC COAST
// 	FC2                 	PUBLIC COAST(TEMPORARY)
// 	FCA                 	MARINE SUPPORT-TESTING & TRAINING
// 	FCA2                	MARINE SUPPORT-TESTING & TRAINING (TEMPORARY)
// 	FCL                 	PRIVATE COAST
// 	FCL2                	PRIVATE COAST (TEMPORARY)
// 	FCU                 	MARINE UTILITY
// 	FCU1                	MARINE UTILITY (MOBILE)
// 	FIS                 	Flight Information Services
// 	FIS1                	Flight Information Services with Hand Held/Mobile
// 	FIS2                	Flight Information Services for Temporary Operatio
// 	FLT                 	Auxilary Test
// 	FLTC                	Auxiliary Test - Interconnect
// 	FLTI                	Auxilary Test - Itinerant
// 	FLTL                	Auxilary Test - Itinerant Interconnect
// 	FLU                 	AVIATION SUPPORT SERVICE
// 	FLU1                	AVIATION SUPPORT SERVICE (MOBILE)
// 	FMA1                	AIRCRAFT FLIGHT TEST STATION
// 	FX                  	Fixed
// 	FX0S                	Operational Fixed - Stand-by
// 	FX1                 	Control
// 	FX1A                	Control - Airport Terminal Use
// 	FX1C                	Control - Interconnect
// 	FX1I                	Control - Itinerant
// 	FX1J                	Control - Temporary Interconnect
// 	FX1K                	Control - Stand-by Interconnect
// 	FX1L                	Control - Itinerant Interconnect
// 	FX1S                	Control - Stand-by
// 	FX1T                	Control - Temporary
// 	FX2                 	Fixed Relay
// 	FX2A                	Fixed Relay - Airport Terminal Use
// 	FX2C                	Fixed Relay - Interconnect
// 	FX2I                	Fixed Relay - Itinerant
// 	FX2J                	Fixed Relay - Temporary Interconnect
// 	FX2K                	Fixed Relay - Stand-by Interconnect
// 	FX2L                	Fixed Relay - Itinerant Interconnect
// 	FX2S                	Fixed Relay - Stand-by
// 	FX2T                	Fixed Relay - Temporary
// 	FX3                 	Secondary Fixed (Tone Signalling)
// 	FX3A                	Secondary Fixed - Airport Terminal Use
// 	FX3C                	Secondary Fixed - Interconnect
// 	FX3I                	Secondary Fixed - Itinerant
// 	FX3J                	Secondary Fixed - Temporary Interconnect
// 	FX3K                	Secondary Fixed - Stand-by Interconnect
// 	FX3L                	Secondary Fixed - Itinerant Interconnect
// 	FX3S                	Secondary Fixed - Stand-by
// 	FX3T                	Secondary Fixed  - Temporary
// 	FX5                 	Temporary Fixed
// 	FXA                 	Fixed - Airport Terminal Use
// 	FXB                 	Primary Permanent Fixed Stations or Links
// 	FXC                 	Fixed - Interconnect
// 	FXCO                	Central Office
// 	FXCT                	Control
// 	FXDI                	Dispatch
// 	FXI                 	Fixed - Itinerant
// 	FXIO                	Inter-office
// 	FXJ                 	Fixed - Temporary Interconnect
// 	FXK                 	Fixed - Stand-by Interconnect
// 	FXL                 	Fixed - Itinerant Interconnect
// 	FXO                 	Operational Fixed
// 	FXOA                	Fixed - Interconnect - Airport Terminal Use
// 	FXOC                	Operational Fixed - Interconnect
// 	FXOI                	Operational Fixed - Itinerant
// 	FXOJ                	Operational Fixed - Temporary Interconnect
// 	FXOK                	Operational Fixed - Stand-by Interconnect
// 	FXOL                	Operational Fixed - Itinerant Interconnect
// 	FXOS                	Operational Fixed - Stand-by
// 	FXOT                	Operational Fixed - Temporary
// 	FXRP                	Repeater
// 	FXRX                	Fixed Relay
// 	FXS                 	Fixed - Stand-by
// 	FXSB                	Fixed Subscriber
// 	FXT                 	Fixed - Temporary
// 	FXTS                	Auxiliary Test
// 	FXV                 	CTS Exceeding 20'
// 	FXW                 	CTS Meeting 20'
// 	FXY                 	Interzone
// 	FXYC                	Interzone - Interconnect
// 	FXZ                 	Zone
// 	FXZC                	Zone - Interconnect
// 	GCO                 	GROUND COMMUNICATIONS OUTLET
// 	LR                  	Radiolocation Land
// 	LRC                 	Radiolocation Land - Interconnect
// 	LRJ                 	Radiolocation Land - Temporary Interconnect
// 	LRK                 	Radiolocation Land - Stand-by Interconnect
// 	LRS                 	Radiolocation Land - Stand-by
// 	LRT                 	Radiolocation Land - Temporary
// 	M08C                	Centralized Trunk Mobile - Interconnect
// 	MFL                 	AERONAUTICAL MULTICOM
// 	MFL1                	AERONAUTICAL MULTICOM (MOBILE)
// 	MFL2                	AERONAUTICAL MULTICOM (TEMPORARY)
// 	MFX                 	MARINE OPS FIXED
// 	MFX2                	MARINE OPS FIXED (TEMPORARY)
// 	MLSB                	Mobile Subscriber
// 	MO                  	Mobile
// 	MO3                 	Mobile/Vehicular Repeater
// 	MO3A                	Mobile/Vehicular Repeater - Airport Terminal Use
// 	MO3C                	Mobile/Vehicular Repeater with Interconnect
// 	MO3I                	Mobile/Vehicular Repeater - Itinerant
// 	MO3J                	Mobile/Vehicular Repeater with Temporary Interconn
// 	MO3K                	Mobile/Vehicular Repeater with Stand-by Interconne
// 	MO3L                	Mobile/Vehicular Repeater with Itinerant Interconn
// 	MO3S                	Stand-by Mobile/Vehicular Repeater
// 	MO3T                	Temporary Mobile/Vehicular Repeater
// 	MO5                 	Mobile & Temporary Fixed
// 	MO6                 	Private Carrier Mobile Op (profit)
// 	MO6A                	Private Carrier Mobile Op (profit) - Airport Termi
// 	MO6C                	Private Carrier Mobile Op (profit) - Interconn
// 	MO6I                	Private Carrier Mobile Op (profit) - Itinerant
// 	MO6J                	Private Carrier Mobile Operation (profit) with Tem
// 	MO6K                	Private Carrier Mobile Op (profit) - Stand-by
// 	MO6L                	Private Carrier Mobile Operation (profit) with Iti
// 	MO6S                	Private Carrier Mobile Op (profit) - Stand-by
// 	MO6T                	Temporary Private Carrier Mobile Operation (profit
// 	MO7                 	Private Carrier Mobile Op (non-profit)
// 	MO7A                	Private Carrier Mobile Op (non-profit) - Airport T
// 	MO7C                	Private Carrier Mobile Op (non-profit) - Inter
// 	MO7I                	Private Carrier Mobile Op (non-profit) - Itine
// 	MO7J                	Private Carrier Mobile Operation (non-profit) with
// 	MO7K                	Private Carrier Mobile Op (non-profit) - Stand
// 	MO7L                	Private Carrier Mobile Operation (non-profit) with
// 	MO7S                	Private Carrier Mobile Op (non-profit) - Stand
// 	MO7T                	Temporary Private Carrier Mobile Operation (non-pr
// 	MO8                 	Centralized Trunk Mobile
// 	MO8A                	Centralized Trunk Mobile - Airport Terminal Use
// 	MO8C                	Centralized Trunk Mobile - Interconnect
// 	MOA                 	Mobile - Airport Terminal Use
// 	MOC                 	Mobile - Interconnect
// 	MOI                 	Mobile - Itinerant
// 	MOJ                 	Mobile with Temporary Interconnect
// 	MOK                 	Mobile with Stand-by Interconnect
// 	MOL                 	Mobile with Itinerant Interconnect
// 	MOS                 	Mobile - Stand-by
// 	MOT                 	Temporary Mobile
// 	MOU1                	AERONAUTICAL UTILITY MOBILE
// 	MR                  	Radiolocation Mobile
// 	MRT                 	MARINE RECEIVER TEST
// 	MRT2                	MARINE RECEIVER TEST (TEMPORARY)
// 	MSC                 	SHORE RADAR TEST
// 	MSC2                	SHORE RADAR TEST (TEMPORARY)
// 	MSR                 	SHORE RADIONAVIGATION
// 	MSR2                	SHORE RADIONAVIGATION (TEMPORARY)
// 	RCO                 	REMOTE COMMUNICATIONS OUTLET
// 	RLA                 	AERONAUTICAL MARKER BEACON
// 	RLA1                	AERONAUTICAL MARKER BEACON (MOBILE)
// 	RLA2                	AERONAUTICAL MARKER BEACON (TEMPORARY)
// 	RLB                 	AERONAUTICAL RADIO BEACON
// 	RLB1                	AERONAUTICAL RADIO BEACON (MOBILE)
// 	RLB2                	AERONAUTICAL RADIO BEACON (TEMPORARY)
// 	RLC                 	SHORE RADIOLOCATION TEST
// 	RLC2                	SHORE RADIOLOCATION TEST (TEMPORARY)
// 	RLD                 	RADAR/RADAR TEST
// 	RLD1                	RADAR/RADAR TEST (MOBILE)
// 	RLD2                	RADAR/RADAR TEST (TEMPORARY)
// 	RLG                 	GLIDE PATH (SLOPE)
// 	RLG1                	GLIDE PATH (SLOPE) (MOBILE)
// 	RLG2                	GLIDE PATH (SLOPE) (TEMPORARY)
// 	RLL                 	LOCALIZER
// 	RLL1                	LOCALIZER (MOBILE)
// 	RLL2                	LOCALIZER (TEMPORARY)
// 	RLO                 	OMNIDIRECTIONAL RADIO RANGE
// 	RLO1                	OMNIDIRECTIONAL RADIO RANGE (MOBILE)
// 	RLO2                	OMNIDIRECTIONAL RADIO RANGE (TEMPORARY)
// 	RLR                 	SHORE RADIOLOCATION/RACON
// 	RLR2                	SHORE RADIOLOCATION/RACON (TEMPORARY)
// 	RLT                 	RADIONAVIGATION LAND TEST
// 	RLT1                	RADIONAVIGATION LAND TEST (MOBILE)
// 	RLT2                	Radionavigation Land Test - Temporary
// 	RNV                 	RADIONAVIGATION LAND
// 	RNV1                	RADIONAVIGATION LAND (MOBILE)
// 	RNV2                	RADIONAVIGATION LAND (TEMPORARY)
// 	RPC                 	RAMP CONTROL
// 	SAR                 	SEARCH AND RESCUE
// 	SAR1                	SEARCH AND RESCUE (MOBILE)
// 	UAT                 	Universal Access Transceiver service - Fixed
// 	UAT1                	Universal Access Transceiver service - Mobile
// 	UAT2                	Universal Access Transceiver service - Temp Fixed
// 	WDX                 	Radiolocation Weather Radar
// 	WDXS                	Radiolocation Weather Radar - Standby
// 	WDXT                	Radiolocation Weather Radar - Temporary

// PA	Country Code
// 	AD	Andorra
// 	AE	United Arab Emirates
// 	AF	Afghanistan
// 	AG	Antigua and Barbuda
// 	AI	Anguilla
// 	AL	Albania
// 	AM	Armenia
// 	AN	Netherlands Antilles
// 	AO	Angola
// 	AQ	Antarctica
// 	AR	Argentina
// 	AS	American Samoa
// 	AT	Austria
// 	AU	Australia
// 	AW	Aruba
// 	AZ	Azerbaijan
// 	BA	Bosnia-Herzegovina
// 	BB	Barbados
// 	BD	Bangladesh
// 	BE	Belgium
// 	BF	Burkina Faso
// 	BG	Bulgaria
// 	BH	Bahrain
// 	BI	Burundi
// 	BJ	Benin
// 	BM	Bermuda
// 	BN	Brunei Darussalam
// 	BO	Bolivia
// 	BR	Brazil
// 	BS	Bahamas
// 	BT	Bhutan
// 	BV	Bouvet Island
// 	BW	Botswana
// 	BY	Belarus
// 	BZ	Belize
// 	CA	Canada
// 	CC	Cocos (Keeling) Islands
// 	CF	Central African Republic
// 	CG	Congo
// 	CH	Switzerland
// 	CI	Cote D'Ivoire (Ivory Coast)
// 	CJ	Congo (Democratic Republic)
// 	CK	Cook Islands
// 	CL	Chile
// 	CM	Cameroon
// 	CN	China
// 	CO	Colombia
// 	CR	Costa Rica
// 	CS	Czech Republic
// 	CU	CSK
// 	CV	Cape Verde
// 	CX	Christmas Island
// 	CY	Cuba
// 	CZ	Cyprus
// 	DE	Germany
// 	DJ	Djibouti
// 	DK	Denmark
// 	DM	Dominica
// 	DO	Dominican Republic
// 	DZ	Algeria
// 	EC	Ecuador
// 	EE	Estonia
// 	EG	Egypt
// 	EH	Western Sahara
// 	ER	Eritrea
// 	ES	Spain
// 	ET	Ethiopia
// 	FI	Finland
// 	FJ	Fiji
// 	FK	Falkland Islands (Malvinas)
// 	FM	Micronesia
// 	FO	Faroe Islands
// 	FR	France
// 	FX	"France, Metropolitan           "
// 	GA	Gabon
// 	GD	Grenada
// 	GE	Georgia
// 	GF	French Guiana
// 	GH	Ghana
// 	GI	Gibraltar
// 	GL	Greenland
// 	GM	Gambia
// 	GN	Guinea
// 	GP	Guadeloupe
// 	GQ	Equatorial Guinea
// 	GR	Greece
// 	GS	S. Georgia/S. Sandwich Isls
// 	GT	Guatemala
// 	GU	Guam
// 	GW	Guinea-Bissau
// 	GY	Guyana
// 	HK	Hong Kong
// 	HM	Heard and McDonald Islands
// 	HN	Honduras
// 	HR	Croatia (Hrvatska)
// 	HT	Haiti
// 	HU	Hungary
// 	ID	Indonesia
// 	IE	Ireland
// 	IL	Israel
// 	IN	India
// 	IO	British Indian Ocean Territory
// 	IQ	Iraq
// 	IR	Iran
// 	IS	Iceland
// 	IT	Italy
// 	JM	Jamaica
// 	JO	Jordan
// 	JP	Japan
// 	KE	Kenya
// 	KG	Kyrgyzstan
// 	KH	Cambodia
// 	KI	Kiribati
// 	KM	Comoros
// 	KN	Saint Kitts and Nevis
// 	KP	Korea (North)
// 	KR	Korea (South)
// 	KW	Kuwait
// 	KY	Cayman Islands
// 	KZ	Kazakhstan
// 	LA	Laos
// 	LB	Lebanon
// 	LC	Saint Lucia
// 	LI	Liechtenstein
// 	LK	Sri Lanka
// 	LR	Liberia
// 	LS	Lesotho
// 	LT	Lithuania
// 	LU	Luxembourg
// 	LV	Latvia
// 	LY	Libya
// 	MA	Morocco
// 	MC	Monaco
// 	MD	Moldova
// 	MG	Madagascar
// 	MH	Marshall Islands
// 	MK	Macedonia
// 	ML	Mali
// 	MM	Myanmar
// 	MN	Mongolia
// 	MO	Macau
// 	MP	Norway
// 	MQ	Martinique
// 	MR	Mauritania
// 	MS	Montserrat
// 	MT	Malta
// 	MU	Mauritius
// 	MV	Maldives
// 	MW	Malawi
// 	MX	Mexico
// 	MY	Malaysia
// 	MZ	Mozambique
// 	NA	Namibia
// 	NC	New Zealand
// 	NE	Nigeria
// 	NF	Northern Mariana Islands
// 	NG	Niue
// 	NI	Niger
// 	NL	Netherlands
// 	NP	Nepal
// 	NR	Nauru
// 	NT	New Caledonia
// 	NU	Norfolk Island
// 	NZ	Nicaragua
// 	OM	Oman
// 	PA	Panama
// 	PE	Peru
// 	PF	French Polynesia
// 	PG	Papua New Guinea
// 	PH	Philippines
// 	PK	Pakistan
// 	PL	Poland
// 	PM	St. Pierre and Miquelon
// 	PN	Pitcairn
// 	PR	Puerto Rico
// 	PT	Portugal
// 	PW	Palau
// 	PY	Paraguay
// 	QA	Qatar RE Reunion
// 	RN	Reunion
// 	RO	Romania
// 	RS	Serbia
// 	RU	Russian Federation
// 	RW	Rwanda
// 	SA	Saudi Arabia
// 	Sb	Solomon Islands
// 	SC	Seychelles
// 	SD	Sudan
// 	SE	Sweden
// 	SG	Singapore
// 	SH	St. Helena
// 	SI	Slovenia
// 	SJ	Svalbard and Jan Mayen Islands
// 	SK	Slovak Republic
// 	SL	Sierra Leone
// 	SM	San Marino
// 	SN	Senegal
// 	SO	Somalia
// 	SR	Suriname
// 	ST	Sao Tome and Principe
// 	SV	El Salvador
// 	SY	Syria
// 	SZ	Swaziland
// 	TC	Turks and Caicos Islands
// 	TD	Chad
// 	TF	French Southern Territories
// 	TG	Togo
// 	TH	Thailand
// 	TJ	Tajikistan
// 	TK	Tokelau
// 	TM	Turkmenistan
// 	TN	Tunisia
// 	TO	Tonga
// 	TP	East Timor
// 	TR	Turkey
// 	TT	Trinidad and Tobago
// 	TV	Tuvalu
// 	TW	Taiwan
// 	TZ	Tanzania
// 	UA	Ukraine
// 	UG	Uganda
// 	UK	United Kingdom
// 	UM	Uruguay
// 	US	United States
// 	UY	United States Minor Islands
// 	UZ	Uzbekistan
// 	VA	Vatican City State
// 	VC	Saint Vincent and Grenadines
// 	VE	Venezuela
// 	VG	Virgin Islands (British)
// 	VI	Virgin Islands (U.S.)
// 	VN	Viet Nam
// 	VU	Vanuatu
// 	WF	Wallis and Futuna Islands
// 	WS	Samoa
// 	YD	Yemen Democratic
// 	YE	Yemen
// 	YT	Mayotte
// 	YU	Yugoslavia
// 	ZA	South Africa
// 	ZM	Zambia
// 	ZR	Zaire
// 	ZW	Zimbabwe

// PA	MAS/DEMS Sub Type of Operation
// 	FD	Fixed Two-Way DEMS
// 	FI	Fixed One-Way Inbound MAS
// 	FO	Fixed One-Way Outbound MAS
// 	FT	Fixed Two-Way MAS
// 	MD	Multiple Two-Way DEMS
// 	MM	Mobile Master
// 	MO	Multiple One-Way Outbound MAS
// 	MT	Multiple Two-Way MAS

// RI	Entity Type Codes
// 	F	Applicant
// 	D	Disclosable Interest Holder (DIH)
// 	A	Affiliate

// RI	Year Sequence ID
// 	1	Most Recent Year
// 	2	One Year Prior
// 	3	Two Years Prior

// SH 	General Class of Ship
// 	MM	Merchant
// 	PL	Pleasure
// 	SV	Rescue
// 	FV	Fishing
// 	GV	Official Service Ship

// SH	Specific Class of Ship
// 	ACV	Air-cushion vehicle
// 	AUX	Auxiliary Ship
// 	CHA	Barge
// 	BLK	Bulk Carrier
// 	CBL	Cable Ship
// 	PMX	Cargo and Passenger
// 	CA?	Cargo Ship
// 	CAB	Coaster
// 	CON	Container Ship
// 	BTA	Factory Ship
// 	FBT	Ferry
// 	PH?	Fishing Vessel
// 	VDT	Hydrofoil
// 	MTB	Motorboat
// 	OIL	Oil Tanker
// 	TPO	Ore Carrier
// 	PA?	Passenger Ship
// 	PLT	Pilot Tender
// 	FRG	Reefer
// 	EXP	Research or Survey Ship
// 	VLR	Sailing Ship
// 	RAM	Salvage Ship
// 	SLO	Sloop
// 	RAV	Supply Vessel
// 	CIT	Tanker
// 	ECO	Training Ship
// 	TRA	Tramp
// 	CHR	Trawler
// 	TUG	Tug
// 	BLN	Whaler
// 	YAT	Yach

// SH	Type of Authorization
// 	F	Fleet License
// 	P	Portable License
// 	R	Regulare License

// 			Updated 10/01/2020
