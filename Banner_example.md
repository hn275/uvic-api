## Get terms

GET Request:

```go
url := "https://banner.uvic.ca/StudentRegistrationSsb/ssb/classSearch/getTerms?searchTerms=&offset=0&max=100"
```

Response:

```json
[
  {
    "code": "202301",
    "description": "Second Term: Jan - Apr 2023"
  },
  {
    "code": "202209",
    "description": "First Term: Sep - Dec 2022"
  },
  {
    "code": "202205",
    "description": "Summer Session: May - Aug 2022 (View Only)"
  },
  {
    "code": "202201",
    "description": "Second Term: Jan - Apr 2022 (View Only)"
  },
  {
    "code": "202109",
    "description": "First Term: Sep - Dec 2021 (View Only)"
  },
  {
    "code": "202105",
    "description": "Summer Session: May - Aug 2021 (View Only)"
  },
  {
    "code": "202101",
    "description": "Second Term: Jan - Apr 2021 (View Only)"
  },
  {
    "code": "202009",
    "description": "First Term: Sep - Dec 2020 (View Only)"
  },
  {
    "code": "202005",
    "description": "Summer Session: May - Aug 2020 (View Only)"
  },
  {
    "code": "202001",
    "description": "Second Term: Jan - Apr 2020 (View Only)"
  }
]
```

## Get subjects

GET request:

```go
url := "https://banner.uvic.ca/StudentRegistrationSsb/ssb/classSearch/get_subject?searchTerm=&offset=1&max=100000&term=202301"
```

Response:

