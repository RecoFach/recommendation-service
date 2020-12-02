# Recommendation

## ToC

Supported methods:

- `post` - create new `recommendation` based on a certain user preferences


## POST 
**Path** - `/recommend`
To create new Subject perform a put request on /subjects path with json-body. All fields are required.

Table of parameters:

| Name                 | Type |
|----------------------|------|
| Software engineering | Int  |
| AI                   | Int  |
| Low-level            | Int  |
| Security             | Int  |
| Web                  | Int  |
| Theoretical          | Int  |


Example:
{ 'Software engineering': 0,
  'AI': 0,
  'Low-level': 0,
  'Security': 1,
  'Web': 0,
  'Theoretical': 0}


Curl command:
`curl --location --request POST 'http://0.0.0.0:2000/recommend?Software%20engineering=0&AI=0&Low-level=0&Security=1&Web=0&Theoretical=0' \
--header ': '`

Response:
`{"0":{"Course name":"Cryptography and
Cryptanalysis","Link":"https:\/\/tu-dresden.de\/ing\/informatik\/sya\/ps\/studium\/lectures\/crypto","sws":4,"german":1,"english":0,"Komplexpraktikum":0,"Seminar":0,"Vorlesung":1,"Software
engineering":0,"AI":0,"Low-level":0,"Security":1,"Web":0,"Theoretical":0,"Sommersemester":1,"Wintersemester":0},"1":{"Course
name":"Security and Cryptography
II","Link":"https:\/\/tu-dresden.de\/ing\/informatik\/sya\/ps\/studium\/lectures\/sac-ii?set_language=en","sws":4,"german":0,"english":1,"Komplexpraktikum":0,"Seminar":0,"Vorlesung":1,"Software
engineering":0,"AI":0,"Low-level":0,"Security":1,"Web":0,"Theoretical":0,"Sommersemester":1,"Wintersemester":0},"2":{"Course
name":"Seminar on
Privacy","Link":"https:\/\/tu-dresden.de\/ing\/informatik\/sya\/ps\/studium\/seminars\/ps-sic\/?set_language=de","sws":2,"german":1,"english":0,"Komplexpraktikum":0,"Seminar":1,"Vorlesung":0,"Software
engineering":0,"AI":0,"Low-level":1,"Security":1,"Web":0,"Theoretical":0,"Sommersemester":1,"Wintersemester":1},"3":{"Course
name":"Advanced Practical Course Cryptography and
Security","Link":"https:\/\/tu-dresden.de\/ing\/informatik\/sya\/ps\/studium\/seminars\/kp-crypto","sws":4,"german":1,"english":1,"Komplexpraktikum":1,"Seminar":0,"Vorlesung":0,"Software
engineering":0,"AI":0,"Low-level":1,"Security":1,"Web":0,"Theoretical":0,"Sommersemester":1,"Wintersemester":1},"4":{"Course
name":"Microkernel-Based Operating
Systems","Link":"https:\/\/tu-dresden.de\/ing\/informatik\/sya\/professur-fuer-betriebssysteme\/studium\/vorlesungen\/mos","sws":3,"german":1,"english":0,"Komplexpraktikum":0,"Seminar":0,"Vorlesung":1,"Software
engineering":0,"AI":0,"Low-level":1,"Security":1,"Web":0,"Theoretical":0,"Sommersemester":0,"Wintersemester":1}}`