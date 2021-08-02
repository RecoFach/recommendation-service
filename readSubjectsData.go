package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// contains all the fields covered by the computer science faculty
// based on which will be performed the recommendations
type UserPreferences struct {
	SOFTWARE    int `json:"software"`
	AI          int `json:"ai"`
	LOWLEVEL    int `json:"lowlevel"`
	SECURITY    int `json:"security"`
	WEB         int `json:"web"`
	THEORETICAL int `json:"theoretical"`
}

//Presenting the Language in which a course can be taught
type Language struct {
	German  string `json:"german"`
	English string `json:"english"`
}

type CourseType struct {
	Komplexpraktikum int `json:"komplexpraktikum"`
	Seminar          int `json:"seminar"`
	Vorlesung        int `json:"vorlesung"`
}

type Semester struct {
	Sommersemester int `json:"sommersemester"`
	Wintersemester int `json:"wintersemester"`
}

//representing a whole subjects with all the fields it contains based on the annotation
type SubjectsCsvStruct struct {
	Index       int    `json:"index"`
	CourseName  string `json:"course_name"`
	Link        string `json:"link"`
	Sws         int    `json:"sws"`
	Semester    Semester
	Coursetype  CourseType
	Language    Language
	Preferences UserPreferences
}

//reading the subjects .csv file with the correct deliminations
func readSubjectsCSV() [][]string {
	csvFile, err := os.Open("subjects.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return csvData
}

//fill the read Data into the SubjectsCsvStruct data structure
func fillStructWithCsvData(csvData [][]string) []SubjectsCsvStruct {
	var allRecords []SubjectsCsvStruct
	var oneRecord SubjectsCsvStruct

	for index, each := range csvData {
		oneRecord.Index = index
		oneRecord.CourseName = each[0]
		oneRecord.Link = each[1]
		oneRecord.Sws, _ = strconv.Atoi(each[2])
		oneRecord.Language.German = each[3]
		oneRecord.Language.English = each[4]
		oneRecord.Coursetype.Komplexpraktikum, _ = strconv.Atoi(each[5])
		oneRecord.Coursetype.Seminar, _ = strconv.Atoi(each[6])
		oneRecord.Coursetype.Vorlesung, _ = strconv.Atoi(each[7])
		oneRecord.Preferences.SOFTWARE, _ = strconv.Atoi(each[8])
		oneRecord.Preferences.AI, _ = strconv.Atoi(each[9])
		oneRecord.Preferences.LOWLEVEL, _ = strconv.Atoi(each[10])
		oneRecord.Preferences.SECURITY, _ = strconv.Atoi(each[11])
		oneRecord.Preferences.WEB, _ = strconv.Atoi(each[12])
		oneRecord.Preferences.THEORETICAL, _ = strconv.Atoi(each[13])
		oneRecord.Semester.Sommersemester, _ = strconv.Atoi(each[14])
		oneRecord.Semester.Wintersemester, _ = strconv.Atoi(each[15])

		allRecords = append(allRecords, oneRecord)
	}
	return allRecords
}

func recordsToJson(allRecords PairList) []byte {
	jsonData, err := json.Marshal(allRecords) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check
	// NOTE : You can stream the JSON data to http service as well instead of saving to file
	fmt.Println(string(jsonData))
	return jsonData

}
