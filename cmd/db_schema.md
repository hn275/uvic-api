# DB Schema for UVic API

## NOTE

Here is my assumption about the scheduling:

- Same time, diff days: ie, 10:30am - 11:20am for TWF
- Need to ask discord, but the following schema would be built based of this schema

**CLASSES TABLE**

- id: 1
- subject: ATWP
- subject_description: Academic and Technical Writing
- course: 305
- course_description: The Rhetoric of Health and Medicine

**SECTIONS TABLE**

- id: 1
- class_id: 1
- instructor_id: 1
- crn: 20280
- section: A01
- type: Lecture
- method: Face-to-face
- capacity: 20
- start_date: Jan 09, 2023
- end_date: Apr 06, 2023

**SCHEDULES TABLE**

- id: 1
- section_id: 1
- building: David Strong Building
- room: C114
- start_time: 10:30am
- end_time: 11:20am
- day: [Monday, Tuesday, Wednesday]

**INSTRUCTORS TABLE**

- id: 1
- name: John Doe
- email: jdoe@uvic.ca