```json
[
  {
    "code": "ATWP",
    "description": "Academic and Technical Writing"
  },
  {
    "code": "ASL",
    "description": "American Sign Language"
  },
  {
    "code": "ANTH",
    "description": "Anthropology"
  },
  {
    "code": "ART",
    "description": "Art"
  },
  {
    "code": "AE",
    "description": "Art Education"
  },
  {
    "code": "AHVS",
    "description": "Art History &amp; Visual Studies"
  },
  {
    "code": "ASTR",
    "description": "Astronomy"
  },
  {
    "code": "BIOC",
    "description": "Biochemistry"
  },
  {
    "code": "BCMB",
    "description": "Biochemistry and Microbiology"
  },
  {
    "code": "BIOL",
    "description": "Biology"
  },
  {
    "code": "BME",
    "description": "Biomedical Engineering"
  },
  {
    "code": "BUS",
    "description": "Business"
  },
  {
    "code": "CS",
    "description": "Canadian Studies"
  },
  {
    "code": "CHEM",
    "description": "Chemistry"
  },
  {
    "code": "CYC",
    "description": "Child and Youth Care"
  },
  {
    "code": "CIVE",
    "description": "Civil Engineering"
  },
  {
    "code": "COM",
    "description": "Commerce"
  },
  {
    "code": "CD",
    "description": "Community Development"
  },
  {
    "code": "CSC",
    "description": "Computer Science"
  },
  {
    "code": "CW",
    "description": "Creative Writing"
  },
  {
    "code": "CH",
    "description": "Cultural Heritage"
  },
  {
    "code": "CSPT",
    "description": "Cultural, Social &amp; Political T"
  },
  {
    "code": "EDCI",
    "description": "Curriculum and Instruction"
  },
  {
    "code": "DR",
    "description": "Dispute Resolution"
  },
  {
    "code": "EOS",
    "description": "Earth and Ocean Sciences"
  },
  {
    "code": "ECON",
    "description": "Economics"
  },
  {
    "code": "ED-D",
    "description": "Ed Psyc and Leadership Studies"
  },
  {
    "code": "ECE",
    "description": "Electrical &amp; Computer Engineer"
  },
  {
    "code": "ENGR",
    "description": "Engineering"
  },
  {
    "code": "ENGL",
    "description": "English"
  },
  {
    "code": "ENT",
    "description": "Entrepreneurship and Small Bus"
  },
  {
    "code": "ER",
    "description": "Environmental Restoration"
  },
  {
    "code": "ES",
    "description": "Environmental Studies"
  },
  {
    "code": "EUS",
    "description": "European Studies"
  },
  {
    "code": "EPHE",
    "description": "Exercise Sc, Phys &amp; Health Ed"
  },
  {
    "code": "FA",
    "description": "Fine Arts"
  },
  {
    "code": "FORB",
    "description": "Forest Biology"
  },
  {
    "code": "FRAN",
    "description": "French (FRAN)"
  },
  {
    "code": "GNDR",
    "description": "Gender Studies"
  },
  {
    "code": "GEOG",
    "description": "Geography"
  },
  {
    "code": "GMST",
    "description": "Germanic Studies"
  },
  {
    "code": "GDS",
    "description": "Global Development Studies"
  },
  {
    "code": "GS",
    "description": "Graduate Studies"
  },
  {
    "code": "GREE",
    "description": "Greek"
  },
  {
    "code": "GRS",
    "description": "Greek and Roman Studies"
  },
  {
    "code": "HINF",
    "description": "Health Information Science"
  },
  {
    "code": "HLTH",
    "description": "Health and Community Services"
  },
  {
    "code": "HSTR",
    "description": "History"
  },
  {
    "code": "HDCC",
    "description": "Human Dimensions - Climate Chg"
  },
  {
    "code": "HUMA",
    "description": "Humanities"
  },
  {
    "code": "IED",
    "description": "Indigenous Education"
  },
  {
    "code": "IGOV",
    "description": "Indigenous Governance"
  },
  {
    "code": "IN",
    "description": "Indigenous Nationhood"
  },
  {
    "code": "INGH",
    "description": "Indigenous Peoples Health"
  },
  {
    "code": "IS",
    "description": "Indigenous Studies"
  },
  {
    "code": "ISP",
    "description": "Intercultural Studies and Prac"
  },
  {
    "code": "INTD",
    "description": "Interdisciplinary"
  },
  {
    "code": "IB",
    "description": "International Business"
  },
  {
    "code": "INTS",
    "description": "International Health Studies"
  },
  {
    "code": "ITAL",
    "description": "Italian"
  },
  {
    "code": "LATI",
    "description": "Latin"
  },
  {
    "code": "LAS",
    "description": "Latin American Studies"
  },
  {
    "code": "LAW",
    "description": "Law"
  },
  {
    "code": "LING",
    "description": "Linguistics"
  },
  {
    "code": "MM",
    "description": "Master in Management"
  },
  {
    "code": "MGB",
    "description": "Masters of Global Business"
  },
  {
    "code": "MBA",
    "description": "Masters-Business Administratio"
  },
  {
    "code": "MATH",
    "description": "Math"
  },
  {
    "code": "MECH",
    "description": "Mechanical Engineering"
  },
  {
    "code": "MEDS",
    "description": "Medical Sciences"
  },
  {
    "code": "MEDI",
    "description": "Medieval Studies"
  },
  {
    "code": "MICR",
    "description": "Microbiology"
  },
  {
    "code": "MUS",
    "description": "Music"
  },
  {
    "code": "NRSC",
    "description": "Neuroscience"
  },
  {
    "code": "NUED",
    "description": "Nurse Educator"
  },
  {
    "code": "NURS",
    "description": "Nursing"
  },
  {
    "code": "NURA",
    "description": "Nursing Advanced Practice"
  },
  {
    "code": "NUNP",
    "description": "Nursing Practitioners"
  },
  {
    "code": "NUHI",
    "description": "Nursing/Health Info Science"
  },
  {
    "code": "PAAS",
    "description": "Pacific and Asian Studies"
  },
  {
    "code": "PHIL",
    "description": "Philosophy"
  },
  {
    "code": "PHYS",
    "description": "Physics"
  },
  {
    "code": "POLI",
    "description": "Political Science"
  },
  {
    "code": "PSYC",
    "description": "Psychology"
  },
  {
    "code": "ADMN",
    "description": "Public Administration"
  },
  {
    "code": "PHSP",
    "description": "Public Health &amp; Social Policy"
  },
  {
    "code": "RCS",
    "description": "Religion, Culture and Society"
  },
  {
    "code": "SCIE",
    "description": "Science"
  },
  {
    "code": "SMGT",
    "description": "Service Management"
  },
  {
    "code": "SLST",
    "description": "Slavic Studies"
  },
  {
    "code": "SDH",
    "description": "Social Dimensions of Health"
  },
  {
    "code": "SJS",
    "description": "Social Justice Studies"
  },
  {
    "code": "SOSC",
    "description": "Social Sciences"
  },
  {
    "code": "SOCW",
    "description": "Social Work"
  },
  {
    "code": "SOCI",
    "description": "Sociology"
  },
  {
    "code": "SENG",
    "description": "Software Engineering"
  },
  {
    "code": "SPAN",
    "description": "Spanish"
  },
  {
    "code": "STAT",
    "description": "Statistics"
  },
  {
    "code": "ED-P",
    "description": "Teacher Ed: Seminar &amp; Pract."
  },
  {
    "code": "TS",
    "description": "Technology and Society"
  },
  {
    "code": "THEA",
    "description": "Theatre"
  },
  {
    "code": "WRIT",
    "description": "Writing"
  }
]
```
